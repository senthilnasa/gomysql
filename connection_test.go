package gomysql_test

import (
	"fmt"
	"testing"

	"github.com/senthilnasa/gomysql"
)

func TestNewMySQLConnectionPool(t *testing.T) {
	errorConfig := gomysql.ErrorLogConfig{ErrorApiurl: "", IsPostRequest: true, ErrorFromFeild: "errData"}
	connectionConfig := gomysql.MySQLConfig{Host: "localhost", Port: 3306, User: "root", Pass: "nasa", DbName: "mysql", Sizeofpool: 100, ErrorLog: errorConfig}
	pool, err := gomysql.NewMySQLConnectionPool(connectionConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pool.CloseAll()
	fmt.Println("Open New Connection")
	connection := pool.Get()
	fmt.Println("Close  Connection")

	pool.Release(connection)

}
