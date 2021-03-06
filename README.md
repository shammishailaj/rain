rain
====

BitTorrent client and library in Go.

[![Build Status](https://travis-ci.org/cenkalti/rain.svg?branch=master)](https://travis-ci.org/cenkalti/rain)
[![Coverage Status](https://coveralls.io/repos/github/cenkalti/rain/badge.svg?branch=master)](https://coveralls.io/github/cenkalti/rain?branch=master)
[![GoDoc](https://godoc.org/github.com/cenkalti/rain?status.svg)](https://godoc.org/github.com/cenkalti/rain)
[![GitHub Release](https://img.shields.io/github/release/cenkalti/rain.svg)](https://github.com/cenkalti/rain/releases)

Features
--------
- [x] [Core protocol](http://bittorrent.org/beps/bep_0003.html)
- [x] [Fast extension](http://bittorrent.org/beps/bep_0006.html)
- [x] [Magnet links](http://bittorrent.org/beps/bep_0009.html)
- [x] [Multiple trackers](http://bittorrent.org/beps/bep_0012.html)
- [x] [UDP trackers](http://bittorrent.org/beps/bep_0015.html)
- [x] [DHT protocol](http://bittorrent.org/beps/bep_0005.html)
- [x] [PEX](http://bittorrent.org/beps/bep_0011.html)
- [x] [Message stream encryption](http://wiki.vuze.com/w/Message_Stream_Encryption)
- [x] Fast resuming
- [x] IP blocklist
- [x] RPC server
- [x] Console UI
- [ ] [WebSeed](http://bittorrent.org/beps/bep_0019.html)

Installing
----------

Get the latest binary from [releases page](https://github.com/cenkalti/rain/releases) or install development version:

`go get github.com/cenkalti/rain`

Usage
-----

- `rain server` command runs a RPC server.
- `rain client add <magnet_or_torrent>` adds a torrent and print it's ID.
- `rain client stats <ID>` prints the stats of the torrent.

Run `rain help` to see other commands.
