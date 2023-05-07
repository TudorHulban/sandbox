# Install

```bash
go get -u github.com/securego/gosec/v2/cmd/gosec
```

Copy from $GOPATH/bin to /usr/local/go/bin.

```bash
gosec -no-fail ./... 2>&1 | tee ~/ram/sast-gosec.log
```
