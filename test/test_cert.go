package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	////公钥解析
	//fmt.Println("================解析公钥================")
	//pkFile, err := os.Open("./ecpublic.pem")
	//handleError(err)
	//fileinfo, err := pkFile.Stat()
	//handleError(err)
	//pkBytes := make([]byte, fileinfo.Size())
	//n, err := pkFile.Read(pkBytes)
	//handleError(err)
	//fmt.Printf("公钥文件内容为：\n%s\n", pkBytes[:n])
	////解析pem文件
	//pkBlock, _ := pem.Decode(pkBytes)
	////解析公钥
	//pk, err := x509.ParsePKIXPublicKey(pkBlock.Bytes)
	//handleError(err)
	//switch pub := pk.(type) {
	//case *ecdsa.PublicKey:
	//	fmt.Printf("ecdsa公钥X：%d\n", pub.X)
	//	fmt.Printf("公钥X长度：%d bit\n", pub.X.BitLen())
	//default:
	//	panic("公钥解析错误！")
	//}
	//pkFile.Close()
	//fmt.Println("================解析公钥成功！================")

	//time.Sleep(1 * time.Second)
	//fmt.Println()

	//证书解析
	fmt.Println("================解析证书================")
	certFile, err := os.Open("/root/go/src/atc/fabric/crypto-config/peerOrganizations/eastchina.atc.com/users/User1@eastchina.atc.com/tls/client.crt")
	handleError(err)
	fileinfo, err := certFile.Stat()
	handleError(err)
	certBytes := make([]byte, fileinfo.Size())
	n, err := certFile.Read(certBytes)
	handleError(err)
	fmt.Printf("证书内容为：\n%s\n", certBytes[:n])
	//解析pem文件
	certBlock, _ := pem.Decode(certBytes)
	//解析证书
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	handleError(err)
	fmt.Println(cert.Issuer)
	switch pub := cert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		fmt.Printf("ecdsa公钥X：%d\n", pub.X)
		fmt.Printf("公钥X长度：%d bit\n", pub.X.BitLen())
	default:
		panic("公钥解析错误！")
	}
	certFile.Close()
	fmt.Println("================解析证书成功！================")
}
