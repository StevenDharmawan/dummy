package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"phase3-gc1-shopping/model/web"
)

func CallPaymentService(url string, request web.PaymentRequest) *web.ErrorResponse {
	reqBody, err := json.Marshal(request)
	if err != nil {
		errResponse := ErrInternalServer(err.Error())
		return &errResponse
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		errResponse := ErrInternalServer(err.Error())
		return &errResponse
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errResponse := ErrInternalServer(err.Error())
		return &errResponse
	}
	defer resp.Body.Close()

	// Log respons yang diterima dari server pembayaran
	fmt.Println("Payment response status code:", resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Payment response body:", string(body))

	if resp.StatusCode != http.StatusCreated {
		errResponse := ErrInternalServer(err.Error())
		return &errResponse

	}

	return nil
}
