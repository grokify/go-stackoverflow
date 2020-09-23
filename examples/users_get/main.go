package main

import (
	"context"
	"fmt"
	"log"

	stackoverflow "github.com/grokify/go-stackoverflow/client"
	"github.com/grokify/gotilla/fmt/fmtutil"
)

func main() {
	apiClient := stackoverflow.NewAPIClient(
		stackoverflow.NewConfiguration())

	site := "stackoverflow"

	info, resp, err := apiClient.UsersApi.GetUsers(
		context.Background(),
		site,
		&stackoverflow.GetUsersOpts{})

	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(info)
	fmt.Println("DONE")
}
