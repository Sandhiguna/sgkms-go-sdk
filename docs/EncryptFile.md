# Encrypt File

## `EncryptMultipleFiles()` Function

    Encrypt a plaintext using last key version of AES-256-GCM or RSA key without metadata.


| **Name**            |**Value**                                                   |
|---------------------|-------------------------------------------------------------|
| **Function Name**    | `EncryptMultipleFiles`                                                   |
| **Request Parameters** | `keyId string`, `OutputDir string`, `inputFiles []string` |
| **Return Type**      | `string`                                  |


### Request Parameters - Encrypt Multiple Files

The parameter type used for encrypt multiple files.

|Name | Type | Algorithm |Description |Required |
|-------------|:-------------:|:-------------:|:-------------:|-------------|
|  **keyId** |**string**| **AES/RSA** |   | **required** |
|  **OutputDir** |**string**| **AES/RSA**| | **required** |
|  **inputFiles** |**[]string**| **AES/RSA** |  | **required** |

### Respons Ciphertext - Encrypt Multiple Files

The encryption multiple file response varies depending on the algorithm used.

|Algorithm | response |
|-------------|-------------|
|  **AES** |string| 
|  **RSA** |stirng| 


### Example Encrypt Multiple Files

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

	//Encrypt Multiple Files
	EncryptMultipleFIle, err := user.EncryptMultipleFiles(keyAes, "sgkms-go-sdk/OutputFile/DecryptMultipleFile", []string{
		"sgkms-go-sdk/InputFile/EncryptMultipleFIle/1.jpeg",
		"sgkms-go-sdk/InputFile/EncryptMultipleFIle/2.pdf",
	})
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Encrypt Multiple File: %+v\n", EncryptMultipleFIle)

}
```
## `DecryptMultipleFiles()` Function

    Decrypt a ciphertext using specific key version (AES-256-GCM or RSA) without metadata.


| **Name**            |**Value**                                           |
|---------------------|-------------------------------------------------------------|
| **Function Name**    | `DecryptMultipleFiles`                                                   |
| **Request Parameters** | `keyId string`, `OutputDir string`, `encryptedFiles []string` |
| **Return Type**      | `string`                                  |


### Request Parameters - Decrypt Multiple Files

The parameter type used for decryption multiple files.

|Name | Type | Algorithm |Description |Required |
|-------------|:-------------:|:-------------:|:-------------:|-------------|
|  **keyId** |**string**| **AES/RSA** |   | **required** |
|  **OutputDir** |**string**| **AES/RSA**| | **required** |
|  **encryptedFiles** |**[]string**| **AES/RSA** |  | **required** |

### Respons Ciphertext - Decrypt


|Algorithm | response |
|-------------|-------------|
|  **AES** |string| 
|  **RSA** |stirng| 


### Example Decrypt Multiple File

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

	//Encrypt Multiple Files
	EncryptMultipleFIle, err := user.EncryptMultipleFiles(keyAes, "sgkms-go-sdk/OutputFile/DecryptMultipleFile", []string{
		"sgkms-go-sdk/InputFile/EncryptMultipleFIle/1.jpeg",
		"sgkms-go-sdk/InputFile/EncryptMultipleFIle/2.pdf",
	})
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Decrypt Multiple File: %+v\n", EncryptMultipleFIle)

	//Decrypt Multiple Files
	DecryptMultipleFile, err := user.DecryptMultipleFile(keyAes, "sgkms-go-sdk/OutputFile/DecryptMultipleFile", []string{
		"sgkms-go-sdk/OutputFile/DecryptMultipleFile/1.sgc",
		"sgkms-go-sdk/OutputFile/DecryptMultipleFile/2.sgc",
	})
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Decrypt Multiple File: %+v\n", DecryptMultipleFile)

}
```

## `SealMultipleFile()` Function

    Encrypt a plaintext using last key version of AES-256-GCM or RSA key without metadata.


| **Name**            |**Value**                                                   |
|---------------------|-------------------------------------------------------------|
| **Function Name**    | `SealMultipleFile`                                                   |
| **Request Parameters** | `keyId string`, `OutputDir string`, `inputFiles []string` |
| **Return Type**      | `string`                                  |


### Request Parameters - Seal Multiple Files

The parameter type used for seal multiple files.

|Name | Type | Algorithm |Description |Required |
|-------------|:-------------:|:-------------:|:-------------:|-------------|
|  **keyId** |**string**| **AES/RSA** |   | **required** |
|  **OutputDir** |**string**| **AES/RSA**| | **required** |
|  **inputFiles** |**[]string**| **AES/RSA** |  | **required** |

### Respons Ciphertext - Seal Multiple Files

The seal multiple file response varies depending on the algorithm used.

|Algorithm | response |
|-------------|-------------|
|  **AES** |string| 
|  **RSA** |stirng| 


### Example Seal Multiple Files

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

	//Seal Multiple Files
	SealMultipleFile, err := user.SealMultipleFile(keyAes, "sgkms-go-sdk/OutputFile/DecryptMultipleFile", []string{
		"sgkms-go-sdk/InputFile/EncryptMultipleFIle/1.jpeg",
		"sgkms-go-sdk/InputFile/EncryptMultipleFIle/2.pdf",
	})
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Seal Multiple File: %+v\n", SealMultipleFile)
}
```

## `UnsealMultipleFile()` Function

    Decrypt a ciphertext using specific key version (AES-256-GCM or RSA) without metadata.


| **Name**            |**Value**                                           |
|---------------------|-------------------------------------------------------------|
| **Function Name**    | `UnsealMultipleFile`                                                   |
| **Request Parameters** | `keyId string`, `OutputDir string`, `encryptedFiles []string` |
| **Return Type**      | `string`                                  |


### Request Parameters - Unseal Multiple File

The parameter type used for unseal multiple files.

|Name | Type | Algorithm |Description |Required |
|-------------|:-------------:|:-------------:|:-------------:|-------------|
|  **keyId** |**string**| **AES/RSA** |   | **required** |
|  **OutputDir** |**string**| **AES/RSA**| | **required** |
|  **encryptedFiles** |**[]string**| **AES/RSA** |  | **required** |

### Respons Ciphertext - Unseal Multiple File


|Algorithm | response |
|-------------|-------------|
|  **AES** |string| 
|  **RSA** |stirng| 


### Example Unseal Multiple File

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

	//Seal Multiple Files
	SealMultipleFile, err := user.SealMultipleFile(keyAes, "sgkms-go-sdk/OutputFile/DecryptMultipleFile", []string{
		"sgkms-go-sdk/InputFile/EncryptMultipleFIle/1.jpeg",
		"sgkms-go-sdk/InputFile/EncryptMultipleFIle/2.pdf",
	})
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Unseal Multiple File: %+v\n", SealMultipleFile)

	//Unseal Multiple Files
	unsealSealMultipleFile, err := user.UnsealMultipleFile(keyAes, "sgkms-go-sdk/OutputFile/DecryptMultipleFile", []string{
		"sgkms-go-sdk/OutputFile/DecryptMultipleFile/1.sgc",
		"sgkms-go-sdk/OutputFile/DecryptMultipleFile/2.sgc",
	})
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Unseal Multiple File: %+v\n", unsealSealMultipleFile)
}
```

## `CompressFiles()` Function

    Encrypt a plaintext using last key version of AES-256-GCM or RSA key without metadata.


| **Name**            |**Value**                                                   |
|---------------------|-------------------------------------------------------------|
| **Function Name**    | `CompressFiles`                                                   |
| **Request Parameters** | `keyId string`, `OutputDir string`, `inputFiles []string` |
| **Return Type**      | `string`                                  |


### Request Parameters - Compress Files

The parameter type used for compress files.

|Name | Type | Algorithm |Description |Required |
|-------------|:-------------:|:-------------:|:-------------:|-------------|
|  **keyId** |**string**| **AES/RSA** |   | **required** |
|  **OutputDir** |**string**| **AES/RSA**| | **required** |
|  **inputFiles** |**[]string**| **AES/RSA** |  | **required** |

### Respons Ciphertext - Compress Files

The compress file response varies depending on the algorithm used.

|Algorithm | response |
|-------------|-------------|
|  **AES** |string| 
|  **RSA** |stirng| 


### Example Compress Files

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

	//Compress Files
	compressFile, err := user.CompressFiles(keyAes, "sgkms-go-sdk/OutputFile/DecryptMultipleFile", []string{
		"sgkms-go-sdk/InputFile/EncryptMultipleFIle/1.jpeg",
		"sgkms-go-sdk/InputFile/EncryptMultipleFIle/2.pdf",
	})
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Compress File: %+v\n", compressFile)
}
```

## `UncompressFile()` Function

    Decrypt a ciphertext using specific key version (AES-256-GCM or RSA) without metadata.


| **Name**            |**Value**                                           |
|---------------------|-------------------------------------------------------------|
| **Function Name**    | `UncompressFile`                                                   |
| **Request Parameters** | `keyId string`, `OutputDir string`, `encryptedFiles []string` |
| **Return Type**      | `string`                                  |


### Request Parameters - Uncompress Files

The parameter type used for uncompress files.

|Name | Type | Algorithm |Description |Required |
|-------------|:-------------:|:-------------:|:-------------:|-------------|
|  **keyId** |**string**| **AES/RSA** |   | **required** |
|  **OutputDir** |**string**| **AES/RSA**| | **required** |
|  **encryptedFiles** |**[]string**| **AES/RSA** |  | **required** |

### Respons Ciphertext - Uncompress File


|Algorithm | response |
|-------------|-------------|
|  **AES** |string| 
|  **RSA** |stirng| 


### Example Uncompress File

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

	//Compress Files
	compressFile, err := user.CompressFiles(keyAes, "/home/sandhiguna/Documents/sgkms-go-sdk/OutputFile/DecryptMultipleFile", []string{
		"/home/sandhiguna/Documents/sgkms-go-sdk/InputFile/EncryptMultipleFIle/1.jpeg",
		"/home/sandhiguna/Documents/sgkms-go-sdk/InputFile/EncryptMultipleFIle/2.pdf",
	})
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Compress File: %+v\n", compressFile)

	//Uncompress File
	uncompressFile, err := user.UncompressFiles(keyAes, "/home/sandhiguna/Documents/sgkms-go-sdk/OutputFile/DecryptMultipleFile", []string{
		"/home/sandhiguna/Documents/sgkms-go-sdk/OutputFile/DecryptMultipleFile/20250724_101407.sgc",
	})
	if err != nil {
		log.Fatalf("Error creating SGKMS object: %v", err)
	}
	fmt.Printf("Uncompress File: %+v\n", uncompressFile)
}
```