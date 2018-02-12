package shot

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

func newSingletonValue(initialize initialize) *singleton {
	return &singleton{once: sync.Once{}, initialize: initialize}
}

type initialize func() (interface{}, error)

type singleton struct {
	once       sync.Once
	value      interface{}
	initialize initialize
}

func (s *singleton) get() (interface{}, error) {
	var err error
	s.once.Do(func() {
		s.value, err = s.initialize()
	})
	if err != nil {
		return nil, err
	}
	return s.value, nil
}

type filledBinding interface {
	ok() error
	get() interface{}
}

func newNoScopeBinding(initialize initialize) filledBinding {
	return &noScopeBinding{initialize}
}

type noScopeBinding struct {
	initialize initialize
}

func (binding *noScopeBinding) ok() error {
	_, err := binding.initialize()
	return err
}

func (binding *noScopeBinding) get() interface{} {
	value, _ := binding.initialize()
	return value
}

func newSingletonBinding(initialize initialize) filledBinding {
	return &singletonBinding{newSingletonValue(initialize)}
}

type singletonBinding struct {
	singleton *singleton
}

func (binding *singletonBinding) ok() error {
	_, err := binding.singleton.initialize()
	return err
}

func (binding *singletonBinding) get() interface{} {
	value, _ := binding.singleton.get()
	return value
}

func newEagerSingletonBinding(initialize initialize) filledBinding {
	return &eagerSingletonBinding{newSingletonValue(initialize)}
}

type eagerSingletonBinding struct {
	singleton *singleton
}

func (binding *eagerSingletonBinding) ok() error {
	_, err := binding.singleton.initialize()
	return err
}

func (binding *eagerSingletonBinding) get() interface{} {
	value, _ := binding.singleton.get()
	return value
}

type binding interface {
	fill(injector Injector) filledBinding
	getScope() Scope
	withScope(scope Scope) binding
	getKey() Key
}

func newUntargettedBinding(key Key) binding {
	return &untargettedBinding{
		key:   key,
		scope: NoScope,
	}
}

type untargettedBinding struct {
	key   Key
	scope Scope
}

func (binding *untargettedBinding) fill(injector Injector) filledBinding {
	return resolveBindingScope(binding.scope, func() (interface{}, error) {
		return buildByStructure(injector, binding.key.Interface())
	})
}

func (binding *untargettedBinding) getScope() Scope {
	return binding.scope
}

func (binding *untargettedBinding) withScope(scope Scope) binding {
	binding.scope = scope
	return binding
}

func (binding *untargettedBinding) getKey() Key {
	return binding.key
}

func newLinkedBinding(key Key, scope Scope, implementation interface{}) binding {
	return &linkedBinding{
		key:            key,
		scope:          scope,
		implementation: implementation,
	}
}

type linkedBinding struct {
	key            Key
	scope          Scope
	implementation interface{}
}

func (binding *linkedBinding) fill(injector Injector) filledBinding {
	return resolveBindingScope(binding.scope, func() (interface{}, error) {
		return buildByStructure(injector, binding.implementation)
	})
}

func (binding *linkedBinding) getScope() Scope {
	return binding.scope
}

func (binding *linkedBinding) withScope(scope Scope) binding {
	binding.scope = scope
	return binding
}

func (binding *linkedBinding) getKey() Key {
	return binding.key
}

func newConstructorBinding(key Key, scope Scope, constructor interface{}) binding {
	return &constructorBinding{
		key:         key,
		scope:       scope,
		constructor: constructor,
	}
}

type constructorBinding struct {
	key         Key
	scope       Scope
	constructor interface{}
}

func (binding *constructorBinding) getScope() Scope {
	return binding.scope
}

func (binding *constructorBinding) withScope(scope Scope) binding {
	binding.scope = scope
	return binding
}

func (binding *constructorBinding) getKey() Key {
	return binding.key
}

func (binding *constructorBinding) fill(injector Injector) filledBinding {
	return resolveBindingScope(binding.scope, func() (interface{}, error) {
		return buildByConstructor(injector, binding.constructor)
	})
}

func newInstanceBinding(key Key, scope Scope, instance interface{}) binding {
	return &instanceBinding{
		key:      key,
		scope:    scope,
		instance: instance,
	}
}

type instanceBinding struct {
	key      Key
	scope    Scope
	instance interface{}
}

func (binding *instanceBinding) getScope() Scope {
	return binding.scope
}

func (binding *instanceBinding) withScope(scope Scope) binding {
	binding.scope = scope
	return binding
}

func (binding *instanceBinding) getKey() Key {
	return binding.key
}

func (binding *instanceBinding) fill(injector Injector) filledBinding {
	return resolveBindingScope(binding.scope, func() (interface{}, error) {
		return binding.instance, nil
	})
}

func resolveBindingScope(scope Scope, initialize initialize) filledBinding {
	switch scope {
	case SingletonInstance:
		return newSingletonBinding(initialize)
	case EagerSingleton:
		return newEagerSingletonBinding(initialize)
	default:
		return newNoScopeBinding(initialize)
	}
}

func buildByStructure(injector Injector, structure interface{}) (interface{}, error) {
	structureType := reflect.TypeOf(structure)

	if structureType == nil {
		return nil, errors.New("can't reflect a struct nil")
	}

	if structureType.Kind() == reflect.Ptr {
		structureType = structureType.Elem()
	}

	if structureType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("can't reflect a struct not struct (type %v)", structureType)
	}

	structureValue := reflect.Indirect(reflect.New(structureType))

	return fillStructure(injector, structureValue)
}

func fillStructure(injector Injector, structureValue reflect.Value) (interface{}, error) {
	for i := 0; i < structureValue.Type().NumField(); i++ {
		structField := structureValue.Type().Field(i)
		_, ok := structField.Tag.Lookup("inject")
		if !ok {
			continue
		}
		structValueField := structureValue.Field(i)
		if !structValueField.CanSet() {
			return nil, errors.New("can't set a private field of struct")
		}
		value := reflect.ValueOf(injector.GetByKey(NewKeyByType(structField.Type)))
		structValueField.Set(value)
	}
	return structureValue.Addr().Interface(), nil
}

func buildByConstructor(injector Injector, constructorFunc interface{}) (interface{}, error) {

	constructor := reflect.ValueOf(constructorFunc)
	constructorType := reflect.TypeOf(constructorFunc)

	if constructorType == nil {
		return nil, errors.New("can't reflect a constructorFunc nil")
	}

	if constructorType.Kind() != reflect.Func {
		return nil, fmt.Errorf("can't reflect a constructorFunc not function (type %v)", constructorType)
	}

	constructorArgs := buildArgs(injector, constructorType)

	return callConstructor(constructor, constructorArgs)
}

func callConstructor(constructor reflect.Value, constructorArgs []reflect.Value) (interface{}, error) {
	values := constructor.Call(constructorArgs)

	if len(values) != 1 {
		return nil, errors.New("a constructor should return only one result")
	}

	return values[0].Interface(), nil
}

func buildArgs(injector Injector, constructorType reflect.Type) []reflect.Value {
	var args []reflect.Value
	for i := 0; i < constructorType.NumIn(); i++ {
		argType := constructorType.In(i)
		value := reflect.ValueOf(injector.GetByKey(NewKeyByType(argType)))
		args = append(args, value)
	}
	return args
}
