package helper

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"
)

// InitPaths: Initializes the paths provided if they dont exist
func InitPaths(paths []string) error {
	for _, p := range paths {
		if filepath.Ext(p) == "" {
			if err := os.MkdirAll(p, 0777); err != nil {
				return err
			}
		} else {
			dir := filepath.Dir(p)
			if err := os.MkdirAll(dir, 0777); err != nil {
				return err
			}
			if _, err := os.Stat(p); os.IsNotExist(err) {
				f, err := os.Create(p)
				if err != nil {
					return err
				}
				f.Close()
			}
		}
	}
	return nil
}

// GenerateCryptoID creates a random 32-character hexadecimal string.
// It generates 16 bytes of random data and encodes it in hexadecimal format.
func GenerateCryptoID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
