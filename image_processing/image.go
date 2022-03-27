package main

import (
	"fmt"
	"image"
	"log"
	"net/http"

	"github.com/disintegration/imaging"
)

func main() {
	// Open a test image.
	// src, err := imaging.Open("image/beef.jpeg")
	// fmt.Println(src)

	// url := "https://gorillamove.s3.ap-southeast-1.amazonaws.com/products/images/323-03-2022-16-47-33PM-D094F7-huggies-baby-diaper-dry-belt-s-up-to-7-kg-36-pcs.webp"
	// resp, err := http.Get("https://d1.awsstatic.com/Startups/Activate/AWS_RGB%20activate.9bfe21e86a3b7c4419087a93b0749e9087017bac.png")
	resp, err := http.Get("https://gorillamove.s3.ap-southeast-1.amazonaws.com/products/images/323-03-2022-17-34-28PM-F53B92-beef.jpeg")
	// resp, err := http.Get(url)
	fmt.Println(resp)
	if err != nil {
		return // handle error somehow
	}
	defer resp.Body.Close()

	src, _, err := image.Decode(resp.Body)
	if err != nil {
		return // handle error somehow
	}
	fmt.Println(src)
	// src, err := imaging.Open(img)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// // Crop the original image to 300x300px size using the center anchor.
	// src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	// // Resize the cropped image to width = 200px preserving the aspect ratio.
	// src = imaging.Resize(src, 200, 0, imaging.Lanczos)

	// // Create a blurred version of the image.
	// img1 := imaging.Blur(src, 5)

	// // Create a grayscale version of the image with higher contrast and sharpness.
	// img2 := imaging.Grayscale(src)
	// img2 = imaging.AdjustContrast(img2, 20)
	// img2 = imaging.Sharpen(img2, 2)

	// // Create an inverted version of the image.
	// img3 := imaging.Invert(src)

	// // Create an embossed version of the image using a convolution filter.
	// img4 := imaging.Convolve3x3(
	// 	src,
	// 	[9]float64{
	// 		-1, -1, 0,
	// 		-1, 1, 1,
	// 		0, 1, 1,
	// 	},
	// 	nil,
	// )

	// // Create a new image and paste the four produced images into it.
	// dst := imaging.New(400, 400, color.NRGBA{0, 0, 0, 0})
	// dst = imaging.Paste(dst, img1, image.Pt(0, 0))
	// dst = imaging.Paste(dst, img2, image.Pt(0, 200))
	// dst = imaging.Paste(dst, img3, image.Pt(200, 0))
	// dst = imaging.Paste(dst, img4, image.Pt(200, 200))
	// Resize and crop the srcImage to fill the 100x100px area.
	dstImageFill := imaging.Fill(src, 150, 150, imaging.Center, imaging.Lanczos)

	// Resize srcImage to size = 128x128px using the Lanczos filter.
	dstImage128 := imaging.Resize(src, 150, 150, imaging.Lanczos)

	// Resize srcImage to size = 256x256px using the Lanczos filter.
	dstImageNearest := imaging.Resize(src, 150, 150, imaging.NearestNeighbor)

	// Save the resulting image as JPEG.
	err = imaging.Save(dstImageFill, "image/dstImageFill.png")
	err = imaging.Save(dstImage128, "image/dstImage128.png")
	err = imaging.Save(dstImageNearest, "image/dstImageNearest.png")
	fmt.Println("test---------------------")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
