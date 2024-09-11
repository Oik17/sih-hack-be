package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type TranslateRequest struct {
	Text       string `json:"text"`
	TargetLang string `json:"targetLang"`
}

type TranslateResponse struct {
	Translations []struct {
		TranslatedText string `json:"translatedText"`
	} `json:"translations"`
}

func Translate(c echo.Context) error {
	var req TranslateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
	}

	requestBody := map[string]interface{}{
		"contents":           []string{req.Text},
		"targetLanguageCode": req.TargetLang,
	}

	requestData, err := json.Marshal(map[string]interface{}{
		"parent":             fmt.Sprintf("projects/%s/locations/global", os.Getenv("GOOGLE_PROJECT_ID")),
		"contents":           requestBody["contents"],
		"targetLanguageCode": requestBody["targetLanguageCode"],
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to marshal request data",
			"error":   err.Error(),
		})
	}

	url := "https://translate.googleapis.com/v3beta1/projects/locations/global:translateText"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to call Google Translate API",
			"error":   err.Error(),
		})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to read response body",
			"error":   err.Error(),
		})
	}

	var translateResp TranslateResponse
	if err := json.Unmarshal(body, &translateResp); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to unmarshal response data",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, translateResp)
}
