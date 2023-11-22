package refund

import (
	"encoding/json"
	"github.com/nanson1998/zlp-sdk/helper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var endpoint = helper.SB

func RefundOrder(req *CreateRefundRequest, isProduction bool) (*CreateRefundResponse, error) {
	t := helper.GetAppTime()
	prefixDate := helper.GetTimeString(t)
	prefixMRefundId := prefixDate + "_" + strconv.Itoa(req.AppID)
	l := 7 + len(strconv.Itoa(req.AppID))
	if req.MRefundID[:l] != prefixMRefundId {
		return &CreateRefundResponse{
			ReturnCode:       2,
			SubReturnCode:    -401,
			ReturnMessage:    "Giao dịch thất bại",
			SubReturnMessage: "m_refund_id phải bắt đầu bằng:" + prefixMRefundId,
		}, nil
	}
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.ZPTransID, req.Amount, req.Description, t)
	req.MacKey = mac
	req.Timestamp = t

	createReq, err := req.ToValues()
	if err != nil {
		return nil, err
	}
	if isProduction == true {
		endpoint = helper.PROD
	}
	res, err := http.PostForm(endpoint+"/v2/refund", createReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var response CreateRefundResponse

	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}
	return &response, nil
}

func QueryRefund(req *QueryRefundRequest, isProduction bool) (*QueryRefundResponse, error) {
	t := helper.GetAppTime()
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.MRefundID, t)
	r := &QueryRefundRequest{
		AppID:     req.AppID,
		MRefundID: req.MRefundID,
		MacKey:    mac,
		Timestamp: t,
	}

	queryReq, err := r.ToValues()
	if err != nil {
		return nil, err
	}
	if isProduction == true {
		endpoint = helper.PROD
	}
	res, err := http.PostForm(endpoint+"/v2/query_refund", queryReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var queryRes QueryRefundResponse
	if err := json.Unmarshal(body, &queryRes); err != nil {
		log.Fatal(err)
	}
	return &queryRes, nil
}
