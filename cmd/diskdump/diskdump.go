package main

import "github.com/ElecTwix/diskdump/pkg/dump"

func main() {
	err := dump.DumpFiles()
	if err != nil {
		panic(err)
	}
}
