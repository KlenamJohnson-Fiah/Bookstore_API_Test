package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// const (
// 	mysql_username = "mysql_username"
// 	mysql_password = "mysql_password"
// 	//mysql_host     = "mysql_host"
// 	mysql_schema = "mysql_schema"
// )

var (
	Client *sql.DB

	// 	username = os.Getenv(mysql_username)
	// 	password = os.Getenv(mysql_password)
	// 	//host     = os.Getenv(mysql_host)
	// 	schema = os.Getenv(mysql_schema)
)

// func readENVvariables(key string) string {

// 	// load .env file
// 	err := godotenv.Load("/Users/klenam/Documents/go/src/bookstore_users-api/utils/.env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	return os.Getenv(key)
// }

//%s:%s@tcp(%s)/%s "sslmode=disable" user:password@/dbname
func init() {
	envErr := godotenv.Load("/Users/klenam/Documents/go/src/bookstore_users-api/utils/.env")

	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}

	datasourceName := fmt.Sprintf("%s:%s@tcp/%s", os.Getenv("mysql_username"), os.Getenv("mysql_password"), os.Getenv("mysql_schema"))
	//"root",
	//"testdatabase",
	//"127.0.0.1",
	//"user_db",
	//)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("UserDB Database successful initiated")
}
