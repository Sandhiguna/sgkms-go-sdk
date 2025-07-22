package common

import (
	"crypto/tls"
	"net/http"
)

var sharedHTTPClient *http.Client

// InitTLSClient membuat client TLS dari file .pem dan .key
func InitTLSClient(certPath, keyPath string, insecure bool) (*http.Client, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: insecure,
	}

	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: tlsConfig},
	}

	return client, nil
}

// Gunakan ini untuk menyimpan 1 instance client reusable (opsional)
func SetSharedHTTPClient(client *http.Client) {
	sharedHTTPClient = client
}

func GetSharedHTTPClient() *http.Client {
	return sharedHTTPClient
}
