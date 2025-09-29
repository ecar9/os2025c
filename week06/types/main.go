package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var id int16
	//var name string
	//var gpa float32

	//id = 999
	//name = "Kim inha"
	//gpa = 3.99

	// var id int16 = 999
	// var name string = "Kim inha"
	// var gpa float32 = 3.99

	// var id = 999
	// var name = "Kim inha"
	// var gpa = 3.99

	id := 999
	name := "Kim inha"
	gpa := 3.99

	fmt.Println("학번은", id, reflect.TypeOf(id), "이름은", name, reflect.TypeOf(id))
	fmt.Println("평점은 :", gpa, reflect.TypeOf(gpa))

}
