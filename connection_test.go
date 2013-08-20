package sql_utils

import (
//  "testing"
  "code.google.com/p/gcfg"
)

type DbConf struct {
  User      string
  Password  string
  DbName    string
}

func getConfiguration() DbConf {
  var cfg DbConf

  err := gcfg.ReadFileInto(&cfg, "config.gcfg")
  if err != nil {
    panic(err)
  }

 return cfg
}
