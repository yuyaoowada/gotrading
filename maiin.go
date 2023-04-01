package main

import (
	"fmt"
	"gotrading/app/controllers"
	"gotrading/app/models"
	"gotrading/config"
	"gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
	//apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	//ticker, _ := apiClient.GetTicker("BTC_USD")
	//fmt.Println(ticker)
	//fmt.Println(ticker.GetMidPrice())
	//fmt.Println(ticker.DateTime())
	//fmt.Println(ticker.TruncateDateTime(time.Hour))
}
