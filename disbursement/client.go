package disbursement

import (
	"encoding/json"
	"github.com/nanson1998/zlp-sdk/helper"
	"github.com/nanson1998/zlp-sdk/helper/rsautil"
	"io/ioutil"
	"log"
	"net/http"
)

var endpoint = helper.SB

func QueryUser(req *QueryUserRequest, isProduction bool) (*QueryUserResponse, error) {
	t := helper.GetAppTime()
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.Phone, t)
	req.MacKey = mac
	req.Time = t
	createReq, err := req.ToValues()
	if err != nil {
		return nil, err
	}
	if isProduction == true {
		endpoint = helper.PROD
	}
	res, err := http.PostForm(endpoint+"/v2/disbursement/user", createReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var createRes QueryUserResponse
	if err := json.Unmarshal(body, &createRes); err != nil {
		log.Fatal(err)
	}
	return &createRes, nil

}

func Topup(req *TopupRequest, isProduction bool) (*TopupResponse, error) {
	t := helper.GetAppTime()
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.PaymentId, req.PartnerOrderId, req.MuID, req.Amount, req.Description, req.PartnerEmbedData, req.ExtraInfo, t)
	sig, err := rsautil.BuildSign(mac, req.RSAPrivateKey)
	req.Time = t
	req.Sig = sig
	queryReq, err := req.ToValues()
	if err != nil {
		return nil, err
	}
	if isProduction == true {
		endpoint = helper.PROD
	}
	res, err := http.PostForm(endpoint+"/v2/disbursement/topup", queryReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
	var createRes TopupResponse
	if err := json.Unmarshal(body, &createRes); err != nil {
		log.Fatal(err)
	}
	return &createRes, nil
}

func QueryOrder(req *QueryOrderRequest, isProduction bool) (*QueryOrderResponse, error) {
	t := helper.GetAppTime()
	// calculate mac
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.PartnerOrderId, t)
	req.MacKey = mac
	req.Time = t
	queryReq, err := req.ToValues()
	if err != nil {
		return nil, err
	}
	if isProduction == true {
		endpoint = helper.PROD
	}
	res, err := http.PostForm(endpoint+"/v2/disbursement/txn", queryReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var createRes QueryOrderResponse
	if err := json.Unmarshal(body, &createRes); err != nil {
		log.Fatal(err)
	}
	return &createRes, nil
}

func Balance(req *BalanceRequest, isProduction bool) (*BalanceResponse, error) {
	t := helper.GetAppTime()
	// calculate mac
	mac := helper.BuildMAC(req.MacKey, "|", req.AppID, req.PaymentId, t)
	req.MacKey = mac
	req.Time = t
	queryReq, err := req.ToValues()
	if err != nil {
		return nil, err
	}
	if isProduction == true {
		endpoint = helper.PROD
	}
	res, err := http.PostForm(endpoint+"/v2/disbursement/balance", queryReq)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
	var createRes BalanceResponse
	if err := json.Unmarshal(body, &createRes); err != nil {
		log.Fatal(err)
	}
	return &createRes, nil
}
