package main

import (
	"fmt"
	"reflect"
)

//MakeMap creates a map object
func MakeMap(fpt interface{}) {
	fnV := reflect.ValueOf(fpt).Elem()
	fnI := reflect.MakeFunc(fnV.Type(), implMap)
	fnV.Set(fnI)
}

//implMap is a map function with reflection as input
func implMap(in []reflect.Value) []reflect.Value {
	// return nil if we are missing the map func or input for it
	if len(in) < 2 {
		return nil
	}

	var res []reflect.Value
	mapFunc := in[0]
	mapIn := in[1]

	// there are only two types we should consider
	switch in[1].Kind() {
	case reflect.Map:
		for _, key := range mapIn.MapKeys() {
			newVal := mapFunc.Call([]reflect.Value{mapIn.MapIndex(key)})
			fmt.Println("NewVal", newVal[0])
			mapIn.SetMapIndex(key, newVal[0])
		}
	case reflect.Slice:
		for i := 0; i < mapIn.Len(); i++ {
			newVal := mapFunc.Call([]reflect.Value{mapIn.Index(i)})
			fmt.Println("NewVal", newVal[0])
			mapIn.Index(i).Set(newVal[0])
		}
	}
	return append(res, mapIn)
}

func main() {
	println("It is said that Go has no generics.\nHowever we have many other ways to implement a generics like library if less smoothly,one is reflect.MakeFunc.\nUnderscore is a very useful js library,and now let's implement part of it-map,it will help you to understand how reflect works.\nPlease finish the 'implMap' function and pass the test.\nhttps://blog.golang.org/laws-of-reflection will help you to understand reflection in go.")
}
