package httpengine

type HttpConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
