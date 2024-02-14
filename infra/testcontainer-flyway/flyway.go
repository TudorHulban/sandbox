package testcontainerflyway

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

type ConfigFlyway struct {
	MigrationsPath string

	DBName   string
	User     string
	Password string
	Port     uint
}

func (c ConfigFlyway) AsURL() string {
	return fmt.Sprintf(
		"jdbc:postgresql://localhost:%d/%s",
		c.Port,
		c.DBName,
	)
}

type ContainerFlyway struct {
	testcontainers.Container

	*ConfigFlyway
}

func RunContainer(ctx context.Context, cfg *ConfigFlyway, options ...testcontainers.ContainerCustomizer) (*ContainerFlyway, error) {
	req := testcontainers.ContainerRequest{
		Image: "flyway/flyway:8.2.2",
		Env: map[string]string{
			"FLYWAY_URL":      cfg.AsURL(),
			"FLYWAY_USER":     cfg.User,
			"FLYWAY_PASSWORD": cfg.Password,
		},
		Files: []testcontainers.ContainerFile{
			{
				HostFilePath:      cfg.MigrationsPath,
				ContainerFilePath: "/flyway/sql",
			},
		},
		Networks: []string{
			"host",
		},
		Cmd: []string{
			"migrate",
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
