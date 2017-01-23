package numeric_tools

import (
	"fmt"
	"reflect"
)

func canConvert(src interface{}, dstType reflect.Type) (reflect.Value, bool) {
	val := reflect.ValueOf(src)
	val = reflect.Indirect(val)
	if !val.Type().ConvertibleTo(dstType) {
		return val, false
	}
	return val, true
}

func EnsureFloat(src interface{}) (float64, error) {
	floatType := reflect.TypeOf(float64(0.0))
	if val, ok := canConvert(src, floatType); ok {
		result := val.Convert(floatType)
		return result.Float(), nil
	} else {
		return 0, fmt.Errorf("Cannot convert %v to float64", src)
	}
}

func EnsureInteger(src interface{}) (int64, error) {
	t := reflect.TypeOf(int64(0))
	if val, ok := canConvert(src, t); ok {
		result := val.Convert(t)
		return result.Int(), nil
	} else {
		return 0, fmt.Errorf("Cannot convert %v to int64", src)
	}
}

func EnsureInt32(src interface{}) (int32, error) {
    v, err := EnsureInteger(src)
    return int32(v), err
}

func EnsureInt(src interface{}) (int, error) {
    v, err := EnsureInteger(src)
    return int(v), err
}
