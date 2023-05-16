package main

import (
	"fmt"
	"github.com/wcharczuk/go-chart/v2"
	"os"
	"strconv"
	"text/tabwriter"
)

func main() {
	n, dataX, dataY := input()
	isPositive := true
	for i := 0; i < n; i++ {
		if dataX[i] <= 0 {
			isPositive = false
		}
		if dataY[i] <= 0 {
			isPositive = false
		}
		break
	}
	_, whereToWrite := output()
	if isPositive {
		fullOutput(n, dataX, dataY, whereToWrite)
	} else {
		smallOutput(n, dataX, dataY, whereToWrite)
	}
}

func fullOutput(n int, dataX, dataY []float64, whereToWrite *os.File) {
	pearsonCoeff, linearApproximation := linealApprox(n, dataX, dataY)
	powerApproximation := powerApprox(n, dataX, dataY)
	expApproximation := expApprox(n, dataX, dataY)
	logApproximation := logApprox(n, dataX, dataY)
	quadraticApproximation := quadraticApprox(n, dataX, dataY)
	cubeApproximation := cubeApprox(n, dataX, dataY)

	allApproxs := []Data{
		linearApproximation,
		powerApproximation,
		expApproximation,
		logApproximation,
		quadraticApproximation,
		cubeApproximation,
	}

	bestApprox := allApproxs[0]
	min := allApproxs[0].meanSquareDeviation
	for _, value := range allApproxs {
		if value.meanSquareDeviation < min {
			bestApprox = value
			min = value.meanSquareDeviation
		}
	}

	tbw := tabwriter.NewWriter(whereToWrite, 1, 1, 1, ' ', 0)

	for _, value := range allApproxs {
		fmt.Fprintln(whereToWrite, value.nameOfApprox+": ")
		var functionString string
		switch value.nameOfApprox {
		case LINEAR:
			functionString = fmt.Sprintf("P(x) = %.3Fx + %.3F; Коэффициент корреляции Пирсона: %.3F",
				value.coefficents[0], value.coefficents[1], pearsonCoeff)
		case QUADRATIC:
			functionString = fmt.Sprintf("P(x) = %.3Fx^2 + %.3Fx + %.3F",
				value.coefficents[0], value.coefficents[1], value.coefficents[2])
		case CUBIC:
			functionString = fmt.Sprintf("P(x) = %.3Fx^3 + %.3Fx^2 + %.3Fx + %.3F",
				value.coefficents[0], value.coefficents[1], value.coefficents[2], value.coefficents[3])
		case EXP:
			functionString = fmt.Sprintf("P(x) = %.3F * e^%.3Fx",
				value.coefficents[0], value.coefficents[1])
		case LOG:
			functionString = fmt.Sprintf("P(x) = %.3Fx + %.3F",
				value.coefficents[0], value.coefficents[1])
		case POWER:
			functionString = fmt.Sprintf("P(x) = %.3F * x^%.3F",
				value.coefficents[0], value.coefficents[1])
		}
		fmt.Fprintln(whereToWrite, functionString)
		meanSqD := strconv.FormatFloat(value.meanSquareDeviation, 'f', 3, 64)
		fmt.Fprintln(whereToWrite, MEAN_SQ_D+meanSqD)
		lineX, lineY, linePhi, lineE := "X:\t", "Y:\t", "Phi:\t", "E:\t"

		for i := 0; i < n; i++ {
			lineX += strconv.FormatFloat(value.xData[i], 'f', 3, 64) + "\t"
			lineY += strconv.FormatFloat(value.yData[i], 'f', 3, 64) + "\t"
			linePhi += strconv.FormatFloat(value.phi[i], 'f', 3, 64) + "\t"
			lineE += strconv.FormatFloat(value.eps[i], 'f', 3, 64) + "\t"
		}
		fmt.Fprintln(tbw, lineX)
		fmt.Fprintln(tbw, lineY)
		fmt.Fprintln(tbw, linePhi)
		fmt.Fprintln(tbw, lineE)
		tbw.Flush()
		fmt.Fprintln(whereToWrite, "\n")
	}
	fmt.Fprintln(whereToWrite, BEST_APPROX+bestApprox.nameOfApprox+"\n")
	fmt.Fprintln(tbw)

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 20,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth: chart.Disabled,
					DotWidth:    5,
				},
				XValues: allApproxs[0].xData,
				YValues: allApproxs[0].yData,
			},
			chart.ContinuousSeries{
				Name:    LINEAR,
				XValues: allApproxs[0].xData,
				YValues: allApproxs[0].phi,
			},
			chart.ContinuousSeries{
				Name:    POWER,
				XValues: allApproxs[1].xData,
				YValues: allApproxs[1].phi,
			},
			chart.ContinuousSeries{
				Name:    EXP,
				XValues: allApproxs[2].xData,
				YValues: allApproxs[2].phi,
			},
			chart.ContinuousSeries{
				Name:    LOG,
				XValues: allApproxs[3].xData,
				YValues: allApproxs[3].phi,
			},
			chart.ContinuousSeries{
				Name:    QUADRATIC,
				XValues: allApproxs[4].xData,
				YValues: allApproxs[4].phi,
			},
			chart.ContinuousSeries{
				Name:    CUBIC,
				XValues: allApproxs[5].xData,
				YValues: allApproxs[5].phi,
			},
		},
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)

}

func smallOutput(n int, dataX, dataY []float64, whereToWrite *os.File) {
	pearsonCoeff, linearApproximation := linealApprox(n, dataX, dataY)
	quadraticApproximation := quadraticApprox(n, dataX, dataY)
	cubeApproximation := cubeApprox(n, dataX, dataY)
	fmt.Fprintln(whereToWrite, CANT_OUTPUT)

	allApproxs := []Data{
		linearApproximation,
		quadraticApproximation,
		cubeApproximation,
	}

	bestApprox := allApproxs[0]
	min := allApproxs[0].meanSquareDeviation
	for _, value := range allApproxs {
		if value.meanSquareDeviation < min {
			bestApprox = value
			min = value.meanSquareDeviation
		}
	}

	tbw := tabwriter.NewWriter(whereToWrite, 1, 1, 1, ' ', 0)

	for _, value := range allApproxs {
		fmt.Fprintln(whereToWrite, value.nameOfApprox+": ")
		var functionString string
		switch value.nameOfApprox {
		case LINEAR:
			functionString = fmt.Sprintf("P(x) = %.3Fx + %.3F; Коэффициент корреляции Пирсона: %.3F",
				value.coefficents[0], value.coefficents[1], pearsonCoeff)
		case QUADRATIC:
			functionString = fmt.Sprintf("P(x) = %.3Fx^2 + %.3Fx + %.3F",
				value.coefficents[0], value.coefficents[1], value.coefficents[2])
		case CUBIC:
			functionString = fmt.Sprintf("P(x) = %.3Fx^3 + %.3Fx^2 + %.3Fx + %.3F",
				value.coefficents[0], value.coefficents[1], value.coefficents[2], value.coefficents[3])
		case EXP:
			functionString = fmt.Sprintf("P(x) = %.3F * e^%.3Fx",
				value.coefficents[0], value.coefficents[1])
		case LOG:
			functionString = fmt.Sprintf("P(x) = %.3Fx + %.3F",
				value.coefficents[0], value.coefficents[1])
		case POWER:
			functionString = fmt.Sprintf("P(x) = %.3F * x^%.3F",
				value.coefficents[0], value.coefficents[1])
		}
		fmt.Fprintln(whereToWrite, functionString)
		meanSqD := strconv.FormatFloat(value.meanSquareDeviation, 'f', 3, 64)
		fmt.Fprintln(whereToWrite, MEAN_SQ_D+meanSqD)
		lineX, lineY, linePhi, lineE := "X:\t", "Y:\t", "Phi:\t", "E:\t"

		for i := 0; i < n; i++ {
			lineX += strconv.FormatFloat(value.xData[i], 'f', 3, 64) + "\t"
			lineY += strconv.FormatFloat(value.yData[i], 'f', 3, 64) + "\t"
			linePhi += strconv.FormatFloat(value.phi[i], 'f', 3, 64) + "\t"
			lineE += strconv.FormatFloat(value.eps[i], 'f', 3, 64) + "\t"
		}
		fmt.Fprintln(tbw, lineX)
		fmt.Fprintln(tbw, lineY)
		fmt.Fprintln(tbw, linePhi)
		fmt.Fprintln(tbw, lineE)
		tbw.Flush()
		fmt.Fprintln(whereToWrite, "\n")
	}
	fmt.Fprintln(whereToWrite, BEST_APPROX+bestApprox.nameOfApprox+"\n")
	fmt.Fprintln(tbw)

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 20,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth: chart.Disabled,
					DotWidth:    5,
				},
				XValues: allApproxs[0].xData,
				YValues: allApproxs[0].yData,
			},
			chart.ContinuousSeries{
				Name:    LINEAR,
				XValues: allApproxs[0].xData,
				YValues: allApproxs[0].phi,
			},
			chart.ContinuousSeries{
				Name:    QUADRATIC,
				XValues: allApproxs[1].xData,
				YValues: allApproxs[1].phi,
			},
			chart.ContinuousSeries{
				Name:    CUBIC,
				XValues: allApproxs[2].xData,
				YValues: allApproxs[2].phi,
			},
		},
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)

}
