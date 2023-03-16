package thailandposttracking

import (
	"time"

	"github.com/pkg/errors"
)

const (
	webhookAPIBaseUrl                 = "https://trackwebhook.thailandpost.co.th/post/api/v1/"
	endpointWebhookGetToken           = "authenticate/token"
	endpointWebhookSubscribeByBarcode = "hook"
	endpointWebhookSubscribeByReceipt = "hook/receipt"
	endpointWebhookSubscribeByProfile = "user-register/user-profile"
)

func (c *Client) getWebhookToken() error {
	// # REQ
	// Authorization: Token (Your-Token-Key)
	// Content-Type: application/json

	// # RESP
	// {
	// 	"expire": "2019-09-28 10:18:20+07:00",
	// 	"token": "eyJ0eXAiOiJKV1QiLCJhbG..."
	// }
	resp := GetTokenResponse{}
	// set token
	c.webhookToken.Expire = time.Time(resp.Expire)
	c.webhookToken.Token = resp.Token

	return nil
}

type (
	SubscribeByBarcodeRequest struct {
		Status                string   `json:"status"`              // เช่น รับฝาก = '103', ทั้งหมด = 'all'
		Language              string   `json:"language"`            // TH,EN,CN
		Barcodes              []string `json:"barcode"`             // หมายเลขสิ่งของ เช่น EY145587896TH
		RequestPreviousStatus bool     `json:"req_previous_status"` // รับสถานะย้อนหลัง = 'true', ไม่รับสถานะย้อนหลัง ='false'
	}
	SubscribeByBarcodeResponse struct {
		Data    SubscribeByBarcodeResponseData `json:"response"`
		Message string                         `json:"message"`
		Status  bool                           `json:"status"`
	}
	SubscribeByBarcodeResponseData struct {
		Items      []SubscribeByBarcodeResponseItem `json:"items"`
		TrackCount SubscribeTrackCount              `json:"track_count"`
	}
	SubscribeByBarcodeResponseItem struct {
		Barcode string `json:"barcode"`
		Status  bool   `json:"status"`
	}
	SubscribeTrackCount struct {
		TrackDate       string `json:"track_date"`
		Count           int    `json:"count_number"`
		TrackCountLimit int    `json:"track_count_limit"`
	}
)

func (c *Client) SubscribeByBarcode() error {
	if isTokenExpired(&c.webhookToken) {
		err := c.getWebhookToken()
		if err != nil {
			return errors.Wrap(err, "[SubscribeByBarcode]")
		}
	}

	return nil
}

type (
	SubscribeByReceiptRequest struct {
		Status                string   `json:"status"`              // เช่น รับฝาก = '103', ทั้งหมด = 'all'
		Language              string   `json:"language"`            // TH,EN,CN
		ReceiptNo             []string `json:"receiptNo"`           // หมายเลขใบเสร็จ เช่น 361101377131
		RequestPreviousStatus bool     `json:"req_previous_status"` // รับสถานะย้อนหลัง = 'true', ไม่รับสถานะย้อนหลัง ='false'
	}
	SubscribeByReceiptResponse struct {
		Data    SubscribeByReceiptResponseData `json:"response"`
		Message string                         `json:"message"`
		Status  bool                           `json:"status"`
	}
	SubscribeByReceiptResponseData struct {
		Receipts   map[string]SubscribeByReceiptResponseReceipts `json:"receipts"`
		TrackCount SubscribeTrackCount                           `json:"track_count"`
	}
	SubscribeByReceiptResponseReceipts []string
)

func (c *Client) SubscribeByReceipt() error {
	if isTokenExpired(&c.webhookToken) {
		err := c.getWebhookToken()
		if err != nil {
			return errors.Wrap(err, "[SubscribeByReceipt]")
		}
	}

	return nil
}

func (c *Client) SubscribeByProfile() error {
	return errors.New("[SubscribeByProfile] Not implemented")

	// if isTokenExpired(&c.webhookToken) {
	// 	err := c.getWebhookToken()
	// 	if err != nil {
	// 		return errors.Wrap(err, "[SubscribeByProfile]")
	// 	}
	// }
	// return nil
}

func (c *Client) UnSubscribeByProfile() error {
	return errors.New("[UnSubscribeByProfile] Not implemented")
	// if isTokenExpired(&c.webhookToken) {
	// 	err := c.getWebhookToken()
	// 	if err != nil {
	// 		return errors.Wrap(err, "[UnSubscribeByProfile]")
	// 	}
	// }
	// return nil
}

type (
	WebhookData struct {
		Items         []ThailandPostItemStatus `json:"items"`
		TrackDateTime string                   `json:"track_datetime"` /// "10/09/2562 10:17+07:00"
	}
)

/*
Authorization: Bearer (Your-Key)
Content-Type: application/json
*/

func (c *Client) WebhookCallback() {

}
