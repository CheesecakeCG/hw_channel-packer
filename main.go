package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func OpenAllImages() (images []image.Image, paths []string, err error) {
	// images = make([]image.Image, len(os.Args)-1)
	for _, arg := range os.Args[1:] {
		if f, err := os.Open(arg); err != nil {
			continue
		} else {
			if img, _, err := image.Decode(f); err != nil {
				continue
			} else {
				images = append(images, img)
				paths = append(paths, arg)
			}
		}
	}
	return
}


func ConvertImageToMaxBitdepth(img image.Image) (output image.Image) {
	output = imaging.Clone(img)
	if img.ColorModel() == color.GrayModel {
		output := image.NewGray16(img.Bounds())
		for x := 0; x < img.Bounds().Dx(); x++ {
			for y := 0; y < img.Bounds().Dy(); y++ {
				old_col := img.At(x, y)
				new_col := color.Gray16Model.Convert(old_col)
				output.Set(x, y, new_col)
			}
		}
	}
	if img.ColorModel() == color.RGBAModel {
		output := image.NewRGBA64(img.Bounds())
		for x := 0; x < img.Bounds().Dx(); x++ {
			for y := 0; y < img.Bounds().Dy(); y++ {
				old_col := img.At(x, y)
				new_col := color.RGBA64Model.Convert(old_col)
				output.Set(x, y, new_col)
			}
		}
	}

	return output
}






func image_append() {
	images, paths, err := OpenAllImages()
	if err != nil {
		log.Fatal(err)
	}
	if len(images) != 2 {
		log.Fatal("Need exactly 2 images.")
	}


	color_img := imaging.Clone(images[0])
	

	alpha_img := ConvertImageToMaxBitdepth(images[1])
	alpha_img = imaging.Resize(alpha_img, images[0].Bounds().Size().X, images[0].Bounds().Size().Y, imaging.Lanczos)
	alpha_img = imaging.Grayscale(alpha_img)
	for _, a := range os.Args[1:] {
		if a == "-invert" || a == "-i" {
			alpha_img = imaging.Invert(alpha_img)
			break
		}
	}
	
	

	outputImage := imaging.Clone(images[0])

	for y := 0; y < color_img.Bounds().Dy(); y++ {
		for x := 0; x < color_img.Bounds().Dx(); x++ {
			col_pix, ok := color.NRGBA64Model.Convert(color_img.At(x, y)).(color.NRGBA64)
			if !ok {
				log.Println("Color image is not RGBA64")
			}
			gray_pix, ok := color.Gray16Model.Convert(alpha_img.At(x, y)).(color.Gray16)
			if !ok {
				log.Println("Alpha image is not Gray16")
			}

			new_pix := col_pix
			new_pix.A = gray_pix.Y
			
			outputImage.Set(x, y, outputImage.ColorModel().Convert(new_pix))
		}
		
	}
	
	imaging.Save(outputImage, paths[0])
 
}
 
func images_channel_concat()  {
	log.Fatalln("Not implemented yet.")
	// images, err := openAllImages()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, img := range images {

	// }
}

func printHelp() {
	println(
`
hare_ware's channnel packer - a tool for manipulating image channels.

Usage:
-a --append
    Puts the second image as the last unfilled channel of the first image.
    "hwchanpack -a rgb.png a.png"
-c --channel
    Each image specifies the channel it should use. If not specified, the the channel will be black.
    "hwchanpack -c r.png b.png g.png a.png out.png"

-i --invert
    When used with -a, the alpha image will be inverted.
    "hwchanpack -a -i rgb.png a.png"

-h --help
    Prints this help message.
    "hwchanpack -h"
`)
}

func main() {

    for _, a := range os.Args[1:] {
		switch a {
		case "--append":
			fallthrough
		case "-a":
			image_append()
			return
			

		case "--channel":
			fallthrough
		case "-c":
			images_channel_concat()
			return
			

		case "--help":
			fallthrough
		case "-h":
			fallthrough
		default:
			printHelp()
			return

		}
		
	}
	printHelp()

}