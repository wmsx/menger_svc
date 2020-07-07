module github.com/wmsx/menger_svc

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.4.0
	google.golang.org/protobuf v1.25.0
)

// 替换为v1.26.0版本的gRPC库
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
