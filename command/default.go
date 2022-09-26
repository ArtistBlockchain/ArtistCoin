package command

import "github.com/ArtistBlockchain/ArtistCoin/server"

const (
	DefaultGenesisFileName = "genesis.json"
	DefaultChainName       = "artistcoin"
	DefaultChainID         = 4849
	DefaultPremineBalance  = "0x33B2E3C9FD0803CE8000000" // 1 billion units of native network currency
	DefaultConsensus       = server.IBFTConsensus
	DefaultGenesisGasUsed  = 458752  // 0x70000
	DefaultGenesisGasLimit = 5242880 // 0x500000
)

const (
	JSONOutputFlag  = "json"
	GRPCAddressFlag = "grpc-address"
	JSONRPCFlag     = "jsonrpc"
)

// GRPCAddressFlagLEGACY Legacy flag that needs to be present to preserve backwards
// compatibility with running clients
const (
	GRPCAddressFlagLEGACY = "grpc"
)
