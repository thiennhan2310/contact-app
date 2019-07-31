# Contact List app (GRPC)

## Requirements
1/ request - response 

- [x] add contact to list 
- [x] get contact detail 
- [x] search contact

2/ request - stream
- export list contact

3/ stream - response
- Import list contact

4/ stream - stream
- Auto complete contact  suggestion


1. Install protoc

Download package [https://github.com/protocolbuffers/protobuf/releases]

For OSX : choose 

Unzip

Move file 

[Binary file]

protoc3/bin/* /usr/local/bin/

[Prebuilt package from google]

protoc3/include/* /usr/local/include/

2. Generate protobu

protoc  --go_out=plugins=grpc:. protos/contact.proto
