package gasapi

import (
	"gopkg.in/h2non/gentleman.v2"
)

var apiUrl string = "https://owlracle.info"

type gasApi struct {
	client    *gentleman.Client
}

type JsonRes struct {
	avgTime float32
	baseFee float32
}

func CreateGasApi() *gasApi {
	client := gentleman.New()
	client.URL(apiUrl)

	return &gasApi{
		client:    client,
	}
}

func (api *gasApi) Get() (float32, float32, error) {
	request := api.client.
		Request().
		AddPath("/eth/gas")

	res, err := request.Send()

	if err != nil {
		return 0, 0, err
	}


  json := JsonRes{}
  err = res.JSON(&json)

  if err != nil {
    return 0, 0, err
  }

	return json.baseFee, json.avgTime, nil
}
