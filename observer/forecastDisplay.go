package observer

import "fmt"

type ForecastDisplay struct {
	Observer
	DisplayElement

	CurrentPressure float64
	LastPressure    float64
	WeatherData     *WeatherData
}

var (
	_ Observer       = (*ForecastDisplay)(nil)
	_ DisplayElement = (*ForecastDisplay)(nil)
)

func (fd *ForecastDisplay) Init(wd *WeatherData) {
	fd.WeatherData = wd
	fd.WeatherData.RegisterObserver(fd)
}

func (fd *ForecastDisplay) Update(temperature, humidity, pressure float64) {
	fd.LastPressure = fd.CurrentPressure
	fd.CurrentPressure = pressure

	fd.Display()
}

func (fd *ForecastDisplay) Display() {
	fmt.Printf("기상 예보:")
	if fd.CurrentPressure > fd.LastPressure {
		fmt.Println("날씨가 좋아지고 있습니다.")
	} else if fd.CurrentPressure == fd.LastPressure {
		fmt.Println("지금과 비슷할 것 같습니다.")
	} else {
		fmt.Println("쌀쌀하며 비가 올 것 같습니다.")
	}
}
