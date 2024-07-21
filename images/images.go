package images

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"golang.org/x/image/draw"
)

func OpenImage(path string) (*image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return &img, nil
}

func DownScale(img *image.Image, width, height int) image.Image {
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(newImg, newImg.Rect, *img, (*img).Bounds(), draw.Over, nil)
	return newImg
}

func ToGrayScale(img *image.Image) *image.Gray {
	gray := image.NewGray((*img).Bounds())
	for x := range (*img).Bounds().Max.X {
		for y := range (*img).Bounds().Max.Y {
			color := (*img).At(x, y)
			gray.Set(x, y, color)
		}
	}
	return gray
}

func ResizeKeepImageRatio(availableWidth, availableHeight, width, height int) (int, int) {
	ratio := float64(width) / float64(height)
	nWidth := int(ratio * 2.3 * float64(availableHeight))
	nHeight := availableHeight
	if nWidth > availableWidth {
		ratio = float64(height) / float64(width)
		nHeight = int(ratio * float64(availableWidth))
		nWidth = availableWidth
	}
	return nWidth, nHeight
}
