package meal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/JoePeach762/PP_project/internal/models"
)

type HTTPClient struct {
	client *http.Client
}

func (c *HTTPClient) FetchProduct(ctx context.Context, name string) (*models.OFFProduct, error) {
	// Поиск по barcode или названию — упрощённо:
	url := fmt.Sprintf("https://world.openfoodfacts.net/cgi/search.pl?search_terms=%s&json=1", url.QueryEscape(name))
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result OFFSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Products) == 0 {
		return nil, errors.New("продукт не найден")
	}

	p := result.Products[0]
	return &OFFProduct{
		Name:         p.ProductName,
		Calories100g: float32(p.EnergyKcal),
		Proteins100g: float32(p.Proteins),
		Fats100g:     float32(p.Fat),
		Carbs100g:    float32(p.Carbohydrates),
	}, nil
}
