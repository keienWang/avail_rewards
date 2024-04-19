package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Define the URL for the web service
const urlLogin = "https://claim-api.availproject.org/check-rewards"

// Define the Infura RPC endpoint
const RPC = "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID"

func main() {

	// Define your private keys
	privateKeys := []string{"YOUR_PRIVATE_KEY_1", "YOUR_PRIVATE_KEY_2"}

	// Create an Ethereum client
	// client, err := ethclient.Dial(RPC)
	// if err != nil {
	// 	fmt.Println("Failed to connect to the Ethereum client:", err)
	// 	return
	// }

	// Loop through each private key
	for _, privateKey := range privateKeys {
		// Convert the private key string to a crypto.PrivateKey
		privateKeyBytes, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			fmt.Println("Failed to parse private key:", err)
			continue
		}

		// Generate the Ethereum address from the private key
		publicKey := privateKeyBytes.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			fmt.Println("Failed to generate public key from private key")
			continue
		}
		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		// Sign the message
		timestamp := time.Now().Unix()
		message := fmt.Sprintf("Greetings from Avail!\n\nSign this message to check your eligibility. This signature will not cost you any fees.\n\nTimestamp: %d", timestamp)
		messageHash := crypto.Keccak256([]byte(message))
		signature, err := crypto.Sign(messageHash, privateKeyBytes)
		if err != nil {
			fmt.Println("Failed to sign message:", err)
			continue
		}

		// Prepare the data for the web service
		requestData := map[string]interface{}{
			"account":       fromAddress.Hex(),
			"signedMessage": common.Bytes2Hex(signature),
			"timestamp":     timestamp,
			"type":          "ETHEREUM",
		}
		requestDataBytes, err := json.Marshal(requestData)
		if err != nil {
			fmt.Println("Failed to marshal request data:", err)
			continue
		}

		// Make the HTTP POST request
		resp, err := http.Post(urlLogin, "application/json", bytes.NewBuffer(requestDataBytes))
		if err != nil {
			fmt.Println("Failed to make HTTP request:", err)
			continue
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			continue
		}

		// Print the response
		fmt.Println("Response:", string(body))
	}
}
