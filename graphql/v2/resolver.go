package v2

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rodrwan/opendata/graphql/gmarcone"
	"github.com/rodrwan/opendata/graphql/transapi"
)

type Resolver struct {
	GMarconeClient *gmarcone.Client
	Transapi       *transapi.Client
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "World", nil
}

func (r *queryResolver) Forecast(ctx context.Context, data *ForecastInput) (*Forecast, error) {
	url := fmt.Sprintf("api/forecast?latitude=%f&longitude=%f&lang=%s", data.Latitude, data.Longitude, data.Lang)
	resp, err := r.GMarconeClient.GET(url)
	if err != nil {
		return nil, err
	}

	var forecast Forecast
	if err := json.Unmarshal(resp, &forecast); err != nil {
		return nil, err
	}

	return &forecast, err
}

func (r *queryResolver) Hearthquake(ctx context.Context, data string) (Earthquake, error) {
	panic("not implemented")
}

func (r *queryResolver) Transantiago(ctx context.Context, data string) (Transantiago, error) {
	resp, err := r.Resolver.Transapi.Get(data)
	if err != nil {
		return Transantiago{}, err
	}

	servicios := make([]Microbus, 0)

	for _, svc := range resp.Services {
		if svc.Valid == 1 {
			servicios = append(servicios, Microbus{
				Valido:    int(svc.Valid),
				Servicio:  svc.Service,
				Patente:   svc.BusPatent,
				Tiempo:    svc.Time,
				Distancia: svc.Distance,
			})
		}
	}

	return Transantiago{
		HoraConsulta: resp.Time,
		Descripcion:  resp.Message,
		Servicios:    servicios,
	}, nil
}
