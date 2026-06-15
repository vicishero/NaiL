// Copyright 2026 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web3

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrInvalidSignature = errors.New("invalid signature")
	ErrInvalidAddress   = errors.New("invalid address")
	ErrNonceMismatch    = errors.New("nonce mismatch")
)

// VerifySignature 验证以太坊签名
// message: 签名的原始消息
// signature: 签名（hex 格式，带 0x 前缀或不带）
// address: 签名者的钱包地址（hex 格式，带 0x 前缀或不带）
func VerifySignature(message, signature, address string) (bool, error) {
	log.Printf("VerifySignature - message=%s, address=%s", message, address)

	// 处理签名，去掉 0x 前缀
	signature = strings.TrimPrefix(signature, "0x")

	// 解码签名
	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		log.Printf("VerifySignature - decode signature failed: %v", err)
		return false, fmt.Errorf("decode signature failed: %w", err)
	}

	// 签名长度应为 65 字节 (r + s + v)
	if len(sigBytes) != 65 {
		log.Printf("VerifySignature - invalid signature length: %d", len(sigBytes))
		return false, ErrInvalidSignature
	}

	log.Printf("VerifySignature - sigBytes v value: %d", sigBytes[64])

	// 调整 v 值（EIP-155 兼容）
	if sigBytes[64] >= 27 {
		sigBytes[64] -= 27
	}

	log.Printf("VerifySignature - adjusted v value: %d", sigBytes[64])

	// 构造以太坊签名消息格式
	// Ethereum signed message: "\x19Ethereum Signed Message:\n" + len(message) + message
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	msgHash := crypto.Keccak256Hash([]byte(prefixedMessage))

	log.Printf("VerifySignature - prefixedMessage: %q", prefixedMessage)
	log.Printf("VerifySignature - msgHash: %x", msgHash.Bytes())

	// 恢复公钥
	pubKey, err := crypto.SigToPub(msgHash.Bytes(), sigBytes)
	if err != nil {
		log.Printf("VerifySignature - recover public key failed: %v", err)
		return false, fmt.Errorf("recover public key failed: %w", err)
	}

	// 从公钥派生地址
	recoveredAddress := crypto.PubkeyToAddress(*pubKey)

	log.Printf("VerifySignature - recoveredAddress: %s, inputAddress: %s", recoveredAddress.Hex(), address)

	// 标准化输入地址
	inputAddress := common.HexToAddress(address)

	// 比较地址
	if !strings.EqualFold(recoveredAddress.Hex(), inputAddress.Hex()) {
		log.Printf("VerifySignature - address mismatch")
		return false, nil
	}

	log.Printf("VerifySignature - signature valid")
	return true, nil
}

// GenerateNonce 生成随机 nonce
func GenerateNonce() string {
	randBytes, err := crypto.GenerateKey()
	if err != nil {
		// fallback 到简单的随机数
		return hex.EncodeToString(crypto.Keccak256([]byte(big.NewInt(0).String())))
	}
	pubKey := randBytes.Public()
	ecdsaPubKey, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return hex.EncodeToString(crypto.Keccak256([]byte(big.NewInt(0).String())))
	}
	addr := crypto.PubkeyToAddress(*ecdsaPubKey)
	return hex.EncodeToString(crypto.Keccak256(addr.Bytes()))
}

// GetNonceMessage 构造用于签名的消息
func GetNonceMessage(nonce string) string {
	return fmt.Sprintf("PaoPao Login: %s", nonce)
}
