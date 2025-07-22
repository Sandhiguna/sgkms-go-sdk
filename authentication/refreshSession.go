package authentication

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Sandhiguna/sgkms-go-sdk/common"
)

func RefreshSession(baseURL string, request RefreshSessionReq) (*LoginRes, error) {
	client := common.GetSharedHTTPClient()
	if client == nil {
		return nil, errors.New("http client belum diinisialisasi")
	}

	jsonBody, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", baseURL+"/agent/refreshSession", bytes.NewBuffer((jsonBody)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, errors.New("refresh session gagal: status " + resp.Status + " |body: " + string(respBody))
	}

	var loginresponse LoginRes
	if err := json.Unmarshal(respBody, &loginresponse); err != nil {
		return nil, err
	}

	return &loginresponse, nil
}
