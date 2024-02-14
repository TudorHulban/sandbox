package testcontainerflyway

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

type ConfigFlyway struct {
	migrationsPath string

	dbName   string
	user     string
	password string
	port     uint
}

func (c ConfigFlyway) AsURL() string {
	return fmt.Sprintf(
		"jdbc:postgresql://localhost:%d/%s",
		c.port,
		c.dbName,
	)
}

type ContainerFlyway struct {
	testcontainers.Container

	*ConfigFlyway
}

func NewContainerFlyway(ctx context.Context, cfg *ConfigFlyway, options ...testcontainers.ContainerCustomizer) (*ContainerFlyway, error) {
	req := testcontainers.ContainerRequest{
		Image: "flyway/flyway:8.2.2",
		Env: map[string]string{
			"URL":      cfg.AsURL(),
			"USER":     cfg.user,
			"PASSWORD": cfg.password,
		},
		Files: []testcontainers.ContainerFile{
			{
				HostFilePath: cfg.migrationsPath,
			},
		},
	}

	genericContainerReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	for _, option := range options {
		option.Customize(
			&genericContainerReq,
		)
	}

	container, errContainer := testcontainers.GenericContainer(
		ctx,
		genericContainerReq,
	)
	if errContainer != nil {
		return nil,
			errContainer
	}

	return &ContainerFlyway{
			Container:    container,
			ConfigFlyway: cfg,
		},
		nil
}
