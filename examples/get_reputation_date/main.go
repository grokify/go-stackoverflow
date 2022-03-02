package main

import (
	"fmt"
	"log"

	"github.com/grokify/mogo/fmt/fmtutil"

	stackexchange "github.com/grokify/go-stackoverflow/client"
	"github.com/grokify/go-stackoverflow/util"
)

func main() {
	apiClient := stackexchange.NewAPIClient(
		stackexchange.NewConfiguration())

	site := "stackoverflow"
	userId := "1908967"

	history, err := util.GetReputationHistoryAll(apiClient, site, userId)
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
