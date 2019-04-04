# vin

CLI for Vim plugins

[![Lang Badge][lang]][Vin]
[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE]

## Overview

Vin is a CLI written in Go to install and update Vim plugins. It plays
nicely with [Janus][].

## Usage

Here's an example usage.

```bash
$ vin install https://github.com/gregsexton/MatchTag
$ vin rm MatchTag
$ vin install https://github.com/matchtags/MatchTags
$ vin ls
The following Janus included plug-ins are disabled:
  syntastic
The following additional plug-ins are installed:
  MatchTags
  ale
  elm-vim
  fzf
  fzf.vim
  tabular
  vim-airline
  vim-closetag
  vim-go
  vim-json
  vim-prettier
  vim-vue
$ vin up
Updating Janus (running rake)
MatchTags
ale
elm-vim
fzf
fzf.vim
tabular
vim-airline
vim-closetag
vim-go
vim-json
vim-prettier
vim-vue

```

## Installation

TODO: Need to update.

## Contributing

To contribute, please fork the repository, create a feature branch, and then
submit a [pull request][].


## License

[Vin][] is released under the MIT license. Please see the [LICENSE][]
file for more information.

[godoc badge]: https://godoc.org/github.com/vimstuff/vin?status.svg
[godoc link]: https://godoc.org/github.com/vimstuff/vin
[Janus]: https://github.com/carlhuda/janus
[lang]: https://img.shields.io/github/languages/top/vimstuff/vin.svg
[LICENSE]: https://github.com/vimstuff/vin/blob/master/LICENSE
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/vimstuff/vin
[report card]: https://goreportcard.com/report/github.com/vimstuff/vin
[Vin]: https://github.com/vimstuff/vin
