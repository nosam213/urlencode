# urlencode
Simple utility for encoding URLs to url-encoded URLs.

### Flags
None currently implemented.


## Usage

### Windows
`urlencode.exe "https://example.com/page?parameter=1&parameter=2"` 

### Unix
`urlencode "https://example.com/page?parameter=1&parameter=2"`


## Compiling
NOTE: Requires network connection for the first compile.

### Windows
```
C:\Users\username\urlencode> dir
go.mod  go.sum  main.go  README.md

C:\Users\username\urlencode> go build -o urlencode.exe -ldflags="-s -w"
go.mod  go.sum  urlencode.exe  main.go  README.md
```

### Unix
```
/home/user/urlencode:~$ ls
go.mod  go.sum  main.go  README.md

/home/user/urlencode:~$ go build -o urldecode -ldflags="-s -w"
go.mod  go.sum  urlencode  main.go  README.md
```
