package php

import "encoding/base64"

// Base64Encode base64_encode()
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64Decode base64_decode()
func Base64Decode(data string) (string, error) {
	switch len(data) % 4 {
	case 2:
		data += "=="
	case 3:
		data += "="
	}

	result, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
