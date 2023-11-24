package zod

var (
	SB   = "https://zlpqc-onlpm-api-gateway.zalopay.vn"
	PROD = "https://apimep.zalopay.vn"
)

type Receiver struct {
	Contact string `json:"contact" url:"contact"`
}

type OrderInfo struct {
	TrackingNumber string `json:"trackingNumber" url:"trackingNumber"`
	Description    string `json:"description" url:"description"`
	Amount         int    `json:"amount" url:"amount"`
}

type CreateInvoiceRequest struct {
	AppID     string      `json:"appId" url:"appId"`
	MCRefID   string      `json:"mcRefId" url:"mcRefId"`
	HubID     string      `json:"hubId" url:"hubId"`
	DriverID  string      `json:"driverId" url:"driverId"`
	Amount    int64       `json:"amount" url:"amount"`
	Receiver  Receiver    `json:"receiver" url:"receiver"`
	OrderInfo []OrderInfo `json:"orderInfo" url:"orderInfo"`
	MCExtInfo string      `json:"mcExtInfo" url:"mcExtInfo"`
	MacKey    string      `json:"mac" url:"mac"`
}

type CreateInvoiceResponse struct {
	OrderUrl string `json:"orderUrl"`
}

type QueryInvoiceRequest struct {
	AppID   string `json:"appId" url:"appId"`
	MCRefID string `json:"mcRefId" url:"mcRefId"`
	MacKey  string `json:"mac" url:"mac"`
}

type QueryInvoiceResponse struct {
	OrderUrl string `json:"orderUrl"`
}

type QueryOrderStatusRequest struct {
	AppID   string `json:"appId" url:"appId"`
	MCRefID string `json:"mcRefId" url:"mcRefId"`
	MacKey  string `json:"mac" url:"mac"`
}

type QueryOrderStatusResponse struct {
	Status    int    `json:"status"`
	Amount    string `json:"amount"`
	ZpTransId string `json:"zpTransId"`
}
