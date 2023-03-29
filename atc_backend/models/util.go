package models

import (
	"crypto/ecdsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"os"
	"strings"
)

type Cert struct {
	Content   string            `json:"content"`
	Issuermap map[string]string `json:"issmap"`
	X         string            `json:"x"`
	Y         string            `json:"y"`
	Xlength   int               `json:"xlength"`
	Ylength   int               `json:"ylength"`
}

func GetCompaniesByRole(role string) *[]Company {
	companies := []Company{}
	err := DBH.QueryAllByField(&companies, "company", "role", role)
	if err != nil {
		return nil
	}

	return &companies
}

func GetPublishersByCompany(company_name string) *[]User {
	users := []User{}
	err := DBH.QueryAllByField(&users, "user", "company", company_name)
	if err != nil {
		return nil
	}

	return &users
}

func GetCert(user_name string) *Cert {
	user := User{}
	err := DBH.QueryOneByField(&user, "user", "userid", user_name)
	if err != nil {
		return nil
	}
	certpath := cryptpath_map[user.Company]
	pubcertpath := certpath + "client.crt"

	cert_ret := Cert{}

	certFile, err := os.Open(pubcertpath)
	if err != nil {
		return nil
	}
	fileinfo, err := certFile.Stat()
	if err != nil {
		return nil
	}
	certBytes := make([]byte, fileinfo.Size())
	n, err := certFile.Read(certBytes)
	if err != nil {
		return nil
	}
	cert_ret.Content = string(certBytes[:n])
	//解析pem文件
	certBlock, _ := pem.Decode(certBytes)
	//解析证书
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil
	}
	cert_ret.Issuermap = getIssusers(cert.Issuer)
	switch pub := cert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		cert_ret.X = pub.X.String()
		cert_ret.Y = pub.Y.String()
		cert_ret.Xlength = pub.X.BitLen()
		cert_ret.Ylength = pub.Y.BitLen()
	default:
		logger.Print("公钥解析错误！")
		return nil
	}
	certFile.Close()

	return &cert_ret
}

func getIssusers(issuers pkix.Name) map[string]string {
	issues := strings.Split(issuers.String(), ",")

	kv_map := map[string]string{}
	for _, issue := range issues {
		key_value := strings.Split(issue, "=")
		kv_map[key_value[0]] = key_value[1]
	}

	return kv_map
}

func GetPrivateKey(user_name string) *[]byte {
	user := User{}
	err := DBH.QueryOneByField(&user, "user", "userid", user_name)
	if err != nil {
		return nil
	}
	certpath := cryptpath_map[user.Company]
	privpath := certpath + "client.key"

	privateKeyBytes, err := ioutil.ReadFile(privpath)

	return &privateKeyBytes
}
