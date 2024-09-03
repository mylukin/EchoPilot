package helper

import (
	"fmt"
	"hash/fnv"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ToInt64 泛型函数，尝试将不同的类型转换为int64。
func ToInt64[T string | int | int32 | int64](v T) int64 {
	var r int64

	switch value := any(v).(type) {
	case string:
		r, _ = strconv.ParseInt(value, 10, 64)
	case int:
		r = int64(value)
	case int32:
		r = int64(value)
	case int64:
		r = value
	}

	return r
}

// ToFloat64 泛型函数，尝试将不同的类型转换为float64。
func ToFloat64[T string | float64](v T) float64 {
	var r float64

	switch value := any(v).(type) {
	case string:
		r, _ = strconv.ParseFloat(value, 64)
	case float64:
		r = value
	}

	return r
}

// ToObjectID 泛型函数，尝试将不同的类型转换为ObjectID。
func ToObjectID[T string | primitive.ObjectID](v T) primitive.ObjectID {
	var r primitive.ObjectID

	switch value := any(v).(type) {
	case string:
		r, _ = primitive.ObjectIDFromHex(value)
	case primitive.ObjectID:
		r = value
	}

	return r
}

// ToString 泛型函数，尝试将不同的类型转换为string。
func ToString[T float32 | float64 | int | int64 | string](v T) string {
	var r string

	switch value := any(v).(type) {
	case float32:
		r = strconv.FormatFloat(float64(value), 'f', -1, 64)
	case float64:
		r = strconv.FormatFloat(value, 'f', -1, 64)
	case int:
		r = strconv.Itoa(value)
	case int64:
		r = strconv.FormatInt(value, 10)
	case string:
		r = value
	default:
		r = fmt.Sprintf("%v", value)
	}

	return r
}

// ConvertibleToUInt32 是一个约束，它匹配所有可以转换为uint32的类型。
func ToUInt32[T string | int64 | primitive.ObjectID](v T) uint32 {
	var r string
	switch value := any(v).(type) {
	case string:
		r = value
	case int64:
		r = strconv.FormatInt(value, 10)
	case primitive.ObjectID:
		r = value.Hex()
	}
	h := fnv.New32a()
	h.Write([]byte(r))
	return h.Sum32()
}
