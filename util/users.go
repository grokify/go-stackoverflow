package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/antihax/optional"
	stackoverflow "github.com/grokify/go-stackoverflow/client"
	"github.com/grokify/gocharts/data/table"
	"github.com/grokify/gotilla/io/ioutilmore"
	"github.com/grokify/gotilla/time/timeutil"
)

const PerPageMax = uint32(100)

// UsersResponseFromFile reads afile into a `stackoverflow.UsersResponse`
// struct.
func UsersResponseFromFile(file string) (stackoverflow.UsersResponse, error) {
	ur := stackoverflow.UsersResponse{}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return ur, err
	}
	err = json.Unmarshal(bytes, &ur)
	return ur, err
}

type UsersSet struct {
	UsersMap map[string]stackoverflow.User `json:"users"`
}

func NewUsersSet() UsersSet {
	return UsersSet{UsersMap: map[string]stackoverflow.User{}}
}

func (us *UsersSet) AddUsers(users []stackoverflow.User) {
	for _, user := range users {
		us.AddUser(user)
	}
}

func (us *UsersSet) AddUser(user stackoverflow.User) {
	idString := strconv.Itoa(int(user.UserId))
	if us.UsersMap == nil {
		us.UsersMap = map[string]stackoverflow.User{}
	}
	us.UsersMap[idString] = user
}

func (us *UsersSet) AddUsersResponseFiles(dir string, rx *regexp.Regexp) error {
	_, filepaths, err := ioutilmore.ReadDirRx(dir, rx, true)
	if err != nil {
		return err
	}
	for _, filepath := range filepaths {
		err := us.AddUsersResponseFile(filepath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (us *UsersSet) AddUsersResponseFile(file string) error {
	ur, err := UsersResponseFromFile(file)
	if err != nil {
		return err
	}
	us.AddUsersResponse(&ur)
	return nil
}

func (us *UsersSet) AddUsersResponse(ur *stackoverflow.UsersResponse) {
	if ur == nil {
		return
	}
	for _, user := range ur.Items {
		idStr := strconv.Itoa(int(user.UserId))
		us.UsersMap[idStr] = user
	}
}

func (us *UsersSet) Slice() []stackoverflow.User {
	users := []stackoverflow.User{}
	for _, usr := range us.UsersMap {
		users = append(users, usr)
	}
	return users
}

func (us *UsersSet) Table(cols []string) *table.Table {
	if len(cols) == 0 {
		cols = []string{"display_name", "reputation", "last_access_date", "location", "link", "about_me"}
	}
	tbl := table.NewTable()
	tbl.Columns = cols
	for _, usr := range us.UsersMap {
		row := []string{}
		for _, col := range cols {
			col = strings.ToLower(strings.TrimSpace(col))
			switch col {
			case "about_me":
				row = append(row, usr.AboutMe)
			case "display_name":
				row = append(row, usr.DisplayName)
			case "last_access_date":
				t := time.Unix(usr.LastAccessDate, 0)
				row = append(row, t.Format(timeutil.DateMDY))
			case "link":
				row = append(row, usr.Link)
			case "location":
				row = append(row, usr.Location)
			case "reputation":
				row = append(row, strconv.Itoa(int(usr.Reputation)))
			default:
				row = append(row, "")
			}
		}
		tbl.Records = append(tbl.Records, row)
	}
	return &tbl
}

func UsersResponseToUsersSet(ur *stackoverflow.UsersResponse) *UsersSet {
	usersSet := UsersSet{}
	if ur == nil {
		return &usersSet
	}
	for _, user := range ur.Items {
		usersSet.AddUser(user)
	}
	return &usersSet
}

func GetUsersMore(apiClient *stackoverflow.APIClient, site string, opts *stackoverflow.GetUsersOpts, perPage, numPages uint32) (UsersSet, error) {
	users := NewUsersSet()
	if perPage == 0 || numPages == 0 {
		return users, nil
	}
	if opts == nil {
		opts = &stackoverflow.GetUsersOpts{}
	}

	for pageNum := uint32(1); pageNum <= numPages; pageNum++ {
		info, resp, err := apiClient.UsersApi.GetUsers(
			context.Background(),
			site,
			&stackoverflow.GetUsersOpts{
				Key:      optional.NewString(os.Getenv("STACKOVERFLOW_KEY")),
				Sort:     optional.NewString("reputation"),
				Pagesize: optional.NewInt32(int32(perPage)),
				Page:     optional.NewInt32(int32(pageNum)),
			},
		)
		if err != nil {
			return users, nil
		} else if resp.StatusCode >= 300 {
			return users, fmt.Errorf("stackoverflow/users.GetUsersMore API StatusCode [%v]", resp.StatusCode)
		}
		users.AddUsers(info.Items)
	}

	return users, nil
}
