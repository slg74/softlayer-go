package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TheWeatherCompany/softlayer-go/common"
	"github.com/TheWeatherCompany/softlayer-go/data_types"
	"github.com/TheWeatherCompany/softlayer-go/softlayer"
	"log"
)

type softlayer_Scale_Group_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Scale_Group_Service(client softlayer.Client) *softlayer_Scale_Group_Service {
	return &softlayer_Scale_Group_Service{
		client: client,
	}
}

func (slsgs *softlayer_Scale_Group_Service) GetName() string {
	return "SoftLayer_Scale_Group"
}

func (slsgs *softlayer_Scale_Group_Service) CreateObject(template data_types.SoftLayer_Scale_Group) (data_types.SoftLayer_Scale_Group, error) {
	parameters := data_types.SoftLayer_Scale_Group_Parameters{
		Parameters: []interface{}{
			template,
		},
	}

	requestBody, err := json.Marshal(parameters)
	log.Printf("[INFO]  ***** request body: %s", requestBody)
	if err != nil {
		return data_types.SoftLayer_Scale_Group{}, err
	}

	data, errorCode, err := slsgs.client.GetHttpClient().DoRawHttpRequest(fmt.Sprintf("%s/createObject", slsgs.GetName()), "POST", bytes.NewBuffer(requestBody))
	if err != nil {
		return data_types.SoftLayer_Scale_Group{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Scale_Group#createObject\nResponse from SoftLayer: %s\nHTTP error code: '%d'", data, errorCode)
		return data_types.SoftLayer_Scale_Group{}, errors.New(errorMessage)
	}

	err = slsgs.client.GetHttpClient().CheckForHttpResponseErrors(data)
	if err != nil {
		return data_types.SoftLayer_Scale_Group{}, err
	}

	softLayer_Scale_Group := data_types.SoftLayer_Scale_Group{}
	err = json.Unmarshal(data, &softLayer_Scale_Group)
	if err != nil {
		return data_types.SoftLayer_Scale_Group{}, err
	}

	return softLayer_Scale_Group, nil
}

func (slsgs *softlayer_Scale_Group_Service) ForceDeleteObject(group int) (bool, error) {
	response, errorCode, err := slsgs.client.GetHttpClient().DoRawHttpRequest(fmt.Sprintf("%s/%d/forceDeleteObject", slsgs.GetName(), group), "GET", new(bytes.Buffer))
	if err != nil {
		return false, err
	}

	if res := string(response[:]); res != "true" {
		return false, fmt.Errorf("Failed to force delete scale group with id '%d', got '%s' as response from the API.", group, res)
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Scale_Group#forceDeleteObject, HTTP error code: '%d'", errorCode)
		return false, errors.New(errorMessage)
	}

	return true, err
}
