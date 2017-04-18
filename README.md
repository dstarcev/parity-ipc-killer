# parity-ipc-killer
A tool that helps to reproduce the bug with Parity IPC getting stuck

### getting started

- [install go](https://golang.org/doc/install#tarball)
- run `go get github.com/dstarcev/parity-ipc-killer`
- start `parity`
- start `parity-ipc-killer`. It will concurrently poll parity with `eth_blockNumber` requests. The concurrency level is 5. See in the output how many of the requests succeed.
- restart `parity-ipc-killer`. Now parity completely stops responding and the output is empty.

### usage

`parity-ipc-killer [ipc endpoint path]`