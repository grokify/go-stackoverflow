package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grokify/go-stackexchange/client"
	"github.com/grokify/gotilla/fmt/fmtutil"
)

func main() {
	apiClient := stackexchange.NewAPIClient(
		stackexchange.NewConfiguration())

	site := "stackoverflow"

	info, resp, err := apiClient.UsersApi.GetUsers(
		context.Background(),
		site,
		&stackexchange.GetUsersOpts{})

	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(info)
	fmt.Println("DONE")
}
