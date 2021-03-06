package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/grokify/go-stackoverflow/util"
	"github.com/grokify/gocharts/data/table"
	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"
)

func main() {
	loaded, err := config.LoadDotEnvSkipEmptyInfo(".env")
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(loaded)

	usersSet := util.NewUsersSet()
	usersSet.AddUsersResponseFiles("../users_get/users2",
		regexp.MustCompile(`users_pg-(\d+)\-(\d+)\.json`))

	fmt.Printf("NUM_USERS [%v]\n", len(usersSet.UsersMap))

	rx := regexp.MustCompile(`(?i)(\bPM\b|\bproduct\s+manage)`)

	usersSet2 := util.NewUsersSet()
	i := 0
	for _, usr := range usersSet.UsersMap {
		if rx.MatchString(usr.AboutMe) {
			i++
			fmt.Printf("[%v] LINK [%v]\n", i, usr.Link)
			usersSet2.AddUser(usr)
		}
	}
	tbl := usersSet2.Table([]string{})
	tbl.Name = "PM"
	table.WriteXLSX("_users.xlsx", tbl)

	fmt.Println("DONE")
}
