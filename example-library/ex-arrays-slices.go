/*
Basically, almost never use arrays in Go
https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html
https://medium.com/@haaawk/i-thought-i-understood-how-iteration-over-an-array-works-but-apparently-not-in-golang-441a7abd6540

  - Arrays in go are values (more on this below)
  - Tons of the features in it involve copying values which is hella expensive
  - Intellisense will capture out of bounds array assignments - neat
  - I don't like the syntax to initialize slices and arrays. They are way to close together for
    the performance implications they can have. Basically if you see a number assigned, it's an array else slice
*/
package examples

// You have to specify the size of the array in the return rather than
// the generic type. This pretty much just ends up throwing this into
// the garbage can
func MyArrayWithValue() [2]string {
	myArr2:= [2]string{"hell", "wrld"}
	myArr2[0] = "hello"
	myArr2[1] = "world"
	return myArr2
}

// Initalize with empty array then fill values
func MyArrayEmpty() [2]string {
	myArr2:= [2]string{}
	myArr2[0] = "hello"
	myArr2[1] = "world"
	return myArr2
}

// Empty Slice using append function
func MySliceEmpty() []int {
	mysl1 := []int{}
	mysl1 = append(mysl1, 1)
	return mysl1
}

// Slice initalize with values
func MySliceValues() []string {
	mysl1 := []string{"slice","w", "values"}
	mysl1 = append(mysl1, "anything")
	return mysl1
}

// Alternate way to spawn a slice and alternate way to assign values
func MySliceMake() []string {

	s := make([]string, 1)
	s[0] = "slicevalue"
	return s
}

func MySliceNested() [][]string {
	s := [][]string{}

	sNested := []string{"this", "has", "three"}
	// Assignment method 1
	// s[0] = sNested
	s = append(s, sNested)
	return s
}

// Declaring things outside of a function
// var exampleArray = [5]int8 {1,2,3,4,5}
// var mySlice = []int{}

/*
// Executing this in main

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

*/