package config

type ApplicationConfig struct {
	Jwt   Jwt         `yaml:"jwt"`
	Mongo MongoConfig `yaml:"mongo"`
}

type Jwt struct {
	Secret   string `yaml:"secret"`
	Audience string `yaml:"audience"`
	Issuer   string `yaml:"issuer"`
}
type MongoConfig struct {
	ConnectionString string            `yaml:"connectionString"`
	Database         string            `yaml:"database"`
	Collection       map[string]string `yaml:"collection"`
}
