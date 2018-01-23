package shot

import (
	"reflect"
)

type Key interface {
	Interface() interface{}
	ReflectType() reflect.Type
}


type key struct {
	reflectType reflect.Type
}

func (key key) Interface() interface{} {
	return reflect.New(key.ReflectType()).Interface()
}

func (key key) ReflectType() reflect.Type {
	return key.reflectType
}

func NewKey(rawType interface{}) Key {
	reflectType := reflect.TypeOf(rawType)
	return NewKeyByType(reflectType)
}

func NewKeyByType(reflectType reflect.Type) Key {
	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	return key{reflectType}
}