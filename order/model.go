package order

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

type CreateOrderRequest struct {
	AppID       int    `json:"app_id" url:"app_id"`
	AppUser     string `json:"app_user" url:"app_user"`
	AppTransID  string `json:"app_trans_id" url:"app_trans_id"`
	Amount      int64  `json:"amount" url:"amount"`
	AppTime     int64  `json:"app_time" url:"app_time"`
	Item        string `json:"item" url:"item"`
	Description string `json:"description" url:"description"`
	EmbedData   string `json:"embed_data" url:"embed_data"`
	BankCode    string `json:"bank_code" url:"bank_code"`
	Mac         string `json:"mac" url:"mac"`
	CallbackURL string `json:"callback_url" url:"callback_url"`
	DeviceInfo  string `json:"device_info" url:"device_info"`
	SubAppID    string `json:"sub_app_id" url:"sub_app_id"`
	Title       string `json:"title" url:"title"`
	Currency    string `json:"currency" url:"currency"`
	Phone       string `json:"phone" url:"phone"`
	Email       string `json:"email" url:"email"`
	Address     string `json:"address" url:"address"`
	ProductCode string `json:"product_code" url:"product_code"`
}

type CreateOrderResponse struct {
	ReturnCode       int    `json:"return_code"`
	SubReturnCode    int    `json:"sub_return_code"`
	ReturnMessage    string `json:"return_message"`
	SubReturnMessage string `json:"sub_return_message"`
	OrderUrl         string `json:"order_url"`
	OrderToken       string `json:"order_token"`
	ZpTransToken     string `json:"zp_trans_token"`
	QrCode           string `json:"qr_code"`
}

type QueryStatusRequest struct {
	AppID      int    `json:"app_id" url:"app_id"`
	AppTransID string `json:"app_trans_id" url:"app_trans_id"`
	Mac        string `json:"mac" url:"mac"`
}

type QueryStatusResponse struct {
	ReturnCode       int32  `json:"return_code"`
	ReturnMessage    string `json:"return_message"`
	SubReturnCode    int32  `json:"sub_return_code"`
	SubReturnMessage string `json:"sub_return_message"`
	Amount           int64  `json:"amount"`
	ZPTransID        int64  `json:"zp_trans_id"`
}

// ToValues convert to url.Values
func (cor *CreateOrderRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}

// ToValues convert to url.Values
func (cor *QueryStatusRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}
