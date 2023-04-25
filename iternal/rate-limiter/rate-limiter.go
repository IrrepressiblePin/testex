package rate_limiter

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"sync"
	"time"
	"tz/pkg/configuration"
	"tz/pkg/utils"
)

type iPRateLimiter struct {
	ips      map[string]*rate.Limiter
	mu       *sync.RWMutex
	ipPrefix string
	second   int
	request  int
}

var limiter *iPRateLimiter

func init() {
	config := configuration.GetConfig()
	limiter = newIPRateLimiter(config.LimitSecondRequest, config.LimitRequest, config.IpPrefix)
}

func newIPRateLimiter(second int, request int, ipPrefix string) *iPRateLimiter {
	return &iPRateLimiter{
		ips:      make(map[string]*rate.Limiter),
		mu:       &sync.RWMutex{},
		ipPrefix: ipPrefix,
		second:   second,
		request:  request,
	}
}

func (i *iPRateLimiter) addIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	l := rate.NewLimiter(rate.Every(time.Duration(i.second)*time.Second), i.request)
	i.ips[ip] = l
	return l
}

func (i *iPRateLimiter) getLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	l, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.addIP(ip)
	}
	i.mu.Unlock()
	return l
}

func RateLimiter() gin.HandlerFunc {
	return func(context *gin.Context) {
		if ip := context.GetHeader("X-Forwarded-For"); ip != "" {
			_, network, err := net.ParseCIDR(utils.Concat(ip, limiter.ipPrefix))
			if err != nil {
				context.AbortWithStatus(http.StatusBadRequest)
				return
			}
			if !limiter.getLimiter(network.String()).Allow() {
				context.AbortWithStatus(http.StatusTooManyRequests)
				return
			}
			context.Next()
			return
		}
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}
}
