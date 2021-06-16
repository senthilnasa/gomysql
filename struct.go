package gomysql

import "sync"

// MySql Error Config
type ErrorLogConfig struct {
	ErrorApiurl    string
	IsPostRequest  bool
	ErrorFromFeild string
}

// MySql Config Struct
type MySQLConfig struct {
	Host       string
	Port       int
	User       string
	Pass       string
	DbName     string
	Sizeofpool int
	ErrorLog   ErrorLogConfig
}

// Mysql Connection Poll
type MySQLConnectionPool struct {
	sync.Mutex
	blocked     map[uint32]bool
	connections []MySQLConnection
}
