# hipcat
pipe data to hipchat

> Check `~/.hipcat` for default configs, otherwise uses passed file and/or args.

```
$ go run hipcat.go -h
NAME:
   hipcat - read file or stdin to hipchat

USAGE:
   hipcat [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --room, -r                   hipchat room [$HIPCAT_ROOM]
   --token, -t                  hipchat api token [$HIPCAT_TOKEN]
   --sender, -s "hipcat"        hipchat sender [$HIPCAT_SENDER]
   --host, -H                   hipchat host [$HIPCAT_HOST]
   --config, -c                 hipcat config file  [$HIPCAT_CONFIG]
   --help, -h                   show help
   --version, -v                print the version
```


**EXAMPLES**

```
# after setting up ~/.hipcat, see .hipcat.sample
$ echo "foo bar bah @jmervine" > example.txt
$ cat example.txt | go run hipcat.go
foo bar bah @jmervine
$ go run hipcat.go example.txt
foo bar bah @jmervine
```
