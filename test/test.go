package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sdk"
	"service"
)

func query(servicesetup *service.ServiceSetup) {
	result, err := servicesetup.FindAtcInfoByID("20d23a24-b9e0-4946-89e1-6d7a8c761584")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var atc service.Atc
		json.Unmarshal(result, &atc)
		fmt.Println("根据身份证号码查询信息成功：")
		fmt.Println(atc)
	}

}

func queryByPublisher(servicesetup *service.ServiceSetup, querystr string) {
	result, err := servicesetup.FindAtcByPublisher(querystr)
	fmt.Println(result)

	var atcs []service.Atc
	if err != nil {
		fmt.Println(err.Error())
	} else {
		//var atc service.Atc
		json.Unmarshal(result, &atcs)
		fmt.Println("根据身份证号码查询信息成功：")
		fmt.Println(atcs)
	}
}

func save(servicesetup *service.ServiceSetup, atc service.Atc) {
	msg, err := servicesetup.SaveAtc(atc)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}
}

func main() {
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

	info := sdk.SdkEnvInfo{
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

	sdkentity, err := sdk.Setup("../config/config_northchina.yaml", &info)

	if err != nil {
		fmt.Println(">> Sdk set error", err)
		os.Exit(-1)
	}

	servicesetup, err := service.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[1], sdkentity)
	if err != nil {
		fmt.Println(">> init chaincode error", err)
		os.Exit(-1)
	}

	//atc := service.Atc{
	//	ID:      "1",
	//	Time:    "111",
	//	User:    "DSDDS",
	//	Company: "AirChina",
	//	Type:    "ACARS",
	//	Info:    "ABCDEFG",
	//}

	//save(servicesetup, atc)
	//query(servicesetup)

	pb_map := map[string]string{"Publisher": "nc"}
	fmt.Println(pb_map)
	fmt.Println(MapToJson(pb_map))
	map_str := MapToJson(pb_map)

	queryByPublisher(servicesetup, map_str)
}

func MapToJson(param map[string]string) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}
