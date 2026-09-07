package course

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	c "github.com/ncostamagna/go_http_client/client"

	"github.com/dasilro/gocourse_domain/domain"
)

type (
	DataResponse struct {
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
	}

	Transport interface {
		Get(id string) (*domain.Course, error)
	}

	clientHTTP struct {
		client c.Transport
	}
)

func NewHttpClient(baseURL, token string) Transport {
	header := http.Header{}

	if token != "" {
		header.Set("Authorization", token)
	}

	return &clientHTTP{
		client: c.New(header, baseURL, 5000*time.Millisecond, true),
	}
}

func (c *clientHTTP) Get(id string) (*domain.Course, error) {
	dataResponse := DataResponse{Data: &domain.Course{}}

	u := url.URL{}
	u.Path += fmt.Sprintf("/courses/%s", id)

	resp := c.client.Get(u.String())

	if resp.Err != nil {
		return nil, resp.Err
	}

	if err := resp.FillUp(&dataResponse); err != nil {
		return nil, fmt.Errorf("%s", resp)
	}

	if resp.StatusCode == 404 {
		return nil, ErrNotFound{fmt.Sprintf("%s", dataResponse.Message)}
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("%s", dataResponse.Message)
	}

	return dataResponse.Data.(*domain.Course), nil
}
