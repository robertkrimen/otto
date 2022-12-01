package otto

import (
	"reflect"
)

func (rt *runtime) newGoMapObject(value reflect.Value) *object {
	obj := rt.newObject()
	obj.class = classObjectName // TODO Should this be something else?
	obj.objectClass = classGoMap
	obj.value = newGoMapObject(value)
	return obj
}

type goMapObject struct {
	value     reflect.Value
	keyType   reflect.Type
	valueType reflect.Type
}

func newGoMapObject(value reflect.Value) *goMapObject {
	if value.Kind() != reflect.Map {
		dbgf("%/panic//%@: %v != reflect.Map", value.Kind())
	}
	return &goMapObject{
		value:     value,
		keyType:   value.Type().Key(),
		valueType: value.Type().Elem(),
	}
}

func (o goMapObject) toKey(name string) reflect.Value {
	reflectValue, err := stringToReflectValue(name, o.keyType.Kind())
	if err != nil {
		panic(err)
	}
	return reflectValue
}

func (o goMapObject) toValue(value Value) reflect.Value {
	reflectValue, err := value.toReflectValue(o.valueType)
	if err != nil {
		panic(err)
	}
	return reflectValue
}

func goMapGetOwnProperty(obj *object, name string) *property {
	goObj := obj.value.(*goMapObject)
	value := goObj.value.MapIndex(goObj.toKey(name))
	if value.IsValid() {
		return &property{obj.runtime.toValue(value.Interface()), 0o111}
	}

	// Other methods
	if method := obj.value.(*goMapObject).value.MethodByName(name); method.IsValid() {
		return &property{
			value: obj.runtime.toValue(method.Interface()),
			mode:  0o110,
		}
	}

	return nil
}

func goMapEnumerate(obj *object, all bool, each func(string) bool) {
	goObj := obj.value.(*goMapObject)
	keys := goObj.value.MapKeys()
	for _, key := range keys {
		if !each(toValue(key).String()) {
			return
		}
	}
}

func goMapDefineOwnProperty(obj *object, name string, descriptor property, throw bool) bool {
	goObj := obj.value.(*goMapObject)
	// TODO ...or 0222
	if descriptor.mode != 0o111 {
		return obj.runtime.typeErrorResult(throw)
	}
	if !descriptor.isDataDescriptor() {
		return obj.runtime.typeErrorResult(throw)
	}
	goObj.value.SetMapIndex(goObj.toKey(name), goObj.toValue(descriptor.value.(Value)))
	return true
}

func goMapDelete(obj *object, name string, throw bool) bool {
	goObj := obj.value.(*goMapObject)
	goObj.value.SetMapIndex(goObj.toKey(name), reflect.Value{})
	// FIXME
	return true
}
