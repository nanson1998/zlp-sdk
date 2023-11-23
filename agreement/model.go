package agreement

import (
	"github.com/google/go-querystring/query"
	"net/url"
)

type CreateBindingRequest struct {
	AppID            int    `json:"app_id" url:"app_id"`
	AppTransID       string `json:"app_trans_id" url:"app_trans_id"`
	Identifier       string `json:"identifier" url:"identifier"`
	BindingType      string `json:"binding_type" url:"binding_type"`
	BindingData      string `json:"binding_data" url:"binding_data"`
	MaxAmount        int    `json:"max_amount" url:"max_amount"`
	RedirectURL      string `json:"redirect_url" url:"redirect_url"`
	RedirectDeepLink string `json:"redirect_deep_link" url:"redirect_deep_link"`
	CallbackURL      string `json:"callback_url" url:"callback_url"`
	ReqDate          int64  `json:"req_date"  url:"req_date"`
	MacKey           string `json:"mac"  url:"mac"`
}

type CreateBindingResponse struct {
	BaseResponse
	BindingToken  string `json:"binding_token"`
	DeepLink      string `json:"deep_link"`
	BindingQRLink string `json:"binding_qr_link"`
	ShortLink     string `json:"short_link"`
}

type QuerybindingRequest struct {
	AppID       int    `json:"app_id" url:"app_id"`
	AppTransID  string `json:"app_trans_id" url:"app_trans_id"`
	BindingType string `json:"binding_type" url:"binding_type"`
	ReqDate     int64  `json:"req_date"  url:"req_date"`
	MacKey      string `json:"mac"  url:"mac"`
}

type QueryBindingResponse struct {
	BaseResponse
	Data Data `json:"data"`
}

type QueryBalanceRequest struct {
	AppID      int    `json:"app_id" url:"app_id"`
	Identifier string `json:"identifier" url:"identifier"`
	PayToken   string `json:"pay_token" url:"pay_token"`
	Amount     int    `json:"amount" url:"amount"`
	ReqDate    int64  `json:"req_date"  url:"req_date"`
	MacKey     string `json:"mac"  url:"mac"`
}

type QueryBalanceResponse struct {
	BaseResponse
	Data           DataBalance `json:"data"`
	DiscountAmount int64       `json:"discount_amount"`
	ReformUrl      string      `json:"reform_url"`
}

type PayByTokenRequest struct {
	AppID        int    `json:"app_id" url:"app_id"`
	Identifier   string `json:"identifier" url:"identifier"`
	ZpTransToken string `json:"zp_trans_token" url:"zp_trans_token"`
	PayToken     string `json:"pay_token" url:"pay_token"`
	BindingType  string `json:"binding_type" url:"binding_type"`
	RedirectUrl  int    `json:"redirect_url" url:"redirect_url"`
	ReqDate      int64  `json:"req_date"  url:"req_date"`
	MacKey       string `json:"mac"  url:"mac"`
}

type PayByTokenResponse struct {
	BaseResponse
	Amount    int64  `json:"amount"`
	ReformUrl string `json:"reform_url"`
	ZpTransId string `json:"zp_trans_id	"`
}

type UnbindRequest struct {
	AppID      int    `json:"app_id" url:"app_id"`
	Identifier string `json:"identifier" url:"identifier"`
	BindingId  string `json:"binding_id" url:"binding_id"`
	ReqDate    int64  `json:"req_date"  url:"req_date"`
	MacKey     string `json:"mac"  url:"mac"`
}

type UnbindResponse struct {
	BaseResponse
}

type Data struct {
	AppID          int    `json:"app_id"`
	AppTransID     string `json:"app_trans_id"`
	BindingId      string `json:"binding_id" `
	BindingData    string `json:"binding_data" `
	PayToken       string `json:"pay_token" `
	ServerTime     int    `json:"server_time"`
	MerchantUserId int    `json:"merchant_user_id"`
	Status         int    `json:"status"`
	MsgType        int    `json:"msg_type"`
	ZpUserId       string `json:"zp_user_id" `
}

type DataBalance struct {
	Channel  int    `json:"channel"`
	Payable  bool   `json:"payable"`
	BankCode string `json:"bank_code" `
}

type BaseResponse struct {
	ReturnCode       int32  `json:"return_code"`
	ReturnMessage    string `json:"return_message"`
	SubReturnCode    int32  `json:"sub_return_code"`
	SubReturnMessage string `json:"sub_return_message"`

	err error
}

// ToValues convert to url.Values
func (cor *CreateBindingRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}

// ToValues convert to url.Values
func (cor *QuerybindingRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}

func (cor *QueryBalanceRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}

func (cor *PayByTokenRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}

func (cor *UnbindRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}
