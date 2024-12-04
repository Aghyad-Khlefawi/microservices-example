.ONESHELL: # 

export GOPATH=/usr/local/go/bin
export GOBIN = /usr/local/go/bin
export PATH = $(shell printenv PATH):/usr/local/go/bin

up: infra-up dev-up
	

dev-up:
	cd deploy
	docker-compose -f docker-compose.yaml up -d corporate-api identity

dev-down:
	cd deploy
	docker-compose -f docker-compose.yaml down corporate-api identity

infra-up:
	cd deploy
	docker-compose -f docker-compose.yaml up -d mongodb servicebus

corporate-dev:
	cd ./src/corporate/
	dotnet run

identity-dev:
	cd ./src/identity/
	go run ./cmd/main.go

identity-build-proto:
	cd ./src/identity/
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	protoc -I ./protos --go_out=. --go-grpc_out=. ./protos/identity.proto



install-proto:
	sudo apt-get install unzip
	cd /usr/local/ 
	wget https://github.com/protocolbuffers/protobuf/releases/download/v29.0/protoc-29.0-linux-x86_64.zip
	sudo unzip protoc-29.0-linux-x86_64.zip -d /usr/local/protoc
	sudo mv /usr/local/protoc/bin/protoc /usr/local/bin/protoc
	sudo mv -v /usr/local/protoc/include/* /usr/local/include/.
	sudo rm -rf /usr/local/protoc
