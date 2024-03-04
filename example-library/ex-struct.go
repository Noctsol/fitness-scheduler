// Example of structure usage in Go
package examples


type MyStruct struct {
	SomeString string
	SomeNumb int8
	SomeFloat float32
	SomeBool bool
	SomeArray []string
}

func NewStruct (name string,
	numb int8, flt float32, 
	booli bool, 
	myarray []string) MyStruct {
	
	// Note the placement of the brackets
	// For some reason this matters
	// What kind of parser is this
	m := MyStruct{
		SomeString: name,
		SomeNumb: numb,
		SomeFloat: flt,
		SomeBool: false}
	
	var muhList []string 
	m.SomeArray = muhList
	
	return m
}


// You can't use the := notation outside of functions
var NoConstructor = MyStruct{
	SomeString: "outsidefunction",
	SomeNumb: 2,
	SomeFloat: 2.3,
	SomeBool: false,
	SomeArray: []string{}}



