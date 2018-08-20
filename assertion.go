package gnose

import (
	"reflect"
	"fmt"
)

type Assert struct {
	logger Testlogger
}

func NewAssert(log Testlogger) Assert {
	return Assert{logger:log}
}

func (as Assert)AssertTrue(value bool, msg ...interface{}) {
	if !value {
		as.logger.Exception("Error : got exception : %s, Except %v, Actual %v", msg, !value, value)
	}
}

func (as Assert)AssertNonCriticalTrue(value bool, msg ...interface{}) (err string) {
	if !value {
		as.logger.Error("Error : got exception : %s, Except %v, Actual %v", msg, !value, value)
	}
	err = fmt.Sprintf("Error : got exception : %s, Except %v, Actual %v", msg, !value, value)
	return
}

func (as Assert)diffLog(isCritical bool, msg string) {
	if isCritical {
		as.logger.Exception(msg)
	} else {
		as.logger.Warning(msg)

	}
}

func (as Assert)valueCheck(v1, v2 reflect.Value, isCritical bool, msg string) (err string) {
	k1 := v1.Kind()
	k2 := v2.Kind()
	if k1 != k2 {
		err = fmt.Sprintf("Error: Different type not allow to compare, type of v1 is %s, v2 is %s", k1.String(), k2.String())
	}
	switch k1{
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v1.Int() != v2.Int() {
			err = fmt.Sprintf("Error: got exception : %s, Except: %d equal %d", msg, v1.Int(), v2.Int())
		}
	case reflect.Float32, reflect.Float64:
		if v1.Float() != v2.Float() {
			err = fmt.Sprintf("Error: got exception : %s, Except: %f equal %f", msg, v1.Float(), v2.Float())
		}
	case reflect.Bool:
		if v1.Bool() != v2.Bool() {
			err = fmt.Sprintf("Error: got exception : %s, Except: %v equal %v", msg, v1.Bool(), v2.Bool())
		}
	case reflect.String:
		if v1.String() != v2.String() {
			err = fmt.Sprintf("Error: got exception : %s, Except: %s equal %s", msg, v1.String(), v2.String())
		}
	case reflect.Slice, reflect.Array:
		factorType1 := v1.Type().Elem().String()
		factorType2 := v2.Type().Elem().String()
		if factorType1 != factorType2 {
			err = fmt.Sprintf("Error: got exception : element type of slice is not matched, v1 : %s, v2 : %s", factorType1, factorType2)
		}
		for i := 0; i < v1.Len(); i++ {
			e1 := v1.Index(i)
			e2 := v2.Index(i)
			if e1 != e2 {
				err = fmt.Sprintf("Error: got exception: %s", msg)
				break
			}
		}
	default:
		err = fmt.Sprintf("Error: not support parameters passed : %v", k1)
	}
	as.diffLog(isCritical, err)
	return
}

func (as Assert)AssertEqual(a, b interface{}, msg string) {
	v1 := reflect.ValueOf(a)
	v2 := reflect.ValueOf(b)
	as.valueCheck(v1, v2, true, msg)
}

func (as Assert)AssertNonCriticalEqual(a, b interface{}, msg string) (err string) {
	v1 := reflect.ValueOf(a)
	v2 := reflect.ValueOf(b)
	err = as.valueCheck(v1, v2, false, msg)
	return
}