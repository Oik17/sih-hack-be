package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oik17/sih-agrihealth/internal/utils"
)

type NewsResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		ID     string `json:"id"`
		Source struct {
			Name string `json:"name"`
		} `json:"source"`
		Author      string `json:"author"`
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		PublishedAt string `json:"publishedAt"`
	} `json:"articles"`
}

func NewsControllers(c echo.Context) error {
	API_KEY := utils.Config("NEWSAPI")
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=Crop+livestock+disease+india&from=2024-08-11&sortBy=popularity&apiKey=%s", API_KEY)

	resp, err := http.Get(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	var newsResponse NewsResponse
	err = json.Unmarshal(body, &newsResponse)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	for i := range newsResponse.Articles {
		newsResponse.Articles[i].ID = uuid.New().String()
	}

	return c.JSON(http.StatusOK, newsResponse.Articles)
}
