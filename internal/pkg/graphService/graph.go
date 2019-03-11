package graph

import (
	log "github.com/sirupsen/logrus"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
	"image/color"
	"math"
	"os"
)

const (
	IMAGE_HEIGHT   = 1
	IMAGE_WIDTH    = 6
	NUM_DATAPOINTS = 75
)

func PriceGraph(prices []float64, times []int64, path string) error {
	// color definitions
	blue := color.RGBA{R: 40, G: 54, B: 142, A: 255}
	lightBlue := color.RGBA{R: 204, G: 223, B: 248, A: 255}

	// downsample data
	factor, index, sumPrices, sumTimes := len(times)/NUM_DATAPOINTS, 0, float64(0), int64(0)
	min, max := math.Inf(1), math.Inf(-1)
	dataPoints := make(plotter.XYs, NUM_DATAPOINTS)

	if factor < 1 {
		factor = 1
	}

	for i := 1; i < len(times) && index < NUM_DATAPOINTS; i++ {
		sumPrices += prices[i]
		sumTimes += times[i]

		if i%factor == 0 {
			dataPoints[index].X = float64(sumTimes / int64(factor))
			dataPoints[index].Y = sumPrices / float64(factor)

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

	// resize datapoint array if necessary
	if index < len(dataPoints) {
		newDatapoints := make(plotter.XYs, index)

		for i := range newDatapoints {
			newDatapoints[i].X = dataPoints[i].X
			newDatapoints[i].Y = dataPoints[i].Y
		}
		dataPoints = newDatapoints
	}

	// Create plot object
	graph, err := plot.New()
	if err != nil {
		return err
	}

	// configure graph
	graph.HideAxes()
	graph.BackgroundColor = color.Transparent

	// Create line object
	line, err := plotter.NewLine(dataPoints)
	if err != nil {
		return err
	}

	// configure line
	line.LineStyle.Color = blue
	line.FillColor = lightBlue

	// add a small margin at the bottom
	graph.Y.Min = min - (max-min)*0.1
	graph.X.Padding = 0
	graph.Y.Padding = 0

	// add line to graph
	graph.Add(line)

	// Draw and save graph
	c := vgimg.PngCanvas{vgimg.NewWith(
		vgimg.UseWH(IMAGE_WIDTH * vg.Centimeter, IMAGE_HEIGHT * vg.Centimeter),
		vgimg.UseBackgroundColor(color.Transparent),
	)}
	graph.Draw(draw.New(c))

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = c.WriteTo(f)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	log.Println("Created", path, "with", index, "points")
	return nil
}
