package ascii

import (
	"errors"
	"fmt"
	"image"

	"github.com/Aaron70/Ascii-Art/images"
)

var DefaultAsciiRunes []rune = []rune{' ', '.', ':', '-', '=', '+', '*', '#', '%', '@'}

type AsciiCanvas[T any] struct {
	Characters []T
	screen     [][]T
	cols, rows int
}

func (c *AsciiCanvas[T]) ToString(newline bool) string {
	str := ""
	for row := range len(c.screen) {
		for col := range len(c.screen[row]) {
			val := c.screen[row][col]
			if r, ok := any(val).(rune); ok {
				str += fmt.Sprintf("%c", r)
			} else {
				str += fmt.Sprintf("%v", val)
			}
		}
		if newline {
			str += "\n"
		}
	}
	return str
}

func (c *AsciiCanvas[T]) ToAsciiArt(img *image.Gray) error {
	if c.cols != img.Bounds().Max.X || c.rows != img.Bounds().Max.Y {
		return errors.New("The given image dimensions doesn't match with the canvas dimensions.")
	}

	for row := range c.rows {
		for col := range c.cols {
			c.screen[row][col] = ramp(c.Characters, (*img).GrayAt(col, row).Y)
		}
	}

	return nil
}

func (c *AsciiCanvas[T]) PreProcessAndConvertToAsciiArt(img *image.Image) {
	image := images.DownScale(img, c.cols, c.rows)
	imgGray := images.ToGrayScale(&image)
	c.ToAsciiArt(imgGray)
}

func CreateAsciiCanvas[T any](cols, rows int, characters []T) AsciiCanvas[T] {
	canvas := AsciiCanvas[T]{
		Characters: characters,
		cols:       cols, rows: rows,
		screen: make([][]T, rows),
	}
	for row := range rows {
		canvas.screen[row] = make([]T, cols)
	}
	return canvas
}

func CreateAsciiCanvasWith[T any](cols, rows int, characters []T, val T) AsciiCanvas[T] {
	canvas := AsciiCanvas[T]{
		Characters: characters,
		cols:       cols, rows: rows,
		screen: make([][]T, rows),
	}
	for row := range rows {
		canvas.screen[row] = make([]T, cols)
		for col := range cols {
			canvas.screen[row][col] = val
		}
	}
	return canvas
}

func ramp[T any](characters []T, color uint8) T {
	return characters[int(float64(len(characters)-1)*(float64(color)/255.0))]
}
