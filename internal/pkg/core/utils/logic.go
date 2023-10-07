package utils

import (
	"fmt"
	"log"
)

// Must 没有错误，如果出现错误抛出异常
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Error(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s:%v", msg, err))
	}
	log.Println("Success:", msg)
}
