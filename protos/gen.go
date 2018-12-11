package wireguard

//go:generate sh -c "protoc -I./ --go_out=plugins=grpc:./ ./*.proto"
