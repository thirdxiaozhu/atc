package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
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

// 生成密钥对
func generateECCKeyPair(c elliptic.Curve, privateKeyFile string, publicKeyFile string) error {
	//使用ECDSA生成密钥对
	privateKey, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		return err
	}
	//使用509
	private, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}
	//pem
	block := pem.Block{
		Type:  "esdsa private key",
		Bytes: private,
	}
	file, err := os.Create(privateKeyFile)
	if err != nil {
		return err
	}
	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}
	file.Close()

	//处理公钥
	public := privateKey.PublicKey

	//x509序列化
	publicKey, err := x509.MarshalPKIXPublicKey(&public)
	if err != nil {
		return err
	}
	//pem
	public_block := pem.Block{
		Type:  "ecdsa public key",
		Bytes: publicKey,
	}
	file, err = os.Create(publicKeyFile)
	if err != nil {
		return err
	}
	//pem编码
	err = pem.Encode(file, &public_block)
	if err != nil {
		return err
	}
	return nil
}

func DERToPrivateKey(der []byte) (key interface{}, err error) {
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

// ECC签名--私钥
func ECCSignature(sourceData []byte, privateKeyFile string) []byte {
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
	privateKey, err := DERToPrivateKey(block.Bytes)
	//哈希运算
	hashText := sha1.Sum(sourceData)
	//数字签名
	signa, err := ecdsa.SignASN1(rand.Reader, privateKey.(*ecdsa.PrivateKey), hashText[:])
	return signa
}

func GetPubKeyFromX509(certPath string) (*ecdsa.PublicKey, error) {
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

func ECCVerify(signa []byte, sourceData []byte, certPath string) bool {
	publicKey, err := GetPubKeyFromX509(certPath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	hashText := sha1.Sum(sourceData)
	res := ecdsa.VerifyASN1(publicKey, hashText[:], signa)
	return res
}

// ECC认证
//func ECCVerify(rText, sText, sourceData []byte, publicKeyFilePath string) bool {
//	//读取公钥文件
//	file, err := os.Open(publicKeyFilePath)
//	if err != nil {
//		panic(err)
//	}
//	info, err := file.Stat()
//	if err != nil {
//		panic(err)
//	}
//	buf := make([]byte, info.Size())
//	file.Read(buf)
//	//pem解码
//	block, _ := pem.Decode(buf)
//
//	//x509
//	publicStream, err := x509.ParsePKIXPublicKey(block.Bytes)
//	if err != nil {
//		panic(err)
//	}
//	//接口转换成公钥
//	publicKey := publicStream.(*ecdsa.PublicKey)
//	hashText := sha1.Sum(sourceData)
//	var r, s big.Int
//	r.UnmarshalText(rText)
//	s.UnmarshalText(sText)
//	//认证
//	res := ecdsa.Verify(publicKey, hashText[:], &r, &s)
//	defer file.Close()
//	return res
//}

func main() {
	//generateECCKeyPair(elliptic.P256(), "./private.pem", "./public.pem")
	rawData := []byte("Hello go")
	//r, s := ECCSignature(rawData, "/root/go/src/atc/fabric/crypto-config/peerOrganizations/eastchina.atc.com/users/User1@eastchina.atc.com/tls/client.key")
	signa := ECCSignature(rawData, "/root/go/src/atc/fabric/crypto-config/peerOrganizations/eastchina.atc.com/users/User1@eastchina.atc.com/tls/client.key")

	encoded := base64.StdEncoding.EncodeToString(signa)
	fmt.Println(encoded)
	// SGVsbG8gd29ybGQuIOS9oOWlve+8jOS4lueVjO+8gQ==

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))

	fmt.Println(decoded)
	//res := ECCVerify(r, s, rawData, "./public.pem")
	res := ECCVerify(decoded, rawData, "/root/go/src/atc/fabric/crypto-config/peerOrganizations/eastchina.atc.com/users/User1@eastchina.atc.com/tls/client.crt")
	if res {
		fmt.Println("Signature is Accepted")
	}

}
