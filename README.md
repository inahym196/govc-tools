# govc-tools
- govmomiライブラリを利用してvCenterAPIを叩くツール群

## Motivation
- golangの練習
  - testの練習
  - govmomiの練習

## How to build
`go build main.go`

## Global setting
```sh
❯ cat .cred.env
GOVCTOOLS_USER=user
GOVCTOOLS_PASS=pass
GOVCTOOLS_HOST=localhost
GOVCTOOLS_PORT=8989
GOVCTOOLS_DISABLE_TLS=false
```
## Usage
<ins>main</ins>
```
❯ go run main.go --help
NAME:
   govc-tools - vcenter cli tools

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   portgroup   portgroupをhogehogeする
   permission  permissionをhogehogeする
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```
<ins>portgroup</ins>
```
❯ go run main.go portgroup --help
NAME:
   govc-tools portgroup - portgroupをhogehogeする

USAGE:
   govc-tools portgroup command [command options] [arguments...]

COMMANDS:
   list  show portgroup list

OPTIONS:
   --help, -h  show help
```

## Example
<ins>portgroup</ins>
```
❯ go build main.go
❯ ./main portgroup list
dvportgroup-11
dvportgroup-13
```
or
```
❯ go run main.go portgroup list
dvportgroup-11
dvportgroup-13
```