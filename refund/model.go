package refund

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

type CreateRefundRequest struct {
	MRefundID   string `json:"m_refund_id" url:"m_refund_id"`
	AppID       int    `json:"app_id" url:"app_id"`
	ZPTransID   string `json:"zp_trans_id" url:"zp_trans_id"`
	Amount      int64  `json:"amount" url:"amount"`
	Timestamp   int64  `json:"timestamp" url:"timestamp"`
	MacKey      string `json:"mac" url:"mac"`
	Description string `json:"description" url:"description"`
}

type CreateRefundResponse struct {
	ReturnCode       int    `json:"return_code"`
	ReturnMessage    string `json:"return_message"`
	SubReturnCode    int    `json:"sub_return_code"`
	SubReturnMessage string `json:"sub_return_message"`
	RefundID         int    `json:"refund_id"`
}

type QueryRefundRequest struct {
	MRefundID string `json:"m_refund_id" url:"m_refund_id"`
	AppID     int    `json:"app_id" url:"app_id"`
	Timestamp int64  `json:"timestamp" url:"timestamp"`
	MacKey    string `json:"mac" url:"mac"`
}

type QueryRefundResponse struct {
	ReturnCode       int    `json:"return_code"`
	ReturnMessage    string `json:"return_message"`
	SubReturnCode    int    `json:"sub_return_code"`
	SubReturnMessage string `json:"sub_return_message"`
}

// ToValues convert to url.Values
func (cor *CreateRefundRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}

// ToValues convert to url.Values
func (cor *QueryRefundRequest) ToValues() (url.Values, error) {
	return query.Values(cor)
}
