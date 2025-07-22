package cryptography

import (
	"encoding/json"
	"fmt"
	"os"

	zipper "github.com/Sandhiguna/sgkms-go-sdk/common"
)

type Response struct {
	Message string `json:"message"`
}

func UncompressFiles(req DecryptMultipleFileReq) (string, error) {

	decryptedJSON, err := DecryptMultipleFiles(req)
	if err != nil {
		return "", fmt.Errorf(`{"error":"decrypt failed: %v"}`, err)
	}

	var decryptedFiles []DecryptedFileJSON
	if err := json.Unmarshal([]byte(decryptedJSON), &decryptedFiles); err != nil {
		return "", fmt.Errorf(`{"error":"parse decrypt result: %v"}`, err)
	}

	// var extractedPaths []string
	for _, file := range decryptedFiles {
		// File.Output = path ke ZIP yang sudah didekripsi
		if err := zipper.Unzip(file.Output, req.OutputDir); err != nil {
			return "", fmt.Errorf(`{"error":"unzip failed: %v"}`, err)
		}
		// Hapus file ZIP setelah berhasil diekstrak
		if err := os.Remove(file.Output); err != nil {
			return "", fmt.Errorf(`{"error":"failed to delete zip file: %v"}`, err)
		}
		// extractedPaths = append(extractedPaths, file.Output)
	}

	result := Response{
		Message: "success Uncompress File",
	}

	finalJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf(`{"error":"marshal result failed: %v"}`, err)
	}

	return string(finalJSON), nil
}
