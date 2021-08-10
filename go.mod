module github.com/JungleMC/login-service

go 1.16

require (
	github.com/JungleMC/sdk v0.0.0-20210810042112-e30cdbe2f38a
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/go-redis/redis/v8 v8.11.2
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.7.0 // indirect
	google.golang.org/protobuf v1.27.1
)

replace github.com/JungleMC/sdk v0.0.0-20210809140359-e8dcfa68f6af => ../sdk
