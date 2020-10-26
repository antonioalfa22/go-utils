package errors

import "fmt"

func Check(e error)  {
	if e != nil {
		panic(e)
	}
}

func CheckSoft(e error) {
	if e != nil {
		fmt.Println("Error: ", e.Error())
	}
}
