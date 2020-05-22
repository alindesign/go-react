package react

import (
	"reflect"
)

func toMap(value interface{}) map[string]interface{} {
	newMap := map[string]interface{}{}

	if value != nil {
		rVal := reflect.Indirect(reflect.ValueOf(value))
		rType := rVal.Type()

		var field reflect.StructField
		var fieldVal reflect.Value

		for i := 0; i < rType.NumField(); i++ {
			field = rType.Field(i)
			fieldVal = rVal.FieldByName(field.Name)
			if !fieldVal.IsZero() {
				newMap[field.Name] = fieldVal.Interface()
			}
		}
	}

	return newMap
}
