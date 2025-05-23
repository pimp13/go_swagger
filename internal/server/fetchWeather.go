package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

type HTTPError struct {
	Message string `json:"message"`
}

func fetchWeather(city string) (*WeatherData, error) {
	const apiKey = "ecfd8aa66dac451426df9b4f6a8ef94e"

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s",
		city,
		apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch weather data failed: %w", err)
	}
	defer resp.Body.Close()

	var weatherData WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("decoding json data for %s failed: %w", city, err)
	}

	return &weatherData, nil
}

// getWeatherByCity godoc
//
//	@Summary		getWeatherByCity
//	@Description	Get the weather for a city
//	@Tags			weather
//	@Accept			json
//	@Produce		json
//	@Param			city	path		string	true	"city name"
//	@Success		200		{object}	WeatherData
//	@Failure		400		{object}	HTTPError
//	@Failure		404		{object}	HTTPError
//	@Failure		500		{object}	HTTPError
//	@Router			/weather/{city} [get]
func getWeatherByCity(c *gin.Context) {
	city := c.Param("city")
	resp, err := fetchWeather(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &HTTPError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
