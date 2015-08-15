# ntsh
The Nice Try shell, meant for honeypots

## Installation
```bash
go get github.com/kd5pbo/ntsh
mkdir /tmp/ntsh
```

## Usage
A telnet server using ncat:
```bash
nohup ncat -ltkp 23 -e $GOPATH/bin/ntsh
```

It should also be usable as a shell in `/etc/passwd`:
```
honey:*:1002:1002::/home/honey:/path/to/ntsh
```
You'll probably have to put it somewhere other than `$GOPATH/bin`, though.

## Adding commands
Each command will need it's own source file in the top-level directory,
probably with the same name as the command.  It should have an `init()`
function which calls `Register()` (from `command.go`).  It's perfectly fine
to put source files for a command in a subdirectory, like with `id`.

Have a look at `id.go`, the `id/` directory, and `rm.go` for examples.
Generally speaking, commands should try to print out "Nice Try" in some
fashion.
