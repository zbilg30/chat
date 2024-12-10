package config

import "time"

type (
	Mongo struct {
		Uri      string
		Database string
		Timeout  time.Duration
	}
	Jwt struct {
		SigningKey []byte `mapstructure:"signing-key"`
		Auth       Auth
	}
	Auth struct {
		Expiry time.Duration
	}
	Config struct {
		Jwt   Jwt
		Mongo Mongo
	}
)

func NewConfig() (Config, error) {

	return Config{Mongo: Mongo{
		Uri:      "",
		Database: "ecommerce",
		Timeout:  10 * time.Second,
	},
		Jwt: Jwt{SigningKey: []byte("secret-key"), Auth: Auth{Expiry: time.Hour * 24}}}, nil
}
