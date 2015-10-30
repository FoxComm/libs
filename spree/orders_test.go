package spree

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyOrdersSuccessEmpty(t *testing.T) {
	payload := `{"count":0,"total_count":0,"current_page":1,"per_page":25,"pages":0,"orders":[]}`
	server := setUpServer(200, payload)
	defer server.Close()

	data, err := SpreeAPI.MyOrders("", server.URL)

	assert.NoError(t, err)

	assert.Equal(t, len(data), 0)
}

func TestMyOrdersSuccessWithData(t *testing.T) {
	orderPayload := `[{"id":112,"number":"R544771758","item_total":"137.94","total":"112.94","ship_total":"0.0","state":"complete","adjustment_total":"-25.0","user_id":31,"created_at":"2015-06-26T16:37:46.785Z","updated_at":"2015-06-26T16:45:36.714Z","completed_at":"2015-06-26T16:45:36.714Z","payment_total":"0.0","shipment_state":"pending","payment_state":"balance_due","email":"eug23@local.local","special_instructions":null,"channel":"spree","included_tax_total":"0.0","additional_tax_total":"0.0","display_included_tax_total":"$0.00","display_additional_tax_total":"$0.00","currency":"USD","newgistics_status":null,"display_item_total":"$137.94","total_quantity":6,"display_total":"$112.94","display_ship_total":"$0.00","display_tax_total":"$0.00","display_adjustment_total":"-$25.00","token":"4c4940604f076b71","checkout_steps":["address","payment","confirm","complete"]},{"id":114,"number":"R891329072","item_total":"91.96","total":"91.96","ship_total":"0.0","state":"complete","adjustment_total":"0.0","user_id":31,"created_at":"2015-06-26T16:45:38.236Z","updated_at":"2015-06-26T16:46:02.587Z","completed_at":"2015-06-26T16:46:02.587Z","payment_total":"0.0","shipment_state":"pending","payment_state":"balance_due","email":"eug23@local.local","special_instructions":null,"channel":"spree","included_tax_total":"0.0","additional_tax_total":"0.0","display_included_tax_total":"$0.00","display_additional_tax_total":"$0.00","currency":"USD","newgistics_status":null,"display_item_total":"$91.96","total_quantity":4,"display_total":"$91.96","display_ship_total":"$0.00","display_tax_total":"$0.00","display_adjustment_total":"$0.00","token":"e6e9817abcff2374","checkout_steps":["address","payment","confirm","complete"]}]`
	payload := fmt.Sprintf(`{"count":2,"total_count":2,"current_page":1,"per_page":25,"pages":1,"orders":%s}`, orderPayload)
	server := setUpServer(200, payload)
	defer server.Close()

	data, err := SpreeAPI.MyOrders("", server.URL)

	assert.NoError(t, err)

	var referenceData []Order
	orderPayloadBytes := []byte(orderPayload)
	json.Unmarshal(orderPayloadBytes, &referenceData)

	assert.Equal(t, len(referenceData), len(data))
	assert.Equal(t, referenceData, data)
}

func TestMyOrdersOrdersError(t *testing.T) {
	payload := `{"error": "something bad happened"}`
	server := setUpServer(422, payload)
	defer server.Close()

	data, err := SpreeAPI.MyOrders("", server.URL)

	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestCompleteOrdersSuccessEmpty(t *testing.T) {
	payload := `{"count":0,"total_count":0,"current_page":1,"per_page":25,"pages":0,"orders":[]}`
	server := setUpServer(200, payload)
	defer server.Close()

	data, err := SpreeAPI.CompleteOrders(42, "", server.URL)

	assert.NoError(t, err)

	assert.Equal(t, len(data), 0)
}

func TestCompleteOrdersSuccessWithData(t *testing.T) {
	orderPayload := `[{"id":112,"number":"R544771758","item_total":"137.94","total":"112.94","ship_total":"0.0","state":"complete","adjustment_total":"-25.0","user_id":31,"created_at":"2015-06-26T16:37:46.785Z","updated_at":"2015-06-26T16:45:36.714Z","completed_at":"2015-06-26T16:45:36.714Z","payment_total":"0.0","shipment_state":"pending","payment_state":"balance_due","email":"eug23@local.local","special_instructions":null,"channel":"spree","included_tax_total":"0.0","additional_tax_total":"0.0","display_included_tax_total":"$0.00","display_additional_tax_total":"$0.00","currency":"USD","newgistics_status":null,"display_item_total":"$137.94","total_quantity":6,"display_total":"$112.94","display_ship_total":"$0.00","display_tax_total":"$0.00","display_adjustment_total":"-$25.00","token":"4c4940604f076b71","checkout_steps":["address","payment","confirm","complete"]},{"id":114,"number":"R891329072","item_total":"91.96","total":"91.96","ship_total":"0.0","state":"complete","adjustment_total":"0.0","user_id":31,"created_at":"2015-06-26T16:45:38.236Z","updated_at":"2015-06-26T16:46:02.587Z","completed_at":"2015-06-26T16:46:02.587Z","payment_total":"0.0","shipment_state":"pending","payment_state":"balance_due","email":"eug23@local.local","special_instructions":null,"channel":"spree","included_tax_total":"0.0","additional_tax_total":"0.0","display_included_tax_total":"$0.00","display_additional_tax_total":"$0.00","currency":"USD","newgistics_status":null,"display_item_total":"$91.96","total_quantity":4,"display_total":"$91.96","display_ship_total":"$0.00","display_tax_total":"$0.00","display_adjustment_total":"$0.00","token":"e6e9817abcff2374","checkout_steps":["address","payment","confirm","complete"]}]`
	payload := fmt.Sprintf(`{"count":2,"total_count":2,"current_page":1,"per_page":25,"pages":1,"orders":%s}`, orderPayload)
	server := setUpServer(200, payload)
	defer server.Close()

	data, err := SpreeAPI.CompleteOrders(42, "", server.URL)

	assert.NoError(t, err)

	var referenceData []Order
	orderPayloadBytes := []byte(orderPayload)
	json.Unmarshal(orderPayloadBytes, &referenceData)

	assert.Equal(t, len(referenceData), len(data))
	assert.Equal(t, referenceData, data)
}

func TestCompleteOrdersError(t *testing.T) {
	payload := `{"error": "something bad happened"}`
	server := setUpServer(422, payload)
	defer server.Close()

	data, err := SpreeAPI.CompleteOrders(42, "", server.URL)

	assert.Error(t, err)
	assert.Nil(t, data)
}

func setUpServer(code int, payload string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, payload)
	}))
}
