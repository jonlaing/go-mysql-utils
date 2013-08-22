package mysql_utils

import (
  "testing"
  "strings"
)

type Model struct {
  Id    int     `mysql:"id INT NOT NULL AUTO_INCREMENT,pk"`
  Name  string  `mysql:"name VARCHAR(20) NOT NULL"`
}

func TestParseFields(t *testing.T) {
  var s MysqlUtil
  fields, pks := s.parseFields(&Model{})
  
  if fields[0] != "id INT NOT NULL AUTO_INCREMENT" {
    t.Error("First field of Model expected \"id INT NOT NULL AUTO_INCREMENT\", but got " + fields[0])
  }

  if fields[1] != "name VARCHAR(20) NOT NULL" {
    t.Error("Second field of Model expected \"name VARCHAR(20) NOT NULL\", but got " + fields[1])
  }
  
  if len(pks) != 1 {
    t.Error("Expected single primary key")
  }

  if pks[0] != "id" {
    t.Error("Expected primary key on \"id\"")
  }
}

func TestBuildingCreateStatement(t *testing.T) {
  var s MysqlUtil
  statement := s.buildCreateTableStatement("model", &Model{})

  if len(statement) < 1 {
    t.Error("Expected to be string with length")
  }

  if !strings.Contains(statement, "CREATE TABLE") {
    t.Error("Expected \"CREATE TABLE\", instead got " + statement)
  }
}

func TestFieldList(t *testing.T) {
  var s MysqlUtil
  list := s.FieldList(&Model{})

  if list[0] != "id" || list[1] != "name" {
    t.Error("Expected field list of \"id\",\"name\" got " + strings.Join(list, ","))
  }
}

func TestTableCreation(t *testing.T) {
  var s MysqlUtil
  s.Conn.GetConfiguration("config.gcfg","test")

  if s.TableExists("model") {
    err := s.DropTable("model")

    if err != nil {
      t.Error("Error removing table for creation")
    }
  }

  err := s.CreateTable("model", &Model{})
  if err != nil {
    t.Error("Error creating table \"model\": ", err)
  }

  if !s.TableExists("model") {
    t.Error("Table should exist")
  }
}

func TestTableDrop(t *testing.T) {
  var s MysqlUtil
  s.Conn.GetConfiguration("config.gcfg", "test")

  if !s.TableExists("model") {
    err := s.CreateTable("model", &Model{})
    if err != nil {
      t.Error("Error creating table \"model\" for deletion: ", err)
    }
  }

  err := s.DropTable("model")

  if err != nil {
    t.Error("Error dropping table \"model\": ", err)
  }

  if s.TableExists("model") {
    t.Error("Table shouldn't exist")
  }
}
