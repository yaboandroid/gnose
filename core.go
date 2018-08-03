package core

import (
	"fmt"
	"log"
	"reflect"
)

func init() {
	fmt.Println("this  is setup methods")
}

func Test() {
	fmt.Println("this is test fun in core")
}

func AssertTrue(value bool, msg ...interface{}) {
	fmt.Println(value)
	if !value {
		//log.Fatalf("Error : got exception : %s, Except %v, Actual %v", msg, !value, value)
		//log.Printf("Error : got exception : %s, Except %v, Actual %v", msg, !value, value)
		log.Panicf("Error : got exception : %s, Except %v, Actual %v", msg, !value, value)
	}
}

func valueCheck(v1,v2 reflect.Value,msg string){
	k1 := v1.Kind()
	k2 := v2.Kind()
	if k1 != k2 {
		panic(fmt.Sprintf("Error: Different type not allow to compare, type of v1 is %s, v2 is %s", k1.String(), k2.String()))
	}
	switch k1{
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v1.Int() != v2.Int() {
			//log.Fatalf("Error: got exception : %s, Except: %d equal %d", msg, v1.Int(), v2.Int())
			//log.Printf("Error: got exception : %s, Except: %d equal %d", msg, v1.Int(), v2.Int())
			log.Panicf("Error: got exception : %s, Except: %d equal %d", msg, v1.Int(), v2.Int())
		}
	case reflect.Float32, reflect.Float64:
		if v1.Float() != v2.Float() {
			log.Fatalf("Error: got exception : %s, Except: %f equal %f", msg, v1.Float(), v2.Float())
		}
	case reflect.Bool:
		if v1.Bool() != v2.Bool() {
			log.Fatalf("Error: got exception : %s, Except: %v equal %v", msg, v1.Bool(), v2.Bool())
		}
	case reflect.String:
		if v1.String() != v2.String() {
			log.Fatalf("Error: got exception : %s, Except: %s equal %s", msg, v1.String(), v2.String())
		}
	case reflect.Slice,reflect.Array:
		factorType1:=v1.Type().Elem().String()
		factorType2:=v2.Type().Elem().String()
		if factorType1!=factorType2{
			log.Fatalf("Error: got exception : element type of slice is not matched, v1 : %s, v2 : %s",factorType1,factorType2)
		}
		for i:=0;i<v1.Len();i++{
			e1:=v1.Index(i)
			e2:=v2.Index(i)
			if e1!=e2{
				log.Fatalf("Error: got exception: %s",msg)
			}
		}
	default:
		log.Fatalf("Error: got exception %s","not support parameters passed")
	}
}


func AssertEqual(a, b interface{}, msg string) {
	v1 := reflect.ValueOf(a)
	v2 := reflect.ValueOf(b)
	valueCheck(v1,v2,msg)

}
