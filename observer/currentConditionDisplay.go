package observer

import "fmt"

type CurrentConditionsDisplay struct {
	Observer
	DisplayElement

	Temperature float64
	Humidity    float64
	WeatherData *WeatherData
}

var (
	_ Observer       = (*CurrentConditionsDisplay)(nil)
	_ DisplayElement = (*CurrentConditionsDisplay)(nil)
)

// Init
// Constructor 대응
func (cc *CurrentConditionsDisplay) Init(wd *WeatherData) {
	cc.WeatherData = wd
	cc.WeatherData.RegisterObserver(cc)
}

func (cc *CurrentConditionsDisplay) Update(temperature, humidity, pressure float64) {
	cc.Temperature = temperature
	cc.Humidity = humidity

	cc.Display()
}

func (cc *CurrentConditionsDisplay) Display() {
	fmt.Println("현재 상태: 온도", cc.Temperature, "F, 습도", cc.Humidity, "%")
}
