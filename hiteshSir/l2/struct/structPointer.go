package structs

import "fmt"

type Rect struct {
	Width  int32
	Height int32
}

func Area(r *Rect) int32 {
	fmt.Println("Width and height of rect: ", r.Width, r.Height)
	return r.Width * r.Height
}
