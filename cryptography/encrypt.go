package cryptography

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Sandhiguna/sgkms-go-sdk/common"
)

func Encrypt(sessionToken string, slotId int, keyId string, plaintextEncrypt []PlaintextEncrypt, baseURL string) (*EncryptRes, error) {
	client := common.GetSharedHTTPClient()
	if client == nil {
		return nil, errors.New("http client belum diinisialisasi")
	}

	// Membuat payload sesuai dengan format request yang benar
	payload := EncryptReq{
		SessionToken: sessionToken,
		SlotID:       slotId,
		KeyID:        keyId,
		Plaintext:    plaintextEncrypt, // Array dari Plaintext
	}

	// Marshal struct EncryptRequest menjadi JSON
	body, _ := json.Marshal(payload)

	// Membuat request POST
	req, _ := http.NewRequest("POST", baseURL+"/encrypt", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Mengirimkan request ke server
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("encrypt gagal: " + string(respBody))
	}

	// Parse response menjadi EncryptResponse
	var result EncryptRes
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, err
	}

	// Mengembalikan seluruh hasil enkripsi dalam EncryptResponse
	return &result, nil
}
