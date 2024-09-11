package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TranslateRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

type TranslateResponse struct {
	Trans              string `json:"trans"`
	SourceLanguageCode string `json:"source_language_code"`
	SourceLanguage     string `json:"source_language"`
	TrustLevel         int    `json:"trust_level"`
}

func Translate(c echo.Context) error {
	var req TranslateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
	}

	payload := map[string]string{
		"from": req.From,
		"to":   req.To,
		"text": req.Text,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to marshal JSON",
			"error":   err.Error(),
		})
	}

	url := "https://google-translate113.p.rapidapi.com/api/v1/translator/text"
	reqBody := bytes.NewReader(payloadBytes)
	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create request",
			"error":   err.Error(),
		})
	}

	httpReq.Header.Add("x-rapidapi-key", "cfe1a66caamsh85d61d2b9b14717p109f9bjsneb6dd9c4fe71")
	httpReq.Header.Add("x-rapidapi-host", "google-translate113.p.rapidapi.com")
	httpReq.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to execute request",
			"error":   err.Error(),
		})
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to read response",
			"error":   err.Error(),
		})
	}

	var translateResponse TranslateResponse
	err = json.Unmarshal(body, &translateResponse)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to parse response",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "successfully translated language",
		"data":    translateResponse.Trans,
	})
}
