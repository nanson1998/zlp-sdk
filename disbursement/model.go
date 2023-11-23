package disbursement

import (
	"github.com/google/go-querystring/query"
	"net/url"
)

type QueryUserRequest struct {
	AppID  int    `json:"app_id" url:"app_id"`
	Phone  string `json:"phone" url:"phone"`
	Time   int64  `json:"time" url:"time"`
	MacKey string `json:"mac"  url:"mac"`
}

type QueryUserResponse struct {
	BaseResponse
	Data UserData `json:"data"`
}

type TopupRequest struct {
	AppID            int    `json:"app_id" url:"app_id"`
	PartnerOrderId   string `json:"partner_order_id" url:"partner_order_id"`
	PaymentId        string `json:"payment_id" url:"payment_id"`
	MuID             string `json:"m_u_id"  url:"m_u_id"`
	PartnerEmbedData string `json:"partner_embed_data"  url:"partner_embed_data"`
	Amount           int64  `json:"amount"  url:"amount"`
	Description      string `json:"description"  url:"description"`
	ReferenceId      string `json:"reference_id"  url:"reference_id"`
	ExtraInfo        string `json:"extra_info"  url:"extra_info"`
	Time             int64  `json:"time"  url:"time"`
	Sig              string `json:"sig"  url:"sig"`
	MacKey           string `json:"mac"  url:"mac"`
	RSAPrivateKey    string `json:"rsa_private_key" url:"rsa_private_key"`
}

type TopupResponse struct {
	BaseResponse
	Data TopupData `json:"data"`
}

type QueryOrderRequest struct {
	AppID          int    `json:"app_id" url:"app_id"`
	PartnerOrderId string `json:"partner_order_id" url:"partner_order_id"`
	Time           int64  `json:"time"  url:"time"`
	MacKey         string `json:"mac"  url:"mac"`
}

type QueryOrderResponse struct {
	BaseResponse
	Data OrderData `json:"data"`
}

type BalanceRequest struct {
	AppID     int    `json:"app_id" url:"app_id"`
	PaymentId string `json:"payment_id" url:"payment_id"`
	Time      int64  `json:"time"  url:"time"`
	MacKey    string `json:"mac"  url:"mac"`
}

type BalanceResponse struct {
	BaseResponse
	Data BalanceData `json:"data"`
}

type UserData struct {
	ReferenceId   string `json:"reference_id"`
	MUId          string `json:"m_u_id"`
	Name          string `json:"name"`
	Phone         string `json:"phone" `
	OnboardingUrl string `json:"onboarding_url" `
}

type TopupData struct {
	OrderID     string `json:"order_id"`
	Status      int    `json:"status"`
	MUID        string `json:"m_u_id"`
	Phone       string `json:"phone"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	PartnerFee  int64  `json:"partner_fee"`
	ZLPFee      int64  `json:"zlp_fee"`
	ExtraInfo   string `json:"extra_info"`
	Time        int64  `json:"time"`
	UpgradeUrl  string `json:"upgrade_url"`
}

type OrderData struct {
	OrderID     string `json:"order_id"`
	Status      int    `json:"status"`
	MUID        string `json:"m_u_id"`
	Phone       string `json:"phone"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	PartnerFee  int64  `json:"partner_fee"`
	ZLPFee      int64  `json:"zlp_fee"`
	ExtraInfo   string `json:"extra_info"`
	Time        int64  `json:"time"`
	ZlpTransId  string `json:"zp_trans_id"`
	ResultUrl   string `json:"result_url"`
}

type BalanceData struct {
	Balance int64 `json:"balance"`
}

type BaseResponse struct {
	ReturnCode       int32  `json:"return_code"`
	ReturnMessage    string `json:"return_message"`
	SubReturnCode    int32  `json:"sub_return_code"`
	SubReturnMessage string `json:"sub_return_message"`
}

// ToValues convert to url.Values
func (cor *QueryUserRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}

// ToValues convert to url.Values
func (cor *TopupRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}

func (cor *QueryOrderRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}

func (cor *BalanceRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}
