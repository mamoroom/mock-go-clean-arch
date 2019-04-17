package config

type EnvConfig interface {
	GetDevEnv() string
}

type EnvConfigImpl struct {
}

func (e *EnvConfigImpl) GetDevEnv() string {
	return "dev"
}
func NewEnvConfig() EnvConfig {
	return &EnvConfigImpl{}
}
