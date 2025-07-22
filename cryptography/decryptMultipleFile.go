package cryptography

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Sandhiguna/sgkms-go-sdk/common"
)

func DecryptMultipleFiles(req DecryptMultipleFileReq) (string, error) {
	client := common.GetSharedHTTPClient()
	if client == nil {
		return "", errors.New("http client belum diinisialisasi")
	}

	var results []DecryptedFileJSON

	for _, encryptedFile := range req.EncryptedFiles {
		rawData, err := os.ReadFile(encryptedFile)
		if err != nil {
			log.Printf("Gagal baca file %s: %v", encryptedFile, err)
			continue
		}

		separator := []byte("::")
		parts := bytes.SplitN(rawData, separator, 2)
		if len(parts) != 2 {
			log.Printf("Format file tidak valid: %s", encryptedFile)
			continue
		}

		var sealedKeys []string
		if err := json.Unmarshal(parts[0], &sealedKeys); err != nil {
			log.Printf("Unmarshal metadata gagal: %s", encryptedFile)
			continue
		}

		reqUnseal := UnsealReq{
			Init: Init{
				BaseURL:      req.Init.BaseURL,
				SessionToken: req.Init.SessionToken,
				SlotID:       req.Init.SlotID,
			},
			Ciphertext: sealedKeys,
		}

		unsealedResp, err := Unseal(reqUnseal)
		if err != nil {
			log.Printf("Unseal gagal: %s: %v", encryptedFile, err)
			continue
		}

		if len(unsealedResp.Result.Plaintext) != 2 {
			log.Printf("Format plaintext tidak sesuai: %s", encryptedFile)
			continue
		}

		originalFileName := unsealedResp.Result.Plaintext[0]
		randomKeyB64 := unsealedResp.Result.Plaintext[1]

		randomKey, err := base64.StdEncoding.DecodeString(randomKeyB64)
		if err != nil {
			log.Printf("Decode key gagal: %s", encryptedFile)
			continue
		}

		block, err := aes.NewCipher(randomKey)
		if err != nil {
			log.Printf("AES cipher gagal: %s", encryptedFile)
			continue
		}

		aesGCM, err := cipher.NewGCM(block)
		if err != nil {
			log.Printf("GCM gagal: %s", encryptedFile)
			continue
		}

		encryptedData := parts[1]
		nonceSize := aesGCM.NonceSize()
		if len(encryptedData) < nonceSize {
			log.Printf("Data terenkripsi tidak valid: %s", encryptedFile)
			continue
		}

		nonce := encryptedData[:nonceSize]
		ciphertext := encryptedData[nonceSize:]

		plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			log.Printf("Dekripsi gagal: %s", encryptedFile)
			continue
		}

		outputPath := filepath.Join(req.OutputDir, originalFileName)
		if err := os.WriteFile(outputPath, plaintext, 0644); err != nil {
			log.Printf("Tulis file gagal: %s", outputPath)
			continue
		}

		results = append(results, DecryptedFileJSON{
			Input:  encryptedFile,
			Output: outputPath,
		})
	}

	if len(results) == 0 {
		return "", fmt.Errorf("tidak ada file berhasil didekripsi")
	}

	jsonBytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal hasil JSON gagal: %w", err)
	}

	return string(jsonBytes), nil
}
