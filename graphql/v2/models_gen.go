// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package v2

type Currently struct {
	Summary           string  `json:"summary"`
	Temperature       float64 `json:"temperature"`
	Humidity          float64 `json:"humidity"`
	WindSpeed         float64 `json:"windSpeed"`
	PrecipProbability float64 `json:"precipProbability"`
}

type Earthquake struct {
	Enlace      string              `json:"enlace"`
	Latitud     float64             `json:"latitud"`
	Longitude   float64             `json:"longitude"`
	Profundidad float64             `json:"profundidad"`
	Magnitude   EarthquakeMagnitude `json:"magnitude"`
	Imagen      string              `json:"imagen"`
}

type EarthquakeMagnitude struct {
	Magnitude float64 `json:"magnitude"`
	Medida    string  `json:"medida"`
	Fuente    string  `json:"fuente"`
}

type Forecast struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Name      string    `json:"name"`
	Timezone  string    `json:"timezone"`
	Currently Currently `json:"currently"`
}

type ForecastInput struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Lang      string  `json:"lang"`
}

type Microbus struct {
	Valido    int    `json:"valido"`
	Servicio  string `json:"servicio"`
	Patente   string `json:"patente"`
	Tiempo    string `json:"tiempo"`
	Distancia string `json:"distancia"`
}

type Transantiago struct {
	HoraConsulta string     `json:"horaConsulta"`
	Descripcion  string     `json:"descripcion"`
	Servicios    []Microbus `json:"servicios"`
}
