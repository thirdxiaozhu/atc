package models

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func GetSignature(sourceString string, privateKeyFile string) string {

	rawData := []byte(sourceString)
	//读取私钥
	file, err := os.Open(privateKeyFile)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解密
	block, _ := pem.Decode(buf)
	//x509解密
	privateKey, err := derToPrivateKey(block.Bytes)
	//哈希运算
	hashText := sha1.Sum(rawData)
	//数字签名
	sign_bytes, err := ecdsa.SignASN1(rand.Reader, privateKey.(*ecdsa.PrivateKey), hashText[:])
	return base64.StdEncoding.EncodeToString(sign_bytes)
}

func VerifySignature(sign_string string, sourceData []byte, certPath string) bool {

	decoded, _ := base64.StdEncoding.DecodeString(sign_string)

	publicKey, err := getPubKeyFromX509(certPath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	hashText := sha1.Sum(sourceData)
	res := ecdsa.VerifyASN1(publicKey, hashText[:], decoded)
	return res
}

func derToPrivateKey(der []byte) (key interface{}, err error) {
	if key, err = x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}

	if key, err = x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey:
			return
		default:
			return nil, errors.New("Found unknown private key type in PKCS#8 wrapping")
		}
	}

	if key, err = x509.ParseECPrivateKey(der); err == nil {
		return
	}

	return nil, errors.New("Invalid key type. The DER must contain an rsa.PrivateKey or ecdsa.PrivateKey")
}

func getPubKeyFromX509(certPath string) (*ecdsa.PublicKey, error) {
	certFile, err := os.Open(certPath)
	if err != nil {
		return nil, err
	}
	fileinfo, err := certFile.Stat()
	if err != nil {
		return nil, err
	}
	certBytes := make([]byte, fileinfo.Size())
	n, err := certFile.Read(certBytes)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("证书内容为：\n%s\n", certBytes[:n])
	certBlock, _ := pem.Decode(certBytes[:n])
	//解析证书
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, err
	}
	certFile.Close()
	switch pub := cert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		return pub, nil
	default:
		fmt.Println("公钥解析错误！")
		return nil, nil
	}
}
