package utils

import "reflect"

func DeepEqual(if1 interface{}, if2 interface{}) bool {
	return reflect.DeepEqual(if1, if2)
}

func TransformStructToMap(obj interface{}) map[string]interface{} {
	refType := reflect.TypeOf(obj)
	refValue := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < refType.NumField(); i++ {
		data[refType.Field(i).Name] = refValue.Field(i).Interface()
	}
	return data
}

func RemoveMapProperties(source map[string]interface{}, keys ...string) {
	for _, key := range keys {
		delete(source, key)
	}
}
