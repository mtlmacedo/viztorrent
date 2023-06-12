package main

import (
	"fmt"
	"os"

	"github.com/mtlmacedo/viztorrent/pkg/constants"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	fmt.Println(inPath, outPath, constants.PeerSize)
}
