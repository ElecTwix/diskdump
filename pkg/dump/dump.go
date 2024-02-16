package dump

import (
	"log"
	"os"
	"path"

	"github.com/ElecTwix/diskdump/pkg/disk"
)

func DumpFiles() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	partitions, err := disk.GetAllDisks()
	if err != nil {
		return err
	}

	for _, partition := range partitions {
		log.Printf("Dumping %s\n", partition.Name)
		path := path.Join(currentDir, partition.Name)
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
		err = partition.CopyAllToPath(path)
		if err != nil {
			return err
		}
	}

	log.Println("Dumping complete!")

	return nil
}
