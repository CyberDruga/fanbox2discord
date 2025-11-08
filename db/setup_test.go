package db

import (
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

func TestSetup(t *testing.T) {

	err := NewClient(":memory:")

	if err != nil {
		t.Error(
			"Something wasn't right",
			"error", err.(*errors.Error).ErrorStack(),
		)
	}
}
