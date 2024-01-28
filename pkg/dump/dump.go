package dump

import (
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
		path := path.Join(currentDir, partition.Name)
		err := os.Mkdir(path, os.ModeDir)
		if err != nil {
			return err
		}
		err = partition.CopyAllToPath(path)
		if err != nil {
			return err
		}
	}

	return nil
}
