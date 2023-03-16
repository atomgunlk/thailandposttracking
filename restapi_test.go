package thailandposttracking_test

import (
	"encoding/json"
	"testing"

	"github.com/atomgunlk/golang-common/pkg/logger"
	"github.com/atomgunlk/thailandposttracking"
	"github.com/stretchr/testify/assert"
)

func TestGetItemsbyBarcode(t *testing.T) {
	response, err := thpostClient.GetItemsbyBarcode(&thailandposttracking.GetItemsbyBarcodeRequest{
		Status:   "all",
		Language: "TH",
		Barcodes: []string{
			"RK151179337TH",
			"RK151179266TH",
			"RK151179385TH",
			"RK151179323TH",
			"RK151179120TH",
			"RK151178756TH",
			"RK151178844TH",
			"RK151178504TH",
		},
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, true, response.Status)

	jsonStr, _ := json.MarshalIndent(response, "", "    ")
	logger.Debugf("RESP : %s\r\n", jsonStr)
}

func TestGetItemsbyReceipt(t *testing.T) {
	response, err := thpostClient.GetItemsbyReceipt(&thailandposttracking.GetItemsbyReceiptRequest{
		Status:   "all",
		Language: "TH",
		ReceiptNo: []string{
			"361101377131",
			"361101377132",
			"361101377133",
		},
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, true, response.Status)

	jsonStr, _ := json.MarshalIndent(response, "", "    ")
	logger.Debugf("RESP : %s\r\n", jsonStr)
}
