package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type ZIPResult struct {
	XMLName      xml.Name `xml:"ZIP_result"`
	Text         string   `xml:",chardata"`
	Result       []Result `xml:"result"`
	ADDRESSValue Address  `xml:"ADDRESS_value"`
}

type Result struct {
	Text              string `xml:",chardata"`
	Name              string `xml:"name,attr"`
	Version           string `xml:"version,attr"`
	RequestURL        string `xml:"request_url,attr"`
	RequestZipNum     string `xml:"request_zip_num,attr"`
	RequestZipVersion string `xml:"request_zip_version,attr"`
	ResultCode        string `xml:"result_code,attr"`
	ResultZipNum      string `xml:"result_zip_num,attr"`
	ResultZipVersion  string `xml:"result_zip_version,attr"`
	ResultValuesCount string `xml:"result_values_count,attr"`
}

type Address struct {
	Text  string  `xml:",chardata"`
	Value []Value `xml:"value"`
}

type Value struct {
	Text        string `xml:",chardata"`
	StateKana   string `xml:"state_kana,attr"`
	CityKana    string `xml:"city_kana,attr"`
	AddressKana string `xml:"address_kana,attr"`
	CompanyKana string `xml:"company_kana,attr"`
	State       string `xml:"state,attr"`
	City        string `xml:"city,attr"`
	Address     string `xml:"address,attr"`
	Company     string `xml:"company,attr"`
}

func main() {

	// Open our xmlFile
	xmlFile, err := os.Open("test.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened test.xml")

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var results ZIPResult
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'results' which we defined above
	err = xml.Unmarshal(byteValue, &results)
	if err != nil {
		fmt.Println(err)
	}
	// print the parsed XML data
	//fmt.Printf("%+v", results)
	//fmt.Println()

	/*for _, result := range results.Result {
		fmt.Printf("Name: %s, Version: %s, Request URL: %s, Result Code: %s, Result Zip Num: %s, Result Zip Version: %s\n", result.Name, result.Version, result.RequestURL, result.ResultCode, result.ResultZipNum, result.ResultZipVersion)
	}
	*/

	for _, value := range results.ADDRESSValue.Value {
		fmt.Printf("State Kana: %s, City Kana: %s, Address Kana: %s, Company Kana: %s, State: %s, City: %s, Address: %s, Company: %s\n", value.StateKana, value.CityKana, value.AddressKana, value.CompanyKana, value.State, value.City, value.Address, value.Company)

	}
}
