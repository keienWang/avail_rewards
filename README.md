# Ethereum Wallet Checker

This project contains a script written in both JavaScript (Node.js) and Go for interacting with Ethereum wallets and a web service. The script checks the eligibility of Ethereum wallets to claim rewards from a specific web service.

## Features

- Interacts with Ethereum wallets using private keys.
- Signs a message with the wallet's private key and sends it to a web service to check eligibility.
- Supports both JavaScript (Node.js) and Go implementations.

## Prerequisites

- Node.js and npm for JavaScript implementation.
- Go Ethereum (`geth`) for Go implementation.

## Usage

1. Clone this repository.

2. Install dependencies:
   - For JavaScript: `npm install`
   - For Go: Ensure you have Go Ethereum (`geth`) installed.

3. Update the code:
   - Replace `"YOUR_PRIVATE_KEY_1"`, `"YOUR_PRIVATE_KEY_2"`, and `"YOUR_INFURA_PROJECT_ID"` with your actual private keys and Infura project ID.

4. Run the code:
   - For JavaScript: `node script.js`
   - For Go: `go run script.go`

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
