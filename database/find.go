package database

// import (
// 	"database/sql/driver"
// 	"errors"
// 	"reflect"
// )

// func (m *Manager) Find(dest any, sql string, args ...any) error {
// 	if reflect.ValueOf(dest).Kind() != reflect.Ptr {
// 		return errors.New("dest is not a pointer")
// 	}
// 	if reflect.ValueOf(dest).Elem().Kind() != reflect.Slice {
// 		return errors.New("dest is not a slice pointer")
// 	}
// 	if err := m.db.Ping(); err != nil {
// 		return err
// 	}
// 	rows, err := m.db.Query(sql, args...)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	columns, err := rows.Columns()
// 	if err != nil {
// 		return err
// 	}

// 	values := make([]any, len(columns))
// 	pointers := make([]any, len(columns))
// 	for i := range columns {
// 		pointers[i] = &values[i]
// 	}

// 	t := reflect.ValueOf(dest).Elem().Type().Elem()
// 	s := reflect.ValueOf(dest).Elem()
// 	for rows.Next() {
// 		rows.Scan(pointers...)
// 		value := reflect.New(t).Elem()
// 		for i, column := range columns {
// 			for j := 0; j < t.NumField(); j++ {
// 				f := t.Field(j)
// 				if tag, ok := f.Tag.Lookup("database"); ok && column == tag {
// 					if kv := reflect.ValueOf(values[i]); kv.IsValid() {
// 						switch v := value.FieldByName(f.Name); v.Kind() {
// 						case reflect.Bool:
// 							if b, err := driver.Bool.ConvertValue(values[i]); err != nil {
// 								return err
// 							} else {
// 								v.SetBool(b.(bool))
// 							}
// 						case reflect.Pointer:
// 							if !v.Elem().CanSet() {
// 								v.Set(reflect.New(v.Type().Elem()))
// 							}
// 							v.Elem().Set(kv.Convert(v.Type().Elem()))
// 						default:
// 							v.Set(kv.Convert(v.Type()))
// 						}
// 					}
// 					break
// 				}
// 			}
// 		}
// 		s.Set(reflect.Append(s, value))
// 	}

// 	return nil
// }
