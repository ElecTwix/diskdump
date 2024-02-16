package copy_test

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/ElecTwix/diskdump/pkg/copy"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createRandomName(minSize, maxSize int) string {
	size := rand.Intn(maxSize-minSize) + minSize
	name := make([]byte, size)
	for i := range name {
		name[i] = letters[rand.Intn(len(letters))]
	}
	return string(name)
}

func CreateFileWithNumberLines(path string, n int) error {
	for i := 0; i < n; i++ {

		filePath := filepath.Join(path, createRandomName(2, 10))
		_, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating file: %w", err)
		}
	}
	return nil
}

func getDirNames(dirPath string) ([]string, error) {
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	var dirNames []string
	for _, entry := range dirEntries {
		dirNames = append(dirNames, entry.Name())
	}

	return dirNames, nil
}

func TestCopyDirectoryToPath(t *testing.T) {
	tempTestDir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(tempTestDir)
	defer os.Remove(tempTestDir)

	src, err := os.MkdirTemp(tempTestDir, "src")
	if err != nil {
		t.Fatal(err)
	}

	dest, err := os.MkdirTemp(tempTestDir, "dest")
	if err != nil {
		t.Fatal(err)
	}

	srcPath, err := filepath.Abs(src)
	if err != nil {
		t.Fatal(err)
	}

	destPath, err := filepath.Abs(dest)
	if err != nil {
		t.Fatal(err)
	}

	err = CreateFileWithNumberLines(srcPath, 10)
	if err != nil {
		t.Fatal(err)
	}

	err = copy.CopyDirectoryToPath(srcPath, destPath)
	if err != nil {
		t.Fatal(err)
	}

	srcNames, err := getDirNames(srcPath)
	if err != nil {
		t.Fatal(err)
	}

	destNames, err := getDirNames(destPath)
	if err != nil {
		t.Fatal(err)
	}

	if len(srcNames) != len(destNames) {
		t.Errorf("expected %d got %d", len(srcNames), len(destNames))
	}

	for i, srcName := range srcNames {
		if srcName != destNames[i] {
			t.Errorf("expected %s got %s", srcName, destNames[i])
		}
	}

	for _, srcName := range srcNames {
		srcFilePath := filepath.Join(srcPath, srcName)
		destFilePath := filepath.Join(destPath, srcName)
		srcFileInfo, err := os.Stat(srcFilePath)
		if err != nil {
			t.Fatal(err)
		}
		destFileInfo, err := os.Stat(destFilePath)
		if err != nil {
			t.Fatal(err)
		}
		if srcFileInfo.Size() != destFileInfo.Size() {
			t.Errorf("expected %d got %d", srcFileInfo.Size(), destFileInfo.Size())
		}
	}

	fmt.Println(tempTestDir)
}
