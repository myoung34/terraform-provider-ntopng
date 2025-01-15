package ntopng

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"time"
)

var validate = validator.New()

type Config struct {
	Host  string
	Token string
}
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Config     Config
}

func ValidateConfig(config Config) (bool, error) {
	configKeys := []string{config.Host, config.Token}
	for _, configVal := range configKeys {
		if len(configVal) == 0 {
			return false, errors.New("required config value not found")
		}
	}
	return true, nil
}

func NewClient(config Config) (*Client, error) {
	_, err := ValidateConfig(config) //nolint:errcheck
	if err != nil {
		return nil, err
	}

	c := Client{
		HTTPClient: &http.Client{},
		HostURL:    fmt.Sprintf("%s/lua/rest/v2", config.Host),
	}

	c.Config = config

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	//TODO Support pagination?
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Millisecond*60000),
	)

	defer cancel()
	req = req.WithContext(ctx)

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.Config.Token))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
