package DefaultContainers

import (
	constants "RCTestSetup/Packages/Constants"
	"RCTestSetup/Packages/DockerSDK"

	"github.com/docker/go-connections/nat"
)

func LaunchMongoDbContainer(sdk DockerSDK.Docker, networkID string) (string, error) {
	constainerID, err := sdk.CreateContainer(
		networkID,
		"mongodb",
		constants.MongoDBImage,
		nat.PortSet{
			"27017/tcp": {},
		},
		nil,
		[]string{
			"MONGODB_REPLICA_SET_MODE=primary",
			"MONGODB_REPLICA_SET_NAME=rs0",
			"MONGODB_PORT_NUMBER=27017",
			"MONGODB_INITIAL_PRIMARY_HOST=mongodb",
			"MONGODB_INITIAL_PRIMARY_PORT_NUMBER=27017",
			"MONGODB_ADVERTISED_HOSTNAME=mongodb",
			"MONGODB_ENABLE_JOURNAL=true",
			"ALLOW_EMPTY_PASSWORD=yes",
		},
		map[string]struct{}{
			"/bitnami/mongodb": {},
		},
		nil,
		[]string{
			"mongodb",
			"mongo",
		},
		nil,
		nil,
	)

	return constainerID, err
}