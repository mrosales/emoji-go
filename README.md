# emoji-go
🔍 Go library to search for Unicode emoji information.

[![goreport][goreport-badge]][goreport]
[![GoDoc][godoc-badge]][godoc]
[![MIT License][license-badge]][license]

Provides go structures to list or search the latest emoji symbols.

* Emoji data is sourced from [iamcal/emoji-data](emoji-data) (MIT)
  and pulled from the [JSDelivr CDN](emoji-jsdelivr).
* Fuzzy searching is provided by [lithammer/fuzzysearch](lithammer-fuzzysearch) (MIT)
* The current version supports [Emoji version 13 (2020)](unicode-emoji-13)

## Library Usage

Exposes a constant array of all emoji symbols and an interface for
performing a fuzzy search on the dataset.

See the [godoc](godoc) for more information.

## CLI Usage

Install the CLI with `go get` outside a go module.

```shell
go get github.com/mrosales/emoji-go/cmd/emoji
```

You can look up an emoji symbol directly
```shell
emoji --limit 1 rocket
🚀
```

Or view all the emoji information
```shell
emoji -f json rocket
[
    {
        "name": "rocket",
        "category": "Travel \u0026 Places",
        "alternate_names": [
            "rocket"
        ],
        "unified": "1f680",
        "character": "🚀",
        "sheet_x": 33,
        "sheet_y": 50,
        "added_in": "0.6",
        "platform_support": {
            "Apple": true,
            "Facebook": true,
            "Google": true,
            "Twitter": true
        }
    }
]

# skin tones are also supported
emoji --skin medium_dark wave
👋🏾
```

## Maintenance

### Updating the dataset

The dataset generation script populates `data.go` in the repo root and can be
run by executing `go generate` in the repo root.

The tag revision in `cmd/emojigen/internal/importer/cdn.go`
must be updated to support new emoji versions.


[emoji-data]: https://github.com/iamcal/emoji-data
[emoji-jsdelivr]: https://www.jsdelivr.com/package/npm/emoji-datasource-apple
[lithammer-fuzzysearch]: https://github.com/lithammer/fuzzysearch
[unicode-emoji-13]: https://unicode.org/emoji/charts-13.0/emoji-released.html
[goreport]: https://goreportcard.com/report/github.com/mrosales/emoji-go
[goreport-badge]: https://goreportcard.com/badge/github.com/mrosales/emoji-go?style=flat-square
[godoc]: https://pkg.go.dev/github.com/mrosales/emoji-go
[godoc-badge]: https://img.shields.io/badge/godoc-reference-blue?style=flat-square
[license]: LICENSE.txt
[license-badge]: https://img.shields.io/github/license/mrosales/emoji-go?style=flat-square
