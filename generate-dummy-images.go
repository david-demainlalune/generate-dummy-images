package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math/rand"
	"os"
	"time"
)

func randRange(min int, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

func generateImage(width int, height int, name string, c chan<- string) {
	im := image.NewRGBA(image.Rect(0, 0, width, height))
	b := im.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			im.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
	file, _ := os.Create(name)
	defer file.Close()

	jpeg.Encode(file, im, &jpeg.Options{jpeg.DefaultQuality})
	c <- fmt.Sprintf("generated %s\n", name)
}

func generateImages(count int, minWidth int, maxWidth int, minHeight int, maxHeight int, baseName string) {
	c := make(chan string)
	tabCount := len(fmt.Sprintf("%d", count))
	for i := 0; i < count; i++ {
		name := fmt.Sprintf("%s%0*d.jpg", baseName, tabCount, i)
		width := randRange(minWidth, maxWidth)
		height := randRange(minHeight, maxHeight)
		go generateImage(width, height, name, c)

	}

	for i := 0; i < count; i++ {
		fmt.Println(<-c)
	}
}

func main() {
	count := flag.Int("count", 0, "count of generated images")
	minWidth := flag.Int("minWidth", 400, "min width of generated image")
	maxWidth := flag.Int("maxWidth", 800, "max width of generated image")
	minHeight := flag.Int("minHeight", 400, "min height of generated image")
	maxHeight := flag.Int("maxHeight", 800, "max height of generated image")
	baseName := flag.String("baseName", "dummyFile_", "root filename")

	flag.Parse()

	if *maxWidth < *minWidth {
		fmt.Println("error maxWidth < minWidth")
		os.Exit(1)
	}

	if *maxHeight < *minHeight {
		fmt.Println("error maxHeight < minHeight")
		os.Exit(1)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	start := time.Now()

	generateImages(*count, *minWidth, *maxWidth, *minHeight, *maxHeight, *baseName)

	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}
