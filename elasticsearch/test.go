package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"strings"
)

func main() {
	var r map[string]interface{}
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://112.74.205.92:9200",
		},
		Username: "elastic",
		Password: "gramyang##88",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("Error creating the client: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		fmt.Printf("Error getting response: %s", err)
	}
	defer res.Body.Close() //一定要关闭
	if res.IsError() {
		fmt.Printf("Error: %s", res.String())
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		fmt.Printf("Error parsing the response body: %s", err)
	}
	fmt.Printf("Client: %s", elasticsearch.Version)
	fmt.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	fmt.Println(strings.Repeat("~", 37))

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"match": map[string]interface{}{
						"description": "channel",
					}},
					{"match": map[string]interface{}{
						"title": "rust",
					}},
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		fmt.Printf("Error encoding query: %s", err)
	}
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("article"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		fmt.Printf("Error getting response: %s", err)
	}
	defer res.Body.Close() //响应体必须关闭
	fmt.Println(res)
}
