/*
(B) Companies Financial Statements API

The Financial Modeling Prep Company Audited Financials API provides company financial statements data as reported to the United States
Securities and Exchange Commission ("SEC") by financial institutions, including income statements, balance sheets, and cash flow
statements, which are the three major statements revealing a company’s financial health. The data is standardized and audited before
making it available for past and present values, allowing you to analyze financial records from previous and present years. The
Corporation Financial Reports API gives you a full list of firm reports as text or direct links to the right company's financial report,
where you can easily download them.

Income Statement API:
Income statements are detailed reports that show how much profit or loss the company has generated over a particular year. The income
statement, also known as profit or loss account, reports the company expenses that arise from the business operations
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
)

func getDailyStatements(symbol string) dailyStatement {

	//Initialize the necessary variables, such as the API key or credentials, endpoint, and method.
	var (
		apiKey = "GG67ANDCMVCCRKPJ"                                                                                             //The API key that I have to get from the API site.
		url    = "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=" + symbol + "&apikey=" + apiKey //The API endpoint that I am going to use for the API request.
		method = "GET"                                                                                                          //The method that I am going to use for the API request.
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
	var data dailyStatement

	//Use the NewDecoder() method from the encoding/json package to read from an io.Reader interface and save it as a variable called decoder.
	decoder := json.NewDecoder(bytes.NewReader(body))

	//Use the decoder variable's Decode() method to decode the JSON data and save it as a variable called err.
	err = decoder.Decode(&data)

	//If there is an error when decoding the JSON data then stop the program and print an error message.
	if err != nil {
		fmt.Printf("The JSON decoding failed with error %s\n", err)
	}

	return data
}

func printDailyStatements(data dailyStatement) {

	//Print the data for each day using tabwriter to format the output.
	w := new(tabwriter.Writer)

	//
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	//get todays date using time.Now() and format it as YYYY-MM-DD
	calendarDay := time.Now().Format("2006-01-02")

	//get dayOfTheWeek of the week using time.Now().Weekday() and format it as a string
	dayOfTheWeek := time.Now().Weekday().String()

	//create a map of the calendarDay and the dayOfTheWeek for the past 7 days using time.Now().AddDate(0, 0, -1).Format("2006-01-02") and save it as a variable called week
	week := map[string]string{
		calendarDay: dayOfTheWeek,
		time.Now().AddDate(0, 0, -1).Format("2006-01-02"): time.Now().AddDate(0, 0, -1).Weekday().String(), //yesterday
		time.Now().AddDate(0, 0, -2).Format("2006-01-02"): time.Now().AddDate(0, 0, -2).Weekday().String(), //2 days ago
		time.Now().AddDate(0, 0, -3).Format("2006-01-02"): time.Now().AddDate(0, 0, -3).Weekday().String(), //3 days ago
		time.Now().AddDate(0, 0, -4).Format("2006-01-02"): time.Now().AddDate(0, 0, -4).Weekday().String(), //4 days ago
		time.Now().AddDate(0, 0, -5).Format("2006-01-02"): time.Now().AddDate(0, 0, -5).Weekday().String(), //5 days ago
		time.Now().AddDate(0, 0, -6).Format("2006-01-02"): time.Now().AddDate(0, 0, -6).Weekday().String(), //6 days ago
	}

	//print "Daily Statements for the past 7 days"
	fmt.Fprintln(w, "Daily Statements for the past 7 days", "\n")

	//loop through the week map and print the data for each day, excluding weekends
	for currentCalendarDay := range week {

		if week[currentCalendarDay] == "Saturday" || week[currentCalendarDay] == "Sunday" {
			//do nothing
		} else {
			//fmt.Fprintln(w, "Date: "+week[currentCalendarDay]+", "+currentCalendarDay)
			fmt.Fprintln(w, "Date: "+week[currentCalendarDay]+", "+currentCalendarDay, "\t", "Open\t", "High\t", "Low\t", "Close\t", "Adjusted Close\t", "Volume\t")
			//print a divider line for each field in the data
			fmt.Fprintln(w, "\t", "------\t", "------\t", "------\t", "------\t", "--------------\t", "--------\t")
			fmt.Fprintln(w, "\t", (data.TimeSeriesDailyAdjusted[currentCalendarDay].Open + "0")[:6], "\t", (data.TimeSeriesDailyAdjusted[currentCalendarDay].High + "0")[:6], "\t", (data.TimeSeriesDailyAdjusted[currentCalendarDay].Low + "0")[:6], "\t", (data.TimeSeriesDailyAdjusted[currentCalendarDay].Close + "0")[:6], "\t", data.TimeSeriesDailyAdjusted[currentCalendarDay].AdjustedClose, "\t", data.TimeSeriesDailyAdjusted[currentCalendarDay].Volume)
			fmt.Fprintln(w)
		}
	}
	w.Flush()
}

// Create a struct that will contain all the information of the JSON data.
type dailyStatement struct {
	MetaData struct {
		Information string `json:"1. Information"`
		Symbol      string `json:"2. Symbol"`
		LastRefresh string `json:"3. Last Refreshed"`
		OutputSize  string `json:"4. Output Size"`
		TimeZone    string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeriesDailyAdjusted map[string]struct {
		Open          string `json:"1. open"`
		High          string `json:"2. high"`
		Low           string `json:"3. low"`
		Close         string `json:"4. close"`
		AdjustedClose string `json:"5. adjusted close"`
		Volume        string `json:"6. volume"`
		Dividend      string `json:"7. dividend amount"`
	} `json:"Time Series (Daily)"`
}
