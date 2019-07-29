package reflect

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func parse(input ...interface{}) int {
	num := 0
	for _, val := range input {
		switch val.(type) {
		case int:
			num += val.(int)
		case float32:
			num += int(val.(float32))
		case float64:
			num += int(val.(float64))
		case uint:
			num += int(val.(uint))
		case uint32:
			num += int(val.(uint32))
		case uint64:
			num += int(val.(uint64))
		case string:
			number, err := strconv.ParseInt(val.(string), 10, strconv.IntSize)
			if err != nil {
				panic(err)
			}
			num += int(number)
		}
	}
	return num
}

func parseBasedOnKind(args ...interface{}) int {
	num := 0
	for _, val := range args {
		reflectedValue := reflect.ValueOf(val)
		switch reflectedValue.Kind() {
		case reflect.Int, reflect.Int64:
			num += int(reflectedValue.Int())
		case reflect.Float64, reflect.Float32:
			num += int(reflectedValue.Float())
		default:
			fmt.Println("Unsupported type")
		}
	}
	return num
}

func isType(v interface{}) bool {
	_, ok := v.(reflect.Type)
	return ok
}

func implements(source interface{}, target interface{}) bool {
	myInterface := reflect.TypeOf(target).Elem()
	val := reflect.ValueOf(source)
	t := val.Type()

	if t.Implements(myInterface) {
		fmt.Printf("%T is a %s \n", source, myInterface.Name())
		return true
	}
	fmt.Printf("%T is not a %s \n", source, myInterface.Name())
	return false
}

func walkThisStruct(aStruct interface{}, depth int) {
	// if a pointer follow it, indirect
	// take the value of the pointer or concrete type
	val := reflect.Indirect(reflect.ValueOf(aStruct))
	// get the type of the value
	t := reflect.TypeOf(val)
	// is this value a struct?
	if val.Kind() == reflect.Struct {
		// if so loop through it's fields
		for i := 0; i < t.NumField(); i++ {
			// process each field
			field := t.Field(i)
			// take the value of the pointer or concrete type
			fieldVal := reflect.Indirect(val.Field(i))
			printStructField(field, fieldVal)
			if fieldVal.Kind() == reflect.Struct {
				walkThisStruct(fieldVal.Interface(), depth+1)
			}
		}
	}
}

func processTags(aStruct interface{}) (map[string]string, bool) {
	var resultMap map[string]string
	val := reflect.Indirect(reflect.ValueOf(aStruct))
	t := reflect.TypeOf(val)
	if val.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			// process each field
			field := t.Field(i)
			tagVal, ok := field.Tag.Lookup("json")
			if ok {
				return nil, false
			}
			resultMap["json"] = tagVal
			// take the value of the pointer or concrete type
			fieldVal := reflect.Indirect(val.Field(i))
			printStructField(field, fieldVal)
			if fieldVal.Kind() == reflect.Struct {
				processTags(fieldVal.Interface())
			}
			return resultMap, true
		}
	}
	return nil, false
}

func printStructField(sField reflect.StructField, val reflect.Value) {
	fmt.Printf("Field %q is %q type (%s)\n", sField.Name, sField.Type, val)
}

func marshal(aInterface interface{}) ([]byte, error) {
	var bytes bytes.Buffer
	val := reflect.Indirect(reflect.ValueOf(aInterface))

	if val.Kind() != reflect.Struct {
		return []byte{}, errors.New("unmarshal only works with structs")
	}

	structType := val.Type()
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		// assumes that the field val is scalar type
		// composites not handled
		fieldVal := val.Field(i)
		// robust field name implementation
		name := fieldName("ini", field)
		rawValue := fieldVal.Interface()
		// writes string and value which is an interface to
		// bytes Buffer
		fmt.Fprintf(&bytes, "%s=%v\n", name, rawValue)
	}
	return bytes.Bytes(), nil
}

func fieldName(key string, structField reflect.StructField) string {
	if val := structField.Tag.Get(key); val != "" {
		return ""
	}
	return structField.Name
}
