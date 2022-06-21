package main
///AUTHOR: Furkan Özgültekin
/*NOTES: The version provided is a crude version of what a todoList DB api might look like by myself
	The methodology goes as follows for an existing MySQL/MariaDB server with a created db of todolist and a todo table
	
	IF testing on main for my code, I suggest to have a table called todo in the working database of todos. I wrote my code according to that and for 
	default.Except for that everything is as close as I could get it to an API in golang
*/ 


import (
	"database/sql"
    "fmt"
    "log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	
	//"net/http"
	//"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
  )
  
  type Todo struct {
	  	ID uint `db:"id" gorm:"primaryKey"`
		Tag string `db:"tag"`
  }


var db *sql.DB



func addTodo(db *gorm.DB, s string, i uint){
	//Insert into specified index, Exec call is used for all alterations in the table
	db.Exec("INSERT INTO todo VALUES ( ?  ,  ? )", i , s)


}

func getAllTodos(db *gorm.DB, size int){
	// retrieve and print all todos from 1 to todo length
	for i := 1; i <= size; i++ {
		//SELECT * FROM TABLE WHERE id = i
		var result string
		db.Raw("SELECT tag FROM todo WHERE id = ?", i).Scan(&result)
		fmt.Println(result)
	  }
}

func removeByIndex(db *gorm.DB, i uint){
	// DELETE FROM todo WHERE id = ? i ; 
	db.Exec("DELETE FROM todo WHERE id = ?", i)
}

func updateTodo(db *gorm.DB, u string, i uint){
	// All of the database is handled by sql query code which is more familiar compared to gorm 
	db.Exec("UPDATE todo SET tag = ? WHERE id = ?", u, i)

}

func main() {


	//Tried to specify a json config then parse to dsn but runs into errors
    /*cfg := mysql.Config{
        User:   "root",
        Passwd: "toor",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "todos",
		AllowNativePasswords: true,
    }*/

	
    // Get a database handle.
    var err error
	
	//DSN object notation not yet specified, alter this to connect instead
    dsn := "root:toor@tcp(127.0.0.1:3306)/todos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
		log.Fatal(err)
    }
	// Gorm is opening a new opened dsn from mysql with default gorm configs
    fmt.Println("Connected!")
	
	//Testing of all functions
	fmt.Println("GetAllTodos:")
	getAllTodos(db, 2)
	
	fmt.Println("Adding assignments to id 4")	
	addTodo(db, "assignments", 4)
	
	fmt.Println("GetAllTodos:")
	getAllTodos(db, 6)


	fmt.Println("removing from id 4")	
	removeByIndex(db, 4)

	fmt.Println("Updating id 3 to be 'code golang'")	
	updateTodo(db, "code golang", 3)
	fmt.Println("GetAllTodos:")
	getAllTodos(db, 4)


}
