package config

type AppConfig struct {
	PsqlConfig *PsqlConfig
	HttpConfig *HttpConfig
}

func New() *AppConfig {
	return &AppConfig{
		PsqlConfig: getPsqlConfig(),
		HttpConfig: getHttpConfig(),
	}
}
