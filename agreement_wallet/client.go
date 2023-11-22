package agreement_wallet

import (
	"encoding/json"
	"github.com/nanson1998/zlp-sdk/helper"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateBinding(req *CreateBindingRequest) (*CreateBindingResponse, error) {
	t := helper.GetAppTime()
	prefixDate := helper.GetTimeString(t)
	if req.AppTransID[:6] != prefixDate {
		return &CreateBindingResponse{
			BaseResponse: BaseResponse{
				ReturnCode:       2,
				SubReturnCode:    -92,
				ReturnMessage:    "Giao dịch thất bại",
				SubReturnMessage: "app_trans_id phải bắt đầu bằng:" + prefixDate,
			},
		}, nil
	}
	if req.BindingData == "" {
		req.BindingData = "{}"
	}
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.AppTransID, req.BindingData, req.BindingType, req.Identifier, req.MaxAmount, t)
	req.MacKey = mac
	req.ReqDate = t
	createReq, err := req.ToValues()
	if err != nil {
		return nil, err
	}
	res, err := http.PostForm("https://sb-openapi.zalopay.vn/v2/agreement/bind", createReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var createRes CreateBindingResponse
	if err := json.Unmarshal(body, &createRes); err != nil {
		log.Fatal(err)
	}
	return &createRes, nil

}

func QueryBinding(req *QuerybindingRequest) (*QueryBindingResponse, error) {
	t := helper.GetAppTime()
	// calculate mac
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.AppTransID, t)
	reqQ := &QuerybindingRequest{
		AppID:      req.AppID,
		AppTransID: req.AppTransID,
		ReqDate:    t,
		MacKey:     mac,
	}
	queryReq, err := reqQ.ToValues()
	if err != nil {
		return nil, err
	}
	res, err := http.PostForm("https://sb-openapi.zalopay.vn/v2/agreement/query", queryReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var createRes QueryBindingResponse
	if err := json.Unmarshal(body, &createRes); err != nil {
		log.Fatal(err)
	}
	return &createRes, nil
}

func QueryBalance(req *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	t := helper.GetAppTime()
	// calculate mac
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.PayToken, req.Identifier, req.Amount, t)
	reqQ := &QueryBalanceRequest{
		AppID:      req.AppID,
		PayToken:   req.PayToken,
		ReqDate:    t,
		Identifier: req.Identifier,
		Amount:     req.Amount,
		MacKey:     mac,
	}
	queryReq, err := reqQ.ToValues()
	if err != nil {
		return nil, err
	}
	res, err := http.PostForm("https://sb-openapi.zalopay.vn/v2/agreement/balance", queryReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var createRes QueryBalanceResponse
	if err := json.Unmarshal(body, &createRes); err != nil {
		log.Fatal(err)
	}
	return &createRes, nil
}

func PayByToken(req *PayByTokenRequest) (*PayByTokenResponse, error) {
	t := helper.GetAppTime()
	// calculate mac
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.Identifier, req.ZpTransToken, req.PayToken, t)
	reqQ := &PayByTokenRequest{
		AppID:        req.AppID,
		PayToken:     req.PayToken,
		ZpTransToken: req.ZpTransToken,
		ReqDate:      t,
		Identifier:   req.Identifier,
		RedirectUrl:  req.RedirectUrl,
		MacKey:       mac,
	}
	queryReq, err := reqQ.ToValues()
	if err != nil {
		return nil, err
	}
	res, err := http.PostForm("https://sb-openapi.zalopay.vn/v2/agreement/pay", queryReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var createRes PayByTokenResponse
	if err := json.Unmarshal(body, &createRes); err != nil {
		log.Fatal(err)
	}
	return &createRes, nil
}
