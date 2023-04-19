package main

import (
	"fmt"
	"os"

	"sdk"
)

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

	sdkentity, err := sdk.Setup("./config/config_test.yaml", &info)

	if err != nil {
		fmt.Println(">> Sdk set error", err)
		os.Exit(-1)
	}

	if err := sdk.CreateChannel(&info); err != nil {
		fmt.Println(">> Create Channel error", err)
		os.Exit(-1)
	}

	if err := sdk.JoinChannel(&info); err != nil {
		fmt.Println(">> Join Channel error", err)
		os.Exit(-1)
	}
	packageID, err := sdk.InstallCC(&info)
	fmt.Println(packageID)

	if err != nil {
		fmt.Println(">> install chaincode error", err)
		os.Exit(-1)
	}

	//for err := sdk.ApproveLifecycle(&info, 1, packageID); err != nil; {
	//	fmt.Println(">> approve chaincode error", err)
	//	//os.Exit(-1)
	//	continue
	//}

	if err := sdk.ApproveLifecycle(&info, 1, packageID); err != nil {
		fmt.Println(">> approve chaincode error", err)
		os.Exit(-1)
	}

	if err := sdk.InitCC(&info, false, sdkentity); err != nil {
		fmt.Println(">> init chaincode error", err)
		os.Exit(-1)
	}
	fmt.Println(">> 链码状态设置完成")

}
