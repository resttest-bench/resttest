# requires xcode developer tools on OSX. run:
# xcode-select --install

dependencies:
	brew install protobuf
	brew install 
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go

proto:
	mkdir -p transactions
	protoc transactions.proto --go_out=plugins=grpc:transactions
