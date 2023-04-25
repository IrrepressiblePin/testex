FROM golang:latest

WORKDIR app/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o tz ./cmd/main/app.go

CMD ["./tz"]
