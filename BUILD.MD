# Building the Go Code

```bash
go env
```
This will give the complete evnironment of the go.

> NOTE - We will need only GOOS for this purpose

1. Build for Windows
```bash 
GOOS="windows" go build
```

2. Build for Linux
```bash
GOOS="linux" go build
```

3. Build for MacOS
```bash
GOOS="darwin" go build
```