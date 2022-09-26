package server

import (
	"github.com/ArtistBlockchain/ArtistCoin/consensus"
	consensusDev "github.com/ArtistBlockchain/ArtistCoin/consensus/dev"
	consensusDummy "github.com/ArtistBlockchain/ArtistCoin/consensus/dummy"
	consensusIBFT "github.com/ArtistBlockchain/ArtistCoin/consensus/ibft"
	"github.com/ArtistBlockchain/ArtistCoin/secrets"
	"github.com/ArtistBlockchain/ArtistCoin/secrets/awsssm"
	"github.com/ArtistBlockchain/ArtistCoin/secrets/gcpssm"
	"github.com/ArtistBlockchain/ArtistCoin/secrets/hashicorpvault"
	"github.com/ArtistBlockchain/ArtistCoin/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
