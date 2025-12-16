package geocoding

import (
	"fmt"
	"net/http"
)

type Response []struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type client struct {
	httpClient http.Client
}

func NewClient(httpClient http.Client) *client {
	return &client{
		httpClient: httpClient,
	}
}

func (c *client) GetCoords(city string) (lat, log float64, err error) {
	res, err := c.httpClient.Get(
		fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%city&count=1&language=ru&dformat=json",
			city,
		))
	if err != nil {
		return 0, 0, err
	}

	if res.StatusCode != http.StatusOK {

		return 0, 0, fmt.Errorf("status code %d", res.StatusCode)
	}

	return 0, 0, err
}
