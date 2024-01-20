package copyutil

// Return copy of files to path

import (
	"log"

	cp "github.com/otiai10/copy"
)

func handleErr(src, dest string, err error) error {
	if err == nil {
		return nil
	}
	log.Println("error copying file:", src, "->", dest, ":", err)
	return nil
}

func CopyDirectoryToPath(dirPath string, outDirPath string) error {
	opt := cp.Options{
		OnError: handleErr,
		Sync:    true,
	}

	// TODO: implement custom copy function
	err := cp.Copy(dirPath, outDirPath, opt)
	return err
}
