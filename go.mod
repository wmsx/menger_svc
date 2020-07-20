module github.com/wmsx/menger_svc

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.14
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.5.0
	github.com/wmsx/xconf v0.0.0-20200720194624-34a4108a8759
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	google.golang.org/protobuf v1.25.0
)

// 替换为v1.26.0版本的gRPC库
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
