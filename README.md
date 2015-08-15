# ntsh
The Nice Try shell, meant for honeypots

## Installation
Installation is easy:
```bash
go get github.com/kd5pbo/ntsh
mkdir /tmp/ntsh (if using defaults)
```

## Usage
The defaults are expected to be safe for most situations, but a `-h` gives
the command-line options:

```
-a=false: Prompt for authorization
-c="": Client's address
-l="/tmp/ntsh/ntsh.log": Logfile
-motd="/tmp/ntsh/motd": MOTD to print on connect
-p="[root@localhost:~]# ": Prompt, may be changed later
-pw="": If set, only allow in this password; implies -a
-u="": If set, only allow in this username; implies -a
```

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
