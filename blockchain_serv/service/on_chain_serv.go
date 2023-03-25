package service

import (
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *ServiceSetup) SaveAtc(atc_info Atc) (string, error) {
	eventID := "eventAddAtc"
	reg, notifier := RegitserEvent(t.ChannelClient, t.ChaincodeID, eventID)
	defer t.ChannelClient.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(atc_info)
	if err != nil {
		return "", fmt.Errorf("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addAtc", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.ChannelClient.Execute(req)
	if err != nil {
		return "", err
	}

	err = EventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) FindAtcInfoByID(entityID string) ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryAtcInfoByID", Args: [][]byte{[]byte(entityID)}}
	respone, err := t.ChannelClient.Query(req)
	l := logs.GetLogger()
	l.Println(respone.Payload)
	l.Println(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) FindAtcByPublisher(qs string) ([]byte, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryAtcByPublisher", Args: [][]byte{[]byte(qs)}}
	respone, err := t.ChannelClient.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}