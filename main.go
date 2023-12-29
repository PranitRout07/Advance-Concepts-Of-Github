package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := divide(10.0, 5.0)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println("The result is ", res)

}
func divide(x, y float64) (float64, error) {
	var res float64
	if y == 0 {
		return res, errors.New("Not divisible by 0")
	}
	res = x / y
	return res, nil
}
