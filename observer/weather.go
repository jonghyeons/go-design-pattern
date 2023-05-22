package observer

type WeatherData struct {
	Subject
	Observers []Observer

	Temperature float64
	Humidity    float64
	Pressure    float64
}

// Init
// Constructor 대응
func (w *WeatherData) Init() {
	w.Observers = []Observer{}
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.Observers = append(w.Observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	for i := range w.Observers {
		if w.Observers[i] == o {
			if i == len(w.Observers)-1 {
				w.Observers = w.Observers[:i]
			} else {
				w.Observers = append(w.Observers[:i], w.Observers[i+1:]...)
			}
		}
	}
}

func (w *WeatherData) NotifyObservers() {
	for i := range w.Observers {
		w.Observers[i].Update(w.Temperature, w.Humidity, w.Pressure)
	}
}

func (w *WeatherData) MeasurementsChanged() {
	w.NotifyObservers()
}

func (w *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	w.Temperature = temperature
	w.Humidity = humidity
	w.Pressure = pressure

	w.MeasurementsChanged()
}

func (w *WeatherData) GetTemperature() float64 {
	return w.Temperature
}

func (w *WeatherData) GetHumidity() float64 {
	return w.Humidity
}

func (w *WeatherData) GetPressure() float64 {
	return w.Pressure
}
