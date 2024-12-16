package imagetodot

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/color"

	_ "image/jpeg"
	_ "image/png"
	"os"
)

type Image struct {
	img   image.Image
	Width int
	Hight int
}

type ImageToSymbols interface {
	imageToRGBArray() *[][]int
	OpenImage(string) *[][]int
}

func Handel(w int, h int, imgname string) *[][]int {
	var img ImageToSymbols
	img = &Image{Width: w, Hight: h}
	ou := img.OpenImage(imgname)
	return ou
}

func (img *Image) imageToRGBArray() *[][]int {
	bounds := img.img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	grayArray := make([][]int, height)

	for y := 0; y < height; y++ {
		grayArray[y] = make([]int, width)
		for x := 0; x < width; x++ {
			c := img.img.At(x, y)

			grayColor := color.GrayModel.Convert(c).(color.Gray)

			grayArray[y][x] = int(grayColor.Y)
		}
	}

	return &grayArray
}

func (img *Image) OpenImage(imgname string) *[][]int {

	file, err := os.Open(imgname)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	i, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)

	}
	if img.Width > 0 && img.Hight > 0 {

		img.img = resize.Resize(uint(img.Width), uint(img.Hight), i, resize.Lanczos3)
	}else{
		img.img = i
	}

	return img.imageToRGBArray()
}
