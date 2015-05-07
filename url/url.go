/*
	Params create url encoded strings from structs
	Supported field types: string, int, float
*/
package url

import(
	"net/url"
	"reflect"
	"errors"
	"strconv"
	//"log"
)

func Encode(s interface{}) (string, error) {
	
	originType := reflect.TypeOf(s)
	isPtr := originType.Kind() == reflect.Ptr
	if isPtr { originType = originType.Elem() }
	
	if originType.Kind() != reflect.Struct {
		return "", errors.New("Cannot encode non-struct: " + string(reflect.TypeOf(s).Kind()))
	}

	var values reflect.Value
	if isPtr {
		values = reflect.ValueOf(s).Elem()
	} else {
		values = reflect.ValueOf(s)
	}
	types := values.Type()
	uv := url.Values{}
 
	for i := 0; i < values.NumField(); i++ {
		
		fieldValue := values.Field(i)
		fieldType := types.Field(i)
		tag := fieldType.Tag.Get("param")
		
		if tag == "_" { continue }
		
		if fieldType.PkgPath == ""  && fieldValue.IsValid() {

			var value string = ""
			kind := fieldValue.Kind()

			switch kind {

				case reflect.String:
					value = fieldValue.String()
					break

				case reflect.Int:
					fallthrough

				case reflect.Int64:
					value = strconv.FormatInt(fieldValue.Int(), 10)
					break

				case reflect.Float32:
					value = strconv.FormatFloat(fieldValue.Float(), 'f', -1, 32)
					break

				case reflect.Float64:
					value = strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
					break

				default:
					return "", errors.New("Data type not supported: " + string(kind))
			}
			
			if value != "" { uv.Add(tag, value) }
		}
	}

	return uv.Encode(), nil
}

func ParamsEncode(p interface{}) string {
	s, err := Encode(p)
	if err != nil { return "" }
	return s
}

