package testcontainerflyway

import (
	"context"
	"fmt"

	"github.com/TudorHulban/GolangSandbox/config"
	"github.com/testcontainers/testcontainers-go"
)

type ConfigFlyway struct {
	MigrationsFiles []testcontainers.ContainerFile

	config.DBInfo
}

func (c ConfigFlyway) AsURL() string {
	return fmt.Sprintf(
		"jdbc:postgresql://%s:%d/%s",
		c.Host,
		c.Port,
		c.DBName,
	)
}

func RunContainerFlyway(ctx context.Context, cfg *ConfigFlyway, options ...testcontainers.ContainerCustomizer) error {
	reqContainer := testcontainers.ContainerRequest{
		Image: "flyway/flyway:8.2.2",
		Env: map[string]string{
			"FLYWAY_URL":      cfg.AsURL(),
			"FLYWAY_USER":     cfg.DBUser,
			"FLYWAY_PASSWORD": cfg.DBPassword,
		},
		Files:       cfg.MigrationsFiles,
		NetworkMode: "host",
		Cmd: []string{
			"migrate",
			"-community",
			fmt.Sprintf("-url=%s", cfg.AsURL()),
			fmt.Sprintf("-user=%s", cfg.DBUser),
			fmt.Sprintf("-password=%s", cfg.DBPassword),
		},
	}

	genericContainerReq := testcontainers.GenericContainerRequest{
		ContainerRequest: reqContainer,
		Started:          true,
	}

	for _, option := range options {
		option.Customize(
			&genericContainerReq,
		)
	}

	_, errContainer := testcontainers.GenericContainer(
		ctx,
		genericContainerReq,
	)
	if errContainer != nil && errContainer.Error() != "container exited with code 0: failed to start container" {
		return errContainer
	}

	return nil
}
