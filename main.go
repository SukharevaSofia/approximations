package main

import (
	"fmt"
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
}
