package observer

import (
	"fmt"
	"testing"
)

func TestWeatherStatus(t *testing.T) {
	wd := new(WeatherData)
	wd.Init()

	ccd := new(CurrentConditionsDisplay)
	ccd.Init(wd)

	sd := new(StatisticsDisplay)
	sd.Init(wd)

	fd := new(ForecastDisplay)
	fd.Init(wd)

	wd.SetMeasurements(80, 65, 30.4)
	fmt.Println()
	wd.SetMeasurements(82, 70, 29.2)
	fmt.Println()
	wd.SetMeasurements(78, 90, 29.2)
	fmt.Println()
}
