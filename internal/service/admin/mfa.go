// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"time"
)

// GenerateMFASecret 生成MFA密钥(20字节随机数->base32编码)
func GenerateMFASecret() (string, error) {
	b := make([]byte, 20)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b), nil
}

// GenerateTOTP 根据密钥和时间生成6位验证码
func GenerateTOTP(secret string, t time.Time) (string, error) {
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(strings.ToUpper(secret))
	if err != nil {
		return "", fmt.Errorf("invalid secret: %v", err)
	}
	counter := t.Unix() / 30
	return hotp(key, counter, 6), nil
}

// VerifyTOTP 验证TOTP验证码(允许±1时间步长)
func VerifyTOTP(secret, code string) (bool, error) {
	now := time.Now()
	for i := -1; i <= 1; i++ {
		generated, err := GenerateTOTP(secret, now.Add(time.Duration(i)*30*time.Second))
		if err != nil {
			return false, err
		}
		if generated == code {
			return true, nil
		}
	}
	return false, nil
}

// GetTOTPURI 生成用于Google Authenticator等应用扫码的URI
func GetTOTPURI(secret, accountName, issuer string) string {
	return fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s&algorithm=SHA1&digits=6&period=30",
		issuer, accountName, secret, issuer)
}

func hotp(key []byte, counter int64, digits int) string {
	h := hmac.New(sha1.New, key)
	binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)

	offset := sum[len(sum)-1] & 0x0f
	binary := binary.BigEndian.Uint32(sum[offset:offset+4])
	otp := (binary & 0x7fffffff) % uint32(math.Pow10(digits))
	return fmt.Sprintf("%0*d", digits, otp)
}
