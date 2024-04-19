package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func getSign(privateKey *ecdsa.PrivateKey, time int64) []byte {

	// privateKey, err := crypto.HexToECDSA(priv)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// timestamp := time.Now().Unix()
	message := fmt.Sprintf("Greetings from Avail!\n\nSign this message to check your eligibility. This signature will not cost you any fees.\n\nTimestamp: %d", time)

	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		fmt.Println("sign error!")
	}
	signatureBytes[64] += 27

	return signatureBytes
}

func GetAvail(privateKey *ecdsa.PrivateKey) {

	// Define the URL for the web service
	const urlLogin = "https://claim-api.availproject.org/check-rewards"

	// Define the Infura RPC endpoint
	const RPC = "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID"

	// Create an Ethereum client
	// client, err := ethclient.Dial(RPC)
	// if err != nil {
	// 	fmt.Println("Failed to connect to the Ethereum client:", err)
	// 	return
	// }

	// Loop through each private key

	// Convert the private key string to a crypto.PrivateKey
	// privateKeyBytes, err := crypto.HexToECDSA(priv)
	// if err != nil {
	// 	fmt.Println("Failed to parse private key:", err)
	// 	return
	// }

	// Generate the Ethereum address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Failed to generate public key from private key")
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Sign the message
	timestamp := time.Now().Unix()

	signature := getSign(privateKey, timestamp)

	// Prepare the data for the web service
	requestData := map[string]interface{}{
		"account":       fromAddress.Hex(),
		"signedMessage": "0x" + common.Bytes2Hex(signature),
		"timestamp":     timestamp,
		"type":          "ETHEREUM",
	}
	// fmt.Println(requestData)
	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Failed to marshal request data:", err)
		return

	}

	// Make the HTTP POST request
	resp, err := http.Post(urlLogin, "application/json", bytes.NewBuffer(requestDataBytes))
	if err != nil {
		fmt.Println("Failed to make HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	// Print the response
	fmt.Println("Response:", string(body))

}

func main() {
	startPrivateKey := big.NewInt(0)
	for {
		startPrivateKey.Add(startPrivateKey, big.NewInt(1))
		privateKeyBytes := startPrivateKey.Bytes()

		// 为字节数组填充前导0，以确保其长度为32字节
		paddedPrivateKey := append(make([]byte, 32-len(privateKeyBytes)), privateKeyBytes...)
		privateKeyECDSA, err := crypto.ToECDSA(paddedPrivateKey)
		if err != nil {
			fmt.Printf("Failed to create private key: %v\n", err)
			return
		}
		fmt.Println(hex.EncodeToString(paddedPrivateKey))
		GetAvail(privateKeyECDSA)
	}

}
