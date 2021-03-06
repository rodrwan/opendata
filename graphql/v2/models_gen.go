// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package v2

import (
	"github.com/rodrwan/opendata/graphql/tyaas"
)

type Currently struct {
	Summary           string  `json:"summary"`
	Temperature       float64 `json:"temperature"`
	Humidity          float64 `json:"humidity"`
	WindSpeed         float64 `json:"windSpeed"`
	PrecipProbability float64 `json:"precipProbability"`
}

type Earthquake struct {
	Enlace      string                 `json:"enlace"`
	Latitud     float64                `json:"latitud"`
	Longitud    float64                `json:"longitud"`
	Profundidad float64                `json:"profundidad"`
	Magnitudes  []*EarthquakeMagnitude `json:"magnitudes"`
	Imagen      string                 `json:"imagen"`
}

type EarthquakeMagnitude struct {
	Magnitud float64 `json:"magnitud"`
	Medida   string  `json:"medida"`
	Fuente   string  `json:"fuente"`
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

type Horoscope struct {
	Titulo    string           `json:"titulo"`
	Horoscopo tyaas.ZodiacSign `json:"horoscopo"`
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

type ZodiacSignData struct {
	Nombre     string `json:"nombre"`
	FechaSigno string `json:"fechaSigno"`
	Amor       string `json:"amor"`
	Salud      string `json:"salud"`
	Dinero     string `json:"dinero"`
	Color      string `json:"color"`
	Numero     string `json:"numero"`
}
