package v2

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rodrwan/opendata/graphql/gmarcone"
)

type Resolver struct {
	GMarconeClient *gmarcone.Client
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
