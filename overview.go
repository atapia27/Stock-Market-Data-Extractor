package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/tabwriter"
)

func getOverview(symbol string) companyOverview {

	//Initialize the necessary variables, such as the API key or credentials, endpoint, and method.
	var (
		url    = "https://www.alphavantage.co/query?function=OVERVIEW&symbol=" + symbol + "&apikey=GG67ANDCMVCCRKPJ"
		method = "GET" //The method that I am going to use for the API request.
	)

	//Make a request to the API endpoint using the HTTP method defined before.
	response, err := http.NewRequest(method, url, nil)

	//If there is an error when making a request to the API endpoint then stop the program and print an error message.
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	//Send the request to the API endpoint.
	client := &http.Client{}
	resp, err := client.Do(response)

	//If there is an error when sending the request to the API endpoint then stop the program and print an error message.
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	//Close the response.
	defer resp.Body.Close()

	//Create a new instance of the io.ReadCloser interface so it can be passed to the json.NewDecoder() method.
	body, _ := io.ReadAll(resp.Body)

	//Create a new decoder that reads from an io.Reader interface and saves it as a variable.
	var data companyOverview

	//Use the NewDecoder() method from the encoding/json package to read from an io.Reader interface and save it as a variable called decoder.
	decoder := json.NewDecoder(bytes.NewReader(body))

	//Decode the JSON-encoded data and store the result in the value pointed to by the variable called data.
	err = decoder.Decode(&data)

	//If there is an error when decoding the JSON-encoded data then stop the program and print an error message.
	if err != nil {
		fmt.Printf("The JSON decoding failed with error %s\n", err)
	}

	//print the data and its type.
	//TitleCards := TitleCard{data.Symbol, data.Name, data.Description, data.CIK, data.PERatio}

	return data
}

func printOverviewStatement(data companyOverview) {
	////print the name of the variable "Symbol" and Symbol as two separate columns using tabwriter
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Symbol:\t", data.Symbol)
	fmt.Fprintln(w, "Asset Type:\t", data.AssetType)
	fmt.Fprintln(w, "Name:\t", data.Name)
	fmt.Fprintln(w, "Description:\t", data.Description)
	fmt.Fprintln(w, "CIK:\t", data.CIK)
	fmt.Fprintln(w, "Exchange:\t", data.Exchange)
	fmt.Fprintln(w, "Currency:\t", data.Currency)
	fmt.Fprintln(w, "Country:\t", data.Country)
	fmt.Fprintln(w, "Sector:\t", data.Sector)
	fmt.Fprintln(w, "Industry:\t", data.Industry)
	fmt.Fprintln(w, "Address:\t", data.Address)
	fmt.Fprintln(w, "Fiscal Year End:\t", data.FiscalYearEnd)
	fmt.Fprintln(w, "Latest Quarter:\t", data.LatestQuarter)
	fmt.Fprintln(w, "Market Capitalization:\t", data.MarketCapitalization)
	fmt.Fprintln(w, "EBITDA:\t", data.EBITDA)
	fmt.Fprintln(w, "PE Ratio:\t", data.PERatio)
	fmt.Fprintln(w, "PEG Ratio:\t", data.PEGRatio)
	fmt.Fprintln(w, "Book Value:\t", data.BookValue)
	fmt.Fprintln(w, "Dividend Per Share:\t", data.DividendPerShare)
	fmt.Fprintln(w, "Dividend Yield:\t", data.DividendYield)
	fmt.Fprintln(w, "EPS:\t", data.EPS)
	fmt.Fprintln(w, "Revenue Per Share TTM:\t", data.RevenuePerShareTTM)
	fmt.Fprintln(w, "Profit Margin:\t", data.ProfitMargin)
	fmt.Fprintln(w, "Operating Margin TTM:\t", data.OperatingMarginTTM)
	fmt.Fprintln(w, "Return On Assets TTM:\t", data.ReturnOnAssetsTTM)
	fmt.Fprintln(w, "Return On Equity TTM:\t", data.ReturnOnEquityTTM)
	fmt.Fprintln(w, "Revenue TTM:\t", data.RevenueTTM)
	fmt.Fprintln(w, "Gross Profit TTM:\t", data.GrossProfitTTM)
	fmt.Fprintln(w, "Diluted EPS TTM:\t", data.DilutedEPSTTM)
	fmt.Fprintln(w, "Quarterly Earnings Growth YOY:\t", data.QuarterlyEarningsGrowthYOY)
	fmt.Fprintln(w, "Quarterly Revenue Growth YOY:\t", data.QuarterlyRevenueGrowthYOY)
	fmt.Fprintln(w, "Analyst Target Price:\t", data.AnalystTargetPrice)
	fmt.Fprintln(w, "Trailing PE:\t", data.TrailingPE)
	fmt.Fprintln(w, "Forward PE:\t", data.ForwardPE)
	fmt.Fprintln(w, "Price To Sales Ratio TTM:\t", data.PriceToSalesRatioTTM)
	fmt.Fprintln(w, "Price To Book Ratio:\t", data.PriceToBookRatio)
	fmt.Fprintln(w, "EV To Revenue:\t", data.EVToRevenue)
	fmt.Fprintln(w, "EV To EBITDA:\t", data.EVToEBITDA)
	fmt.Fprintln(w, "Beta:\t", data.Beta)
	fmt.Fprintln(w, "52 Week High:\t", data.X52WeekHigh)
	fmt.Fprintln(w, "52 Week Low:\t", data.X52WeekLow)
	fmt.Fprintln(w, "50 Day Moving Average:\t", data.X50DayMovingAverage)
	fmt.Fprintln(w, "200 Day Moving Average:\t", data.X200DayMovingAverage)
	fmt.Fprintln(w, "Shares Outstanding:\t", data.SharesOutstanding)
	fmt.Fprintln(w, "Dividend Date:\t", data.DividendDate)
	fmt.Fprintln(w, "Ex Dividend Date:\t", data.ExDividendDate)
	w.Flush()
}

// Create a struct that will hold the data from the API endpoint.
type companyOverview struct {
	Symbol                     string `json:"Symbol"`
	AssetType                  string `json:"AssetType"`
	Name                       string `json:"Name"`
	Description                string `json:"Description"`
	CIK                        string `json:"CIK"`
	Exchange                   string `json:"Exchange"`
	Currency                   string `json:"Currency"`
	Country                    string `json:"Country"`
	Sector                     string `json:"Sector"`
	Industry                   string `json:"Industry"`
	Address                    string `json:"Address"`
	FiscalYearEnd              string `json:"FiscalYearEnd"`
	LatestQuarter              string `json:"LatestQuarter"`
	MarketCapitalization       string `json:"MarketCapitalization"`
	EBITDA                     string `json:"EBITDA"`
	PERatio                    string `json:"PERatio"`
	PEGRatio                   string `json:"PEGRatio"`
	BookValue                  string `json:"BookValue"`
	DividendPerShare           string `json:"DividendPerShare"`
	DividendYield              string `json:"DividendYield"`
	EPS                        string `json:"EPS"`
	RevenuePerShareTTM         string `json:"RevenuePerShareTTM"`
	ProfitMargin               string `json:"ProfitMargin"`
	OperatingMarginTTM         string `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          string `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          string `json:"ReturnOnEquityTTM"`
	RevenueTTM                 string `json:"RevenueTTM"`
	GrossProfitTTM             string `json:"GrossProfitTTM"`
	DilutedEPSTTM              string `json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY string `json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  string `json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         string `json:"AnalystTargetPrice"`
	TrailingPE                 string `json:"TrailingPE"`
	ForwardPE                  string `json:"ForwardPE"`
	PriceToSalesRatioTTM       string `json:"PriceToSalesRatioTTM"`
	PriceToBookRatio           string `json:"PriceToBookRatio"`
	EVToRevenue                string `json:"EVToRevenue"`
	EVToEBITDA                 string `json:"EVToEBITDA"`
	Beta                       string `json:"Beta"`
	X52WeekHigh                string `json:"52WeekHigh"`
	X52WeekLow                 string `json:"52WeekLow"`
	X50DayMovingAverage        string `json:"50DayMovingAverage"`
	X200DayMovingAverage       string `json:"200DayMovingAverage"`
	SharesOutstanding          string `json:"SharesOutstanding"`
	DividendDate               string `json:"DividendDate"`
	ExDividendDate             string `json:"ExDividendDate"`
}
