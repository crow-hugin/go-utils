package format

import (
	"fmt"
	"reflect"
)

// ToInterfaceArr 将任意切片转换成interface切片
func ToInterfaceArr(slice interface{}) []interface{} {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return nil
	}
	sliceValue := reflect.ValueOf(slice)
	retSlice := make([]interface{}, sliceValue.Len())
	for k := 0; k < sliceValue.Len(); k++ {
		retSlice[k] = sliceValue.Index(k).Interface()
	}
	return retSlice
}

// StructToMap 将结构体转换成map
func StructToMap(obj interface{}, ignore bool) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if !IsStructPtr(t) {
		panic(fmt.Sprintf("%v must be  a struct pointer", obj))
	}
	t = t.Elem()
	v = v.Elem()

	var ret = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		fieldV := v.Field(i)
		if ignore && (!fieldV.IsValid() || fieldV.IsZero()) {
			continue
		}
		fieldT := t.Field(i)
		get := fieldT.Tag.Get("json")
		var tag string
		if len(get) == 0 {
			tag = fieldT.Name
		} else {
			tag = get
		}
		ret[tag] = fieldV.Interface()
	}
	return ret
}

// StructToFieldSlice 提取结构体字段
func StructToFieldSlice(obj interface{}) []string {
	t := reflect.TypeOf(obj)
	var FieldData []string
	for i := 0; i < t.NumField(); i++ {
		FieldData = append(FieldData, t.Field(i).Name)
	}
	return FieldData
}

func IsStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
