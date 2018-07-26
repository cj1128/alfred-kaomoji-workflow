# Alfred Kaomoji Workflow

[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](http://mit-license.org/2018)

Alfred workflow to quickly search kaomojis \\\(۶•̀ᴗ•́)۶//／／.

<p align="center">
  <img src="http://ww1.sinaimg.cn/large/9b85365dgy1ftn5wj39m7g20mg0fcx6q" />
</p>

## Installation o(≧∇≦o)

click to download latest version of [Kaomoji.alfredworkflow](https://github.com/fate-lovely/alfred-kaomoji-workflow/releases/download/v1.0.0/Kaomoji.alfredworkflow).

## Usage (´◑ω◐`)

press `Enter` to copy kaomoji.

## Development

kaomoji database `lib.json` is fetched from [asciilib].

1. generate `kaomojis.go`

    ```bash
    $ go run scripts/generate_data.go
    ```

2. bundle assets to alfred workflow

    ```bash
    $ make bundle
    ```

then, we get `Kaomoji.alfredworkflow` in `tmp` dir.

[asciilib]: https://github.com/iansinnott/asciilib
