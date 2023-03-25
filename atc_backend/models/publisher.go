package models

import (
	"fmt"
	"service"
	"time"

	"github.com/google/uuid"
)

type HistoryItem struct {
	TxId string
	Atc  service.Atc
}

type AtcChaincode struct {
}

func PostAtc(userid, msgtype, content string) (string, error) {
	logger.Println(userid, msgtype, content)
	user := User{}
	company := Company{}
	DBH.QueryOneByField(&user, "user", "name", userid)
	company.ID = uint(user.Company)
	DBH.QueryByPK(&company)

	sign := GetSignature(content, cryptpath_map[user.Company]+"client.key")
	u, _ := uuid.NewRandom()

	address, err := UploadIPFS(content)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	Atc := &service.Atc{
		ID:        u.String(),
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		Publisher: userid,
		Company:   company.Name,
		Type:      msgtype,
		Address:   address,
		Signature: sign,
	}
	logger.Println(Atc, setup_map[userid], setup_map)

	transid, err := service.Save(setup_map[userid], *Atc)
	if err != nil {
		logger.Println(err)
		return "", err
	}
	return transid, nil
}
