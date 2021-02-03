package internal

type Config struct{
	PortVarName string
	AddVarName string
	PortDefault string
	AddDefault string

	DBHost string
	DBUser string
	DBPassword string
	DBName string
	DBPort string

	EtcdHost string
	EtcdPort string

	LogHost string
	LogPort string
}

func GetConfig() *Config{
	c := Config{}

	c.PortVarName = "PUERTO"
	c.AddVarName = "ADDRESS"
	c.PortDefault = "8080"
	c.AddDefault = "localhost"

	c.DBHost = "DB_HOST" 
	c.DBUser = "DB_USER"
	c.DBPassword = "DB_PASSWORD"
	c.DBName = "DB_NAME"
	c.DBPort = "DB_PORT"

	c.EtcdHost = "ETCD_HOST"
	c.EtcdPort = "ETCD_PORT"

	c.LogHost = "LOG_HOST"
	c.LogPort = "LOG_PORT"

	
	return &c
}