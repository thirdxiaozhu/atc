docker stop $(docker ps -a | awk '($2 ~ /hyperledger*/) {print $1}')
docker stop $(docker ps -a | awk '($2 ~ /dev-.*.atc.*/) {print $1}')
docker rm $(docker ps -a | awk '($2 ~ /hyperledger*/) {print $1}')
docker rm $(docker ps -a | awk '($2 ~ /dev-.*.atc.*/) {print $1}')
#docker rm $(docker ps -a -q)
#docker rmi $(docker images -q)
docker volume rm $(docker volume ls -qf dangling=true)

sudo docker network prune
sudo docker volume prune
cd fabric && docker-compose up -d
docker start mysql
docker start ipfs
cd ..
go build mainStage.go
./mainStage