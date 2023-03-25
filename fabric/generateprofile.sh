rm -rf ./channel-artifacts
rm -rf ./crypto-config

cryptogen generate --config=crypto-config.yaml
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block -channelID test-channel
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mychannel
configtxgen -outputAnchorPeersUpdate ./channel-artifacts/EastChinaMSPanchors.tx -profile TwoOrgsChannel -channelID mychannel -asOrg EastChina
configtxgen -outputAnchorPeersUpdate ./channel-artifacts/NorthChinaMSPanchors.tx -profile TwoOrgsChannel -channelID mychannel -asOrg NorthChina
configtxgen -outputAnchorPeersUpdate ./channel-artifacts/AirChinaMSPanchors.tx -profile TwoOrgsChannel -channelID mychannel -asOrg AirChina
configtxgen -outputAnchorPeersUpdate ./channel-artifacts/CsairMSPanchors.tx -profile TwoOrgsChannel -channelID mychannel -asOrg Csair
