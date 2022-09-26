## Artistcoin
* Name - ArtistCoin
* Symbol - ARTC
* Supply - 1 billion
* Blocktime - 3 seconds
* Consensus - PoS
* P2P Port - 1478
* JSON-RPC Port - 8545 
* EVM Compatible

## Official Links


### Build from Source (Ubuntu 20.04)
Requirements - `Go >=1.17`

#### Setup Go Path
```
sudo nano ~/.profile
```
Paste this into the bottom of the file
```
export GOPATH=$HOME/work
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```
```
source ~/.profile
```

#### Install Go
```
wget https://go.dev/dl/go1.17.13.linux-amd64.tar.gz
sudo tar -xvf go1.17.13.linux-amd64.tar.gz
sudo mv go /usr/local && rm go1.17.13.linux-amd64.tar.gz
```
Check that it's installed
```
go version
```
You should see something like this:
```
go version go1.17.13 linux/amd64
```

#### Build ArtistCoin
```
git clone https://github.com/ArtistBlockchain/ArtistCoin.git
cd ArtistCoin/
go build -o artistcoin main.go
```

#### Running a Full Validating Node
After you have [downloaded](https://github.com/ArtistBlockchain/ArtistCoin/releases/latest) the binaries or [built from source](https://github.com/ArtistBlockchain/ArtistCoin#build-from-source), enter the `ArtistCoin` directory and run the following:
```
./artistcoin server --data-dir ~/.artistcoin --chain mainnet-genesis.json --seal --max-slots 40960 --grpc 0.0.0.0:9632 --libp2p 0.0.0.0:1478 --jsonrpc 0.0.0.0:8545 --max-inbound-peers 128 --max-outbound-peers 16
```

#### Running a non-Validating node
```
./artistcoin server --data-dir ~/.artistcoin --chain mainnet-genesis.json --libp2p 0.0.0.0:1478 --nat <public_or_private_ip>
```


---
```
Copyright 2022 Polygon Technology
Copyright 2022 ArtistCoin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
