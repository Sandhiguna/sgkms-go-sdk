# Encrypt & Decrypt

## `Encrypt()` Function

    Encrypt a plaintext using last key version of AES-256-GCM or RSA key without metadata.


| **Name**            |**Value**                                                   |
|---------------------|-------------------------------------------------------------|
| **Function Name**    | `Encrypt`                                                   |
| **Request Parameters** | `keyId string`, `plaintext []cryptography.PlaintextEncrypt` |
| **Return Type**      | `*cryptography.EncryptRes`                                  |


### Request Parameters - Encrypt

The parameter type used for encryption.

|Name | Type | Algorithm |Description |Required |
|-------------|:-------------:|:-------------:|:-------------:|-------------|
|  **keyId** |**string**| **AES/RSA** |   | **required** |
|  **plaintext** |**[]cryptography.PlaintextEncrypt**| **AES**| **Text, AAD(optional)** | **required** |
|  **plaintext** |**[]cryptography.PlaintextEncrypt**| **RSA** | **Text** | **required** |

### Respons Ciphertext - Encrypt

The encryption response varies depending on the algorithm used.

|Algorithm | Text | AAD |MAC |IV |
|-------------|:-------------:|:-------------:|:-------------:|-------------|
|  **AES** |**available**|  **not available**| **available**  | **available** |
|  **RSA** |**available**| **not available**| **not available** | **not available** |


### Example Encrypt

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
	keyAes   :=  os.Getenv("KEY_AES")
	keyRsa   :=  os.Getenv("KEY_RSA")

	user, err := sgkms.New(certPath, keyPath, password, baseURL, slotID, true)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}

	//Symmetric Encryption Algorithm AES
	plaintextAes := []cryptography.PlaintextEncrypt{
		{
			Text: "Budi Setiawan",
			AAD:  "nama", //optional
		}, {
			Text: "JL Kaliurang km 10",
			AAD:  "alamat", //optional
		},
	}
	encryptAES, err := user.Encrypt(keyAes, plaintextAes)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Encrypt AES: %+v\n", encryptAES)

	//Asimmetric Encrypt Algorithm RSA
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
}
```

## `Decrypt()` Function

    Decrypt a ciphertext using specific key version (AES-256-GCM or RSA) without metadata.


| **Name**            |**Value**                                                   |
|---------------------|-------------------------------------------------------------|
| **Function Name**    | `Decrypt`                                                   |
| **Request Parameters** | `keyId string`, `keyVersion *int`, `Ciphertext []cryptography.Ciphertext` |
| **Return Type**      | `*cryptography.DecryptRes, error`                                  |


### Request Parameters - Decrypt

The parameter type used for decryption

|Name | Type | Algorithm |Description |Required |
|-------------|:-------------:|:-------------:|:-------------:|-------------|
|  **keyId** |**string**| **AES/RSA** |   | **required** |
|  **keyVersion** |**int**| **AES** |   | **required** |
|  **keyVersion** |**nil**| **RSA** |   ||
|  **ciphertext** |**[]cryptography.Ciphertext**| **AES**| **Text, AAD, MAC, IV** | **required** |
|  **ciphertext** |**[]cryptography.Ciphertext**| **RSA** | **Text** | **required** |

### Respons Ciphertext - Decrypt


|Algorithm | Ciphertext |
|-------------|-------------|
|  **AES** |[]string| 
|  **RSA** |[]stirng| 


### Example Decrypt

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
	keyAes   :=  os.Getenv("KEY_AES")
	keyRsa   :=  os.Getenv("KEY_RSA")

	user, err := sgkms.New(certPath, keyPath, password, baseURL, slotID, true)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}

	//Symmetric Encryption Algorithm AES
	plaintextAes := []cryptography.PlaintextEncrypt{
		{
			Text: "Budi Setiawan",
			AAD:  "nama", //optional
		}, {
			Text: "JL Kaliurang km 10",
			AAD:  "alamat", //optional
		},
	}
	encrypt, err := user.Encrypt(keyAes, plaintextAes)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Encrypt AES: %+v\n", encrypt)

	ciphertextAes := encrypt.Result.Ciphertext
	for i := range ciphertextAes {
		if i < len(plaintextAes) {
			ciphertextAes[i].AAD = plaintextAes[i].AAD
		}
	}
	decrypt, err := user.Decrypt(keyAes, &encrypt.Result.KeyVersion, ciphertextAes)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Decrypt AES: %+v\n", decrypt)


	//Asimmetric Encrypt Algorithm RSA
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

	decryptRsa, err := user.Decrypt(keyRsa, nil, encryptRsa.Result.Ciphertext)
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Decrypt RSA: %+v\n", decryptRsa)
	}
```