package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"go.zenithar.org/kingdom/internal/services/internal/value"
)

// cbc implements encryption at rest of the provided values given a cipher.Block algorithm.
type cbc struct {
	block cipher.Block
}

// NewCBCTransformer takes the given block cipher and performs encryption and decryption on the given
// data.
func NewCBCTransformer(block cipher.Block) value.Transformer {
	return &cbc{block: block}
}

var (
	errInvalidBlockSize    = fmt.Errorf("the stored data is not a multiple of the block size")
	errInvalidPKCS7Data    = errors.New("invalid PKCS7 data (empty or not padded)")
	errInvalidPKCS7Padding = errors.New("invalid padding on input")
)

func (t *cbc) TransformFromStorage(data []byte) ([]byte, bool, error) {
	blockSize := aes.BlockSize
	if len(data) < blockSize {
		return nil, false, fmt.Errorf("the stored data was shorter than the required size")
	}
	iv := data[:blockSize]
	data = data[blockSize:]

	if len(data)%blockSize != 0 {
		return nil, false, errInvalidBlockSize
	}

	result := make([]byte, len(data))
	copy(result, data)
	mode := cipher.NewCBCDecrypter(t.block, iv)
	mode.CryptBlocks(result, result)

	// remove and verify PKCS#7 padding for CBC
	c := result[len(result)-1]
	paddingSize := int(c)
	size := len(result) - paddingSize
	if paddingSize == 0 || paddingSize > len(result) {
		return nil, false, errInvalidPKCS7Data
	}
	for i := 0; i < paddingSize; i++ {
		if result[size+i] != c {
			return nil, false, errInvalidPKCS7Padding
		}
	}

	return result[:size], false, nil
}

func (t *cbc) TransformToStorage(data []byte) ([]byte, error) {
	blockSize := aes.BlockSize
	paddingSize := blockSize - (len(data) % blockSize)
	result := make([]byte, blockSize+len(data)+paddingSize)
	iv := result[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("unable to read sufficient random bytes")
	}
	copy(result[blockSize:], data)

	// add PKCS#7 padding for CBC
	copy(result[blockSize+len(data):], bytes.Repeat([]byte{byte(paddingSize)}, paddingSize))

	mode := cipher.NewCBCEncrypter(t.block, iv)
	mode.CryptBlocks(result[blockSize:], result[blockSize:])
	return result, nil
}
