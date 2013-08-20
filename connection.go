package connection

import (
  "database/sql"
  "database/sql/driver"
  _ "github.com/go-sql-driver/mysql"
)

var Db driver.Conn

func Open(username string, password string, dbname string) *driver.Conn {
  var err error

  Db, err = sql.Open("mysql", connectionString(username, password, dbname) )
  
  if err != nil {
    panic(err)
  }

  return &Db
}
