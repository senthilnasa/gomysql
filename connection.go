package gomysql

import (
	"database/sql"
	"fmt"
	"hash/fnv"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xlab/closer"
)

func newMySQLConfig(param MySQLConfig) *MySQLConfig {
	// fmt.Println(fmt)
	return &MySQLConfig{param.Host, param.Port, param.User, param.Pass, param.DbName, param.Sizeofpool, param.ErrorLog}
}

type MySQLConnection struct {
	hash     uint32
	Config   *MySQLConfig
	dataBase *sql.DB
}

func CreateMySQLConnection(config *MySQLConfig) (MySQLConnection, error) {
	db := MySQLConnection{Config: config, hash: generateHash()}
	err := db.connect()
	if err != nil {
		// fmt.Println(err)
		return db, err
	}
	closer.Bind(db.Close)
	return db, nil
}

func (db *MySQLConnection) checkConnection() {
	err := db.dataBase.Ping()
	if err != nil {
		fmt.Println("DB reconnect")
		_ = db.connect()
	}
}

func (db *MySQLConnection) connect() error {
	dataBase, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		db.Config.User, db.Config.Pass, db.Config.Host, db.Config.Port, db.Config.DbName))
	if err != nil {
		return err
	}

	err = dataBase.Ping()
	if err != nil {
		return err
	}

	dataBase.SetConnMaxLifetime(10 * time.Minute)
	db.dataBase = dataBase
	return nil
}

func (db *MySQLConnection) Close() {
	_ = db.dataBase.Close()
}

func NewMySQLConnectionPool(param MySQLConfig) (*MySQLConnectionPool, error) {
	config := newMySQLConfig(param)
	pool := &MySQLConnectionPool{connections: make([]MySQLConnection, param.Sizeofpool), blocked: make(map[uint32]bool)}
	for i := 0; i < param.Sizeofpool; i++ {
		connection, err := CreateMySQLConnection(config)
		if err == nil {
			pool.connections[i] = connection
			pool.blocked[connection.hash] = false
		} else {
			connection.queryLog("Database Connection Error =>" + err.Error())
			return nil, err
		}
	}
	return pool, nil
}

func (pool *MySQLConnectionPool) Get() *MySQLConnection {

	pool.Lock()
	defer pool.Unlock()

	for _, con := range pool.connections {
		if !pool.blocked[con.hash] {
			pool.blocked[con.hash] = true
			con.checkConnection()
			return &con
		}
	}
	// Wait for 0.5 Second and check for avaliable Connection
	time.Sleep(500 * time.Millisecond)
	return pool.Get()
}

func (pool *MySQLConnectionPool) Release(connection *MySQLConnection) {
	pool.Lock()
	defer pool.Unlock()
	pool.blocked[connection.hash] = false
}

func (pool *MySQLConnectionPool) CloseAll() {
	for _, c := range pool.connections {
		pool.blocked[c.hash] = true
		c.Close()
	}
	fmt.Println("Closing all Mysql Pool")
}

// Generate hash value for all connections
func generateHash() uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	return hash.Sum32()
}
