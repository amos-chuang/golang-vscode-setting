package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var esQueryRangeHr = 24
	if len(os.Args) > 1 {
		temp, err := strconv.ParseInt(os.Args[1], 10, 32)
		if err == nil && temp > 0 {
			esQueryRangeHr = int(temp)
		}
	}
	var es Es
	var condition EsSearchConditionModel
	condition.Query.BoolCondition.Must.QueryString.Query = "*clearsale* AND *send* AND *invalid*"
	condition.Query.BoolCondition.Filter.Range.Timestamp.Gte = int(time.Now().UnixNano()/1000000) - int(esQueryRangeHr*60*60*1000)
	result := es.Search(condition)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Hits.Total : " + strconv.Itoa(result.Hits.Total))
	for _, hit := range result.Hits.Hits {
		var jsonStr = hit.Source.MsgList[1]
		var beginWith = "Request Payload: "
		var begin = strings.Index(jsonStr, beginWith) + len(beginWith)
		jsonStr = jsonStr[begin:]
		beginWith = "\"ID\":\""
		begin = strings.Index(jsonStr, beginWith) + len(beginWith)
		var orderID = jsonStr[begin:]
		orderID = orderID[:strings.Index(orderID, "\"")]
		fmt.Println("ID : " + orderID)
	}
}
