package post

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

func TestGettingPages(t *testing.T) {

	pages, err := PaginateCreator(PaginateCreatorParams{
		CreatorId: "henyathegenius",
	})

	if err != nil {
		t.Error(errors.New(err).ErrorStack())
	}

	response, err := json.Marshal(pages)

	if err != nil {
		t.Error(errors.New(err).ErrorStack())
	}

	t.Logf("Result: %s", string(response))
}

func TestGettingPosts(t *testing.T) {

	posts, err := ListCreator(ListCreatorParams{
		CreatorId: "henyathegenius",
		FirstId:   "10861097",
	})

	if err != nil {
		t.Errorf("%v", string(errors.New(err).ErrorStack()))
	}

	response, err := json.Marshal(posts)

	if err != nil {
		t.Errorf("%v", errors.New(err).ErrorStack())
	}

	t.Logf("Result: %s", string(response))
}

func TestPostMustNotBeEmpty(t *testing.T) {

	posts, _ := ListCreator(ListCreatorParams{
		CreatorId: "henyathegenius",
		FirstId:   "10861097",
	})

	if len(posts.Body) == 0 {
		t.Error("No post came out")
	}

}
