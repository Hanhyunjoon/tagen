#chaincode insall
docker exec cli peer chaincode install -n genedata -v 1.0 -p github.com/genedata
#chaincode instatiate
docker exec cli peer chaincode instantiate -n genedata -v 1.0 -C mychannel -c '{"Args":["Init"]}' -P 'OR ("Org1MSP.member", "Org2MSP.member")'
sleep 5
#chaincode query a
docker exec cli peer chaincode query -n genedata -C mychannel -c '{"Args":["getGene","a"]}'
#chaincode invoke b
docker exec cli peer chaincode invoke -n genedata -C mychannel -c '{"Args":["addGene","b","1111","1222","333"]}'
sleep 5
#chaincode query b
docker exec cli peer chaincode query -n genedata -C mychannel -c '{"Args":["getGene","b"]}'

echo '-------------------------------------END-------------------------------------'
