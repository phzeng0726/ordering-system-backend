package utils

// import (
// 	"errors"
// 	"reflect"
// 	"strconv"
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
// 				if v, ok := f.Tag.Lookup("database"); ok && column == v {
// 					kv := reflect.ValueOf(values[i])
// 					if kv.IsValid() {
// 						fv := value.FieldByName(f.Name)
// 						switch fv.Kind() {
// 						case reflect.Bool:
// 							boolValue, err := strconv.ParseBool(string(kv.Interface().([]uint8)))
// 							if err != nil {
// 								return err
// 							}
// 							fv.SetBool(boolValue)
// 						default:
// 							fv.Set(kv.Convert(fv.Type()))
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
