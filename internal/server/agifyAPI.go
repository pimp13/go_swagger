package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AgifyData struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type Genderize struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

func fetchAgifyAPI(name string) (*AgifyData, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in status code %d", resp.StatusCode)
	}

	var data AgifyData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

// getAgify godoc
//
//	@Summary		getAgify
//	@Description	Age prediction based on age
//	@Tags			agify
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"your name"
//	@Success		200		{object}	AgifyData
//	@Failure		400		{object}	HTTPError
//	@Failure		404		{object}	HTTPError
//	@Failure		500		{object}	HTTPError
//	@Router			/agify/{name} [get]
func getAgify(c *gin.Context) {
	name := c.Param("name")
	resp, err := fetchAgifyAPI(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &HTTPError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func fetchGenderize() {
}
