package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sdk"
	"service"
)

func del(servicesetup *service.ServiceSetup, ID string) {
	transid, err := servicesetup.DelAtc(ID)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息发布成功, 交易编号为: " + transid)
	}
}

func modify(servicesetup *service.ServiceSetup, atc service.Atc) {
	transid, err := servicesetup.ModifyAtc(atc)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息发布成功, 交易编号为: " + transid)
	}
}

func query(servicesetup *service.ServiceSetup, id string) {
	result, err := servicesetup.FindAtcInfoByID(id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var atc service.Atc
		json.Unmarshal(result, &atc)
		fmt.Println("根据身份证号码查询信息成功：")
		fmt.Println(atc)
	}
}

func queryByString(servicesetup *service.ServiceSetup, querystr string) {
	result, err := servicesetup.FindAtcByQueryString(querystr)

	var atcs []service.Atc

	if err != nil {
		fmt.Println(err.Error())
	} else {
		//var atc service.Atc
		json.Unmarshal(result, &atcs)
		fmt.Println("根据身份证号码查询信息成功：")
		fmt.Println(atcs)
		fmt.Println(atcs[0].Content, atcs[0].Address, atcs[0].IsValid)
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
	//	ID:        "3e340f55-3864-4ac7-8a40-1609d5b30aae",
	//	Time:      "1678654",
	//	Address:   "QmeiiXbdqNRKCeHkTYVTKVfrRsETJ6dksb2t28mPh91sSu",
	//	Signature: "MEQCIHdvSOZWYL6o5cwfRX5VNp8NcXRI6SoCTl8bj5f/TWZ+AiAegMJOh1/L5g/DYSvLKCxhEF9ebXgGW7o0d96pGb34+A==",
	//}

	//save(servicesetup, atc)
	//modify(servicesetup, atc)
	query(servicesetup, "55a7f921-3dfd-41f0-b6c0-8cad448d93e4")

	//pb_map := map[string]string{"Publisher": "ec", "Company": "eastchina"}
	//fmt.Println(pb_map)
	//fmt.Println(MapToJson(pb_map))
	////map_str := MapToJson(pb_map)

	////queryByString(servicesetup, map_str)
	////queryByString(servicesetup, "{\"Publisher\": \"ec\"}")
	//queryByString(servicesetup, "{\"Company\": \"eastchina\"}")
	//queryByString(servicesetup, "{\"Publisher\": \"ec\"}")
	//del(servicesetup, "55a7f921-3dfd-41f0-b6c0-8cad448d93e4")
}

func MapToJson(param map[string]string) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}
