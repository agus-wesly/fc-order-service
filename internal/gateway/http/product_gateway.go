package httpgateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/internal/model"

	"github.com/gofiber/fiber/v2"
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

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 404 {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Product not found")
		}

		return nil, fiber.ErrInternalServerError
	}

	var result model.ProductResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
