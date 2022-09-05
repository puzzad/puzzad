package puzzad

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GetHash(plain string) (string, error) {
	saltLen := 32
	time := 1
	memory := 64 * 1024
	threads := 4
	keyLen := 64
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(plain), salt, uint32(time), uint32(memory), uint8(threads), uint32(keyLen))
	return fmt.Sprintf("argon2id$%d$%d$%d$%s$%s", memory, time, threads, base64.RawStdEncoding.EncodeToString(salt), base64.RawStdEncoding.EncodeToString(hash)), nil
}

func CheckHash(plain, hash string) (bool, error) {
	parts := strings.Split(hash, "$")
	if len(parts) != 6 {
		return false, fmt.Errorf("unable to parse hash")
	}
	if parts[0] != "argon2id" {
		return false, fmt.Errorf("unknown hash type")
	}
	memory, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, err
	}
	time, err := strconv.Atoi(parts[2])
	if err != nil {
		return false, err
	}
	threads, err := strconv.Atoi(parts[3])
	if err != nil {
		return false, err
	}
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}
	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	hashToCompare := argon2.IDKey([]byte(plain), salt, uint32(time), uint32(memory), uint8(threads), uint32(len(decodedHash)))
	return subtle.ConstantTimeCompare(decodedHash, hashToCompare) == 1, nil
}
