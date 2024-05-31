# 🏳 FLAG

FLAG is a command-line wrapper for a small portion of the Shortcut REST API.
Its primary focus is searching for stories using Shortcutʼs [search operators].

Search results are returned as raw JSON data. Combined with other tools, you
might find some interesting use cases for such functionality.

## Installation

Grab the latest release, or any release of your choosing really, from the
[Releases page](https://github.com/ngscheurich/flag/releases).

## Usage

Generate a [Shortcut API token], then log in:

```sh
flag login $SHORTCUT_API_TOKEN
```

Check `flag help` for further usage info.

## Examples

FLAG is most useful when combined with other CLI programs. Following are some
examples of things you might want to do.

Read the description of a story by processing FLAGʼs output with [`jq`] and
piping the results to [Glow]:

```sh
flag search -q 'id:206082' | jq '.data.[].description' -r | glow
```

Get a list of stories you own and format the output using [`jq`]:

```sh
flag search | jq '.data[] | "[sc-\(.id)] \(.name)"'
```

1. Get a list of your in-development stories.
2. Format the list items using `jq`.
3. Select a story using using [Gum] ([fzf] would also work nicely).
4. Start tracking work on that story with [Timewarrior].

```sh
flag search -q 'owner:nscheurich state:"In Development"' |
    jq '.data[] | "[sc-\(.id)] \(.name)"' -r |
    gum choose |
    xargs -I % timew start %
```

[`jq`]: https://jqlang.github.io/jq/
[fzf]: https://github.com/junegunn/fzf
[glow]: https://github.com/charmbracelet/glow
[go]: https://go.dev/doc/install
[gum]: https://github.com/charmbracelet/gum
[search operators]: https://help.shortcut.com/hc/en-us/articles/360000046646-Searching-in-Shortcut-Using-Search-Operators
[shortcut api token]: https://help.shortcut.com/hc/en-us/articles/205701199-Shortcut-API-Tokens
[timewarrior]: https://timewarrior.net/
