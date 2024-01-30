package main

import "fmt"

type ConfigPostgres struct {
	DBName   string `env:"POSTGRES_DB,required" envDefault:"sandbox" valid:"required"`
	User     string `env:"POSTGRES_USER" envDefault:"postgres" valid:"required"`
	Password string `env:"POSTGRES_PASSWORD,required" envDefault:"password"`
	Host     string `env:"DBHOST,required" envDefault:"localhost" valid:"required"`
	Port     int    `env:"DBPORT,required" envDefault:"5432" valid:"range(1024|65534)"`

	MaxIdleConnections           int  `env:"MAX_IDLE_CONNECTIONS" envDefault:"10" valid:"range(1|50)"`
	MaxOpenConnections           int  `env:"MAX_OPEN_CONNECTIONS" envDefault:"10" valid:"range(1|50)"`
	MaxConnectionLifetimeMinutes int  `env:"MAX_CONNECTION_LIFETIME_MINUTES" envDefault:"10" valid:"range(1|300)"`
	UseDefaultTransaction        bool `env:"USE_DEFAULT_TRANSACTION" envDefault:"false"`
}

func (cfg ConfigPostgres) AsDSNGORM() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)
}

func (cfg ConfigPostgres) AsDSNPGX() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
}

func NewConfigPostgres(testContainersDSN string) (*ConfigPostgres, error) {
	return nil, nil
}
