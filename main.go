package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

func decryptBlob(hexInput string, key string) (string, error) { //this function performs XOR + NOT decryption on a hex-encoded string using a given key. there is no need to modify this function.
	data, err := hex.DecodeString(hexInput)
	if err != nil {
		return "", fmt.Errorf("Invalid hex input: %v", err)
	}

	keyBytes := []byte(key)
	keyAlt := byte(len(keyBytes))
	for _, b := range keyBytes {
		keyAlt ^= b
	}

	var result []byte
	for _, b := range data {
		decrypted := ^(b ^ keyAlt)
		result = append(result, decrypted)
	}

	return string(result), nil
}

type BlobEntry struct {
	Label    string `json:"label"`
	HexInput string `json:"hex_input"`
	Decoded  string `json:"decoded"`
}

func main() {
	blobs := []string{ //you should be able to add as many encoded strings as you need, just put them in "quotes" and end with a comma ,. Remove the "0x"...
		"ENCODEDSTRING",
		"ENCODEDSTRING",
		"ENCODEDSTRING",
	}

	key := "DECRYPTIONKEY" //enter the decryption key
	var entries []BlobEntry

	for i, blob := range blobs {
		decoded, err := decryptBlob(blob, key)
		if err != nil {
			fmt.Printf("Blob %d: error - %v\n", i+1, err)
			continue
		}

		entry := BlobEntry{
			Label:    fmt.Sprintf("blob_%d", i+1),
			HexInput: blob,
			Decoded:  decoded,
		}
		entries = append(entries, entry)
	}

	outputFile := "decrypted_blobs.json"
	jsonData, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %v\n", err)
		return
	}

	if err := os.WriteFile(outputFile, jsonData, 0644); err != nil {
		fmt.Printf("Failed to write file: %v\n", err)
		return
	}

	fmt.Printf("All blobs written to %s\n", outputFile)
}
