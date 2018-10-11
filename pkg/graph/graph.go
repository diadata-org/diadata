package graph

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	IMAGE_HEIGHT   = 1
	IMAGE_WIDTH    = 6
	NUM_DATAPOINTS = 75
)

func PriceGraph(prices []float64, times []int64, path string) error {
	graph, err := plot.New()
	if err != nil {
		return err
	}

	// downsample data
	factor, index, sumPrices, sumTimes := len(times)/NUM_DATAPOINTS, 0, float64(0), int64(0)
	if factor < 1 {
		factor = 1
	}
	min, max := math.Inf(1), math.Inf(-1)
	dataPoints := make(plotter.XYs, NUM_DATAPOINTS)

	for i := 1; i < len(times) && index < NUM_DATAPOINTS; i++ {
		sumPrices += prices[i]
		sumTimes += times[i]

		if i%factor == 0 {
			dataPoints[index].X = float64(sumTimes / int64(factor))
			dataPoints[index].Y = sumPrices / float64(factor)

			log.Println("X: ", dataPoints[index].X, "; Y:", dataPoints[index].Y)
			if dataPoints[index].Y < min {
				min = dataPoints[index].Y
			}

			if dataPoints[index].Y > max {
				max = dataPoints[index].Y
			}

			sumPrices = 0
			sumTimes = 0
			index++
		}
	}

	// insert data in line
	line, err := plotter.NewLine(dataPoints)
	if err != nil {
		return err
	}

	// change presentation
	graph.HideAxes()
	graph.BackgroundColor = color.RGBA{R: 0, G: 0, B: 0, A: 0}
	line.LineStyle.Color = color.RGBA{R: 40, G: 54, B: 142, A: 255}

	line.ShadeColor = new(color.Color)
	*line.ShadeColor = color.RGBA{R: 204, G: 223, B: 248, A: 255}

	graph.Y.Min = min - (max-min)*0.1
	graph.X.Padding = 0
	graph.Y.Padding = 0

	// add line to graph
	graph.Add(line)

	// Save graph
	err = graph.Save(IMAGE_WIDTH*vg.Centimeter, IMAGE_HEIGHT*vg.Centimeter, path)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	numPoints := 30 * 24 * 7
	prices := make([]float64, numPoints)
	times := make([]int64, numPoints)

	//rand.Seed(35131515351)
	rand.Seed(time.Now().Unix())
	var x = 0.0
	for i := 0; i < len(prices); i++ {
		prices[i] = math.Sin(float64(rand.Float64()))
		x += 0.1
	}

	weekAgo := time.Now().Add(-time.Hour * 7 * 24)
	for i := range times {
		times[i] = weekAgo.Add(time.Minute * time.Duration(2*i)).Unix()
	}

	err := PriceGraph(prices, times, "example.png")
	if err != nil {
		log.Println(err)
	}
}