package main

import (
	http_request "broadband-speed-monitor/src/http"
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type requestBody struct{}
type responseBody struct {
	StatusRate ArrayVal `xml:"status_rate"`
}

type ArrayVal struct {
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

func RouterRequest() (int, int) {

	// Make HTTP request to router
	var body requestBody
	var result responseBody
	if err := http_request.HttpRequest("http://192.168.1.254/nonAuth/wan_conn.xml", http_request.GET, body, &result); err != nil {
		log.Println(err)
		return 0, 0
	}

	// Parse Value and Get Up and Down rates
	parsedValue, err := parseArrayVal(result.StatusRate.Value)
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	if len(parsedValue) < 2 || len(parsedValue[1]) < 1 {
		log.Println("parsedValue out of range")
		return 0, 0
	}
	rateStr := parsedValue[1][0]
	rates := strings.Split(rateStr, ";")
	upRate, _ := strconv.Atoi(rates[0])
	downRate, _ := strconv.Atoi(rates[1])

	return upRate / 1000000, downRate / 1000000
}

func parseArrayVal(raw string) ([][]string, error) {
	// Step 1: URL decode
	decoded, err := url.QueryUnescape(raw)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert Python-style to JSON-style
	// Replace single quotes with double quotes
	jsonLike := strings.ReplaceAll(decoded, "'", `"`)

	// Step 3: Unmarshal into Go slice
	var result [][]string
	if err := json.Unmarshal([]byte(jsonLike), &result); err != nil {
		return nil, err
	}

	return result, nil
}
