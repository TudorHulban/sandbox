package config

import "fmt"

type DBInfo struct {
	DBName     string `env:"POSTGRES_DB,required" envDefault:"sandbox" valid:"required"`
	DBUser     string `env:"POSTGRES_USER" envDefault:"postgres" valid:"required"`
	DBPassword string `env:"POSTGRES_PASSWORD,required" envDefault:"password"`
	Host       string `env:"DBHOST,required" envDefault:"localhost" valid:"required"`
	Port       uint   `env:"DBPORT,required" envDefault:"5432" valid:"range(1024|65534)"`
}

type ConfigPostgres struct {
	DBInfo

	MaxIdleConnections           int  `env:"MAX_IDLE_CONNECTIONS" envDefault:"10" valid:"range(1|50)"`
	MaxOpenConnections           int  `env:"MAX_OPEN_CONNECTIONS" envDefault:"10" valid:"range(1|50)"`
	MaxConnectionLifetimeMinutes int  `env:"MAX_CONNECTION_LIFETIME_MINUTES" envDefault:"10" valid:"range(1|300)"` //nolint:lll
	UseDefaultTransaction        bool `env:"USE_DEFAULT_TRANSACTION" envDefault:"false"`
}

func (cfg DBInfo) AsDSNGORM() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
}

func (cfg DBInfo) AsDSNPGX() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
}
