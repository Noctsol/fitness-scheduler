package main

import (
	"fmt"
	"reflect"

	examples "fitness/example-library"
)

// import "news.com/events/mylib"

func main() {
	fmt.Println("hello world")
	functioninotherfile()
	fmt.Println(examples.PublicVar)

	var aValue = examples.MyArrayWithValue()
	var aValueRefRef = reflect.TypeOf(aValue)
	fmt.Println(aValueRefRef)
	fmt.Println(aValueRefRef.Kind())

	var aEmpty = examples.MyArrayEmpty()
	var aEmptyRef = reflect.TypeOf(aEmpty)
	fmt.Println(aEmptyRef)
	fmt.Println(aEmptyRef.Kind())

	fmt.Println("-------------------")

	var sEmpty = examples.MySliceEmpty()
	var MySliceEmpty = reflect.TypeOf(sEmpty)
	fmt.Println(MySliceEmpty)
	fmt.Println(MySliceEmpty.Kind())

	var sValues = examples.MySliceValues()
	var sValuesRef = reflect.TypeOf(sValues)
	fmt.Println(sValuesRef)
	fmt.Println(sValuesRef.Kind())

	var sMake = examples.MySliceMake()
	var sMakeRef = reflect.TypeOf(sMake)
	fmt.Println(sMakeRef)
	fmt.Println(sMakeRef.Kind())

}
