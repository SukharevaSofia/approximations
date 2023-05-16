package main

import "math"

func linealApprox(n int, dataX, dataY []float64) (float64, Data) {
	sx, sxx, sy, sxy := 0., 0., 0., 0.
	for i := 0; i < len(dataX); i++ {
		sx += dataX[i]
		sxx += dataX[i] * dataX[i]
		sy += dataY[i]
		sxy += dataX[i] * dataY[i]
	}

	answKramer := GetKramer([][]float64{{sxx, sx}, {sx, float64(n)}}, []float64{sxy, sy})

	//a & b are coefficents of the y=ax+b
	a, b := answKramer[0], answKramer[1]
	coefficents := []float64{a, b}
	f := func(element float64) float64 { return a*element + b }
	phi := getPhi(dataX, f)
	eps := getEps(phi, dataY)
	meanSqDev := meanSquareDeviation(f, dataX, dataY)
	pearsonsCoef := pearsonsCoefficient(dataX, dataY, sx, sy, n)

	return pearsonsCoef, Data{f, coefficents, meanSqDev, dataX, dataY, phi, eps, LINEAR}
}

func pearsonsCoefficient(xData, yData []float64, sx, sy float64, n int) float64 {
	meanX := sx / float64(n)
	meanY := sy / float64(n)
	sum := 0.
	for i := 0; i < n; i++ {
		sum += (xData[i] - meanX) * (yData[i] - meanY)
	}
	return sum / math.Sqrt(squareDeviation(xData, meanX)*squareDeviation(yData, meanY))
}
