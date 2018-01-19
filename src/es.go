package main

import (
	"bytes"
	"encoding/json"
	"time"
)

type Es struct {
}

type EsSearchConditionModel struct {
	Query struct {
		BoolCondition struct {
			Must struct {
				QueryString struct {
					Query string `json:"query"`
				} `json:"query_string"`
			} `json:"must"`
			Filter struct {
				Range struct {
					Timestamp struct {
						Gte int `json:"gte"`
					} `json:"@timestamp"`
				} `json:"range"`
			} `json:"filter"`
		} `json:"bool"`
	} `json:"query"`
}

type EsSearchResultModel struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total    int `json:"total"`
		MaxScore int `json:"max_score"`
		Hits     []struct {
			Index  string `json:"_index"`
			Type   string `json:"_type"`
			ID     string `json:"_id"`
			Score  int    `json:"_score"`
			Source struct {
				Headers struct {
					ContentType        string `json:"content_type"`
					HTTPAcceptEncoding string `json:"http_accept_encoding"`
					RequestPath        string `json:"request_path"`
					HTTPVersion        string `json:"http_version"`
					HTTPConnection     string `json:"http_connection"`
					RequestMethod      string `json:"request_method"`
					HTTPHost           string `json:"http_host"`
					RequestURI         string `json:"request_uri"`
					ContentLength      string `json:"content_length"`
					HTTPUserAgent      string `json:"http_user_agent"`
				} `json:"headers"`
				Timestamp           time.Time `json:"@timestamp"`
				Version             string    `json:"@version"`
				Host                string    `json:"host"`
				Index               string    `json:"index"`
				MethodName          string    `json:"methodName"`
				ClassName           string    `json:"className"`
				From                string    `json:"from"`
				MsgList             []string  `json:"msgList"`
				Type                string    `json:"type"`
				ElapsedMilliseconds int       `json:"elapsedMilliseconds"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (es Es) Search(condition EsSearchConditionModel) EsSearchResultModel {
	var result EsSearchResultModel
	var url = "http://domain:9200/prod-ecgw*/prod-ecgw/_search"
	jsonBytes, err := json.Marshal(condition)
	if err != nil {
		// handle error
		//log.Fatal(err)
	} else {
		var client HttpClient
		resultBytes := client.Post(url, bytes.NewReader(jsonBytes))
		json.Unmarshal(resultBytes, &result)
	}
	return result
}
