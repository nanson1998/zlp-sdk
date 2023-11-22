package order

import (
	"encoding/json"
	"github.com/nanson1998/zlp-sdk/helper"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateOrder(req *CreateOrderRequest) (*CreateOrderResponse, error) {
	t := helper.GetAppTime()
	prefixDate := helper.GetTimeString(t)
	log.Printf(req.AppTransID[:6])
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
	mac := helper.BuildMAC(req.Mac, "|", req.AppID, req.AppTransID, req.AppUser, req.Amount, t, req.EmbedData, req.Item)
	req.Mac = mac
	req.AppTime = t

	createReq, err := req.ToValues()
	if err != nil {
		return nil, err
	}
	log.Println("Req:", createReq)
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
