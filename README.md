# tfc-transfer-validator
## Documentation
protobuff for the `.protoc` file
`https://developers.google.com/protocol-buffers/docs/proto3`

## Running requirements
Necessary install for the micro code generation plugin 
`go get github.com/micro/protoc-gen-micro`

_this plugin change the way to implement the server code, see this documentation [https://github.com/micro/protoc-gen-micro]()_

If imports fail initialise the go module 
`go mod init github.com/the-final-codedown/tfc-transfer-validator`

To update the proto generated code run the following
`protoc -I. --go_out=plugins=micro:. --micro_out=. transfer_service\transfer-
service.proto`

To run the server you must will need the go micro api to provide a gateway
`go get github.com/micro/micro`

To start it :
`micro api`

## Dependency
The project depends on the tfc-cap-updater of the same project to work
A server should be running and available for this to work

## Testing
To run a mongo db should be running on localhost:27017
Carefull as editions will be done, the running databse should be scrappable

A tfc cap updater should also be available
