package cryptography

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Sandhiguna/sgkms-go-sdk/common"
)

func Decrypt(baseURL string, request DecryptReq) (*DecryptRes, error) {
	client := common.GetSharedHTTPClient()
	if client == nil {
		return nil, errors.New("http client belum diinisialisasi")
	}

	jsonBody, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", baseURL+"/decrypt", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, errors.New("decrypt gagal: status " + resp.Status + " | body: " + string(respBody))
	}

	var decryptResponse DecryptRes
	if err := json.Unmarshal(respBody, &decryptResponse); err != nil {
		return nil, err
	}

	return &decryptResponse, nil
}
