#gpProto

This is a simple go micro services project. 
It uses gRPC for interprocess communication


#gRPC

The code for the gRPC communication  is generated using protobuf and the gRPC plugin. To regenerate those files, just protoc the following command from the root of the project. For example to regenerate the customer code:

```bash
 protoc -I=. --go_out=plugins=grpc:.  --go_opt=paths=source_relative customer/customer.proto 
```
That will place the generated customer.pb.go file in the customer directory.

Note that both protobuf and the gRPC pulgin must be installed locally to be able to run the protoc command.
To install these in a Mac:

1. ```brew install protobuf```
2. ```go get github.com/golang/protobuf/protoc-gen-go```
3. ```go install github.com/golang/protobuf/protoc-gen-go```
4. ```export PATH="$PATH:$(go env GOPATH)/bin"```

