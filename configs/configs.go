package configs

var Config Configuration

type MQTTServer struct {
	Server   string
	UserName string
	Password string
	Topic    string
}

type Monitor struct {
	Port string
	Url  string
}

type Logger struct {
	Path string
}

type DatabaseConfiguration struct {
	Host            string
	Port            int
	User            string
	Password        string
	DB              string
	ConnMaxLifetime int
	MaxIdleConns    int
	MaxOpenConns    int
}

type Configuration struct {
	MQTTServer MQTTServer
	Database   DatabaseConfiguration
	Monitor    Monitor
	Logger     Logger
}
