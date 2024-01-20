package disk

import (
	"github.com/ElecTwix/diskdump/pkg/copyutil"
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
		disks[index] = Disk{
			Name:       partition.Device,
			FsType:     partition.Fstype,
			Mountpoint: partition.Mountpoint,
		}
	}

	return []Disk{}, nil
}

func (d *Disk) CopyAllToPath(path string) error {
	return copyutil.CopyDirectoryToPath(d.Mountpoint, path)
}
