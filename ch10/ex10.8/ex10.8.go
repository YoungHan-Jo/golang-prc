package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	myColor := getMyFavoriteColor()
	fmt.Println(colorToString(myColor))
}
