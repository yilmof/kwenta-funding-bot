package main

import (
	"flag"
	"fmt"
	"kwenta_rates/contracts"
	"kwenta_rates/dydx"
	"log"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
)

type Signal struct {
	Symbol        string
	KwentaSide    string
	DydxSide      string
	KwentaFunding float64
	DydxFunding   float64
	KwentaPrice   float64
	DydxPrice     float64

	ContractInstance *contracts.Contracts
}

type Position struct {
	Symbol      string
	Amount      float64
	KwentaPrice float64
	DydxPrice   float64

	KwentaFunding float64
	DydxFunding   float64

	// profit/loss
	KwentaPL float64
	DydxPL   float64

	FundingCalculated time.Time
	PositionEntered   time.Time

	ContractInstance *contracts.Contracts
}

type Account struct {
	Balance   float64
	Positions []Position
}

func (p *Position) UpdatePrices(kwentaPrice, dydxPrice float64) {
	p.KwentaPrice = kwentaPrice
	p.DydxPrice = dydxPrice
}

func (p *Position) UpdateFunding(kwentaFunding, dydxFunding float64) {
	p.KwentaFunding = kwentaFunding
	p.DydxFunding = dydxFunding
	p.FundingCalculated = time.Now()
}

func (p *Position) UpdatePL(kwentaPL, dydxPL float64) {
	p.KwentaPL = kwentaPL
	p.DydxPL = dydxPL
}

func (a *Account) EnterPosition(symbol string, amount, kwentaPrice, dydxPrice, kwentaFunding, dydxFunding float64, instance *contracts.Contracts) {
	cost := amount * kwentaPrice
	if cost > a.Balance {
		log.Println("Balance not enough")
		return
	}
	if len(a.Positions) > 0 {
		log.Println("Already in position")
		return
	}
	a.Balance = a.Balance - cost
	a.Positions = append(a.Positions, Position{
		Symbol:            symbol,
		Amount:            amount,
		KwentaPrice:       kwentaPrice,
		DydxPrice:         dydxPrice,
		KwentaFunding:     kwentaFunding,
		DydxFunding:       dydxFunding,
		KwentaPL:          0,
		DydxPL:            0,
		FundingCalculated: time.Now(),
		PositionEntered:   time.Now(),
		ContractInstance:  instance,
	})
	log.Printf("Entered %s, amount %.2f @ %.2f on Kwenta and %.2f on DYDX\nFunding kwenta: %.4f, funding dydx: %.4f\n", symbol, amount, kwentaPrice, dydxPrice, kwentaFunding, dydxFunding)
}

func (a *Account) ExitPosition(pos *Position) {
	totalProfit := pos.KwentaPL + pos.DydxPL
	if pos.KwentaPL < 0 && pos.DydxPL < 0 {
		a.Balance = a.Balance - pos.KwentaPL - pos.DydxPL
	} else if pos.KwentaPL > 0 && pos.DydxPL < 0 {
		a.Balance = a.Balance + totalProfit
	} else if pos.KwentaPL < 0 && pos.DydxPL > 0 {
		a.Balance = a.Balance + totalProfit
	} else {
		a.Balance = a.Balance + pos.KwentaPL + pos.DydxPL
	}
	a.Positions = []Position{}
	log.Printf("Exited %s with P/L of %.2f\nBalance: %.2f", pos.Symbol, totalProfit, a.Balance)
}

func main() {
	f, err := os.OpenFile("trades.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	isDocker := flag.Bool("docker", false, "use if running with docker")
	flag.Parse()

	if !*isDocker {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// env variables
	RPC_URL := os.Getenv("RPC_URL")
	//DISCORD_WEBHOOK_URL := os.Getenv("DISCORD_WEBHOOK_URL")
	//REDIS_URL := os.Getenv("REDIS_URL")

	// geth client init
	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	// discord webhook client init
	// webhookClient, err := webhook.NewWithURL(DISCORD_WEBHOOK_URL)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // redis init
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     REDIS_URL,
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	// _ = rdb

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

	//previousNotificationTime := make(map[string]time.Time)

	account := Account{
		Balance:   2000,
		Positions: []Position{},
	}
	log.Printf("Initialized paper balance of $%.2f\n", account.Balance)

	signals := make(chan Signal)

	fmt.Println("Checking funding rates...")

	go func() {
		for signal := range signals {
			if len(account.Positions) == 0 {
				bal := account.Balance * 0.95
				amount := bal / signal.KwentaPrice
				account.EnterPosition(signal.Symbol, amount, signal.KwentaPrice, signal.DydxPrice, signal.KwentaFunding, signal.DydxFunding, signal.ContractInstance)
				go func() {
					for {
						kwentaPrice := getKwentaPrice(&account)
						dydxPrice, err := getDydxPrice(&account)
						if err != nil {
							log.Fatal(err)
						}
						account.Positions[0].UpdatePrices(kwentaPrice, dydxPrice)
						log.Println("Updated prices for current position")
						currentBalance := account.Balance + account.Positions[0].KwentaPL + account.Positions[0].DydxPL
						log.Printf("Current balance: $%.2f\n", currentBalance)
						time.Sleep(5 * time.Minute)
					}
				}()

				go func() {
					for {
						kwentaFunding := getKwentaFunding(&account)
						dydxFunding, err := getDydxFunding(&account)
						if err != nil {
							log.Fatal(err)
						}
						account.Positions[0].UpdateFunding(kwentaFunding, dydxFunding)
						profit := calculateDailyProfit(kwentaFunding, dydxFunding, 1000)
						if profit < 0 {
							log.Printf("Current profit/1000$ %.2f, exiting position...\n", profit)
							account.ExitPosition(&account.Positions[0])
							break
						}
						log.Println("Updated funding rates for current position")
						time.Sleep(1 * time.Hour)
					}
				}()

			}
		}
	}()

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
			if formattedRate >= 0.01 || formattedRate <= -0.01 {
				dydxRate, err := dydx.GetFunding(symbol)
				if err != nil {
					log.Fatal(err)
				}

				dydxPrice, err := dydx.GetPrice(symbol)
				if err != nil {
					log.Fatal(err)
				}

				price, err := instance.AssetPrice(nil)
				if err != nil {
					log.Fatal(err)
				}
				floatPrice, _ := (big.NewFloat(0).SetInt(price.Price)).Float64()
				formattedPrice := floatPrice / 1000000000000000000

				// dailyProfitLeveraged := calculateDailyProfit(formattedRate, dydxRate, 2500)
				// dailyProfitPercentageLeveraged := (dailyProfitLeveraged / 2500) * 100
				// dailyProfitThousand := calculateDailyProfit(formattedRate, dydxRate, 1000)

				shortSkewPrcnt, err := getSkew(instance)
				if err != nil {
					log.Fatal(err)
				}
				longSkewPrcnt := 1 - shortSkewPrcnt
				shortRounded := math.Round(shortSkewPrcnt * 100)
				longRounded := math.Round(longSkewPrcnt * 100)

				if shortRounded > 60 || longRounded > 60 {
					// return false for positive, true for negative
					if math.Signbit(formattedRate) && formattedRate < dydxRate {
						signals <- Signal{
							Symbol:        symbol,
							KwentaSide:    "long",
							DydxSide:      "short",
							KwentaFunding: formattedRate / 100,
							DydxFunding:   dydxRate / 100,
							KwentaPrice:   formattedPrice,
							DydxPrice:     dydxPrice,

							ContractInstance: instance,
						}
					} else if !math.Signbit(formattedRate) && formattedRate > dydxRate {
						signals <- Signal{
							Symbol:        symbol,
							KwentaSide:    "short",
							DydxSide:      "long",
							KwentaFunding: formattedRate / 100,
							DydxFunding:   dydxRate / 100,
							KwentaPrice:   formattedPrice,
							DydxPrice:     dydxPrice,

							ContractInstance: instance,
						}
					}

				}

				// if time.Since(previousNotificationTime[symbol]) >= (time.Hour * 6) {
				// 	fundingMsg := ""
				// 	if longRounded > shortRounded {
				// 		fundingMsg = "⬆️"
				// 	} else if longRounded < shortRounded {
				// 		fundingMsg = "⬇️"
				// 	} else {
				// 		fundingMsg = "🔃"
				// 	}

				// 	embedList := []discord.Embed{}
				// 	embedList = append(embedList, discord.NewEmbedBuilder().
				// 		SetColor(0x005eff).
				// 		SetURLf("https://kwenta.eth.limo/market/?asset=s%s&accountType=isolated_margin", strings.ToUpper(symbol)).
				// 		SetTitle(strings.ToUpper(symbol)).
				// 		AddField("Kwenta funding", fmt.Sprintf("%.4f%%", formattedRate), true).
				// 		AddField("DYDX funding", fmt.Sprintf("%.4f%%", dydxRate), true).
				// 		AddField("Potential daily profit (2.5x lev)", fmt.Sprintf("%.2f%%", dailyProfitPercentageLeveraged), true).
				// 		AddField("Skew Long", strconv.Itoa(int(longRounded))+"%", false).
				// 		AddField("Skew Short", strconv.Itoa(int(shortRounded))+"%", true).
				// 		AddField("Funding", fundingMsg, true).
				// 		AddField("daily profit for 1000$", fmt.Sprintf("$%.2f", dailyProfitThousand), true).
				// 		SetTimestamp(time.Now()).
				// 		SetAuthorName("Kwenta").
				// 		SetFooterText("queried").
				// 		Build())

				// 	_, err = webhookClient.CreateEmbeds(embedList, rest.WithDelay(5*time.Second))
				// 	if err != nil {
				// 		log.Fatal(err)
				// 	}
				// 	fmt.Printf("%s: %f\n", symbol, formattedRate)
				// 	previousNotificationTime[symbol] = time.Now()
				// } else {
				// 	fmt.Printf("%s: %f (Already sent notification within 30 minutes)\n", symbol, formattedRate)
				// }
			}
		}
		time.Sleep(time.Minute)
	}
}

func getKwentaPrice(account *Account) float64 {
	price, err := account.Positions[0].ContractInstance.AssetPrice(nil)
	if err != nil {
		log.Fatal(err)
	}
	floatPrice, _ := (big.NewFloat(0).SetInt(price.Price)).Float64()
	formattedPrice := floatPrice / 1000000000000000000
	return formattedPrice
}

func getKwentaFunding(account *Account) float64 {
	rate, err := account.Positions[0].ContractInstance.CurrentFundingRate(nil)
	if err != nil {
		log.Fatal(err)
	}
	floatRate, _ := (big.NewFloat(0).SetInt(rate)).Float64()
	formattedRate := (floatRate / 24) / 1000000000000000000
	return formattedRate
}

func getDydxPrice(account *Account) (float64, error) {
	dydxPrice, err := dydx.GetPrice(account.Positions[0].Symbol)
	if err != nil {
		return 0, err
	}
	return dydxPrice, nil
}

func getDydxFunding(account *Account) (float64, error) {
	dydxRate, err := dydx.GetFunding(account.Positions[0].Symbol)
	if err != nil {
		return 0, err
	}
	return dydxRate, nil
}

func calculateDailyProfit(kwentaRate, dydxRate, base float64) float64 {
	// return false for positive, true for negative
	if math.Signbit(kwentaRate) {
		if math.Signbit(dydxRate) {
			kwentaProfit := (math.Abs(kwentaRate) / 100) * 24 * base
			dydxProfit := (math.Abs(dydxRate) / 100) * 24 * base
			var totalProfit float64
			if kwentaProfit >= dydxProfit {
				totalProfit = kwentaProfit - dydxProfit
			} else {
				totalProfit = dydxProfit - kwentaProfit
			}
			return totalProfit
		} else {
			kwentaProfit := (math.Abs(kwentaRate) / 100) * 24 * base
			dydxProfit := (math.Abs(dydxRate) / 100) * 24 * base
			return kwentaProfit + dydxProfit
		}
	} else {
		if math.Signbit(dydxRate) {
			kwentaProfit := (math.Abs(kwentaRate) / 100) * 24 * base
			dydxProfit := (math.Abs(dydxRate) / 100) * 24 * base
			return kwentaProfit + dydxProfit
		} else {
			kwentaProfit := (math.Abs(kwentaRate) / 100) * 24 * base
			dydxProfit := (math.Abs(dydxRate) / 100) * 24 * base
			var totalProfit float64
			if kwentaProfit >= dydxProfit {
				totalProfit = kwentaProfit - dydxProfit
			} else {
				totalProfit = dydxProfit - kwentaProfit
			}
			return totalProfit
		}
	}
}

func getSkew(instance *contracts.Contracts) (float64, error) {
	skew, err := instance.MarketSizes(nil)
	if err != nil {
		return 0, err
	}
	skewLong, _ := (big.NewFloat(0).SetInt(skew.Long)).Float64()
	skewShort, _ := (big.NewFloat(0).SetInt(skew.Short)).Float64()
	skewShortPrcnt := skewShort / (skewLong + skewShort)
	return skewShortPrcnt, nil
}
