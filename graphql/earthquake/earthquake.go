package earthquake

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	errInvalidDate         = errors.New("Invalid date")
	errInternalServerError = errors.New("Internal server error")
	baseURL                = url.URL{
		Scheme: "https",
		Host:   "api.adderou.cl",
		Path:   "sismo/",
	}
)

// Client ...
type Client struct {
	client  *http.Client
	baseURL string
}

// NewClient ...
func NewClient() *Client {
	c := &http.Client{}

	return &Client{
		client:  c,
		baseURL: baseURL.String(),
	}
}

type Earthquake struct {
	Enlace      string                `json:"enlace"`
	Latitud     float64               `json:"latitud"`
	Longitud    float64               `json:"longitud"`
	Profundidad float64               `json:"profundidad"`
	Magnitudes  []EarthquakeMagnitude `json:"magnitudes"`
	Imagen      string                `json:"imagen"`
}

type EarthquakeMagnitude struct {
	Magnitud float64 `json:"magnitud"`
	Medida   string  `json:"medida"`
	Fuente   string  `json:"fuente"`
}

// Do execute an http request.
func (c *Client) Do(req *http.Request) ([]Earthquake, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, errInternalServerError
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := make([]Earthquake, 0)
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return nil, errInvalidDate
	}

	return response, nil
}

// Get http GET method.
func (c *Client) Get(date string) ([]Earthquake, error) {
	req, err := http.NewRequest("GET", c.setURL(""), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("fecha", date)
	req.URL.RawQuery = q.Encode()
	return c.Do(req)

}

func (c *Client) setURL(path string) string {
	return c.baseURL
}
