package utilities

import (
	"cimble/constants"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDatabase() (*gorm.DB, error) {
	var err error

	dbUrl := ConstructDatabaseUrl()
	db, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	return db, err
}

func DisconnectDatabase() {
	fmt.Println(`Closing db connection...`)
	dbInstance, err := db.DB()
	if err != nil {
		fmt.Printf(`Error closing db connection: %v`, err.Error())
		return
	}

	dbInstance.Close()
}

func GetDatabase() *gorm.DB {
	return db
}

func ConstructDatabaseUrl() (dbUrl string) {
	dbHost := GetEnvironmentVariableString(string(constants.DB_HOST))
	dbPort := GetEnvironmentVariableString(string(constants.DB_PORT))
	dbUser := GetEnvironmentVariableString(string(constants.DB_USER))
	dbPassword := GetEnvironmentVariableString(string(constants.DB_PASSWORD))
	dbName := GetEnvironmentVariableString(string(constants.DB_NAME))

	dbUrl = fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`, dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Printf("DB_URL: %s\n", dbUrl)
	return
}
