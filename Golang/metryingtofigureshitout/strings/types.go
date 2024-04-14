package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "Hello"

	fmt.Println(reflect.TypeOf(s))    //string
	fmt.Println(reflect.TypeOf(s[0])) //uint8, interesting
	fmt.Println(reflect.TypeOf("ה"))  //Unicode is also a string, makes sense

	//But how about this then?
	s1 := "ᾨ"
	fmt.Println(reflect.TypeOf(s1[0])) //also uint8 :o

}
