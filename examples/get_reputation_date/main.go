package main

import (
	"fmt"
	"log"

	"github.com/grokify/gotilla/fmt/fmtutil"

	"github.com/grokify/go-stackexchange/client"
	"github.com/grokify/go-stackexchange/clientutil"
)

func main() {
	apiClient := stackexchange.NewAPIClient(
		stackexchange.NewConfiguration())

	site := "stackoverflow"
	userId := "1908967"

	history, err := stackexchangeutil.GetReputationHistoryAll(apiClient, site, userId)
	if err != nil {
		log.Fatal(err)
	}

	fmtutil.PrintJSON(history)
	fmt.Printf("COUNT [%v]\n", len(history.Items))

	repByDay := history.ReputationChangeByDayMap()
	fmtutil.PrintJSON(repByDay)

	dayRep, err := history.DateForReputation(int32(2000))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DATE [%v]\n", dayRep.Day)

	fmt.Println("DONE")
}
