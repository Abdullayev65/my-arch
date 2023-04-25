package config

import "time"

type JwtToken struct {
	ExpiringDuration time.Duration `yaml:"expiring_duration"`
	SecreteKey       string        `yaml:"secrete_key"`
}
