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
	_, whereToWrite := output()
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

//
//func generateLineItems(name string, args []float64, data Data) []opts.LineData {
//	var items []opts.LineData
//	var minPoint, maxPoint float64 = 9999999, -9999999
//	for i := 0; i < len(data.xData); i++ {
//		if data.xData[i] < minPoint {
//			minPoint = data.xData[i]
//		}
//		if data.xData[i] > maxPoint {
//			maxPoint = data.xData[i]
//		}
//	}
//	for i := minPoint; i < maxPoint; i += (maxPoint - minPoint) / float64(len(data.xData)) {
//		val := 0.
//		switch name {
//		case LINEAR:
//			val = args[0]*i + args[1]
//		case QUADRATIC:
//			val = args[0] + args[1]*i + args[2]*i*i
//		case CUBIC:
//			val = args[0] + args[1]*i + args[2]*i*i + args[3]*i*i*i
//		case EXP:
//			val = args[0] * math.Exp(args[1]*i)
//		case LOG:
//			if i > 0 {
//				val = args[0]*math.Log(i) + args[1]
//				items = append(items, opts.LineData{Value: val})
//			}
//		case POWER:
//			val = args[0] * math.Pow(i, args[1])
//			if !math.IsNaN(val) {
//				items = append(items, opts.LineData{Value: val})
//			}
//		default:
//			val = 0
//		}
//		if name != "log" && name != "pow" {
//			items = append(items, opts.LineData{Value: val})
//		}
//	}
//	return items
//}
