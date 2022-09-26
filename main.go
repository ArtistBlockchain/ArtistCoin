package main

import (
	_ "embed"

	"github.com/ArtistBlockchain/ArtistCoin/command/root"
	"github.com/ArtistBlockchain/ArtistCoin/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
