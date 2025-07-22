package cryptography

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Sandhiguna/sgkms-go-sdk/common"
)

func Unseal(req UnsealReq) (*UnsealRes, error) {
	client := common.GetSharedHTTPClient()
	if client == nil {
		return nil, errors.New("http client belum diinisialisasi")
	}

	body, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", req.Init.BaseURL+"/unseal", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		respData, _ := io.ReadAll(resp.Body)
		return nil, errors.New("unseal gagal: status " + resp.Status + " - " + string(respData))
	}

	var unsealRes UnsealRes
	if err := json.NewDecoder(resp.Body).Decode(&unsealRes); err != nil {
		return nil, err
	}

	return &unsealRes, nil

}
