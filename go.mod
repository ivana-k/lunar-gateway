module gateway

go 1.21.3

require (
	github.com/fullstorydev/grpcurl v1.8.8
	github.com/goccy/go-json v0.10.2
	github.com/gorilla/mux v1.8.0
	github.com/jhump/protoreflect v1.15.2
	google.golang.org/grpc v1.57.0
	gopkg.in/yaml.v3 v3.0.1
	iam-service v1.0.0
)

require (
	github.com/bufbuild/protocompile v0.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/c12s/oort => ../oort

replace github.com/c12s/magnetar => ../magnetar

replace iam-service => ../iam-service/iam-service
