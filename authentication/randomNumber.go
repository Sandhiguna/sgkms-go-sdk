package authentication

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Sandhiguna/sgkms-go-sdk/common"
)

func RandomNumber(length, slotId int, sessionToken, baseURL string) (*RandomResult, error) {
	client := common.GetSharedHTTPClient()
	if client == nil {
		return nil, errors.New("http client belum diinisialisasi")
	}

	body := map[string]interface{}{
		"sessionToken": sessionToken,
		"slotId":       slotId,
		"length":       length,
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", baseURL+"/rng", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("randomNumber gagal: " + string(raw))
	}

	var rngResp RngResponse
	err = json.Unmarshal(raw, &rngResp)
	if err != nil {
		return nil, err
	}

	return &rngResp.Result, nil
}
