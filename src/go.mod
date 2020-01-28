module github.com/the-final-codedown/tfc-transfer-validator

go 1.13

replace github.com/the-final-codedown/tfc-cap-updater => ../tfc-cap-updater

require (
	github.com/golang/protobuf v1.3.2
	github.com/jnewmano/grpc-json-proxy v0.0.0-20190711184636-c105eed9ab4a
	github.com/segmentio/kafka-go v0.3.4
	github.com/the-final-codedown/tfc-cap-updater v0.0.0-20191025113802-b27c9114dc82
	go.mongodb.org/mongo-driver v1.1.2
	google.golang.org/grpc v1.24.0
)
