package root

import (
	"fmt"
	"os"

	"github.com/ArtistBlockchain/ArtistCoin/command/backup"
	"github.com/ArtistBlockchain/ArtistCoin/command/genesis"
	"github.com/ArtistBlockchain/ArtistCoin/command/helper"
	"github.com/ArtistBlockchain/ArtistCoin/command/ibft"
	"github.com/ArtistBlockchain/ArtistCoin/command/license"
	"github.com/ArtistBlockchain/ArtistCoin/command/loadbot"
	"github.com/ArtistBlockchain/ArtistCoin/command/monitor"
	"github.com/ArtistBlockchain/ArtistCoin/command/peers"
	"github.com/ArtistBlockchain/ArtistCoin/command/secrets"
	"github.com/ArtistBlockchain/ArtistCoin/command/server"
	"github.com/ArtistBlockchain/ArtistCoin/command/status"
	"github.com/ArtistBlockchain/ArtistCoin/command/txpool"
	"github.com/ArtistBlockchain/ArtistCoin/command/version"
	"github.com/ArtistBlockchain/ArtistCoin/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "ArtistCoin is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		loadbot.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
