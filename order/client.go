package order

import (
	"encoding/json"
	"github.com/nanson1998/zlp-sdk/helper"
	"io/ioutil"
	"log"
	"net/http"
)

var endpoint = helper.SB

func CreateOrder(req *CreateOrderRequest, isProduction bool) (*CreateOrderResponse, error) {
	t := helper.GetAppTime()
	prefixDate := helper.GetTimeString(t)
	if req.AppTransID[:6] != prefixDate {
		return &CreateOrderResponse{
			ReturnCode:       2,
			SubReturnCode:    -92,
			ReturnMessage:    "Giao dịch thất bại",
			SubReturnMessage: "app_trans_id phải bắt đầu bằng:" + prefixDate,
		}, nil
	}

	if req.Item == "" {
		req.Item = "[]"
	}
	if req.EmbedData == "" {
		req.EmbedData = "{}"
	}
	// calculate mac
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.AppTransID, req.AppUser, req.Amount, t, req.EmbedData, req.Item)
	req.MacKey = mac
	req.AppTime = t

	createReq, err := req.ToValues()
	if err != nil {
		return nil, err
	}
	if isProduction == true {
		endpoint = helper.PROD
	}
	res, err := http.PostForm(endpoint+"/v2/create", createReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var response CreateOrderResponse

	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}

	return &response, nil
}

func QueryOrder(req *QueryStatusRequest, isProduction bool) (*QueryStatusResponse, error) {
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.AppTransID, req.MacKey)
	req.MacKey = mac
	queryReq, err := req.ToValues()
	if err != nil {
		return nil, err
	}
	if isProduction == true {
		endpoint = helper.PROD
	}
	res, err := http.PostForm(endpoint+"/v2/query", queryReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var queryRes QueryStatusResponse
	if err := json.Unmarshal(body, &queryRes); err != nil {
		log.Fatal(err)
	}
	return &queryRes, nil
}
