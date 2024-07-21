package main

import (
	"fmt"
	"image"

	"github.com/Aaron70/Ascii-Art/ascii"
	"github.com/Aaron70/Ascii-Art/images"
	"github.com/Aaron70/Ascii-Art/terminal"
)

func main() {
	imagePath := "./resources/mamibw.jpg"
	widht, height, err := terminal.GetScreenDimensions()
	if err != nil {
		panic(fmt.Sprintf("An error occurred: %v\n", err))
	}
	img, err := images.OpenImage(imagePath)
	if err != nil {
		panic(fmt.Sprintf("An error occurred: %v\n", err))
	}
	convert(*img, widht, height, true)
}

func convert(img image.Image, availableWidth, availableHeight int, newline bool) {
	var DefaultAsciiRunes []rune = []rune{' ', '-', '=', '+', '#', '&', '%', '@'}
	//var DefaultAsciiRunes []rune = []rune{' ', '.', ':', '-', '=', '+', '#', '&', '%', '@'}
	//var DefaultAsciiStrings []string = []string{" ", " . ", ":", "-", "=", "+", "#", "&", "%", "@"}
	//var DefaultAsciiStrings []string = []string{"▫️", "〰️", "➿", "➰", "✖️", "🟰", "💲", "◾"}
	//var DefaultAsciiHearts []string = []string{"🤍", "🤎", "🖤"}
	//var DefaultAsciiHearts []string = []string{"🤍", "🩵", "💚", "❤️", "🩶", "🤎", "🖤"}
	//var DefaultAsciiInt []int = []int{0, 1, 3, 4, 5, 6, 7, 8}
	imgWidth, imgHeight := img.Bounds().Max.X, img.Bounds().Max.Y
	width, height := images.ResizeKeepImageRatio(availableWidth, availableHeight, imgWidth, imgHeight)
	asciiArt := ascii.CreateAsciiCanvas(width, height, DefaultAsciiRunes)

	asciiArt.PreProcessAndConvertToAsciiArt(&img)
	fmt.Print(asciiArt.ToString(newline))
}

//func manualPreprocessingImage() {
//	imagePath := "./resources/selfie2.jpg"
//	width, height, err := terminal.GetScreenDimensions()
//	if err != nil {
//		panic(fmt.Sprintf("An error occurred: %v\n", err))
//	}
//	img, err := images.OpenImage(imagePath)
//	if err != nil {
//		panic(fmt.Sprintf("An error occurred: %v\n", err))
//	}
//	imgWidth, imgHeight := (*img).Bounds().Max.X, (*img).Bounds().Max.Y
//	width, height = images.ResizeKeepImageRatio(width, height, imgWidth, imgHeight)
//	asciiArt := ascii.CreateAsciiCanvasWith[rune](width, height, )
//
//	asciiArt.PreProcessAndConvertToAsciiArt(img)
//	fmt.Print(asciiArt.ToString(true))
//}
//
//func preBuildProcessingImage() {
//	imagePath := "./resources/Jake.png"
//	img, err := images.OpenImage(imagePath)
//	if err != nil {
//		panic(fmt.Sprintf("An error occurred: %v\n", err))
//	}
//	imgWidth, imgHeight := (*img).Bounds().Max.X, (*img).Bounds().Max.Y
//	width, height, err := terminal.GetScreenDimensions()
//	if err != nil {
//		panic(fmt.Sprintf("An error occurred: %v\n", err))
//	}
//	width, height = images.ResizeKeepImageRatio(width, height, imgWidth, imgHeight)
//	asciiArt := ascii.CreateAsciiCanvasWith(width, height, '*')
//	asciiArt.PreProcessAndConvertToAsciiArt(img)
//	fmt.Print(asciiArt.ToString(true))
//}
