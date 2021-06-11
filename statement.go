package gomysql

import "fmt"

// Return as String Array
func (db *MySQLConnection) Sread(query string, params ...interface{}) ([]map[string]string, error) {
	db.checkConnection()
	stmt, err := db.dataBase.Prepare(query)
	if err != nil {
		// queryLog("Create Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)
	if err != nil {
		// queryLog("Pass Parameters to Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		// queryLog("Pass Parameters to Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return nil, err
	}
	count := len(columns)
	if count == 0 {
		return nil, nil
	}
	tableData := make([]map[string]string, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for i := 0; i < count; i++ {
		valuePtrs[i] = &values[i]
	}
	for rows.Next() {
		rows.Scan(valuePtrs...)
		entry := make(map[string]string)
		for i, col := range columns {
			var v string
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = fmt.Sprintf("%v", val)
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	return tableData, nil
}

// Return as Dynamic Array
func (db *MySQLConnection) Read(query string, params ...interface{}) ([]map[string]interface{}, error) {
	db.checkConnection()
	paramPointers := make([]interface{}, len(params))
	for i := range params {
		paramPointers[i] = &params[i]
	}
	stmt, err := db.dataBase.Prepare(query)
	if err != nil {
		// queryLog("Create Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(paramPointers...)
	if err != nil {
		// queryLog("Pass Parameters to Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	if count == 0 {
		return nil, nil
	}
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	return tableData, nil
}

//Insert and return the inserted status
func (db *MySQLConnection) Insert(query string, params ...interface{}) (bool, error) {
	db.checkConnection()
	stmt, err := db.dataBase.Prepare(query)
	if err != nil {
		// queryLog("Create Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(params...)
	if err != nil {
		// queryLog("Execute Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return false, err
	}
	a, err := res.RowsAffected()
	if err != nil {
		// queryLog("RowsAffected Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return false, err
	}
	if a > 0 {
		return true, nil
	}
	return false, nil
}

//Insert and return the  Auto Generated Id
func (db *MySQLConnection) InsertAndGetAutoGenId(query string, params ...interface{}) (int64, error) {
	db.checkConnection()
	stmt, err := db.dataBase.Prepare(query)
	if err != nil {
		// queryLog("Create Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return -1, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(params...)
	if err != nil {
		// queryLog("Execute Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return -1, err
	}

	result, err := res.LastInsertId()
	if err != nil {
		// queryLog("LastInsertId Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return -1, err
	}
	return result, nil
}

//Update andd return theas bool
func (db *MySQLConnection) Update(query string, params ...interface{}) (bool, error) {
	db.checkConnection()
	stmt, err := db.dataBase.Prepare(query)
	if err != nil {
		// queryLog("Create Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(params...)
	if err != nil {
		// queryLog("Execute Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return false, err
	}
	a, err := res.RowsAffected()
	if err != nil {
		// queryLog("RowsAffected Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return false, err
	}
	if a > 0 {
		return true, nil
	}
	return false, nil
}

//Delete and return the Delete status
func (db *MySQLConnection) Delete(query string, params ...interface{}) (bool, error) {
	db.checkConnection()
	stmt, err := db.dataBase.Prepare(query)
	if err != nil {
		// queryLog("Create Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(params...)
	if err != nil {
		// queryLog("Execute Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return false, err
	}
	a, err := res.RowsAffected()
	if err != nil {
		// queryLog("RowsAffected Statement  \n Query => "+query+" \n Error => "+ err.Error())
		return false, err
	}
	if a > 0 {
		return true, nil
	}
	return false, nil
}
