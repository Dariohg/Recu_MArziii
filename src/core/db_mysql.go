package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Conn_MySQL struct {
	DB  *sql.DB
	Err string
}

func GetDBPool() *Conn_MySQL {
	error := ""
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error al cargar el archivo .env: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USUARIO")
	dbPass := os.Getenv("DB_PASSWORD")
	dbSchema := os.Getenv("DB_NOMBRE")

	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbUser == "" {
		dbUser = "root"
	}
	if dbSchema == "" {
		dbSchema = "productos_db"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		dbUser, dbPass, dbHost, dbSchema)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		error = fmt.Sprintf("error al abrir la base de datos: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * 60)

	if err := db.Ping(); err != nil {
		db.Close()
		error = fmt.Sprintf("error al verificar la conexión: %v", err)
	}

	return &Conn_MySQL{DB: db, Err: error}
}

func (conn *Conn_MySQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	return result, nil
}

func (conn *Conn_MySQL) FetchRows(query string, values ...interface{}) *sql.Rows {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		fmt.Printf("error al ejecutar la consulta SELECT: %v", err)
	}
	return rows
}
