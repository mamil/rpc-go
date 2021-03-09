#!bash
#grpc
go get -u google.golang.org/grpc

# Protocol Buffers v3
wget https://github.com/google/protobuf/releases/download/v3.5.1/protobuf-all-3.5.1.zip
unzip protobuf-all-3.5.1.zip
cd protobuf-3.5.1/
./configure
make
make install

# Protoc Plugin
go get -u github.com/golang/protobuf/protoc-gen-go



