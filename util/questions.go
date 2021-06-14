package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/antihax/optional"
	stackoverflow "github.com/grokify/go-stackoverflow/client"
	"github.com/grokify/gocharts/data/statictimeseries"
	"github.com/pkg/errors"
)

func GetQuestionsAll(apiClient *stackoverflow.APIClient, site string, opts *stackoverflow.GetQuestionsOpts) ([]stackoverflow.Question, error) {
	if apiClient == nil {
		apiClient = stackoverflow.NewAPIClient(
			stackoverflow.NewConfiguration())
	}
	site = strings.TrimSpace(site)
	if len(site) == 0 {
		site = SiteStackOverflow
	}
	if opts == nil {
		opts = &stackoverflow.GetQuestionsOpts{
			Page:     optional.NewInt32(1),
			Pagesize: optional.NewInt32(int32(PerPageMax))}
	}
	questions := []stackoverflow.Question{}
	page := int32(1)

	info, resp, err := apiClient.QuestionsApi.GetQuestions(
		context.Background(), site, opts)
	if err != nil {
		httpBody := APIErrorBody(err)
		return questions, errors.Wrap(err, string(httpBody))
	} else if resp.StatusCode > 299 {
		httpBody := APIErrorBody(err)
		return questions, fmt.Errorf("%s API HTTP status [%d]", string(httpBody), resp.StatusCode)
	}
	questions = append(questions, info.Items...)

	for info.HasMore {
		page++
		opts.Page = optional.NewInt32(page)
		info, resp, err = apiClient.QuestionsApi.GetQuestions(
			context.Background(), site, opts)
		if err != nil {
			return questions, err
		} else if resp.StatusCode > 299 {
			return questions, fmt.Errorf("API status [%d]", resp.StatusCode)
		}
		questions = append(questions, info.Items...)
	}

	return questions, nil
}

func ReadQuestionsFile(file string) ([]stackoverflow.Question, error) {
	questions := []stackoverflow.Question{}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return questions, err
	}
	err = json.Unmarshal(bytes, &questions)
	if err != nil {
		return questions, err
	}
	questions = QuestionsDedupe(questions)
	return questions, nil
}

func QuestionsDedupe(questions []stackoverflow.Question) []stackoverflow.Question {
	deduped := []stackoverflow.Question{}
	seen := map[int32]int{}
	for _, q := range questions {
		if _, ok := seen[q.QuestionId]; !ok {
			deduped = append(deduped, q)
			seen[q.QuestionId] = 1
		}
	}
	return deduped
}

func QuestionsToDataSeries(dataseriesName string, questions []stackoverflow.Question) (statictimeseries.DataSeries, error) {
	ds := statictimeseries.NewDataSeries()
	questions = QuestionsDedupe(questions)

	for _, q := range questions {
		creationDt := time.Unix(int64(q.CreationDate), 0)
		ds.AddItem(statictimeseries.DataItem{
			SeriesName: dataseriesName,
			Time:       creationDt.UTC(),
			Value:      1})
	}
	return ds, nil
}
