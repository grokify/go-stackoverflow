package util

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/grokify/go-stackexchange/client"
	"github.com/grokify/gotilla/time/timeutil"
	"github.com/grokify/gotilla/type/maputil"
)

type History struct {
	Items []stackexchange.ReputationHistory
}

func (history *History) ReputationChangeByDayMap() map[string]int32 {
	repByDay := map[string]int32{}
	for _, i := range history.Items {
		t := timeutil.UnixToDay(int64(i.CreationDate))
		ymd := t.Format(timeutil.RFC3339FullDate)
		if _, ok := repByDay[ymd]; !ok {
			repByDay[ymd] = 0
		}
		repByDay[ymd] += i.ReputationChange
	}
	return repByDay
}

func (history *History) ReputationChangeByDaySlice() []DayReputation {
	days := []DayReputation{}
	data := history.ReputationChangeByDayMap()
	dates := maputil.StringKeysSorted(data)
	for _, date := range dates {
		if rep, ok := data[date]; ok {
			dayRep := DayReputation{
				Day:           date,
				DayReputation: rep}
			l := len(days)
			if l > 0 {
				dayRep.TotalReputation = days[l-1].TotalReputation + dayRep.DayReputation
			} else {
				dayRep.TotalReputation = dayRep.DayReputation
			}
			days = append(days, dayRep)
		}
	}
	return days
}

func (history *History) DateForReputation(rep int32) (DayReputation, error) {
	days := history.ReputationChangeByDaySlice()
	for _, day := range days {
		if day.TotalReputation >= rep {
			return day, nil
		}
	}
	return DayReputation{}, fmt.Errorf("Reputation did not exceed [%v]", rep)
}

type DayReputation struct {
	Day             string
	DayReputation   int32
	TotalReputation int32
}

func GetReputationHistoryAll(apiClient *stackexchange.APIClient, site, userIds string) (History, error) {
	history := History{Items: []stackexchange.ReputationHistory{}}
	nextPage := int32(1)

	getNextPage := true

	for getNextPage {
		info, resp, err := apiClient.UsersApi.GetUsersReputationHistory(
			context.Background(),
			userIds,
			site,
			&stackexchange.GetUsersReputationHistoryOpts{
				Page:     optional.NewInt32(nextPage),
				Pagesize: optional.NewInt32(100)})

		if err != nil {
			return history, err
		} else if resp.StatusCode >= 300 {
			apiErr := fmt.Errorf("Error Status: %v", resp.StatusCode)
			return history, apiErr
		}
		history.Items = append(history.Items, info.Items...)
		nextPage += 1
		getNextPage = info.HasMore
	}

	return history, nil
}
