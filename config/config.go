package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	Secret struct {
		AccessSecretKey  string `env-required:"true" yaml:"access_secret_key" env:"ACCESS_SECRET_KEY"`
		RefreshSecretKey string `env-required:"true" yaml:"refresh_secret_key" env:"REFRESH_SECRET_KEY"`
	}

	Domain struct {
		Name string `env-required:"true" yaml:"name" env:"DOMAIN_NAME"`
	} `yaml:"domain"`
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	GRPC struct {
		Port string `env-required:"true" yaml:"port" env:"GRPC_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	Postgres struct {
		Host     string `env-required:"true" yaml:"host" env:"PG_HOST"`
		Port     string `env-required:"true" yaml:"port" env:"PG_PORT"`
		User     string `env-required:"true" yaml:"user" env:"PG_USER"`
		Password string `env-required:"true" yaml:"password" env:"PG_PASSWORD"`
		DBName   string `env-required:"true" yaml:"dbname" env:"PG_DBNAME"`
		PoolMax  int    `env-required:"false" yaml:"pool_max" env:"PG_POOL_MAX"`
	} `yaml:"postgres"`

	Oauth2Google struct {
		ClientID     string `env-required:"false" yaml:"client_id"`
		ClientSecret string `env-required:"false" yaml:"client_secret"`
		RedirectURI  string `env-required:"false" yaml:"redirect_uri"`
	} `yaml:"oauth2google"`
	MailSender struct {
		From     string `env-required:"false" yaml:"from"`
		Password string `env-required:"false" yaml:"password"`
		SMTPHost string `env-required:"false" yaml:"smtp_host"`
		SMTPPort string `env-required:"false" yaml:"smtp_port"`
	} `yaml:"mail_sender"`
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config.dev.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
