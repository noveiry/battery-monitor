package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"battery-monitor/pkg/models"
)

type Client struct {
	http *http.Client

	telemetryURL string
	alertURL     string
}

func NewClient(
	telemetryURL string,
	alertURL string,
) *Client {

	return &Client{
		http: &http.Client{
			Timeout: 5 * time.Second,
		},
		telemetryURL: telemetryURL,
		alertURL:     alertURL,
	}
}

func (c *Client) Batteries() (
	[]models.BatteryTelemetry,
	error,
) {

	resp, err := c.http.Get(
		fmt.Sprintf(
			"%s/batteries",
			c.telemetryURL,
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result []models.BatteryTelemetry

	err = json.NewDecoder(
		resp.Body,
	).Decode(&result)

	return result, err
}

func (c *Client) Stats() (
	map[string]any,
	error,
) {

	resp, err := c.http.Get(
		fmt.Sprintf(
			"%s/stats",
			c.telemetryURL,
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result map[string]any

	err = json.NewDecoder(
		resp.Body,
	).Decode(&result)

	return result, err
}

func (c *Client) Alerts() (
	[]models.BatteryAlert,
	error,
) {

	resp, err := c.http.Get(
		fmt.Sprintf(
			"%s/alerts",
			c.alertURL,
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result []models.BatteryAlert

	err = json.NewDecoder(
		resp.Body,
	).Decode(&result)

	return result, err
}
