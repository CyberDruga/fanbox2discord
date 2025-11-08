package generic

import (
	"encoding/json"
	"log/slog"
	"os"
	"testing"

	"github.com/go-errors/errors"
)

func init() {

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)

	slog.SetDefault(slog.New(handler))

}

func TestGetLatestPostId(t *testing.T) {

	creator := "henyathegenius"

	postId, err := GetLatestPostId(creator)

	if err != nil {
		t.Errorf("%v", err.(*errors.Error).ErrorStack())
	}

	if err != nil {
		t.Error(errors.New("Couldn't turn result into json for output"))
	}

	t.Logf("Result: %v", postId)

}

func TestGetLatestPosts(t *testing.T) {

	creatorId := "henyathegenius"

	posts, err := GetLatestPosts(creatorId)

	if err != nil {
		t.Error(errors.New("Couldn't get posts"))
	}

	if len(posts) == 0 {
		t.Error(errors.New("Somehow [posts] is empty"))
	}

	result, err := json.Marshal(posts)

	if err != nil {
		t.Error(
			"Couldn't turn result into json for output",
			"error", err.Error(),
			"value", string(result),
		)
	}

	t.Logf("result: %v", string(result))
}
