package service

import (
	"fmt"
	"log"
	"sdk"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type Atc struct {
	//ObjectType string `json:"docType"`
	ID        string `json:"ID"`
	Time      string `json:"Time"`
	Publisher string `json:"Publisher"`
	Company   string `json:"Company"`
	Type      string `json:"Type"`
	Address   string `json:"Address"`
	Signature string `json:"Signature"`

	Historys []HistoryItem // 当前edu的历史记录
}

type HistoryItem struct {
	TxId string
	Atc  Atc
}

type ServiceSetup struct {
	ChaincodeID   string
	ChannelClient *channel.Client
	LedgerClient  *ledger.Client
}

var (
	info   sdk.SdkEnvInfo
	logger *log.Logger
)

func init() {
	cc_name := "atc_chaincode"
	cc_version := "1.0.0"

	orgs := []*sdk.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "EastChina",
			OrgMspId:      "EastChinaMSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: "/root/go/src/atc/fabric/channel-artifacts/EastChinaMSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "NorthChina",
			OrgMspId:      "NorthChinaMSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: "/root/go/src/atc/fabric/channel-artifacts/NorthChinaMSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "AirChina",
			OrgMspId:      "AirChinaMSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: "/root/go/src/atc/fabric/channel-artifacts/AirChinaMSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Csair",
			OrgMspId:      "CsairMSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: "/root/go/src/atc/fabric/channel-artifacts/CsairMSPanchors.tx",
		},
	}

	info = sdk.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    "/root/go/src/atc/fabric/channel-artifacts/channel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "ATCOrderer",
		OrdererEndpoint:  "orderer.atc.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    "/root/go/src/atc/chaincode/",
		ChaincodeVersion: cc_version,
	}

	logger = logs.GetLogger()
}

func RegitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败:", err)
	}
	return reg, notifier
}

func EventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}

func InitService(chaincodeID, channelID string, org *sdk.OrgInfo, sdk *fabsdk.FabricSDK) (*ServiceSetup, error) {
	handler := &ServiceSetup{
		ChaincodeID: chaincodeID,
	}
	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(org.OrgUser), fabsdk.WithOrg(org.OrgName))
	channel_client, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new channel client: %s", err)
	}
	ledger_client, err := ledger.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new ledger client: %s", err)
	}

	handler.ChannelClient = channel_client
	handler.LedgerClient = ledger_client
	return handler, nil
}

func InitSetup(configpath string) *ServiceSetup {
	sdkentity, err := sdk.Setup("../config/config_northchina.yaml", &info)

	if err != nil {
		logger.Println(">> Sdk set error", err)
		return nil
	}

	servicesetup, err := InitService(info.ChaincodeID, info.ChannelID, info.Orgs[1], sdkentity)
	if err != nil {
		logger.Println(">> init chaincode error", err)
		return nil
	}

	return servicesetup
}

func Save(servicesetup *ServiceSetup, atc Atc) (string, error) {
	msg, err := servicesetup.SaveAtc(atc)
	if err != nil {
		logger.Println(err.Error())
		return "", err
	}
	logger.Println("信息发布成功, 交易编号为: " + msg)
	return msg, nil
}
