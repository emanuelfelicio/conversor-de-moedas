package exchange

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/emanuelfelicio/conversor-de-moedas/internal/domain"
)

type ExchangeRateResponse struct {
	Data map[string]float64 `json:"data"`
}

func Rate(apiURL string, apiKey string, input *domain.ConversionInput) (float64, error) {

	baseUrl, err := url.Parse(apiURL)
	if err != nil {
		return 0, fmt.Errorf("url:%v is not correct", baseUrl)
	}

	queryParams := baseUrl.Query()
	queryParams.Add("base_currency", string(input.From()))
	queryParams.Add("currencies", string(input.To()))
	baseUrl.RawQuery = queryParams.Encode()

	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("erro in request")
	}
	req.Header.Set("apikey", apiKey)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return 0, fmt.Errorf("erro in request HTTP: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("erro in request status %d", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		return 0, fmt.Errorf("erro in ready body response %v", err)
	}

	var exchangeResp ExchangeRateResponse
	err = json.Unmarshal(bodyBytes, &exchangeResp)

	if err != nil {
		return 0, fmt.Errorf("error encoding response in json %v", err)

	}

	rate := exchangeResp.Data[string(input.To())]

	return rate, nil

}
