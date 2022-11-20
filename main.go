package main

import (
	"fmt"
)

func main() {
	const (
		imageWidth  = 256
		imageHeight = 256
	)

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / (imageWidth - 1)
			g := float64(j) / (imageHeight - 1)
			b := 0.25

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
