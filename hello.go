package main

import (
	"errors"
	"fmt"
	"log"
)

func area(length float64, width float64) (float64, error) {
	if length < 0 || width < 0 {
		err := errors.New("you cannot have a negative length/width")
		return -1, err
	}
	result := length * width
	return result, nil
}

func main() {
	length := 5.5
	width := 10
	result, err := area(length, float64(width))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
