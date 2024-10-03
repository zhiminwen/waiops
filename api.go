package waiops

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type API struct {
	BaseUrl string

	tenantID string
	zenToken string
}

func NewAPI(baseUrl, apiUser, apiKey string) *API {
	zenKey := fmt.Sprintf("%s:%s", apiUser, apiKey)
	zenToken := base64.StdEncoding.EncodeToString([]byte(zenKey))
	return &API{
		zenToken: zenToken,
		tenantID: "cfd95b7e-3bc7-4006-a4a8-a73a79c71255", //fixed as of now
		BaseUrl:  baseUrl,
	}
}

func (a *API) CreateRequest() *resty.Request {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	request := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-TenantID", a.tenantID).
		SetHeader("Authorization", "ZenApiKey "+a.zenToken).
		SetHeader("Accept", "application/json")

	return request
}

func (a *API) CallAPI(uri, method string, payloads ...any) (*resty.Response, error) {
	request := a.CreateRequest()

	var resp *resty.Response
	var err error

	switch method {
	case "POST":
		resp, err = request.SetBody(payloads[0]).Post(a.BaseUrl + uri)
	case "GET":
		resp, err = request.Get(a.BaseUrl + uri)
	case "PATCH":
		resp, err = request.SetBody(payloads[0]).Patch(a.BaseUrl + uri)
	default:
		return nil, fmt.Errorf("unsupported method: %s", method)
	}
	if err != nil {
		return resp, err
	}

	if resp.StatusCode() >= 300 {
		return resp, fmt.Errorf("non 200 status")
	}

	return resp, nil
}
