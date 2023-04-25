package config

type DB struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
	Port     string `yaml:"port"`
}
