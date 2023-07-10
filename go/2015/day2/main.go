package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var fileContent, showPosition, err = loadFile()

	if err != nil {
		fmt.Println("File not found!")
		os.Exit(1)
	}

	if showPosition == "" {
		calculateWrappingPaper(fileContent)
	}

	calculateRibbon(fileContent)

}

func loadFile() ([]byte, string, error) {
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	var showPosition string
	if len(os.Args) > 2 {
		showPosition = os.Args[2]
	}

	var fileContent, err = os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}

	return fileContent, showPosition, nil
}

func handleError(err error) {
	if err != nil {
		fmt.Println("It was not possible to convert string to number!")
		os.Exit(2)
	}
}

func calculateWrappingPaper(fileContent []byte) {
	var wrapping_paper float64 = 0
	for _, lines := range strings.Split(string(fileContent), "\n") {
		box := strings.Split(lines, "x")
		if len(box) == 3 {
			l, err := strconv.ParseFloat(box[0], 64)
			handleError(err)
			w, err := strconv.ParseFloat(box[1], 64)
			handleError(err)
			h, err := strconv.ParseFloat(box[2], 64)
			handleError(err)

			l_w := l * w
			w_h := w * h
			h_l := h * l

			min_value := l_w
			if min_value > math.Min(w_h, h_l) {
				min_value = math.Min(w_h, h_l)
			}

			// 2*(l*w) + 2*(w*h) + 2*(h*l)
			wrapping_paper += (2 * l_w) + (2 * w_h) + (2 * h_l) + min_value
		}
	}

	fmt.Printf("%f", wrapping_paper)
	os.Exit(0)
}

func calculateRibbon(fileContent []byte) {
	var ribbon float64 = 0
	for _, lines := range strings.Split(string(fileContent), "\n") {
		box := strings.Split(lines, "x")
		if len(box) == 3 {
			l, err := strconv.ParseFloat(box[0], 64)
			handleError(err)
			w, err := strconv.ParseFloat(box[1], 64)
			handleError(err)
			h, err := strconv.ParseFloat(box[2], 64)
			handleError(err)

			box_size := []float64{l, w, h}
			sort.Float64s(box_size)

			sm_size := box_size[0]
			min_per := box_size[1]

			ribbon += (l * w * h) + ((sm_size * 2) + (min_per * 2))
		}
	}

	fmt.Printf("%f", ribbon)
	os.Exit(0)
}
