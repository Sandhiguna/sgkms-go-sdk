# Seal & Unseal

## `Seal()` Function

    Encrypt a plaintext with metadata using AES-256-GCM key or RSA with session key.


| **Name**            |**Value**                                                   |
|---------------------|-------------------------------------------------------------|
| **Function Name**    | `Seal`                                                   |
| **Request Parameters** | `keyId string`, `plaintext []string` |
| **Return Type**      | `*cryptography.SealRes`                                  |

### Request Parameters - Seal

The parameter type used for seal.

|Name | Type | Algorithm |Description |Required |
|-------------|:-------------:|:-------------:|:-------------:|-------------|
|  **keyId** |**string**| **AES/RSA** |   | **required** |
|  **plaintext** |**[]string**| **AES/RSA**| | **required** |


### Respons Ciphertext - Seal

The Seal response varies depending on the algorithm used.

|Algorithm | Plaintext |
|-------------|-------------|
|  **AES** |[]string| 
|  **RSA** |[]stirng| 



### Example Seal

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Sandhiguna/sgkms-go-sdk"
)

func main() {
	slotIDStr := os.Getenv("SGKMS_SLOT_ID")
	slotID, _ := strconv.Atoi(slotIDStr)
	password := os.Getenv("SGKMS_PASSWORD")
	baseURL := os.Getenv("SGKMS_BASE_URL")
	certPath := os.Getenv("SGKMS_CERT_PATH")
	keyPath := os.Getenv("SGKMS_KEY_PATH")
	keyAes := os.Getenv("KEY_AES")

	user, err := sgkms.New(certPath, keyPath, password, baseURL, slotID, true)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}

	//Seal
	seal, err := user.Seal(keyAes, []string{"data encrypt 1", "data encrypt 2"})
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Seal: %+v\n", seal)
}
```

## `Unseal()` Function

    Decrypt a ciphertext with metadata.

| **Name**            |**Value**                                                   |
|---------------------|-------------------------------------------------------------|
| **Function Name**    | `Unseal`                                                   |
| **Request Parameters** | `keyId string`, `ciphertext []string` |
| **Return Type**      | `*cryptography.UnsealRes`                                  |

### Request Parameters - Unseal

The parameter type used for unseal.

|Name | Type | Algorithm |Description |Required |
|-------------|:-------------:|:-------------:|:-------------:|-------------|
|  **keyId** |**string**| **AES/RSA** |   | **required** |
|  **ciphertext** |**[]string**| **AES/RSA**| | **required** |


### Respons Ciphertext - Unseal

The Unseal response varies depending on the algorithm used.

|Algorithm | Ciphertext |
|-------------|-------------|
|  **AES** |[]string| 
|  **RSA** |[]stirng| 



### Example Seal

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Sandhiguna/sgkms-go-sdk"
)

func main() {
	slotIDStr := os.Getenv("SGKMS_SLOT_ID")
	slotID, _ := strconv.Atoi(slotIDStr)
	password := os.Getenv("SGKMS_PASSWORD")
	baseURL := os.Getenv("SGKMS_BASE_URL")
	certPath := os.Getenv("SGKMS_CERT_PATH")
	keyPath := os.Getenv("SGKMS_KEY_PATH")
	keyAes := os.Getenv("KEY_AES")

	user, err := sgkms.New(certPath, keyPath, password, baseURL, slotID, true)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}

	//Seal
	seal, err := user.Seal(keyAes, []string{"data encrypt 1", "data encrypt 2"})
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Unseal: %+v\n", seal)

	//Unseal
	unseal,err := user.Unseal(keyAes, seal.Result.Ciphertext)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Unseal: %+v\n", unseal)
}
```