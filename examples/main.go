package main

import (
	"github.com/SvenWiltink/volumecontrol"
	"flag"
)
func main() {
	var volume int
	flag.IntVar(&volume,"v", 0, "Volume")
	flag.Parse()
	volumecontrol.ChangeVolume(volume)
}
