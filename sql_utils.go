// A group of utitlities to deal with MySQL in a more Object Oriented way.
package mysql_utils

import (
  "reflect"
  "strings"
  "regexp"
  "database/sql"
)

type MysqlUtil struct {
  Conn Connection // A helper to connect to the MySQL DB
}


/*
Create a MySql table based on the fields of a struct.

Example:
  type Model struct {
    Id    int     `mysql:"id INT NOT NULL AUTO_INCREMENT,pk"`
    Name  string  `mysql:"name VARCHAR(20) NOT NULL"`
  }

  var s MysqlUtil
  s.Conn.GetConfiguration("config.gcfg","test")

  err := s.CreateTable("model", &Model{})
  if err != nil {
    panic(err)
  }

The `mysql` tags in the struct are necessary. The portion before the comma will be directly translated to the table. To denote a field as a primary key, simply split the statement with a comma and add `pk` at the end, as shown in the `Id` field above.

This function will only create the table if it doesn not already exist.
*/
func (s MysqlUtil) CreateTable(tablename string, i interface{}) error {
  db := s.Conn.Open() 
  defer db.Close()

  create_statement := s.buildCreateTableStatement(tablename, i)

  _, err := db.Exec(create_statement)
  return err
}

// Drop a MySql table. This function will only drop the table if it
// does not already exist
func (s MysqlUtil) DropTable(tablename string) error {
  db := s.Conn.Open()
  defer db.Close()

  _, err := db.Exec("DROP TABLE IF EXISTS " + tablename + ";")
  return err
}

// Check if the table is ready to take queries. This can be used to check
// for a table's existence, more or less, since there really isn't a 
// sure-fire way to do that in MySQL.
func (s MysqlUtil) TableReady(tablename string) bool {
  db := s.Conn.Open()
  defer db.Close()

  _, err := db.Query("SELECT * FROM "+tablename+";")

  if err != nil {
    return false
  }

  return true
}

// Simply a wrapper for Go Lang's `sql.Query`.
func (s MysqlUtil) Query(query string, args ...interface{}) (*sql.Rows, error) {
  db := s.Conn.Open()
  defer db.Close()
  rows, err := db.Query(query, args)

  return rows, err
}

// Simply a wrapper for Go Lang's `sql.QueryRow`.
func (s MysqlUtil) QueryRow(query string, args ...interface{}) *sql.Row {
  db := s.Conn.Open()
  defer db.Close()
  row := db.QueryRow(query, args)

  return row
}

/*
A function used for parsing out the MySQL fields of a struct.

Example:
  type Model struct {
    Id    int     `mysql:"id INT NOT NULL AUTO_INCREMENT,pk"`
    Name  string  `mysql:"name VARCHAR(20) NOT NULL"`
  }

  var s MysqlUtil
  list := s.FieldList(&Model{})

  // list[0]  = "id INT NOT NULL AUTO_INCREMENT"
  // list[1]  = "name VARCAHR(20) NOT NULL"
*/
func (s MysqlUtil) FieldList(i interface{}) (list []string) {
  fields, _ := s.parseFields(i)

  for _, field := range fields {
    list = append(list, fieldName(field))
  }

  return
}

/*
A function used for parsing out the MySQL primary keys of a struct.

Example:
  type Model struct {
    Id    int     `mysql:"id INT NOT NULL AUTO_INCREMENT,pk"`
    Name  string  `mysql:"name VARCHAR(20) NOT NULL"`
  }

  var s MysqlUtil
  list := s.PrimaryKeys(&Model{})

  // list[0]  = "id"
*/
func (s MysqlUtil) PrimaryKeys(i interface{}) []string {
  _, pks := s.parseFields(i)

  return pks
}

// Builds the `CREATE TABLE` string.
func (s MysqlUtil) buildCreateTableStatement(tablename string, i interface{}) (create_statement string) {
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

// Parses out the fields of a struct to MySQL friendly fields.
func (s MysqlUtil) parseFields(i interface{}) (sqlfields, primaryKeys []string) {
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


