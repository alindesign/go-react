package react

import (
	"reflect"
)

func toMap(value interface{}) map[string]interface{} {
	newMap := map[string]interface{}{}

	if value != nil {
		rVal := reflect.Indirect(reflect.ValueOf(value))
		rType := rVal.Type()

		if rType.Kind() == reflect.Map {
			return rVal.Interface().(map[string]interface{})
		} else if rType.Kind() == reflect.Slice {
			for i := 0; i < rVal.Len(); i++ {
				for k, v := range toMap(rVal.Index(i).Interface()) {
					newMap[k] = v
				}
			}

			return newMap
		}

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
