package config

type EnvConfig interface {
	GetDevEnv() string
}

type EnvConfigImpl struct {
}

func (e *EnvConfigImpl) GetDevEnv() string {
	return "dva"
}
func NewEnvConfig() EnvConfig {
	return &EnvConfigImpl{}
}
