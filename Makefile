# requires xcode developer tools on OSX. run:
# xcode-select --install

dependencies:
	brew install protobuf
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/Pallinder/go-randomdata

proto:
	mkdir -p transactions
	protoc transactions.proto --go_out=plugins=grpc:transactions
	
run-server:
	go run server/main.go	

run-client:
	go run client/main.go	
