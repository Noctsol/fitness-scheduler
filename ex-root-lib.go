// Example of accessing files in the same root dir as the file with the main()
// Function. You need to list out the package main statement at the top
// This basically compiles like terraform which makes sense given that
// tf is based on go
package main

import "fmt"

func functioninotherfile() {
	fmt.Println("otherfunction")
}
