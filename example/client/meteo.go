package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
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

type MeteoDataDetector struct {
	MinTemperature float64
	MaxTemperature float64
	MinCO2         float64
	MaxCO2         float64
	MinHumidity    float64
	MaxHumidity    float64
}

func NewMeteoDataDetector() *MeteoDataDetector {
	detector := &MeteoDataDetector{}
	detector.MinTemperature, detector.MaxTemperature = genRange(minTemperature, maxTemperature)
	detector.MinCO2, detector.MaxCO2 = genRange(minCO2, maxCO2)
	detector.MinHumidity, detector.MaxHumidity = genRange(minHumidity, maxHumidity)
	return detector
}

func (d *MeteoDataDetector) genTemperature() float64 {
	return round(rand.Float64()*(d.MaxTemperature-d.MinTemperature)+d.MinTemperature, 2)
}

func (d *MeteoDataDetector) genCO2() float64 {
	return round(rand.Float64()*(d.MaxCO2-d.MinCO2)+d.MinCO2, 2)
}

func (d *MeteoDataDetector) genHumidity() float64 {
	return round(rand.Float64()*(d.MaxHumidity-d.MinHumidity)+d.MinHumidity, 2)
}

func (d *MeteoDataDetector) AnalyzeAir() map[string]interface{} {
	timestamp := time.Now().Unix()
	return map[string]interface{}{
		"temperature": d.genTemperature(),
		"humidity":    d.genHumidity(),
		"datetime":    timestamp,
	}
}

func (d *MeteoDataDetector) AnalyzePollution() map[string]interface{} {
	timestamp := time.Now().Unix()
	return map[string]interface{}{
		"co2":      d.genCO2(),
		"datetime": timestamp,
	}
}

func genRange(min, max float64) (float64, float64) {
	return rand.Float64()*(max-min) + min, rand.Float64()*(max-min) + min
}

func round(val float64, decimalPlaces int) float64 {
	shift := math.Pow(10, float64(decimalPlaces))
	return math.Round(val*shift) / shift
}

func marshalAirData(temperature, humidity float64) map[string]interface{} {
	return map[string]interface{}{
		"temperature": temperature,
		"humidity":    humidity,
	}
}

func marshalPollutionData(co2 float64) map[string]interface{} {
	return map[string]interface{}{
		"co2": co2,
	}
}

type Task struct {
	Type     string
	Payload  json.RawMessage
	Priority string
	Metadata map[string]string
}

func genTask(airData, pollutionData map[string]interface{}) (*Task, error) {
	airDataJSON, err := json.Marshal(airData)
	if err != nil {
		return nil, err
	}

	pollutionDataJSON, err := json.Marshal(pollutionData)
	if err != nil {
		return nil, err
	}

	return &Task{
		Type: "MeteoData",
		Payload: json.RawMessage(fmt.Sprintf(`{
			"AirData": %s,
			"PollutionData": %s
		}`, airDataJSON, pollutionDataJSON)),
		Priority: "High",
		Metadata: map[string]string{
			"TimestampAir":       fmt.Sprintf("%v", time.Now().Unix()),
			"TimestampPollution": fmt.Sprintf("%v", time.Now().Unix()),
		},
	}, nil
}

func taskToJSON(task Task) ([]byte, error) {
	taskJSON, err := json.Marshal(task)
	if err != nil {
		return nil, err
	}
	return taskJSON, nil
}

func sendTask(t *Task) {
	url := "http://localhost:25520/new-job"
	method := "POST"
	client := &http.Client{}

	payload, err := taskToJSON(*t)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
}

func main() {
	for {
		time.Sleep(5 * time.Second)
		detector := NewMeteoDataDetector()
		airData := marshalAirData(detector.genTemperature(), detector.genHumidity())
		pollutionData := marshalPollutionData(detector.genCO2())

		task, _ := genTask(airData, pollutionData)
		sendTask(task)
		fmt.Println("Task sent")
	}
}
