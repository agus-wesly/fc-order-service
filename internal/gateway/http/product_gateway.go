package httpgateway

import (
	"order-service/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProductGateway struct {
	baseURL string
}

func NewProductGateway(baseURL string) *ProductGateway {
	return &ProductGateway{baseURL: baseURL}
}


func (c *ProductGateway) GetProductInfo(productId string) (*model.ProductResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/products/%s", c.baseURL, productId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result model.ProductResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
