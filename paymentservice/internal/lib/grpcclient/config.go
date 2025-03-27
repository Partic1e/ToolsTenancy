package grpcclient

type GrpcConfig struct {
	Port             int           `mapstructure:"port"`
	RentService      ServiceConfig `mapstructure:"rent_service"`
	GuaranteeService ServiceConfig `mapstructure:"guarantee_service"`
}

type ServiceConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
