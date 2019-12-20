package main

import (
	"log"
)

var config Config

func main() {
	//config := InitConfig()
	//switch config.action {
	//case "optionScan":
	//	// equityQuotes := LoadOHLCFiles(config)
	//	//log.Println(len(equityQuotes))
	//	//ProcessOptionsFile(config, equityQuotes)
	//	log.Println("Not implemented yet")
	//case "ohlcLoad":
	//	namesCache := LoadNamesFiles(config)
	//
	//	LoadAllOHLCFiles(config, namesCache)
	//}

	bstest()
}

func bstest() {
	S0 := 50.0

	// K := 640.0
	// right := "P"
	// price := 2.4

	K := 100.0
	right := "C"
	price := -1.0

	// vol := 0.2939
	// price := -1.0
	vol := 0.25

	r := 0.05 // risk free rate
	eval_date := "20150115"
	exp_date := "20160115"

	opt := NewOption(right, S0, K, eval_date, exp_date, r, vol, price)

	log.Println("CALL")
	log.Printf("T: %v\n", opt.T)
	log.Printf("Price: %v\n", opt.price)
	log.Printf("Delta: %v\n", opt.delta)
	log.Printf("Theta: %v\n", opt.theta)
	log.Printf("Gamma: %v\n", opt.gamma)
	log.Printf("Volatility: %v\n", opt.sigma)
}
