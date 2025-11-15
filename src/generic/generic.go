package generic

import (
	"fmt"
	"log/slog"
	"regexp"

	"github.com/CyberDruga/fanbox2discord/src/api/post"
	"github.com/CyberDruga/fanbox2discord/src/models/fanbox"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-errors/errors"
)

func GetLatestPostId(creatorId string) (result string, err error) {
	if creatorId == "" {
		panic("[channelId] should never be empty")
	}

	slog.Debug("Getting")
	pages, err := post.PaginateCreator(post.PaginateCreatorParams{
		CreatorId: creatorId,
	})

	if err != nil {
		errors.Join(err, errors.New("Can't get the latest posts"))
		return
	}

	if len(pages.Body) < 1 {
		err = errors.New("Couldn't get any pages from Fanbox")
		return
	}

	firstPage := pages.Body[0]

	slog.Debug("Compiling Regex")
	regex, err := regexp.Compile(`firstId=(\d+)`)
	if err != nil {
		err = errors.Join(err, errors.New("Couldn't compile regex"))
		return
	}

	slog.Debug("Executing Regex")
	matches := regex.FindStringSubmatch(firstPage)

	slog.Debug(fmt.Sprintf("Matches: %v", matches))

	if matches == nil {
		err = errors.New("Couldn't find matching string")
		return
	}

	if len(matches) < 2 {
		err = errors.New("Couldn't find matching string")
		return
	}

	result = matches[1]

	if result == "" {
		err = errors.Errorf("Somehow result is empty")
	}

	return

}

func GetLatestPosts(creatorId string) (result []fanbox.Post, err error) {

	if creatorId == "" {
		panic("[cretorId] can never be empty")
	}

	slog.Debug("Getting latest post ID from: " + creatorId)
	postId, err := GetLatestPostId(creatorId)

	if err != nil {
		err = errors.Join(errors.Errorf("Couldn't get post id"), err)
	}

	slog.Debug("Getting posts from: " + creatorId)
	posts, err := post.ListCreator(post.ListCreatorParams{
		CreatorId: creatorId,
		FirstId:   postId,
	})

	if err != nil {
		err = errors.Join(errors.Errorf("Couldn't get posts"), err)
	}

	slog.Debug("After getting posts", "posts", spew.Sdump(posts.Body))

	result = posts.Body

	return
}
