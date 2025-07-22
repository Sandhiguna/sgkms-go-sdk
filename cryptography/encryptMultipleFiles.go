package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Sandhiguna/sgkms-go-sdk/authentication"
	"github.com/Sandhiguna/sgkms-go-sdk/common"
)

func EncryptMultipleFiles(req FileReq) (string, error) {
	client := common.GetSharedHTTPClient()
	if client == nil {
		return "", errors.New("http client belum diinisialisasi")
	}

	randomResp, err := authentication.RandomNumber(32, req.Init.SlotID, req.Init.SessionToken, req.Init.BaseURL)
	if err != nil {
		return "", fmt.Errorf("failed random key: %w", err)
	}
	randomKey := randomResp.Random

	var results []EncryptedFileJSON

	for _, inFile := range req.InputFiles {
		filename := filepath.Base(inFile)
		plaintexts := []string{filename, base64.StdEncoding.EncodeToString(randomKey)}

		sealReq := SealReq{
			Init: Init{
				BaseURL: req.Init.BaseURL,
			},
			SessionToken: req.Init.SessionToken,
			SlotID:       req.Init.SlotID,
			KeyID:        req.KeyID,
			Plaintext:    plaintexts,
		}

		sealedKeys, err := Seal(sealReq) // menghasikan encrypt nama file dan kunci per file
		if err != nil {
			return "", fmt.Errorf("failed seal %s: %w", filename, err)
		}

		plaintext, err := os.ReadFile(inFile)
		if err != nil {
			return "", fmt.Errorf("baca file %s gagal: %w", inFile, err)
		}

		block, err := aes.NewCipher(randomKey)
		if err != nil {
			return "", fmt.Errorf("buat cipher gagal: %w", err)
		}

		aesGCM, err := cipher.NewGCM(block)
		if err != nil {
			return "", fmt.Errorf("buat aesGCM gagal: %w", err)
		}

		nonce := make([]byte, aesGCM.NonceSize())
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return "", fmt.Errorf("generate nonce gagal: %w", err)
		}

		ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)
		encryptedData := append(nonce, ciphertext...)

		nameOnly := strings.TrimSuffix(filename, filepath.Ext(filename))
		outputFile := filepath.Join(req.OutputDir, nameOnly+".sgc")

		metaBytes, err := json.Marshal(sealedKeys.Result.Ciphertext)
		if err != nil {
			return "", fmt.Errorf("marshal seal metadata gagal: %w", err)
		}
		separator := []byte("::")
		fileData := append(metaBytes, separator...)
		fileData = append(fileData, encryptedData...)

		if err := os.WriteFile(outputFile, fileData, 0644); err != nil {
			return "", fmt.Errorf("tulis file %s gagal: %w", outputFile, err)
		}

		results = append(results, EncryptedFileJSON{
			Input:      inFile,
			Output:     outputFile,
			SealedKeys: sealedKeys.Result.Ciphertext,
		})
	}

	jsonBytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal hasil JSON gagal: %w", err)
	}

	return string(jsonBytes), nil
}
