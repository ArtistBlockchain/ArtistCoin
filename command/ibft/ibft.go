package ibft

import (
	"github.com/ArtistBlockchain/ArtistCoin/command/helper"
	"github.com/ArtistBlockchain/ArtistCoin/command/ibft/candidates"
	"github.com/ArtistBlockchain/ArtistCoin/command/ibft/propose"
	"github.com/ArtistBlockchain/ArtistCoin/command/ibft/quorum"
	"github.com/ArtistBlockchain/ArtistCoin/command/ibft/snapshot"
	"github.com/ArtistBlockchain/ArtistCoin/command/ibft/status"
	_switch "github.com/ArtistBlockchain/ArtistCoin/command/ibft/switch"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
