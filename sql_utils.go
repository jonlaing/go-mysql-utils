package sql_utils

import (
  "reflect"
  "strings"
  "regexp"
)

type SqlUtil struct {
  Conn Connection
}


func (s SqlUtil) CreateTable(tablename string, i interface{}) error {
  db := s.Conn.Open() 
  defer db.Close()

  create_statement := s.buildCreateTableStatement(tablename, i)

  _, err := db.Exec(create_statement)
  return err
}

func (s SqlUtil) DropTable(tablename string) error {
  db := s.Conn.Open()
  defer db.Close()

  _, err := db.Exec("DROP TABLE IF EXISTS " + tablename + ";")
  return err
}

func (s SqlUtil) TableExists(tablename string) bool {
  db := s.Conn.Open()
  defer db.Close()

  _, err := db.Query("SELECT * FROM "+tablename+";")

  if err != nil {
    return false
  }

  return true
}

func (s SqlUtil) FieldList(i interface{}) (list []string) {
  fields, _ := s.parseFields(i)

  for _, field := range fields {
    list = append(list, fieldName(field))
  }

  return
}

func (s SqlUtil) buildCreateTableStatement(tablename string, i interface{}) (create_statement string) {
  fields, pks := s.parseFields(i)
  create_statement = "CREATE TABLE IF NOT EXISTS " + tablename + " ("
  create_statement += strings.Join(fields, ",\n")
  
  if len(pks) > 0 {
    create_statement += ",\nPRIMARY KEY ("
    create_statement += strings.Join(pks, ",")
    create_statement += ")"
  }

  create_statement += ");"
  return
}

func (s SqlUtil) parseFields(i interface{}) (sqlfields, primaryKeys []string) {
  var fieldsql string

  f := reflect.TypeOf(i).Elem()
  for i := 0; i < f.NumField(); i++ {
    fieldsql = f.Field(i).Tag.Get("mysql")

    if strings.Contains(fieldsql, ",") {
      split := strings.Split(fieldsql, ",")
      sqlfields = append(sqlfields, split[0])

      if strings.Contains(split[1], "pk") {
        primaryKeys = append(primaryKeys, fieldName(split[0]))
      }

    } else {
      sqlfields = append(sqlfields, fieldsql)
    }
  }

  return
}

func fieldName(s string) string {
  fieldNameRegexp,_ := regexp.Compile("^\\w+")
  return fieldNameRegexp.FindString(s)
}


