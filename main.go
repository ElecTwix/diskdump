package main

import (
	"os"
	"path"

	"github.com/ElecTwix/diskdump/pkg/disk"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	partitions, err := disk.GetAllDisks()
	if err != nil {
		panic(err)
	}

	for _, partition := range partitions {
		path := path.Join(currentDir, partition.Name)
		err := os.Mkdir(path, os.ModeDir)
		if err != nil {
			panic(err)
		}
		err = partition.CopyAllToPath(path)
		if err != nil {
			panic(err)
		}

	}
}
