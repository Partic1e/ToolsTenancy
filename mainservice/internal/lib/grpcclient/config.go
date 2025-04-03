package grpcclient

type GrpcConfig struct {
	Port           string        `mapstructure:"port"`
	UserService    ServiceConfig `mapstructure:"UserService"`
	PaymentService ServiceConfig `mapstructure:"PaymentService"`
	AdService      ServiceConfig `mapstructure:"AdService"`
	RentService    ServiceConfig `mapstructure:"RentService"`
}

type ServiceConfig struct {
	Host string `mapstructure:"Host"`
	Port string `mapstructure:"Port"`
}
