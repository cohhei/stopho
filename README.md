# stopho

`stopho` is a command line interface tool to generate search URLs from stock photos web sites.

<!-- TOC -->

- [stopho](#stopho)
  - [Demo](#demo)
  - [Usage](#usage)
    - [JP mode :jp:](#jp-mode-jp)
  - [Install](#install)
  - [Contribution](#contribution)
  - [Licence](#licence)
  - [Author](#author)

<!-- /TOC -->

## Demo

## Usage
To make sure [available sites list](config/sites.yml), use `list` command:

```bash
$ stopho list
```

To generate search URLs from stock photo sites, use `search` command:

```bash
$ stopho search $KEYWORD
```

You can add `-q` option to `stopho list` and `stopho search` to display URLs only. If you use `stopho` on **macOS**, open the URLs on your browser with following commands:

```bash
$ open $(stopho list -q)
$ open $(stopho search -q $KEYWORD)
```

### JP mode :jp:
`stopho` has a mode for JP. It can generate URLs from [popular sites in Japan](config/sitesjp.yml). If you are familiar with Japanese culture and internet meme, add `alias stopho='stopho --jp'` to your `.bashrc`.

## Install

To install, use `go get`:

```bash
$ go get -d github.com/kohei-kimura/stopho
```

or use homebrew:

```bash
$ brew tap kohei-kimura/homebrew-stopho
$ brew install stopho
```

or download from [releases page](https://github.com/kohei-kimura/stopho/releases).

## Contribution

1. Fork ([https://github.com/kohei-kimura/stopho/fork](https://github.com/kohei-kimura/stopho/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run `gofmt -s`
1. Create a new Pull Request

## Licence

[MIT](https://choosealicense.com/licenses/mit/)

## Author

[Kohei Kimura](https://kohei-kimura.github.io)

