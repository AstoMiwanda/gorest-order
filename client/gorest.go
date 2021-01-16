package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	ph "github.com/kitabisa/perkakas/v2/httpclient"
	"github.com/pushm0v/gorest-order/model"
)

type GorestClient interface {
	GetCustomerByID(id int) (cust *model.Customer, err error)
}

type gorestClient struct {
	httpClient *ph.HttpClient
	gorestUrl  string
}

func NewGorestClient(gorestUrl string) GorestClient {
	conf := new(ph.HttpClientConf)
	conf.BackoffInterval = 2 * time.Millisecond       // 2ms
	conf.MaximumJitterInterval = 5 * time.Millisecond // 5ms
	conf.Timeout = 15000 * time.Millisecond           // 15s
	conf.RetryCount = 3                               // 3 times

	phClient := ph.NewHttpClient(conf)

	return &gorestClient{
		httpClient: phClient,
		gorestUrl:  gorestUrl,
	}
}

func (c *gorestClient) GetCustomerByID(id int) (cust *model.Customer, err error) {
	resp, err := c.httpClient.Client.Get(fmt.Sprintf("%s/api/v1/customers/%d", c.gorestUrl, id), http.Header{})
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return cust, fmt.Errorf("Error from gorest service : [%d] %v", resp.StatusCode, string(body))
	}

	cust = &model.Customer{}
	err = json.NewDecoder(resp.Body).Decode(cust)

	return
}
