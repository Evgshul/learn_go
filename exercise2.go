package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"strconv"
)

func Pic(dx, dy int, choice int) [][]uint8 {
	picture := make([][]uint8, dy)

	for y := 0; y<dy; y++ {
		row := make([]uint8, dx)
		for x:=0; x<dx; x++ {

			row[x] = uint8(generatePic(x, y, choice))
		}
		picture[y] = row
	}
	return picture
}

func generatePic(x , y, choice int) uint8 {
	switch choice {
	case 1:
		return uint8(x * y)
	case 2:
		return uint8((x + y) / 2)
	case 3:
		return uint8(x ^ y)
	default:
		return 0
	}
}

func main() {

	fmt.Println("Choose function for the image: ")
	fmt.Println("1. x * y")
	fmt.Println("2. (x + y) / 2")
	fmt.Println("3. x ^ y")

	var input string
	fmt.Print("Enter 1, 2, or 3: ")
	_, err2 := fmt.Scanln(&input)
	if err2 != nil {
		return 
	}

	fmt.Println("Your choice: ", input)

	// convert input to int
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > 3 {
		fmt.Println("Invalid input, please input 1,2 or 3")
		return
	}

	pic.Show(func(dx,dy int) [][]uint8 {
		return Pic(dx, dy, choice)
	})
}

