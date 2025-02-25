package official

//go:generate protoc --proto_path=../ --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=Mhelloworld.proto=github.com/amikai/go_prac/official --go-grpc_opt=Mhelloworld.proto=github.com/amikai/go_prac/official helloworld.proto
