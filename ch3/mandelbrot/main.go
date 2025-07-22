// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			plot(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	plot(os.Stdout)
}

func plot(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, superSample(x, y))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func calcPixelColor(x float64, y float64) color.RGBA {
	z := complex(x, y)
	return mandelbrot(z)
}

func superSample(x, y float64) color.RGBA {
	var subPixelColors = [4]color.RGBA{}
	for ix, dx := range [...]float64{-0.5 / 1024, 0.5 / 1024} {
		for iy, dy := range [...]float64{-0.5 / 1024, 0.5 / 1024} {
			subPixelColors[ix+iy] = calcPixelColor(x+dx, y+dy)
		}
	}
	return calcColorAvg(subPixelColors)
}

func calcColorAvg(colors [4]color.RGBA) color.RGBA {
	var avg color.RGBA
	for i, c := range colors {
		if i == 1 {
			avg = c
		} else {
			avg = color.RGBA{
				R: (c.R + avg.R) / 2,
				G: (c.G + avg.G) / 2,
				B: (c.B + avg.B) / 2,
				A: 0xff,
			}

		}
	}
	return avg
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, 0, 0, 0xff}
		}
	}
	return color.RGBA{
		R: 0xff,
		G: 0,
		B: 0,
		A: 0xff,
	}
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
