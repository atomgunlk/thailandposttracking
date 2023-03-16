package thailandposttracking

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/atomgunlk/golang-common/pkg/logger"
	"github.com/atomgunlk/golang-common/pkg/request"
	"github.com/pkg/errors"
)

const (
	restAPIBaseUrl                = "https://trackapi.thailandpost.co.th/post/api/v1/"
	endpointRestGetToken          = "authenticate/token"
	endpointRestGetItemsbyBarcode = "track"
	endpointRestGetItemsbyReceipt = "receipt/track"
	endpointRestRequestItems      = "track/batch"
)

type GetTokenResponse struct {
	Expire TPDateTime `json:"expire"`
	Token  string     `json:"token"`
}

func (c *Client) getAPIToken() error {
	url := fmt.Sprintf("%s%s", restAPIBaseUrl, endpointRestGetToken)
	res, err := c.httpClient.Post(url, request.SendOptions{
		"headers": map[string]interface{}{
			"Authorization": fmt.Sprintf("Token %s", c.secretToken),
			"Content-Type":  "application/json",
		},
	}, nil)

	if err != nil {
		logger.WithError(err).Error("[ThailandPost.getAPIToken]: unable to get access token")
		return errors.Wrap(err, "[ThailandPost.getAPIToken]: unable to get token")
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("[ThailandPost.getAPIToken]: Status code %d", res.StatusCode)
	}

	response := &GetTokenResponse{}
	if err := json.Unmarshal(res.Body, response); err != nil {
		logger.WithError(err).Errorf("[ThailandPost.getAPIToken]: resp body %s", string(res.Body))

		return errors.Wrap(err, "[ThailandPost.getAPIToken]: unable to unmarshal response")
	}

	// set token
	c.apiToken.Expire = time.Time(response.Expire)
	c.apiToken.Token = response.Token

	return nil
}

type (
	GetItemsbyBarcodeRequest struct {
		Status   string   `json:"status"`   // เช่น รับฝาก = '103', ทั้งหมด = 'all'
		Language string   `json:"language"` // TH,EN,CN
		Barcodes []string `json:"barcode"`  // หมายเลขสิ่งของ เช่น EY145587896TH
	}
	GetItemsbyBarcodeResponse struct {
		Data GetItemsbyBarcodeResponseData `json:"response"`
	}
	GetItemsbyBarcodeResponseData struct {
		Items map[string][]ThailandPostItemStatus `json:"items"`
	}
	ThailandPostItemStatus struct {
		Barcode             string `json:"barcode"`              // "EF023395845TH",
		Status              string `json:"status"`               // "103",
		StatusDescription   string `json:"status_description"`   //  "รับฝาก",
		StatusDate          string `json:"status_date"`          //  "04/06/2562 15:53:22+07:00",
		Location            string `json:"location"`             //  "ศูนย์ศิลปาชีพบางไทร",
		Postcode            string `json:"postcode"`             //  "13290",
		DeliveryStatus      string `json:"delivery_status"`      //  null, | "S",
		DeliveryDescription string `json:"delivery_description"` //  null, | "ผู้รับได้รับสิ่งของเรียบร้อยแล้ว",
		DeliveryDateTime    string `json:"delivery_datetime"`    //  null, | "16/03/2566 17:41:44+07:00",
		ReceiverName        string `json:"receiver_name"`        //  null, | "วรรณเพ็ญ วสิน/ผู้รับรับเอง",
		Signature           string `json:"signature"`            //  null | "https://trackimage.thailandpost.co.th/f/signature/QDc4ODQ0YjVzMGx1VDMz/QGI1c1JLMGx1VDMx/QGI1czBsVEh1VDM0/QGI1czBsdTE1MTFUMzI="
	}
)

func (c *Client) GetItemsbyBarcode(req *GetItemsbyBarcodeRequest) (*GetItemsbyBarcodeResponse, error) {
	if isTokenExpired(&c.apiToken) {
		err := c.getAPIToken()
		if err != nil {
			return nil, errors.Wrap(err, "[ThailandPost.GetItemsbyBarcode]")
		}
	}

	url := fmt.Sprintf("%s%s", restAPIBaseUrl, endpointRestGetItemsbyBarcode)
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "[ThailandPost.GetItemsbyBarcode]: unable to marshal req")
	}
	res, err := c.httpClient.Post(url, request.SendOptions{
		"headers": map[string]interface{}{
			"Authorization": fmt.Sprintf("Token %s", c.apiToken.Token),
			"Content-Type":  "application/json",
		},
	}, reqBody)
	if err != nil {
		logger.WithError(err).Error("[ThailandPost.GetItemsbyBarcode]:")
		return nil, errors.Wrap(err, "[ThailandPost.GetItemsbyBarcode]:")
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[ThailandPost.GetItemsbyBarcode]: Status code %d", res.StatusCode)
	}

	response := &GetItemsbyBarcodeResponse{}
	if err := json.Unmarshal(res.Body, response); err != nil {
		logger.WithError(err).Errorf("[ThailandPost.GetItemsbyBarcode]: resp body %s", string(res.Body))
		return nil, errors.Wrap(err, "[ThailandPost.GetItemsbyBarcode]: unable to unmarshal response")
	}

	return response, nil
}
func (c *Client) GetItemsbyReceipt() error {
	// # REQ
	// Authorization: Token (Your-Token-Key)
	// Content-Type: application/json

	// # RESP
	// {
	// 	"expire": "2019-09-28 10:18:20+07:00",
	// 	"token": "eyJ0eXAiOiJKV1QiLCJhbG..."
	// }
	if isTokenExpired(&c.apiToken) {
		err := c.getAPIToken()
		if err != nil {
			return errors.Wrap(err, "[ThailandPost.GetItemsbyBarcode]")
		}
	}

	return nil
}
func (c *Client) RequestItems() error {
	// # REQ
	// Authorization: Token (Your-Token-Key)
	// Content-Type: application/json

	// # RESP
	// {
	// 	"expire": "2019-09-28 10:18:20+07:00",
	// 	"token": "eyJ0eXAiOiJKV1QiLCJhbG..."
	// }
	if isTokenExpired(&c.apiToken) {
		err := c.getAPIToken()
		if err != nil {
			return errors.Wrap(err, "[ThailandPost.GetItemsbyBarcode]")
		}
	}

	return nil
}
