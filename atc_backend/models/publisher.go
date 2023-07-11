package models

import (
	"encoding/json"
	"fmt"
	"service"
	"strings"

	"github.com/google/uuid"
)

type HistoryItem struct {
	TxId string
	Atc  service.Atc
}

type AtcChaincode struct {
}

type CheckStruct struct {
	Publisher string
	Company   string
	Starttime string
	Endtime   string
	Flight    string
}

type MultiAtcPrase struct {
	Type    string `json:"Type"`
	Message string `json:"Message"`
	Flight  string `json:"Flight"`
}

var paraTags map[int]string

func init() {
	paraTags = map[int]string{0: "publisher", 1: "company", 2: "starttime", 3: "endtime"}
}

func PostAtc(userid, msgtype, content, timestamp, flight string) (string, error) {
	user := User{}
	DBH.QueryOneByField(&user, "user", "name", userid)

	sign := GetSignature(content, cryptpath_map[user.Company]+"client.key")

	address := ""
	on_chain_content := ""
	is_ipfs := false
	var err error

	if len(content) >= 32 {
		address, err = UploadIPFS(content)
		is_ipfs = true
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	} else {
		on_chain_content = content
	}
	u, _ := uuid.NewRandom()

	Atc := &service.Atc{
		ID:        u.String(),
		Time:      timestamp,
		Publisher: userid,
		Company:   user.Company,
		Type:      msgtype,
		Address:   address,
		Signature: sign,
		Content:   on_chain_content,
		Flight:    flight,
		IsIPFS:    is_ipfs,
	}

	transid, err := service.Save(setup_map[userid], *Atc)
	if err != nil {
		logger.Println(err)
		return "", err
	}
	return transid, nil
}

func PostMultiAtc(userid, msgs, timestamp string) ([]string, error) {
	var multimsgs []MultiAtcPrase
	var transids []string

	logger.Print(msgs)

	if err := json.Unmarshal([]byte(msgs), &multimsgs); err != nil {
		logger.Print(err)
		return nil, err
	}

	for _, msg := range multimsgs {
		transid, err := PostAtc(userid, msg.Type, msg.Message, timestamp, msg.Flight)
		if err != nil {
			return nil, err
		}
		transids = append(transids, transid)
	}

	return transids, nil
}

func GetAtcs(userid string, paralist []string) *[]service.Atc {

	atcs_p := searchByCondition(userid, paralist)

	atcs := *atcs_p

	for index, _ := range atcs {
		atc := atcs[index]
		atc.Content, atc.IsValid = VerifyandGetContent(atc)

		for index, _ := range atc.Historys {
			atc_h := atc.Historys[index].Atc
			atc_h.Content, atc_h.IsValid = VerifyandGetContent(atc_h)
			atc.Historys[index].Atc = atc_h
		}
		atcs[index] = atc
	}

	return &atcs
}

func EditAtc(paralist []string) string {

	user := User{}
	DBH.QueryOneByField(&user, "user", "name", paralist[0])

	sign := GetSignature(paralist[1], cryptpath_map[user.Company]+"client.key")
	//address, err := UploadIPFS(paralist[1])
	//if err != nil {
	//	logger.Println(err)
	//	return ""
	//}

	address := ""
	on_chain_content := ""
	is_ipfs := false
	var err error

	if len(paralist[1]) >= 32 {
		address, err = UploadIPFS(paralist[1])
		is_ipfs = true
		if err != nil {
			fmt.Println(err)
			return ""
		}
	} else {
		on_chain_content = paralist[1]
	}

	atc := service.Atc{
		ID:        paralist[2],
		Time:      paralist[3],
		Address:   address,
		Signature: sign,
		Content:   on_chain_content,
		IsIPFS:    is_ipfs,
	}
	transid, err := service.Modify(setup_map[paralist[0]], atc)
	if err != nil {
		logger.Println(err)
		return ""
	}

	return transid
}

func DeleteAtc(userid, atcid string) string {
	transid, err := service.Delete(setup_map[userid], atcid)
	if err != nil {
		logger.Println(err)
		return ""
	}

	return transid

}

func getQueryString(cs CheckStruct) string {
	parastr := "{"

	count := 0

	if cs.Publisher != "" {
		parastr += judgeComma(count)
		parastr += fmt.Sprintf("\"Publisher\": \"%s\"", cs.Publisher)
		count++
	}
	if cs.Company != "" {
		parastr += judgeComma(count)
		parastr += fmt.Sprintf("\"Company\": \"%s\"", cs.Company)
		count++
	}
	if cs.Flight != "" {
		parastr += judgeComma(count)
		parastr += fmt.Sprintf("\"Flight\": \"%s\"", cs.Flight)
		count++
	}

	parastr += "}"

	logger.Println(parastr)

	return parastr
}

func judgeComma(count int) string {
	if count == 0 {
		return ""
	} else {
		return ","
	}
}

func VerifyandGetContent(atc service.Atc) (string, bool) {

	var content string

	if atc.IsIPFS == false {
		content = atc.Content
	} else {
		content = CatIPFS(atc.Address)
	}

	isValid := VerifySignature(atc.Signature, content, cryptpath_map[atc.Company]+"client.crt")

	return content, isValid
}

func searchByCondition(userid string, paralist []string) *[]service.Atc {

	publisher_names := strings.Split(paralist[0], ",")
	company_names := strings.Split(paralist[1], ",")
	start_time := paralist[2]
	end_time := paralist[3]
	flight := paralist[4]

	query_strings := []string{}

	//如果publisher为空，那么根据company字段查找,否则根据publisher查找
	if len(publisher_names) == 1 && publisher_names[0] == "" {
		if len(company_names) == 1 && company_names[0] == "" {
			//从company查找到所有role为0的字段，查询全部
			companies := []Company{}
			//用空数组替换company_names
			company_names = []string{}
			err := DBH.QueryAllByField(&companies, "company", "role", "0")
			if err != nil {
				logger.Println("Cant find companies.")
				return nil
			}
			for _, company := range companies {
				company_names = append(company_names, company.Name)
			}
		}

		for _, company_name := range company_names {
			querystring := getQueryString(CheckStruct{
				Publisher: "",
				Company:   company_name,
				Flight:    flight,
			})

			query_strings = append(query_strings, querystring)
		}

	} else {
		for _, publisher_name := range publisher_names {
			querystring := getQueryString(CheckStruct{
				Publisher: publisher_name,
				Company:   "",
				Flight:    flight,
			})

			query_strings = append(query_strings, querystring)
		}
	}

	atcs := []service.Atc{}
	for _, querystring := range query_strings {
		atcs_p := judgeTimeSpan(service.QueryByString(setup_map[userid], querystring), start_time, end_time)

		atcs = append(atcs, *atcs_p...)
	}

	return &atcs
}

func judgeTimeSpan(atcs_p *[]service.Atc, start, end string) *[]service.Atc {

	var atcs []service.Atc

	if start == "" || end == "" {
		return atcs_p
	}

	for _, atc := range *atcs_p {
		if atc.Time > start && atc.Time < end {
			atcs = append(atcs, atc)
		}
	}

	return &atcs
}
