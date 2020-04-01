package config

// TypeDb DB数据结构
type TypeDb struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// Db 数据库配置
var Db TypeDb

func dbInit() {
	switch Current {
	case Dev:
		Db = TypeDb{
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "123456",
			Database: "go2fe-demo",
		}
	case Test:
		Db = TypeDb{
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "123456",
			Database: "go2fe-demo",
		}
	case Prod:
		fallthrough
	default:
		Db = TypeDb{
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "123456",
			Database: "go2fe-demo",
		}
	}
}
