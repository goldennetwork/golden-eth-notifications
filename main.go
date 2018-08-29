package main

import (
	"log"
	"os"
	"time"

	. "github.com/ethereum/go-ethereum/rpc"
	"github.com/goldennetwork/golden-eth-notifications/subscribe"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error: Load .env file.")
	}

	// session := db.NewMongoClient()
	// defer session.Close()

	client, err := Dial(os.Getenv("URL_WS_GOLDEN_RINKEBY"))
	if err != nil {
		log.Fatalln(err)
	}

	rpcQuery := NewRPCQuery(client)
	etherSub := subscribe.NewEthereumSubscribe(time.Second * 2)
	etherSub.StartEtherSub(client, rpcQuery)
}
