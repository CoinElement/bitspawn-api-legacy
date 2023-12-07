build:
	go build -o ./target/

update:
	go mod vendor

fix-eth-dep:
	cp -r "${GOPATH}/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1" "vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"
	cp -r "${GOPATH}/src/github.com/karalabe/usb/" "vendor/github.com/karalabe/"

integration:
	go test ./integration -tags integration
