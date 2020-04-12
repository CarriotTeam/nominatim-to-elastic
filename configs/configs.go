package configs

var Config Configuration

type System struct {
	Url   string
	Threads        int
	DataPerRequest int
	Lng string
}

type Monitor struct {
	Port string
	Url  string
}

type Logger struct {
	Path string
}

type Elastic struct {
	Host     string
	Port     string
	Password string
	Topic    string
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
	Elastic  Elastic
	Database DatabaseConfiguration
	Monitor  Monitor
	Logger   Logger
	System   System
}
