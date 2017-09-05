package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HTTPStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
	configuration struct {
		Server, MySQLDBHost, MySQLDBUser, MySQLDBPwd, Database string
	}
)

// Display AppError
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HTTPStatus: code,
	}
	log.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[logAppConfig]: %s\n", err)
	}
}

// Session holds the MySQL session for database access
var session *gorm.DB

// get database session
func GetDbSession() *gorm.DB {
	if session == nil {
		var err error
		sqlConnection := AppConfig.MySQLDBUser + ":" + AppConfig.MySQLDBPwd + "@tcp(" + AppConfig.MySQLDBHost + ")/" + AppConfig.Database + "?charset=utf8&parseTime=True&loc=Local"
		session, err = gorm.Open("mysql", sqlConnection)
		if err != nil {
			log.Fatalf("[GetDbSession]: %s\n", err)
		}
	}
	return session
}

// Create database session
func createDbSession() {
	var err error
	sqlConnection := AppConfig.MySQLDBUser + ":" + AppConfig.MySQLDBPwd + "@tcp(" + AppConfig.MySQLDBHost + ")/" + AppConfig.Database + "?charset=utf8&parseTime=True&loc=Local"
	session, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		log.Fatalf("[CreateDbSession]: %s\n", err)
	}
}
