package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/antihax/optional"
	"github.com/pkg/errors"

	stackoverflow "github.com/grokify/go-stackoverflow/client"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/io/ioutilmore"
	"github.com/grokify/mogo/os/executil"
	"github.com/grokify/mogo/os/osutil"
	"github.com/grokify/mogo/type/stringsutil"
)

func main() {
	loaded, err := config.LoadDotEnvSkipEmptyInfo(".env")
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(loaded)

	apiClient := stackoverflow.NewAPIClient(
		stackoverflow.NewConfiguration())

	site := "stackoverflow"

	rename := false
	getFiles := true
	if 1 == 0 {
		rename = true
		getFiles = false
	}
	if rename {
		files, err := osutil.ReadDirMore(".", regexp.MustCompile(`users.+\.json`), false, true, false)
		if err != nil {
			log.Fatal(err)
		}
		rx := regexp.MustCompile(`users_pg-(\d+)\-(\d+)\.json`)
		for _, fi := range files {
			fmt.Println(fi.Name())
			name := fi.Name()
			m := rx.FindAllStringSubmatch(name, -1)
			if len(m) > 0 {
				fmtutil.PrintJSON(m)
				pageNum := m[0][1]
				numPages := m[0][2]
				if len(pageNum) == len(numPages) {
					continue
				}
				lenNum := len(m[0][2])
				fmt.Println(lenNum)

				output := stringsutil.PadLeft(pageNum, "0", lenNum)
				fmt.Printf("OUTPUT [%v]\n", output)

				newfile := fmt.Sprintf("users_pg-%v-%v.json", output, numPages)
				fmt.Println(newfile)
				cmd := "mv " + fi.Name() + " " + newfile
				fmt.Println(cmd)
				_, _, err = executil.ExecSimple(cmd)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("SUCC %v\n", cmd)
				//panic("A")
			}
		}
	}
	// b_t)r81Rlo6RfD
	// b_t)r81Rlo6RfD
	//!b_t)r81Rlo6RfD
	// !b_t)r81Rlo6Riz
	if getFiles {
		perPage := uint32(100)
		numPages := uint32(1500)
		opts := &stackoverflow.GetUsersOpts{
			Key:      optional.NewString(os.Getenv("STACKOVERFLOW_KEY")),
			Sort:     optional.NewString("reputation"),
			Order:    optional.NewString("desc"),
			Filter:   optional.NewString("!b_t)r81Rlo6RfD"),
			Pagesize: optional.NewInt32(100),
			Page:     optional.NewInt32(1),
		}
		// opts.Filter = optional.NewString("%21b_t%29r81Rlo6RfD")
		fmtutil.PrintJSON(opts)
		err := GetMore(apiClient, site, opts, perPage, numPages)
		if err != nil {
			log.Fatal(err)
		}
	}
	if 1 == 0 {
		info, resp, err := apiClient.UsersApi.GetUsers(
			context.Background(),
			site,
			&stackoverflow.GetUsersOpts{
				Key:      optional.NewString(os.Getenv("STACKOVERFLOW_KEY")),
				Sort:     optional.NewString("reputation"),
				Order:    optional.NewString("desc"),
				Pagesize: optional.NewInt32(100),
				Page:     optional.NewInt32(1),
			},
		)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode >= 300 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	}
	fmt.Println("DONE")
}

func GetMore(apiClient *stackoverflow.APIClient, site string, opts *stackoverflow.GetUsersOpts, perPage, numPages uint32) error {
	if opts == nil {
		opts = &stackoverflow.GetUsersOpts{}
	}
	opts.Pagesize = optional.NewInt32(int32(perPage))
	fmtutil.PrintJSON(opts)
	fmt.Println(opts.Filter.Value())
	//panic("A")
	// https://api.stackexchange.com/2.2/users?order=desc&sort=reputation&site=stackoverflow&filter=!b_t)r81Rlo6RfD
	// https://api.stackexchange.com/2.2/users?order=desc&sort=reputation&site=stackoverflow&filter=%21b_t%29r81Rlo6RfD
	i := 0
	for pageNum := uint32(1497); pageNum <= numPages; pageNum++ {
		opts.Page = optional.NewInt32(int32(pageNum))
		info, resp, err := apiClient.UsersApi.GetUsers(
			context.Background(), site, opts)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Page Num [%v]", pageNum))
		} else if resp.StatusCode >= 300 {
			return fmt.Errorf("Page Num [%v] Status [%v]", pageNum, resp.StatusCode)
		}
		filename := fmt.Sprintf("users_pg-%v-%v.json", pageNum, numPages)
		err = ioutilmore.WriteFileJSON(filename, info, 0644, "", "  ")
		if err != nil {
			return fmt.Errorf("Page [%v] Error [%v]", pageNum, err.Error())
		}
		fmt.Printf("WROTE [%v]", filename)
		i++
		if i == 178 {
			break
		}
	}
	return nil
}
