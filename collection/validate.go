package collection

import (
	"errors"
	"reflect"
	"strings"
)

//IsStruct - Checks if doc is a struct
func IsStruct(doc interface{}) (bool, error) {

	v := reflect.ValueOf(doc)
	err := isPtr(v)
	if err != nil {
		return false, err
	}

	if v.Elem().Kind() != reflect.Struct {
		return false, errors.New("document is not a struct")
	}
	return true, nil
}

//IsSlice - Check if data is slice of structs
func IsSlice(data interface{}) (bool, error) {

	v := reflect.ValueOf(data)

	if v.Elem().Kind() == reflect.Slice {
		if v.Elem().Type().Elem().Kind() != reflect.Struct {
			return false, errors.New("collection is not a slice of structs")
		}
	}

	return true, nil
}

func isPtr(v reflect.Value) error {

	if v.Kind() != reflect.Ptr {
		return errors.New("expected ptr to slice")
	}
	return nil
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

	v := reflect.ValueOf(doc).Elem()

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

// func requiredFields(v reflect.Value) (string, string, error) {
// 	var id string
// 	var rev string

// 	for i := 0; i < v.NumField(); i++ {

// 		var str string
// 		var exists bool

// 		str, exists = v.Type().Field(i).Tag.Lookup("json")
// 		if !exists {
// 			str, exists = v.Type().Field(i).Tag.Lookup("bson")
// 		}
// 		if exists {

// 			if strings.HasPrefix(str, "_id") && v.Field(i).Kind() == reflect.String {
// 				id = v.Field(i).Interface().(string)
// 			}

// 			if strings.HasPrefix(str, "_rev") && v.Field(i).Kind() == reflect.String {
// 				rev = v.Field(i).Interface().(string)
// 			}
// 		}
// 	}

// 	return id, rev, nil
// }
