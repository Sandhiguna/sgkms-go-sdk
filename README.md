# SGKMS Go SDK

The SGKMS Go SDK is designed to make it easy for Go developers to use the powerful SGKMS REST API features.

This reference guide is meant for product architects and developers to have a basic understanding and overview of Sandhiguna REST API Service.

This SGKMS Go SDK document is intended for SG-KMS version 1.0.0012-GA.

SGKMS Go SDK is a Software Development Kit specifically designed for application development using the Go programming language. This SDK provides various libraries, functions, and tools that help developers access and interact with the SGKMS service.

The client application must connect to SG-KMS with mutual authentication (mTLS) by using certificate issued by CCEV. Certificate is the agent that represent application identity.


# Getting Started

## Installation

Install sgkms-go-sdk with:

```shellxendit
go get github.com/Sandhiguna/sgkms-go-sdk
```

Place the package inside your project directory and include the following line in your import section:

```golang
import sgkms "github.com/Sandhiguna/sgkms-go-sdk"
```
Setting environment sgkms-go-sdk:

```golang
export SGKMS_SLOT_ID=      //slotId for login
export SGKMS_PASSWORD=     //password for login
export SGKMS_BASE_URL=     //host, example https://127.0.0.1:7008/v1.0
export SGKMS_CERT_PATH=    //path certificate with format .pem
export SGKMS_KEY_PATH=     //path certificate with format .key
```

### Example
```go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/Sandhiguna/sgkms-go-sdk"
	"github.com/Sandhiguna/sgkms-go-sdk/cryptography"
)

func main() {
    slotIDStr := os.Getenv("SGKMS_SLOT_ID")
    slotID, _ := strconv.Atoi(slotIDStr)
	password := os.Getenv("SGKMS_PASSWORD")
	baseURL  := os.Getenv("SGKMS_BASE_URL")
	certPath := os.Getenv("SGKMS_CERT_PATH")
	keyPath  := os.Getenv("SGKMS_KEY_PATH")
    keyRsa   :=  os.Getenv("KEY_RSA")
    
	user, err := sgkms.New(certPath, keyPath, password, baseURL, slotID, true)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}

     //Encrypt Asimetric Algorithm RSA
    plaintextRsa := []cryptography.PlaintextEncrypt{
		{
			Text: "Budi Setiawan",
		}, {
			Text: "JL Kaliurang km 10",
		},
	}
	encryptRsa, err := user.Encrypt(keyRsa, plaintextRsa)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Encrypt RSA: %+v\n", encryptRsa)

    //Decrypt Asimmetric Algorithm RSA
	decryptRsa, err := user.Decrypt(keyRsa, nil, encryptRsa.Result.Ciphertext)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Decrypt RSA: %+v\n", decryptRsa)
}
```

# Documentation SGKMS Go SDK

A detailed explanation of how to use SGKMS Go SDK is provided in the list below,

* [Encrypt & Decrypt](docs/Encrypt.md)
