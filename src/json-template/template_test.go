package json_template

import (
	"encoding/json"
	"log/slog"
	"os"
	"testing"

	"github.com/CyberDruga/fanbox2discord/src/models/fanbox"
)

func init() {

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)

	slog.SetDefault(slog.New(handler))

}

func TestJsonTemplate(t *testing.T) {

	var post fanbox.Post

	json.Unmarshal([]byte(` 
		{
      "id": "10791674",
      "title": "🙇💦【全体公開 / public post】",
      "feeRequired": 0,
      "publishedDatetime": "2025-10-25T21:18:40+09:00",
      "updatedDatetime": "2025-10-25T21:18:40+09:00",
      "tags": [],
      "isLiked": false,
      "likeCount": 76,
      "isCommentingRestricted": false,
      "commentCount": 29,
      "isRestricted": false,
      "user": {
        "userId": "111014427",
        "name": "HenyaTheGenius",
        "iconUrl": "https://pixiv.pximg.net/c/160x160_90_a2_g5/fanbox/public/images/user/111014427/icon/DNXfKOedpycPsepx9EmBazm5.jpeg"
      },
      "creatorId": "henyathegenius",
      "hasAdultContent": false,
      "cover": {
        "type": "cover_image",
        "url": "https://pixiv.pximg.net/c/1200x630_90_a2_g5/fanbox/public/images/post/10791674/cover/h0G7edhn5pWOBCWf7Hc2GIu2.jpeg"
      },
      "excerpt": "I am so sorry i havent updated FANBOX for 2 weeks!\nI have been busy and stuff, but next week I will upload new wallpaper! I hope you guys will enjoy dayo! Thank you so much dayo🙇\nand im sorry for the inconvenience dayo🙇\n2週間ほどFANBOX更新できてなくて申し訳ない...！\nいろいろ忙しかったりで手が回...",
      "isPinned": false
    }`), &post)

	jsonStr := `
	{
	"content": "{{ .excerpt }}"
	}
	`

	message, err := ApplyTemplate(jsonStr, post)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if result, err := json.Marshal(message); err == nil {
		t.Logf("result: %v", string(result))
	}

}
