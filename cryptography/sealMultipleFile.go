package cryptography

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Sandhiguna/sgkms-go-sdk/common"
)

func SealMultipleFile(req FileReq) (string, error) {

	client := common.GetSharedHTTPClient()
	if client == nil {
		return "", errors.New("http client belum diinisialisasi")
	}

	results := make(map[string]string)

	for _, inputPath := range req.InputFiles {
		rawData, err := os.ReadFile(inputPath)
		if err != nil {
			return "", fmt.Errorf(`{"error":"gagal baca file %s: %v"}`, inputPath, err)
		}

		ciphertexts := []string{}

		fileName := filepath.Base(inputPath)
		sealReqName := SealReq{
			Init: Init{
				BaseURL: req.Init.BaseURL,
			},
			SessionToken: req.Init.SessionToken,
			SlotID:       req.Init.SlotID,
			KeyID:        req.KeyID,
			Plaintext:    []string{fileName},
		}

		sealedName, err := Seal(sealReqName)
		if err != nil {
			return "", fmt.Errorf("seal gagal untuk file %s: %w", fileName, err)
		}

		ciphertexts = append(ciphertexts, sealedName.Result.Ciphertext...)

		chunkSize := 512 * 1024
		totalChunks := (len(rawData) + chunkSize - 1) / chunkSize

		for i := 0; i < totalChunks; i++ {
			start := i * chunkSize
			end := start + chunkSize
			if end > len(rawData) {
				end = len(rawData)
			}
			chunk := rawData[start:end]

			encoded := base64.StdEncoding.EncodeToString(chunk)

			sealReq := SealReq{
				Init: Init{
					BaseURL: req.Init.BaseURL,
				},
				SessionToken: req.Init.SessionToken,
				SlotID:       req.Init.SlotID,
				KeyID:        req.KeyID,
				Plaintext:    []string{encoded},
			}
			sealedChunk, err := Seal(sealReq)
			if err != nil {
				return "", fmt.Errorf(`{"error":"seal gagal pada file %s chunk %d: %v"}`, inputPath, i, err)
			}
			ciphertexts = append(ciphertexts, sealedChunk.Result.Ciphertext...)
		}

		nameOnly := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		outputFileName := nameOnly + ".sgc"
		outputPath := filepath.Join(req.OutputDir, outputFileName)

		sealBytes, _ := json.Marshal(ciphertexts)
		if err := os.WriteFile(outputPath, sealBytes, 0644); err != nil {
			return "", (fmt.Errorf(`{"error":"gagal simpan file %s: %v"}`, outputPath, err))
		}

		results[inputPath] = outputPath
	}

	jsonBytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal hasil JSON gagal: %w", err)
	}

	return string(jsonBytes), nil
}
