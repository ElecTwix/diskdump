package disk

import (
	"strings"

	"github.com/ElecTwix/diskdump/pkg/copy"
	rawdisk "github.com/shirou/gopsutil/disk"
)

type Disk struct {
	Name       string
	FsType     string
	Mountpoint string
}

func GetAllDisks() ([]Disk, error) {
	partitions, err := rawdisk.Partitions(false)
	if err != nil {
		return nil, err
	}

	disks := make([]Disk, len(partitions))

	for index, partition := range partitions {
		partition.Device = strings.ReplaceAll(partition.Device, "/", "")
		partition.Device = strings.ReplaceAll(partition.Device, ":", "_")
		disks[index] = Disk{
			Name:       partition.Device,
			FsType:     partition.Fstype,
			Mountpoint: partition.Mountpoint,
		}
	}

	return disks, nil
}

func (d *Disk) CopyAllToPath(path string) error {
	return copy.CopyDirectoryToPath(d.Mountpoint, path)
}
