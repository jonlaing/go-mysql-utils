package sql_utils

import (
  "testing"
)


func TestGettingConfiguration(t *testing.T) {
  var cfg Connection

  if cfg.Username != "" {
    t.Error("Username should be blank")
  }

  if cfg.Password != "" {
    t.Error("Password should be blank")
  }

  if cfg.Address != "" {
    t.Error("Address should be blank")
  }

  if cfg.Dbname != "" {
    t.Error("Dbname should be blank")
  }

  cfg.GetConfiguration("config.gcfg", "test")

  if cfg.Username != "test" {
    t.Error("Username should be \"test\", but got " + cfg.Username)
  }

  if cfg.Password != "test" {
    t.Error("Password should be \"test\", but got " + cfg.Password)
  }

  if cfg.Address != "" {
    t.Error("Address should be blank, but got " + cfg.Address)
  }

  if cfg.Dbname != "go_sql_test" {
    t.Error("Dbname should be \"go_sql_test\", but got " + cfg.Dbname )
  }
}

func TestConnectToSql(t *testing.T) {
  var c Connection

  c.GetConfiguration("config.gcfg", "test")

  db := c.Open()
  defer db.Close()
}
