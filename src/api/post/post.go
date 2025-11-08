package post

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"dario.cat/mergo"
	"github.com/CyberDruga/fanbox2discord/src/models/fanbox"
	"github.com/go-errors/errors"
)

type SortType string

type PaginateCreatorParams struct {
	CreatorId string
	Sort      string
}

/*
Gets the available pages from the creator's Fanbox page.

Default values:
  - Sort: "newest"
*/
func PaginateCreator(params PaginateCreatorParams) (result fanbox.Pages, err error) {

	defaultArgs := PaginateCreatorParams{
		Sort: "newest",
	}

	mergo.Merge(&params, defaultArgs)

	if params.CreatorId == "" {
		panic("CreatorId should never be empty")
	}

	url := fmt.Sprintf("https://api.fanbox.cc/post.paginateCreator?creatorId=%s&sort=%s", params.CreatorId, params.Sort)

	slog.Debug("Getting pages: " + url)

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Origin", "https://"+params.CreatorId+".fanbox.cc")

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		err = errors.Join(errors.New("Couldn't get data from post.paginateCreator"), err)
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)

	if result.Error != "" {
		err = errors.New(fmt.Sprintf("Api returned an error: %s", result.Error))
		slog.Warn(string(err.Error()))
		return
	}

	return
}

type ListCreatorParams struct {
	CreatorId              string
	FirstPublishedDatetime string
	FirstId                string
	Sort                   string
	Limit                  string
}

/*
Lists each post from the creator

Default values:
  - FirstPublishedDatetime: "2025-11-08%2012%3A00%3A00" // auto adjusted to current time
  - Sort:                   "newest"
  - Limit:                  "10"
*/
func ListCreator(params ListCreatorParams) (result fanbox.Posts, err error) {

	defaultArgs := ListCreatorParams{
		FirstPublishedDatetime: url.PathEscape(time.Now().Format(time.DateTime)),
		Sort:                   "newest",
		Limit:                  "10",
	}

	mergo.Merge(&params, defaultArgs)
	if params.CreatorId == "" {
		panic("CreatorId should never be empty")
	}

	if params.FirstId == "" {
		panic("FirstId should never be empty")
	}

	url := fmt.Sprintf(
		"https://api.fanbox.cc/post.listCreator?creatorId=%s&firstPublishedDatetime=%s&firstId=%s&sort=%s&limit=%s",
		params.CreatorId,
		params.FirstPublishedDatetime,
		params.FirstId,
		params.Sort,
		params.Limit,
	)

	slog.Debug("Getting pages: " + url)

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Origin", "https://"+params.CreatorId+".fanbox.cc")

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		err = errors.Join(errors.New("Couldn't get data from post.listCreator"), err)
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)

	if result.Error != "" {
		err = errors.New(fmt.Sprintf("Api returned an error: %s", result.Error))
		slog.Warn(err.Error())
		return
	}

	if len(result.Body) == 0 {
		err = errors.New("Api returned no posts")
		slog.Warn(err.Error())
		return
	}

	return

}
