package mysql_utils

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "code.google.com/p/gcfg"
)

type Connection struct {
  Username  string
  Password  string
  Address   string
  Dbname    string
}

func (c Connection) Open() *sql.DB{
  var err error

  db, err := sql.Open("mysql", connectionString(c.Username, c.Password, c.Address, c.Dbname) )
  
  if err != nil {
    panic(err)
  }

  return db
}

func connectionString(username string, password string, address string, dbname string) string {
  return username + ":" + password + "@" + address + "/" + dbname
}

func (c *Connection) GetConfiguration(filename string, env string) {
  cfg := struct {
    Env map[string]*struct {
      Username  string
      Password  string
      Address   string
      Dbname    string
    }
  }{}

  err := gcfg.ReadFileInto(&cfg, filename)
  if err != nil {
    panic(err)
  }

  c.Username  = cfg.Env[env].Username
  c.Password  = cfg.Env[env].Password
  c.Address   = cfg.Env[env].Address
  c.Dbname    = cfg.Env[env].Dbname
}
