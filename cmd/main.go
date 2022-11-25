package main

import (
	"flag"
)

var conf = flag.String("config", "./conf/app.toml", "配置文件路径")

func main() {
	flag.Parse()
}
