package zod

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nanson1998/zlp-sdk/helper"
	"io/ioutil"
	"net/http"
)

var endpoint = SB

func CreateInvoice(req *CreateInvoiceRequest, isProduction bool) (*CreateInvoiceResponse, error) {
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.MCRefID, req.Amount, req.MCExtInfo)
	req.MacKey = mac
	postData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if isProduction == true {
		endpoint = PROD
	}
	res, err := http.Post(endpoint+"/v2/zod", "appication/json", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response CreateInvoiceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err

	}
	return &response, nil
}

func QueryInvoice(req *QueryInvoiceRequest, isProduction bool) (*QueryInvoiceResponse, error) {
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.MCRefID)
	req.MacKey = mac
	postData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if isProduction == true {
		endpoint = PROD
	}
	url := fmt.Sprintf("%s/v2/zod/invoice?appId=%v&mcRefId=%v&mac=%v", endpoint, req.AppID, req.MCRefID, mac)
	r, err := http.NewRequest("GET", url, bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err

	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var queryRes QueryInvoiceResponse
	if err := json.Unmarshal(body, &queryRes); err != nil {
		return nil, err
	}
	return &queryRes, nil
}

func QueryOrderStatus(req *QueryOrderStatusRequest, isProduction bool) (*QueryOrderStatusResponse, error) {
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.MCRefID)
	req.MacKey = mac
	postData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if isProduction == true {
		endpoint = PROD
	}
	url := fmt.Sprintf("%s/v2/zod/status?appId=%v&mcRefId=%v&mac=%v", endpoint, req.AppID, req.MCRefID, mac)
	r, err := http.NewRequest("GET", url, bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err

	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var queryRes QueryOrderStatusResponse
	if err := json.Unmarshal(body, &queryRes); err != nil {
		return nil, err
	}
	return &queryRes, nil
}
