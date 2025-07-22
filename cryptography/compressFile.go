package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Sandhiguna/sgkms-go-sdk/authentication"
	zipper "github.com/Sandhiguna/sgkms-go-sdk/common"
)

func CompressFiles(req FileReq) (string, error) {

	// Generate nama ZIP berdasarkan timestamp
	timestamp := time.Now().Format("20060102_150405")
	zipName := fmt.Sprintf("%s.zip", timestamp)
	zipPath := filepath.Join(req.OutputDir, zipName)

	// Proses ZIP file-file input
	if err := zipper.CreateZipFromFiles(req.InputFiles, zipPath); err != nil {
		return "", fmt.Errorf(`{"error":"zip failed: %v"}`, err)
	}

	defer os.Remove(zipPath)

	randomResp, err := authentication.RandomNumber(32, req.Init.SlotID, req.Init.SessionToken, req.Init.BaseURL)
	if err != nil {
		return "", fmt.Errorf("gagal dapat random key: %w", err)
	}
	randomKey := randomResp.Random
	plaintexts := []string{zipName, base64.StdEncoding.EncodeToString(randomKey)}

	sealReq := SealReq{
		Init: Init{
			BaseURL: req.Init.BaseURL,
		},
		SessionToken: req.Init.SessionToken,
		SlotID:       req.Init.SlotID,
		KeyID:        req.KeyID,
		Plaintext:    plaintexts,
	}

	sealedKeys, err := Seal(sealReq)
	if err != nil {
		return "", fmt.Errorf("seal gagal untuk file %s: %w", zipName, err)
	}

	plaintext, err := os.ReadFile(zipPath)
	if err != nil {
		return "", fmt.Errorf("baca file %s gagal: %w", zipPath, err)
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

	nameOnly := strings.TrimSuffix(zipName, filepath.Ext(zipName))
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

	results := EncryptedFileJSON{
		Input:      zipName,
		Output:     outputFile,
		SealedKeys: sealedKeys.Result.Ciphertext,
	}

	jsonBytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal hasil JSON gagal: %w", err)
	}

	return string(jsonBytes), nil

	// // Proses enkripsi ZIP
	// results, err := encryption.EncryptMultipleFiles(
	// 	BaseURL, CertPath, KeyPath,
	// 	keySeal, password,zipPath
	// 	insecure, slotID,
	// 	[]string{}, outputDir,
	// )
	// if err != nil {
	// 	return "", fmt.Errorf(`{"error":"encrypt failed: %v"}`, err)
	// }
}
