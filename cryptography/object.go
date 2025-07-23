package cryptography

type Init struct {
	BaseURL string `json:"baseurl"`
	SessionToken string        `json:"sessionToken"`
	SlotID       int           `json:"slotId"`
}


type PlaintextEncrypt struct {
	Text string `json:"text"`
	AAD  string `json:"aad,omitempty"`
}

type EncryptReq struct {
	SessionToken string             `json:"sessionToken"`
	SlotID       int                `json:"slotId"`
	KeyID        string             `json:"keyId"`
	Plaintext    []PlaintextEncrypt `json:"plaintext"`
}

type Ciphertext struct {
	Text string `json:"text"`
	AAD  string `json:"aad,omitempty"`
	MAC  string `json:"mac,omitempty"`
	IV   string `json:"iv,omitempty"`
}

type EncryptResult struct {
	KeyVersion int              `json:"keyVersion"`
	Ciphertext []Ciphertext `json:"ciphertext"`
}

type EncryptRes struct {
	Result EncryptResult `json:"result"`
}

// type CipherBlock struct {
// 	Text string `json:"text"`
// 	AAD  string `json:"aad"`
// 	MAC  string `json:"mac"`
// 	IV   string `json:"iv"`
// }

type DecryptReq struct {
	SessionToken string        `json:"sessionToken"`
	SlotID       int           `json:"slotId"`
	KeyID        string        `json:"keyId"`
	KeyVersion   *int           `json:"keyVersion,omitempty"`
	Ciphertext   []Ciphertext `json:"ciphertext"`
}

type DecryptRes struct {
	Result struct {
		Plaintext []string `json:"plaintext"`
	} `json:"result"`
}

type SealReq struct {
	Init
	SessionToken string   `json:"sessionToken"`
	SlotID       int      `json:"slotId"`
	KeyID        string   `json:"keyId"`
	Plaintext    []string `json:"plaintext"`
}

type SealRes struct {
	Result struct {
		Ciphertext []string `json:"ciphertext"`
	} `json:"result"`
}

type SealConfig struct {
	BaseURL      string
	CertPath     string
	KeyPath      string
	InsecureSkip bool
}

type UnsealRes struct {
	Result struct {
		Plaintext []string `json:"plaintext"`
	} `json:"result"`
}

type UnsealReq struct {
	Init
	Ciphertext []string `json:"ciphertext"`
}

// type EncryptFileReq struct {
// 	Init
// 	KeyId string `json:"keyId"`
// 	InputFiles []string `json:"inputFiles"`
// 	OutputDir string `json:"otputDir"`
// }

type EncryptedFileJSON struct {
	Input      string   `json:"input"`
	Output     string   `json:"output"`
	SealedKeys []string `json:"sealed_keys"`
}

type UsealRequest struct {
	SessionToken string   `json:"sessionToken"`
	SlotID       int      `json:"slotId"`
	KeyID        string   `json:"keyId"`
	Ciphertext   []string `json:"cipertext"`
}

type DecryptedFileJSON struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type DecryptMultipleFileReq struct {
	Init
	KeyID          string   `json:"keyId"`
	EncryptedFiles []string `json:"encryptedFiles"`
	OutputDir      string   `json:"outputdir"`
}

type FileReq struct {
	Init
	KeyID      string   `json:"keyId"`
	InputFiles []string `json:"inputFiles"`
	OutputDir  string   `json:"outputDir"`
}