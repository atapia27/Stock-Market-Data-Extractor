package main

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
)

// downloadData downloads the data from the API
func downloadData() error {
	url := "https://www.alphavantage.co/query?function=LISTING_STATUS&apikey=GG67ANDCMVCCRKPJ"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	f, err := os.Create("csvFiles/currentListings.csv")
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// getSymbolMatrix() returns a matrix of the data from the csv file
func getSymbolMatrix() [][]string {
	f, err := os.Open("csvFiles/currentListings.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	var symbol []string
	var name []string
	var exchange []string
	var assetType []string

	for i, record := range records {

		if i == 0 {
		} else {
			symbol = append(symbol, record[0])
			name = append(name, record[1])
			exchange = append(exchange, record[2])
			assetType = append(assetType, record[3])
		}
	}

	matrix := [][]string{symbol, name, exchange, assetType}
	return matrix
}
