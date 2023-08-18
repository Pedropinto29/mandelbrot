package main

import (
	"fmt"
	"html/template"
	"image/color"
	"net/http"
	"sync"

	"github.com/fogleman/gg"
)

// Calculate the number of iterations needed to see if the point is inside or outside the Mandelbrot set.
func Color(n_max int, Cx float64, Cy float64) int {
	var xn, yn, tempx, tempy float64 = 0, 0, 0, 0
	var n int = 0

	//make sure the point doesn't go to infinite
	for (xn*xn+yn*yn) < 4 && n < n_max {
		tempx = xn
		tempy = yn

		//Mandelbrot equations
		xn = tempx*tempx - tempy*tempy + Cx
		yn = 2*tempx*tempy + Cy
		n++
	}

	return n
}

// Allows the multithreading color calculations
func handlerGG(hd *sync.WaitGroup, id, max_Iter int, Cx, Cy float64, i, y int, dc *gg.Context) {
	defer hd.Done()

	color := getColor(Color(max_Iter, Cx, Cy), max_Iter)
	dc.SetRGB255(int(color.R), int(color.G), int(color.B))
	dc.SetPixel(i, y)
}

// Give the color of a pixel based on the number of iterations
func getColor(iter, max_Iter int) color.NRGBA {
	if iter == max_Iter {
		return color.NRGBA{0, 0, 0, 255} // Black color for inside of the Mandelbrot
	}

	if iter >= max_Iter-10 {
		return color.NRGBA{255, 0, 0, 255} // Red shadow for just outside the Mandelbrot
	}

	hue := float64(iter) / float64(max_Iter-10)
	saturation := 1.0
	value := 1.0

	var r, g, b uint8
	hi := int(hue * 6)
	f := hue*6 - float64(hi)
	p := uint8(value * (1 - saturation) * 255)
	q := uint8(value * (1 - f*saturation) * 255)
	t := uint8(value * (1 - (1-f)*saturation) * 255)
	switch hi {
	case 0:
		r, g, b = uint8(value*255), t, p
	case 1:
		r, g, b = q, uint8(value*255), p
	case 2:
		r, g, b = p, uint8(value*255), t
	case 3:
		r, g, b = p, q, uint8(value*255)
	case 4:
		r, g, b = t, p, uint8(value*255)
	default:
		r, g, b = uint8(value*255), p, q
	}

	return color.NRGBA{r, g, b, 255}
}

func main() {
	fmt.Println("Starting web server")

	const img_width, img_height, max_Iter = 800, 800, 100

	const x_min, x_max, y_min, y_max float64 = -2, 0.75, -1.5, 1.25
	var cx, cy, zoom float64 = 0, 0, 0
	var x_center, y_center float64 = img_width / 2, img_height / 2

	// Display image
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dc := gg.NewContext(img_width, img_height)

		var hd sync.WaitGroup
		var workers = 4

		for i := 0; i < img_width; i++ {
			for y := 0; y < img_height; y = y + workers {
				for z := 0; z < workers; z++ {
					cx = ((float64(i)*(x_max-x_min))/(float64(img_width)+cx) - 2)
					cy = ((float64(y+z)*(y_min-y_max))/(float64(img_height)+cy) + 1.25)

					hd.Add(1)
					go handlerGG(&hd, z, max_Iter, cx, cy, i, y+z, dc)
				}

				hd.Wait() //wait until all the threads have finished for each line
			}
		}

		w.Header().Set("Content-Type", "image/png")
		err := dc.EncodePNG(w)
		if err != nil {
			http.Error(w, "Error generating Mandelbrot image", http.StatusInternalServerError)
			return
		}
	})

	//Display interactive image
	http.HandleFunc("/interactive", func(w http.ResponseWriter, r *http.Request) {
		zoomChange := 1.0
		move := 5.0

		if r.FormValue("zoom") == "in" {
			zoom *= zoomChange
		} else if r.FormValue("zoom") == "out" {
			zoom /= zoomChange
		}

		if r.FormValue("move") == "left" {
			x_center -= move * zoom
		} else if r.FormValue("move") == "right" {
			x_center += move * zoom
		} else if r.FormValue("move") == "up" {
			y_center -= move * zoom
		} else if r.FormValue("move") == "down" {
			y_center += move * zoom
		}

		// TO-DO: Implement calculations to display the zoomed or moved image

		webpage, err := template.ParseFiles("interactive.html")
		if err != nil {
			fmt.Println("Couldn't parse html file")
		}
		w.Header().Set("Content-Type", "text/html")
		webpage.Execute(w, nil)
	})

	fmt.Println("Connected to http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting web server:", err)
	}

}
