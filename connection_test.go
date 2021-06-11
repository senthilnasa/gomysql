package gomysql

import (
	"fmt"
	"testing"
)

func TestNewMySQLConnectionPool(t *testing.T) {

	config := NewMySQLConfig("localhost", 3306, "root", "root", "gomysql")
	poolSize := 100

	pool, err := NewMySQLConnectionPool(poolSize, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pool.CloseAll()
	fmt.Println("Open New Connection")
	connection1 := pool.Get()
	fmt.Println("Close  Connection")

	pool.Release(connection1)

}
