package main

import (
	"fmt"

	"github.com/senthilnasa/gomysql"
)

func main() {

	errorConfig := gomysql.ErrorLogConfig{ErrorApiurl: "https://senthilnasa.me/api/mysqlgo/error_log/", ErrorFromFeild: "error", IsPostRequest: true}

	config := gomysql.MySQLConfig{Host: "localhost", Port: 3306, User: "root", Pass: "nasa", DbName: "gomysql", Sizeofpool: 100, ErrorLog: errorConfig}
	pool, err := gomysql.NewMySQLConnectionPool(config)
	if err != nil {
		panic(err)
	}
	defer pool.CloseAll()
	connection := pool.Get()

	// Select  Query
	data, err := connection.Read("SELECT * FROM student")
	if err != nil {
		fmt.Println("Error =>" + err.Error())
	} else {
		for _, v := range data {

			fmt.Println(v)
		}
	}

	// Select Query as String
	data1, err := connection.Sread("SELECT * FROM student")
	if err != nil {
		fmt.Println("Error =>", err.Error())
	} else {
		for _, v := range data1 {

			fmt.Println(v)
		}
	}

	// Insert Query
	isInserted, err := connection.Insert("insert into student(name,dob) values(?,?)", "senthil", "10-06-2021")
	if isInserted {
		fmt.Println("Student added successfully! ")
	} else {
		fmt.Println("Error in adding Student", err.Error())
	}

	// Insert and Get Auto Incremented Id
	StudentID, err := connection.InsertAndGetAutoGenId("insert into student(name,dob) values(?,?)", "senthil Nasa", "10-06-2020")
	if StudentID > 0 {
		fmt.Println("Student added successfully! \n Student Id=>", StudentID)
	} else {
		fmt.Println("Error in adding Student" + err.Error())
	}

	// Update Student
	isUpdated, err := connection.Update("UPDATE student SET name=? WHERE student_id=?", "New Name", StudentID)
	if isUpdated {
		fmt.Println("Student Updated successfully! ")
	} else {
		fmt.Println("Error in Updating Student", err.Error())
	}

	// Delete Student
	isDeleted, err := connection.Delete("DELETE FROM student WHERE student_id=? ", StudentID)
	if isDeleted {
		fmt.Println("Student Deleted successfully! ")
	} else {
		fmt.Println("Error in Deleting Student", err.Error())
	}

	pool.Release(connection)
}
