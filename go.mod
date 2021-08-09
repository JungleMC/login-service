module github.com/JungleMC/login-service

go 1.16

require (
	github.com/JungleMC/sdk v0.0.0-20210808090323-174ad2884541
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/go-redis/redis/v8 v8.11.2
	github.com/google/uuid v1.3.0
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/JungleMC/sdk => ../sdk
