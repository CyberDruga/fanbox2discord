# Fanbox2Discord

This application keeps track of fanbox accounts, and send a discord webhook message whenever there's a new post.

The message is customizable, and it uses Golang's templ language syntax.

There's also some command arguments to help set the bot up.


## Command Arguments

### `--print-example`

After having your `config.toml` well configured, you can call the application with this argument to get the last post
of each account configured.

It'll then post the content of the `fanbox.Post` struct as the output.

This is an example setting one of the accounts to `sushidog`:

```
(fanbox.Post) {
 PostId: (string) (len=8) "10871254",
 Title: (fanbox.JsonString) (len=16) "お知らせ/FYI",
 FeeRequired: (int) 0,
 PublishedDatetime: (string) (len=25) "2025-11-08T14:57:31+09:00",
 UpdatedDatetime: (string) (len=25) "2025-11-09T20:06:06+09:00",
 Tags: ([]string) {
 },
 LikeCount: (int) 5,
 IsCommentingRestricted: (bool) false,
 CommentCount: (int) 6,
 IsRestricted: (bool) false,
 User: (fanbox.User) {
  UserId: (string) (len=8) "49386651",
  Name: (string) (len=8) "sushidog",
  IconUrl: (string) (len=112) "https://pixiv.pximg.net/c/160x160_90_a2_g5/fanbox/public/images/user/49386651/icon/LAqzIAAMFXo8OPEH3rzVC3RF.jpeg"
 },
 CreatorId: (string) (len=8) "sushidog",
 Cover: (fanbox.Cover) {
  Type: (string) (len=10) "post_image",
  Url: (string) (len=85) "https://downloads.fanbox.cc/images/post/10871254/w/1200/8IKn0JQ8ZXEJOmgrMkrUDXA1.jpeg"
 },
 Excerpt: (fanbox.JsonString) (len=411) "ディスコードで告知する間にこちらでご報告！\\nご要望が多かった タロットカードを再販しました！ 🔮✨（数量限定）\\n個展のグッズやアートブックなども ご購入できるようにしました！\\n欲しかった方はこの機会にどうぞ💖\\n📣 Announcement!\\nMy  Tarot Cards are back in stock! 🔮✨(super limited amount)\\nIf you wanted on...",
 IsPinned: (bool) false
}
```

With that you can now setup your message template like this 

```toml
new-message-template = """
{
  "embeds": [
    {
      "author": {
        "name": "{{ .User.Name }}",
        "icon_url": "{{ .User.IconUrl }}",
        "url": "https://{{ .CreatorId }}.fanbox.cc"
      },
      "title": "{{ .Title }}",
      "color": 3542783,
      "url": "https://{{ .CreatorId }}.fanbox.cc/posts/{{ .PostId }}",
      "description": "{{ .Excerpt }}",
      "timestamp": "{{ .PublishedDatetime }}",
      "image": {
        "url": "{{ .Cover.Url }}"
      },
      "footer": {
        "text": "Pledge: {{ .FeeRequired }} JYP"
      }
    }
  ],
  "username": "Sushidog's Fanbox Bot"
}
"""
```

More on templates later.

## `--populate`

This makes the first run of the application to just populate the database, and skip posting new messages to Discord.

If you enabled the Repeat mode on the `config.toml` this will do the same only on the first run. On next runs it'll
send messages normally.

If you're setting this up on a server it's recommended to just run this once by hand, and setting the application 
without this argument later, to avoid problems.

## `--config [config-file]`

This argument sets the path of the config file to use. By default the application will look for a file called
`config.toml` on it's working directory, but with this argument it'll use the informed config file instead. 

Ex.: 

```sh
fanbox2discord --config ~/configs/myconfig.toml
```
### `--json`

This makes all the logs show in json format, in case you want to have that.

### `--debug`

This argument is just there to make developing things easier. The code is litered with `slog.Debug()` calls to see 
what's going on.


### `--show-source`

This makes all the logs show the source of where the logs are being called from. Useful for debuggin.

## Config File Options 

You can find a better documentation looking at the file `config_example.toml`, where it explains each field.


## Message Templating

In this project I really wanted to attempt some form of message customization, because on previous projects, every time
I wanted to add something I had to change the code, or add a flag for it, which was annoying. So this time I leveraged
the power of Golang's templating system.

For that, you can use the argument `--print-example` to get the last post from the account you've configured

It'll output something like this:

```
(fanbox.Post) {
 PostId: (string) (len=8) "10871254",
 Title: (fanbox.JsonString) (len=16) "お知らせ/FYI",
 FeeRequired: (int) 0,
 PublishedDatetime: (string) (len=25) "2025-11-08T14:57:31+09:00",
 UpdatedDatetime: (string) (len=25) "2025-11-09T20:06:06+09:00",
 Tags: ([]string) {
 },
 LikeCount: (int) 5,
 IsCommentingRestricted: (bool) false,
 CommentCount: (int) 6,
 IsRestricted: (bool) false,
 User: (fanbox.User) {
  UserId: (string) (len=8) "49386651",
  Name: (string) (len=8) "sushidog",
  IconUrl: (string) (len=112) "https://pixiv.pximg.net/c/160x160_90_a2_g5/fanbox/public/images/user/49386651/icon/LAqzIAAMFXo8OPEH3rzVC3RF.jpeg"
 },
 CreatorId: (string) (len=8) "sushidog",
 Cover: (fanbox.Cover) {
  Type: (string) (len=10) "post_image",
  Url: (string) (len=85) "https://downloads.fanbox.cc/images/post/10871254/w/1200/8IKn0JQ8ZXEJOmgrMkrUDXA1.jpeg"
 },
 Excerpt: (fanbox.JsonString) (len=411) "ディスコードで告知する間にこちらでご報告！\\nご要望が多かった タロットカードを再販しました！ 🔮✨（数量限定）\\n個展のグッズやアートブックなども ご購入できるようにしました！\\n欲しかった方はこの機会にどうぞ💖\\n📣 Announcement!\\nMy  Tarot Cards are back in stock! 🔮✨(super limited amount)\\nIf you wanted on...",
 IsPinned: (bool) false
}
```

Each of these fields can be used in our template, which is just a json of a webhook message.

For example, this is one message

```json
{
  "embeds": [
    {
      "author": {
        "name": "{{ .User.Name }}",
        "icon_url": "{{ .User.IconUrl }}",
        "url": "https://{{ .CreatorId }}.fanbox.cc"
      },
      "title": "{{ .Title }}",
      "color": 3542783,
      "url": "https://{{ .CreatorId }}.fanbox.cc/posts/{{ .PostId }}",
      "description": "{{ .Excerpt }}",
      "timestamp": "{{ .PublishedDatetime }}",
      "image": {
        "url": "{{ .Cover.Url }}"
      },
      "footer": {
        "text": "Pledge: {{ .FeeRequired }} JYP"
      }
    }
  ],
  "username": "Sushidog's Fanbox Bot"
}
```


The Syntax is very simple: everything inside a `{{ }}` is a value from the `fanbox.Post` object.

On the output of `--print-example` we got this: 

```
 User: (fanbox.User) {
  UserId: (string) (len=8) "49386651",
  Name: (string) (len=8) "sushidog",
  IconUrl: (string) (len=112) "https://pixiv.pximg.net/c/160x160_90_a2_g5/fanbox/public/images/user/49386651/icon/LAqzIAAMFXo8OPEH3rzVC3RF.jpeg"
 },
```

Notice how we have the field `User`, and under it we have the fields `Name` and `IconUrl` that we want, so on our Json
we write this: 

```json
...
    {
      "author": {
        "name": "{{ .User.Name }}",
        "icon_url": "{{ .User.IconUrl }}",
...
```

More over, we want to link the post in the embed. A link to a post in fanbox is displayed as this:
`https://sushidog.fanbox.cc/posts/10871254`.

Back at our `--print-example` output, we have the following fields: 

```
(fanbox.Post) {
 PostId: (string) (len=8) "10871254",
...
 CreatorId: (string) (len=8) "sushidog",
...
}
```

So for our embed url we can do something like this: 

```json
{
  "embeds": [
    {
    ...
      "url": "https://{{ .CreatorId }}.fanbox.cc/posts/{{ .PostId }}",
    ...
    }
  ],
  "username": "Sushidog's Fanbox Bot"
}
```


to make things easier for making the template, you can use tools such as [discohook](https://discohook.app), and
after you get something you're happy with, you can get the Json from it on `Options` -> `Json Editor`, and start 
adding the required options.
