package copy

// Return copy of files to path

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type CopyManager struct {
	srcDir  string
	destDir string
	maxSize int64
}

func copyDirWithWalker(srcDir, destDir string, maxSize int64) error {
	manager := new(CopyManager)
	manager.srcDir = srcDir
	manager.destDir = destDir
	manager.maxSize = maxSize
	return filepath.WalkDir(srcDir, manager.walkerHandler)
}

func (c *CopyManager) walkerHandler(path string, d fs.DirEntry, err error) error {
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("Permission error:", err)
			return filepath.SkipDir
		}
		fmt.Println("Error:", err)
		return err
	}

	file, err := d.Info()
	if err != nil {
		return err
	}

	relPath, _ := filepath.Rel(c.srcDir, path)
	destPath := filepath.Join(c.destDir, relPath)

	if d.IsDir() {
		if strings.HasPrefix(file.Name(), "$") {
			return filepath.SkipDir
		}
		err := os.MkdirAll(destPath, os.ModePerm)
		fmt.Println("Copying", path)
		if err != nil {
			return err
		}
		return nil
	}

	info, err := d.Info()
	if err != nil {
		return err
	}

	if info.Size() > c.maxSize {
		log.Println("File is too big:", path)
		return nil
	}
	go copyFile(path, destPath)
	return nil
}

func copyFile(src, dest string) {
	_, err := os.Stat(src)
	if err != nil {
		fmt.Println("Error stating file", src)
		return
	}

	data, err := os.ReadFile(src)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Permission error", err)
			return
		}
		log.Println("Error opening file", src)
		return
	}

	destFile, err := os.Create(dest)
	if err != nil {
		fmt.Println("Error creating file", dest)
		return
	}

	defer destFile.Close()

	_, err = destFile.Write(data)
	if err != nil {
		log.Println("Error writing to file", dest)
		return
	}

	err = destFile.Sync()
	if err != nil {
		log.Println("Error syncing file", dest)
		return
	}
}

func CopyDirectoryToPath(dirPath string, outDirPath string) error {
	/*
		opt := cp.Options{
			OnError:      handleErr,
			NumOfWorkers: 8,
		}
	*/

	log.Println("started copying")

	hunderedMB := int64(100 * 1024 * 1024)

	err := copyDirWithWalker(dirPath, outDirPath, hunderedMB)
	// err := cp.Copy(dirPath, outDirPath, opt)
	return err
}
