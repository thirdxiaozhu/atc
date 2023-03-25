package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	str := "from A to B 500RMB"
	base64Sig, _ := RSASign([]byte(str), "/root/go/src/atc/fabric/crypto-config/peerOrganizations/eastchina.atc.com/users/User1@eastchina.atc.com/tls/client.key")
	fmt.Println("签名后的信息：", base64Sig)

	err := RSAVerify([]byte(str), base64Sig, "/root/go/src/atc/fabric/crypto-config/peerOrganizations/eastchina.atc.com/users/User1@eastchina.atc.com/tls/client.crt")
	if err == nil {
		fmt.Println("验证成功！")
	} else {
		fmt.Println("验证失败！")
	}

}

// 私钥签名过程
func RSASign(data []byte, filename string) (string, error) {
	// 1、选择hash算法，对需要签名的数据进行hash运算
	myhash := crypto.SHA256
	hashInstance := myhash.New()
	hashInstance.Write(data)
	hashed := hashInstance.Sum(nil)
	// 2、读取私钥文件，解析出私钥对象
	privateKey, err := ReadParsePrivaterKey(filename)
	if err != nil {
		return "", err
	}
	// 3、RSA数字签名（参数是随机数、私钥对象、哈希类型、签名文件的哈希串），生成base64编码的签名字符串
	bytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, myhash, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// 公钥验证签名过程
func RSAVerify(data []byte, base64Sig, filename string) error {
	// 1、对base64编码的签名内容进行解码，返回签名字节
	bytes, err := base64.StdEncoding.DecodeString(base64Sig)
	if err != nil {
		return err
	}
	// 2、选择hash算法，对需要签名的数据进行hash运算
	myhash := crypto.SHA256
	hashInstance := myhash.New()
	hashInstance.Write(data)
	hashed := hashInstance.Sum(nil)
	// 3、读取公钥文件，解析出公钥对象
	publicKey, err := ReadParsePublicKey(filename)
	if err != nil {
		return err
	}
	// 4、RSA验证数字签名（参数是公钥对象、哈希类型、签名文件的哈希串、签名后的字节）
	return rsa.VerifyPKCS1v15(publicKey, myhash, hashed, bytes)
}

// 读取公钥文件，解析出公钥对象
func ReadParsePublicKey(filename string) (*rsa.PublicKey, error) {
	// 1、读取公钥文件，获取公钥字节
	publicKeyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// 2、解码公钥字节，生成加密块对象
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return nil, errors.New("公钥信息错误！")
	}
	// 3、解析DER编码的公钥，生成公钥接口
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 4、公钥接口转型成公钥对象
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey, nil
}

// 读取私钥文件，解析出私钥对象
func ReadParsePrivaterKey(filename string) (*rsa.PrivateKey, error) {
	// 1、读取私钥文件，获取私钥字节
	privateKeyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// 2、对私钥文件进行编码，生成加密块对象
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	// 3、解析DER编码的私钥，生成私钥对象
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
