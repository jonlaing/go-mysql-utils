package sql_utils

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Connection struct {
  Username  string
  Password  string
  DbName    string
}

func (c Connection) Open() *sql.DB{
  var err error

  db, err := sql.Open("mysql", connectionString(c.Username, c.Password, c.DbName) )
  
  if err != nil {
    panic(err)
  }

  return db
}

func connectionString(username string, password string, dbname string) string {
  return username + ":" + password + "@/" + dbname
}
