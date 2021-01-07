package collection

import (
	"reflect"
	"strings"
)

//IsStruct - Checks if doc is a struct
func IsStruct(doc interface{}) (bool, error) {

	v := reflect.ValueOf(doc)
	isptr := isPtr(v)

	if (isptr && v.Elem().Kind() != reflect.Struct) || (!isptr && v.Kind() != reflect.Struct) {
		return false, ErrNotStruct
	}
	return true, nil
}

//IsSlice - Check if data is slice of structs
func IsSlice(data interface{}) (bool, error) {

	v := reflect.ValueOf(data)

	if v.Elem().Kind() == reflect.Slice {
		if v.Elem().Type().Elem().Kind() != reflect.Struct {
			return false, ErrNotSliceOfStructs
		}
	}

	return true, nil
}

func isPtr(v reflect.Value) bool {

	if v.Kind() != reflect.Ptr {
		return false
	}
	return true
}

//RequiredFields - Check if document is a struct and contains required fields
func RequiredFields(doc interface{}) (id string, rev string, err error) {

	_, err = IsStruct(doc)
	if err != nil {
		return "", "", err
	}

	id, rev, err = requiredFields(doc)

	return
}

func requiredFields(doc interface{}) (string, string, error) {

	var id string
	var rev string

	v := reflect.ValueOf(doc)
	if ok := isPtr(v); ok {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {

		str, exists := v.Type().Field(i).Tag.Lookup("json")
		if exists {

			if strings.HasPrefix(str, "_id") && v.Field(i).Kind() == reflect.String {
				id = v.Field(i).Interface().(string)
			}

			if strings.HasPrefix(str, "_rev") && v.Field(i).Kind() == reflect.String {
				rev = v.Field(i).Interface().(string)
			}
		}
	}

	return id, rev, nil
}
