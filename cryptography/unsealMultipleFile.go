package cryptography

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func UnsealMultipleFile(req DecryptMultipleFileReq) (string, error) {

	results := make(map[string]string)

	for _, inputPath := range req.EncryptedFiles {
		data, err := os.ReadFile(inputPath)
		if err != nil {
			return "", fmt.Errorf(`{"error":"gagal baca file %s: %v"}`, inputPath, err)
		}

		var ciphertexts []string
		if err := json.Unmarshal(data, &ciphertexts); err != nil {
			return "", fmt.Errorf(`{"error":"parse sgc file %s failed: %v"}`, inputPath, err)
		}

		if len(ciphertexts) < 2 {
			return "", fmt.Errorf(`{"error":"sgc file %s tidak valid (minimal 2 elemen)"}`, inputPath)
		}

		outputPath := ""
		contentChunks := []byte{}

		for i, c := range ciphertexts {
			unsealReq := UnsealReq{
				Init: Init{
					BaseURL:      req.Init.BaseURL,
					SessionToken: req.Init.SessionToken,
					SlotID:       req.Init.SlotID,
				},
				Ciphertext: []string{c},
			}
			resUnseal, err := Unseal(unsealReq)
			if err != nil {
				return "", fmt.Errorf(`{"error":"unseal index %d gagal di file %s: %v"}`, i, inputPath, err)
			}

			if len(resUnseal.Result.Plaintext) == 0 {
				return "", fmt.Errorf(`{"error":"hasil unseal kosong di index %d pada file %s"}`, i, inputPath)
			}

			if i == 0 {
				fileName := resUnseal.Result.Plaintext[0]
				outputPath = filepath.Join(req.OutputDir, fileName)
			} else {
				decoded, err := base64.StdEncoding.DecodeString(resUnseal.Result.Plaintext[0])
				if err != nil {
					return "", fmt.Errorf(`{"error":"decode base64 gagal di chunk %d: %v"}`, i, err)
				}
				contentChunks = append(contentChunks, decoded...)
			}
		}

		if err := os.WriteFile(outputPath, contentChunks, 0644); err != nil {
			return "", fmt.Errorf(`{"error":"gagal simpan output %s: %v"}`, outputPath, err)
		}

		results[inputPath] = outputPath
	}

	jsonBytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal hasil JSON gagal: %w", err)
	}

	return string(jsonBytes), nil
}
