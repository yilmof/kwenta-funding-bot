package main

import (
	"flag"
	"fmt"
	"kwenta_rates/contracts"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gtuk/discordwebhook"
	"github.com/joho/godotenv"
)

func main() {
	isDocker := flag.Bool("docker", false, "use if running with docker")
	flag.Parse()

	if !*isDocker {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	RPC_URL := os.Getenv("RPC_URL")
	DISCORD_WEBHOOK_URL := os.Getenv("DISCORD_WEBHOOK_URL")

	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	kwenta_contracts := map[string]string{
		"eth":   "0x9363c080Ca0B16EAD12Fd33aac65c8D0214E9d6D",
		"aave":  "0x8C14dB69b1778c7Bbb0683B2deA21F79b9B5f059",
		"atom":  "0x407e5a7548869ce520354d8eca127754625cc9c8",
		"avax":  "0x69d255473d0d15dc087cbf4962a65839aa28a2c5",
		"btc":   "0xb0a1d2c68bf4d0980402dc220ca6ddeb8dbfbc56",
		"doge":  "0xea46a4dfa7d2767ff4bae2b76f5c6bd80057c723",
		"link":  "0xcD077EAb8efBcb94aa04F2055D7A9216F23697A6",
		"matic": "0xF23c5eC62eC4398302eFd84587eb8BA26f21B155",
		"near":  "0x1EaEd534ee8D25DA73a4E21cDE96e4Fca9C46187",
		"sol":   "0x6eE44d9e0f868833A5543bcABc3Bd1a7d843eDB8",
		"uni":   "0xda3a5e9502b23Eaedc8cC048998893013e09787d",
	}

	previousNotificationTime := make(map[string]time.Time)

	fmt.Println("Checking funding rates...")
	for {
		for symbol, contract_address := range kwenta_contracts {
			address := common.HexToAddress(contract_address)
			instance, err := contracts.NewContracts(address, client)
			if err != nil {
				log.Fatal(err)
			}

			rate, err := instance.CurrentFundingRate(nil)
			if err != nil {
				log.Fatal(err)
			}
			floatRate, _ := (big.NewFloat(0).SetInt(rate)).Float64()
			formattedRate := (floatRate / 24) / 10000000000000000
			var username = "Arb bot"
			if formattedRate >= 0.01 || formattedRate <= -0.01 {
				if time.Since(previousNotificationTime[symbol]) >= (time.Minute * 30) {
					content := fmt.Sprintf("Funding rate on %s is now %f", symbol, formattedRate)
					message := discordwebhook.Message{
						Username: &username,
						Content:  &content,
					}
					err := discordwebhook.SendMessage(DISCORD_WEBHOOK_URL, message)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("%s: %f\n", symbol, formattedRate)
					previousNotificationTime[symbol] = time.Now()
				} else {
					fmt.Printf("%s: %f (Already sent notification within 30 minutes)\n", symbol, formattedRate)
				}
			}
		}
		time.Sleep(time.Minute)
	}
}
