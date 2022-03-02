package main

import (
	"fmt"
	"log"

	"github.com/antihax/optional"
	stackoverflow "github.com/grokify/go-stackoverflow/client"
	"github.com/grokify/go-stackoverflow/util"
	"github.com/grokify/simplego/io/ioutilmore"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Tag        string `short:"t" long:"tag" description:"Tag"`
	Site       string `short:"s" long:"site" description:"Site"`
	OutputFile string `short:"o" long:"output" description:"Output file" required:"true"`
	Retrieve   []bool `short:"r" long:"retrieve" description:"Retrieve via API"`
}

func main() {
	args := Options{}
	_, err := flags.Parse(&args)
	if err != nil {
		log.Fatal(err)
	}

	if len(args.Retrieve) > 0 {
		opts := &stackoverflow.GetQuestionsOpts{
			Tagged: optional.NewString(args.Tag)}
		questions, err := util.GetQuestionsAll(nil, args.Site, opts)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutilmore.WriteFileJSON(args.OutputFile, questions, 0644, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
	}

	questions, err := util.ReadFileQuestions(args.OutputFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("QUESTIONS [%d]\n", len(questions))

	ds := util.QuestionsToDataSeries("questions by month", questions)
	ds = ds.ToMonth(true)
	err = ds.WriteXLSX("_question_count_month.xlsx", "Q per month", "month", "count")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DONE")
}
