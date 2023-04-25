package rate_limiter

import (
	"tz/iternal/rate-limiter/dto"
	"tz/pkg/configuration"
	"tz/pkg/logging"
)

type Service struct {
	log    *logging.Logger
	config configuration.ApplicationConfig
}

func NewService() *Service {
	return &Service{
		log:    logging.GetLogger(),
		config: configuration.GetConfig(),
	}
}

func (s *Service) Update(data dto.UpdateDTO) {
	if data.Second == 0 {
		data.Second = s.config.LimitSecondRequest
	}

	if data.Request == 0 {
		data.Request = s.config.LimitRequest
	}

	limiter = newIPRateLimiter(data.Second, data.Request, s.config.IpPrefix)
}
