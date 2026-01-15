package meal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/JoePeach762/PP_project/internal/models"
	offutils "github.com/JoePeach762/PP_project/internal/services/meal/off_utils"
)

type HTTPClient struct {
	client    *http.Client
	userAgent string
}

func NewHTTPClient(userAgent string) *HTTPClient {
	transport := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	return &HTTPClient{
		client:    client,
		userAgent: userAgent,
	}
}

func (c *HTTPClient) FetchProduct(ctx context.Context, name string) (*models.MealTemplate, error) {
	baseURL := "https://world.openfoodfacts.org/cgi/search.pl"
	params := url.Values{}
	params.Add("search_terms", name)
	params.Add("search_simple", "1")
	params.Add("action", "process")
	params.Add("json", "1")
	params.Add("page_size", "1")
	params.Add("fields", "product_name,nutriments")

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OFF returned %d", resp.StatusCode)
	}

	var result offutils.OFFSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Products) == 0 {
		return nil, errors.New("продукт не найден")
	}

	p := result.Products[0]
	n := p.Nutriments

	if n.EnergyKcal == 0 && n.Proteins == 0 && n.Fat == 0 && n.Carbohydrates == 0 {
		return nil, errors.New("нет данных о нутриентах")
	}

	return &models.MealTemplate{
		Name:         p.ProductName,
		Calories100g: n.EnergyKcal,
		Proteins100g: n.Proteins,
		Fats100g:     n.Fat,
		Carbs100g:    n.Carbohydrates,
	}, nil
}
