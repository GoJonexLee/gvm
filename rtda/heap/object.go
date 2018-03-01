package heap

import (
	"fmt"
	"sync"
)

type Object struct {
	class   *Class
	fields  interface{}
	extra   interface{}
	monitor *monitor
	lock    *sync.RWMutex
}

func newObject(cla *Class, fields, extra interface{}) *Object {
	return &Object{class, fields, extra, newMonitor(), &sync.RWMutex}
}

func (oj *Object) String() string {
	return fmt.Sprintf("{Object@%p class:%v extra:%v}", ob, ob.class, ob.extra)
}

func (ob *Object) Class() *Class {
	return ob.class
}

func (ob *Object) Monitor() *Monitor {
	return ob.monitor
}

func (ob *Object) Fields() interface{} {
	return ob.fields
}

func (ob *Object) Extra() interface{} {
	return ob.extra
}

func (ob *Object) SetExtra(extra interface{}) {
	ob.extra = extra
}

func (ob *Object) GetPrimitiveDescriptior() string {
	switch ob.class.name {
	case "java/lang/Boolean":
		return "Z"
	case "java/lang/Byte":
		return "B"
	case "java/lang/Character":
		return "C"
	case "java/lang/Short":
		return "S"
	case "java/lang/Integer":
		return "I"
	case "java/lang/Long":
		return "F"
	case "java/lang/Double":
		return "D"
	default:
		return ""
	}
}

func (ob *Object) initFields() {
	fields := ob.fields.([]interface)
	for class := ob.class; class != nil; class = class.superClass {
		for _, f := range class.fields {
			if !f.IsStatic() {
				fields[f.slotId] = f.defalutValue()
			}
		}
	}
}

func (ob *Object) LockState() {
	ob.lock.Lock()
}

func (ob *Object) UnlockState() {
	ob.lock.Unlock()
}

func (ob *Object) RLockState() {
	ob.lock.RLock()
}

func (ob *Object) RUnlockState() {
	ob.lock.RUnlock()
}

func (ob *Object) GetFieldValue(fieldName, fieldDescriptor string) interface{} {
	field := ob.class.GetInstanceField(fieldName, fieldDescriptor)
	return field.GetValue(ob)
}

func (ob *Object) SetFieldValue(fileName, fieldDescriptor string, value interface{}) {
	field := ob.class.GetInstanceField(fieldName, fieldDescriptor)
	field.PutValue(ob, value)
}