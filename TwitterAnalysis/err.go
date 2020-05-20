package main

import (
	"fmt"
	"log"
)

func ErrorRaise(m string) error {
	return fmt.Errorf(m)
}

func ErrorCheck(err error)  {
	if err != nil{
		log.Fatalln(err)
	}
}
