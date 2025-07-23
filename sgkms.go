package sgkms

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sandhiguna/sgkms-go-sdk/authentication"
	"github.com/Sandhiguna/sgkms-go-sdk/common"
	"github.com/Sandhiguna/sgkms-go-sdk/cryptography"
)

type SGKMS struct {
	client            *http.Client
	baseUrl           string
	slotID            int
	password          string
	SessionToken      string
	LastUsedTime      int64
	IdleTimeoutInMins int
}

func New(certPath, keyPath, password, baseURL string, slotID int, insecure bool) (*SGKMS, error) {
	// Catat waktu mulai
	// startTime := time.Now()

	client, _ := common.InitTLSClient(certPath, keyPath, insecure)

	common.SetSharedHTTPClient(client)

	s := &SGKMS{
		client:   client,
		slotID:   slotID,
		password: password,
		baseUrl:  baseURL,
	}

	if s.SessionToken == "" {
		resp, err := authentication.Login(s.slotID, s.password, s.baseUrl)
		if err != nil {
			return nil, fmt.Errorf("authentication failed: %w", err)
		}

		s.SessionToken = resp.Result.SessionToken
		s.LastUsedTime = resp.Result.LastUsedTime
		s.IdleTimeoutInMins = resp.Result.IdleTimeoutInMins
	}

	// // Catat waktu selesai
	// elapsedTime := time.Since(startTime)
	// fmt.Printf("Waktu Login: %v detik\n", elapsedTime.Seconds())
	return s, nil
}

func (s *SGKMS) ensureSession() error {
	// startTime := time.Now()
	idleTimeout := time.Duration(s.IdleTimeoutInMins) * time.Minute

	fmt.Println("ExpiredAt (time):", time.UnixMilli(s.LastUsedTime).Add(idleTimeout))
	// fmt.Println("ExpiredAt (time):", time.UnixMilli(s.ExpiredAt))
	// fmt.Println(s.IdleTimeoutInMins)

	if s.SessionToken == "" || time.Now().After(time.UnixMilli(s.LastUsedTime).Add(idleTimeout)) {
		fmt.Println("Session expired or missing, refreshing...")

		request := authentication.RefreshSessionReq{
			SlotID:       s.slotID,
			SessionToken: s.SessionToken,
		}

		newToken, err := authentication.RefreshSession(s.baseUrl, request)
		if err != nil {
			// Refresh gagal, fallback ke login
			fmt.Println("Refresh failed, logging in again...")

			resp, err := authentication.Login(s.slotID, s.password, s.baseUrl)
			if err != nil {
				return fmt.Errorf("login failed: %w", err)
			}
			s.SessionToken = resp.Result.SessionToken
			s.LastUsedTime = resp.Result.LastUsedTime
		} else {
			s.SessionToken = newToken.Result.SessionToken
			s.LastUsedTime = newToken.Result.LastUsedTime
		}
	}

	// // Catat waktu selesai
	// elapsedTime := time.Since(startTime)
	// fmt.Printf("Waktu cek session: %v detik\n", elapsedTime.Seconds())
	return nil
}

func (s *SGKMS) RandomNumber(length int) (*authentication.RandomResult, error) {

	if err := s.ensureSession(); err != nil {
		return nil, err
	}

	res, err := authentication.RandomNumber(length, s.slotID, s.SessionToken, s.baseUrl)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SGKMS) Encrypt(keyId string, plaintext []cryptography.PlaintextEncrypt) (*cryptography.EncryptRes, error) {

	if err := s.ensureSession(); err != nil {
		return nil, err
	}

	res, err := cryptography.Encrypt(s.SessionToken, s.slotID, keyId, plaintext, s.baseUrl)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SGKMS) Decrypt(keyId string, keyVersion *int, Ciphertext []cryptography.Ciphertext) (*cryptography.DecryptRes, error) {
	// Catat waktu mulai
	startTime := time.Now()

	if err := s.ensureSession(); err != nil {
		return nil, err
	}

	req := cryptography.DecryptReq{
		SessionToken: s.SessionToken,
		SlotID:       s.slotID,
		KeyID:        keyId,
		Ciphertext:   Ciphertext,
	}

	if keyVersion != nil {
		req.KeyVersion = keyVersion
	}

	res, err := cryptography.Decrypt(s.baseUrl, req)
	if err != nil {
		return nil, err
	}

	// Catat waktu selesai
	elapsedTime := time.Since(startTime)
	fmt.Printf("Waktu Decrypt: %v detik\n", elapsedTime.Seconds())

	return res, nil
}

func (s *SGKMS) Seal(keyId string, plaintext []string) (*cryptography.SealRes, error) {

	if err := s.ensureSession(); err != nil {
		return nil, err
	}

	req := cryptography.SealReq{
		Init: cryptography.Init{
			BaseURL: s.baseUrl,
		},
		SessionToken: s.SessionToken,
		SlotID:       s.slotID,
		KeyID:        keyId,
		Plaintext:    plaintext,
	}

	res, err := cryptography.Seal(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SGKMS) Unseal(keyId string, ciphertext []string) (*cryptography.UnsealRes, error) {

	if err := s.ensureSession(); err != nil {
		return nil, err
	}

	req := cryptography.UnsealReq{
		Init: cryptography.Init{
			BaseURL:      s.baseUrl,
			SessionToken: s.SessionToken,
			SlotID:       s.slotID,
		},
		Ciphertext: ciphertext,
	}

	res, err := cryptography.Unseal(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SGKMS) EncryptMultipleFiles(keyId, OutputDir string, inputFiles []string) (string, error) {

	if err := s.ensureSession(); err != nil {
		return "", err
	}

	req := cryptography.FileReq{
		Init: cryptography.Init{
			BaseURL:      s.baseUrl,
			SessionToken: s.SessionToken,
			SlotID:       s.slotID,
		},
		KeyID:      keyId,
		InputFiles: inputFiles,
		OutputDir:  OutputDir,
	}

	res, err := cryptography.EncryptMultipleFiles(req)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *SGKMS) DecryptMultipleFile(keyId, OutputDir string, encryptedFiles []string) (string, error) {

	if err := s.ensureSession(); err != nil {
		return "", err
	}

	req := cryptography.DecryptMultipleFileReq{
		Init: cryptography.Init{
			BaseURL:      s.baseUrl,
			SessionToken: s.SessionToken,
			SlotID:       s.slotID,
		},
		KeyID:          keyId,
		EncryptedFiles: encryptedFiles,
		OutputDir:      OutputDir,
	}

	res, err := cryptography.DecryptMultipleFiles(req)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *SGKMS) SealMultipleFile(keyId, OutputDir string, inputFiles []string) (string, error) {

	if err := s.ensureSession(); err != nil {
		return "", err
	}

	req := cryptography.FileReq{
		Init: cryptography.Init{
			BaseURL:      s.baseUrl,
			SessionToken: s.SessionToken,
			SlotID:       s.slotID,
		},
		KeyID:      keyId,
		InputFiles: inputFiles,
		OutputDir:  OutputDir,
	}

	res, err := cryptography.SealMultipleFile(req)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *SGKMS) UnsealMultipleFile(keyId, OutputDir string, encryptedFiles []string) (string, error) {

	if err := s.ensureSession(); err != nil {
		return "", err
	}

	req := cryptography.DecryptMultipleFileReq{
		Init: cryptography.Init{
			BaseURL:      s.baseUrl,
			SessionToken: s.SessionToken,
			SlotID:       s.slotID,
		},
		KeyID:          keyId,
		EncryptedFiles: encryptedFiles,
		OutputDir:      OutputDir,
	}

	res, err := cryptography.UnsealMultipleFile(req)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *SGKMS) CompressFiles(keyId, OutputDir string, inputFiles []string) (string, error) {

	if err := s.ensureSession(); err != nil {
		return "", err
	}

	req := cryptography.FileReq{
		Init: cryptography.Init{
			BaseURL:      s.baseUrl,
			SessionToken: s.SessionToken,
			SlotID:       s.slotID,
		},
		KeyID:      keyId,
		InputFiles: inputFiles,
		OutputDir:  OutputDir,
	}

	res, err := cryptography.CompressFiles(req)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *SGKMS) UncompressFiles(keyId, OutputDir string, encryptedFiles []string) (string, error) {

	if err := s.ensureSession(); err != nil {
		return "", err
	}

	req := cryptography.DecryptMultipleFileReq{
		Init: cryptography.Init{
			BaseURL:      s.baseUrl,
			SessionToken: s.SessionToken,
			SlotID:       s.slotID,
		},
		KeyID:          keyId,
		EncryptedFiles: encryptedFiles,
		OutputDir:      OutputDir,
	}

	res, err := cryptography.UncompressFiles(req)
	if err != nil {
		return "", err
	}

	return res, nil
}
