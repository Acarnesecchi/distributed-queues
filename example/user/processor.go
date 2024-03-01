package main

import (
	"math"
	"math/rand"
	"time"
)

const (
	minTemperature     = -10
	maxTemperature     = 45
	minCO2             = 300
	maxCO2             = 500
	minHumidity        = 20
	maxHumidity        = 70
	optimalTemperature = 20
	optimalCO2         = 400
	optimalHumidity    = 40
	minProcessTime     = 0.5
	maxProcessTime     = 3.5
)

type MeteoDataProcessor struct {
	TemperatureSpace []float64
	TemperatureVals  []float64
	CO2Space         []float64
	CO2Vals          []float64
	HumiditySpace    []float64
	HumidityVals     []float64
}

func NewMeteoDataProcessor() *MeteoDataProcessor {
	processor := &MeteoDataProcessor{}
	processor.TemperatureSpace, processor.TemperatureVals = genDistribution(minTemperature, maxTemperature, optimalTemperature)
	processor.CO2Space, processor.CO2Vals = genDistribution(minCO2, maxCO2, optimalCO2)
	processor.HumiditySpace, processor.HumidityVals = genDistribution(minHumidity, maxHumidity, optimalHumidity)
	return processor
}

func (p *MeteoDataProcessor) ProcessMeteoData(temperature, humidity float64) float64 {
	temperatureWellness := valueFromDistribution(p.TemperatureSpace, p.TemperatureVals, temperature)
	humidityWellness := valueFromDistribution(p.HumiditySpace, p.HumidityVals, humidity)
	airWellness := round(2/(1/temperatureWellness+1/humidityWellness), 2)
	p.simulateExecutionTime()
	return airWellness
}

func (p *MeteoDataProcessor) ProcessPollutionData(co2 float64) float64 {
	co2Wellness := valueFromDistribution(p.CO2Space, p.CO2Vals, co2)
	return round(co2Wellness, 2)
}

func (p *MeteoDataProcessor) simulateExecutionTime() {
	time.Sleep(time.Duration(rand.Float64()*(maxProcessTime-minProcessTime)+minProcessTime) * time.Second)
}

// helper functions
func genDistribution(min, max, opt float64) ([]float64, []float64) {
	location := opt
	scale := getScale(min, max)
	x := linspace(min, max, 1000)
	p := skewNormPDF(x, location, scale)
	return x, normalizeData(p)
}

func skewNormPDF(x []float64, center, scale float64) []float64 {
	t := make([]float64, len(x))
	for i, val := range x {
		t[i] = (val - center) / scale
	}
	return multiplyArrays(normPDF(t), 0.5)
}

func normPDF(x []float64) []float64 {
	result := make([]float64, len(x))
	for i, val := range x {
		result[i] = normPDFSingle(val)
	}
	return result
}

func normPDFSingle(x float64) float64 {
	return math.Exp(-0.5*x*x) / math.Sqrt(2*math.Pi)
}

func getScale(min, max float64) float64 {
	return ((max - min) / 25) * 8
}

func normalizeData(data []float64) []float64 {
	minVal, maxVal := minMax(data)
	result := make([]float64, len(data))
	for i, val := range data {
		result[i] = (val - minVal) / (maxVal - minVal)
	}
	return result
}

func minMax(data []float64) (float64, float64) {
	minVal, maxVal := data[0], data[0]
	for _, val := range data {
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return minVal, maxVal
}

func valueFromDistribution(space, values []float64, x float64) float64 {
	position := searchSorted(space, x)
	if position == len(space) {
		position--
	}
	value := values[position]
	if value == 0 {
		value = 0.001
	}
	return value
}

func searchSorted(space []float64, x float64) int {
	for i, val := range space {
		if val >= x {
			return i
		}
	}
	return len(space)
}

func round(val float64, decimalPlaces int) float64 {
	shift := math.Pow(10, float64(decimalPlaces))
	return math.Round(val*shift) / shift
}

func linspace(min, max float64, n int) []float64 {
	result := make([]float64, n)
	step := (max - min) / float64(n-1)
	for i := 0; i < n; i++ {
		result[i] = min + step*float64(i)
	}
	return result
}

func multiplyArrays(arr []float64, scalar float64) []float64 {
	result := make([]float64, len(arr))
	for i, val := range arr {
		result[i] = val * scalar
	}
	return result
}
