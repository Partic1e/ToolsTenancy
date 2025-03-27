package rabbitclient

type RabbitMQConfig struct {
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Channel  string `yaml:"channel"`
}
