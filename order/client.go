package order

import (
	"encoding/json"
	helper2 "github.com/nanson1998/zlp-sdk/helper"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateOrder(req *CreateOrderRequest) (*CreateOrderResponse, error) {
	t := helper2.GetAppTime()
	prefixDate := helper2.GetTimeString(t)
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
	mac := helper2.BuildMAC(req.MacKey, "|", req.AppID, req.AppTransID, req.AppUser, req.Amount, t, req.EmbedData, req.Item)
	req.MacKey = mac
	req.AppTime = t

	createReq, err := req.ToValues()
	if err != nil {
		return nil, err
	}
	res, err := http.PostForm("https://sb-openapi.zalopay.vn/v2/create", createReq)
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

func QueryOrder(req *QueryStatusRequest) (*QueryStatusResponse, error) {
	mac := helper2.BuildMAC(req.MacKey, "|", req.AppID, req.AppTransID, req.MacKey)

	r := &QueryStatusRequest{
		AppID:      req.AppID,
		AppTransID: req.AppTransID,
		MacKey:     mac,
	}

	queryReq, err := r.ToValues()
	if err != nil {
		return nil, err
	}
	res, err := http.PostForm("https://sb-openapi.zalopay.vn/v2/query", queryReq)
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
