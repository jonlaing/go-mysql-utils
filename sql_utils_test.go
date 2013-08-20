package sql_utils

import (
  "testing"
  "strings"
)

type Model struct {
  Id    int     `mysql:"id INT NOT NULL AUTO_INCREMENT,pk"`
  Name  string  `mysql:"name VARCHAR(20) NOT NULL"`
}

var sqlUtil SqlUtil

func TestParseFields(t *testing.T) {
  fields, pks := sqlUtil.parseFields(&Model{})
  
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
  statement := sqlUtil.buildCreateTableStatement("model", &Model{})

  if len(statement) < 1 {
    t.Error("Expected to be string with length")
  }

  if !strings.Contains(statement, "CREATE TABLE") {
    t.Error("Expected \"CREATE TABLE\", instead got " + statement)
  }
}

func TestFieldList(t *testing.T) {
  list := sqlUtil.FieldList(&Model{})

  if list[0] != "id" || list[1] != "name" {
    t.Error("Expected field list of \"id\",\"name\" got " + strings.Join(list, ","))
  }
}
