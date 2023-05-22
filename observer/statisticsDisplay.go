package observer

import "fmt"

type StatisticsDisplay struct {
	Observer
	DisplayElement

	MaxTemperature float64
	MinTemperature float64
	TemperatureSum float64
	NumReading     int
	WeatherData    *WeatherData
}

var (
	_ Observer       = (*StatisticsDisplay)(nil)
	_ DisplayElement = (*StatisticsDisplay)(nil)
)

// Init
// Constructor 대응
func (sd *StatisticsDisplay) Init(wd *WeatherData) {
	sd.MaxTemperature = 0.0
	sd.MinTemperature = 200
	sd.TemperatureSum = 0.0
	sd.WeatherData = wd
	sd.WeatherData.RegisterObserver(sd)
}

func (sd *StatisticsDisplay) Update(temperature, humidity, pressure float64) {
	sd.TemperatureSum = temperature
	sd.NumReading++
	if temperature > sd.MaxTemperature {
		sd.MaxTemperature = temperature
	}

	if temperature < sd.MinTemperature {
		sd.MinTemperature = temperature
	}

	sd.Display()
}

func (sd *StatisticsDisplay) Display() {
	fmt.Println("평균/최고/최저 온도 =", sd.TemperatureSum/float64(sd.NumReading), "/", sd.MaxTemperature, "/", sd.MinTemperature)
}
