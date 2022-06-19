package gasapi

import (
	"fmt"
	"strconv"

	"encoding/json"

	"gopkg.in/h2non/gentleman.v2"
)

var apiUrl string = "https://api.etherscan.io/api?module=gastracker&action=gasoracle"

type gasApi struct {
	client *gentleman.Client
}

type jsonRes struct {
	Status string
	// baseFee float32 `json:"result.suggestedBaseFee"`
	Message string
	Result  json.RawMessage
}

type resultString string
type resultStruct struct {
	BaseFee string `json:"suggestBaseFee"`
}

func CreateGasApi() *gasApi {
	client := gentleman.New()
	client.URL(apiUrl)

	return &gasApi{
		client: client,
	}
}

func (api *gasApi) Get() (float32, error) {
	request := api.client.
		Request()

	res, err := request.Send()

	if err != nil {
		return 0, err
	}

	// fmt.Printf("%#v\n", res.String())
	jsonRes := jsonRes{}
	err = res.JSON(&jsonRes)
	var result interface{}

	if err != nil {
		return 0, err
	}
	// fmt.Printf("%#v\n", jsonRes)
	if jsonRes.Status == "0" {
		result = new(resultString)
	} else {
		result = new(resultStruct)
	}

	err = json.Unmarshal(jsonRes.Result, &result)
	if err != nil {
		return 0, err
	}

	switch typedRes := result.(type) {
	case *resultString:
		return 0, fmt.Errorf("api error: %s", *typedRes)
	case *resultStruct:
		baseFee, err := strconv.ParseFloat(typedRes.BaseFee, 32)
		if err != nil {
			return 0, nil
		}
		return float32(baseFee), nil
	}

	return 0, nil
}
