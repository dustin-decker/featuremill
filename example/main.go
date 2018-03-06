package main

import (
	"github.com/dustin-decker/featuremill"
)

func main() {
	featuremill.ExtractIP("src_ip", "127.0.0.1")
}
