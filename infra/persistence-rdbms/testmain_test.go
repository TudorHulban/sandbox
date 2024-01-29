package main

import (
	"testing"

	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMain(m *testing.M) {
	coPostgres, err := GenericContainer(ctx, GenericContainerRequest{
		ContainerRequest: ContainerRequest{
			Image:        "nginx:1.17.6",
			ExposedPorts: []string{"80/tcp"},
			WaitingFor:   wait.ForListeningPort("80/tcp"),
			Files: []ContainerFile{
				{
					HostFilePath:      "./testdata/hello.sh",
					ContainerFilePath: "/copies-hello.sh",
					FileMode:          0o700,
				},
			},
		},
		Started: false,
	})
}
