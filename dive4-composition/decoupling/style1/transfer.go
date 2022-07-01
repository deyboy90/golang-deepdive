package style1

import (
	"fmt"
	"io"
	"math/rand"
)

type Data struct {
	Line string
}

type FileSystem struct {
	Files []string
}

type BackupDrive struct {
	Backup []string
}

func (fs *FileSystem) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 5:
		return io.EOF
	default:
		d.Line = "Data"
		fmt.Println("Pulling: ", d.Line)
		return nil
	}
}

func (bd *BackupDrive) Store(d *Data) error {
	fmt.Println("Storing: ", d.Line)
	return nil
}

func pull(fs *FileSystem, data []Data) (int, error) {
	// Range over the slice of data and share each element with the FileSystem's Pull method.
	for i := range data {
		if err := fs.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

func store(bd *BackupDrive, data []Data) (int, error) {
	for i := range data {
		if err := bd.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Wraps FileSystem and BackupDirve type into a single system
type TimeMachine struct {
	FileSystem
	BackupDrive
}

// Copy knows how to pull and store data from the System.
// Now we can call the pull and store functions, passing Xenia and Pillar through.
func Copy(tm *TimeMachine, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(&tm.FileSystem, data)
		if i > 0 {
			if _, err := store(&tm.BackupDrive, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func Style1Demo() {
	fmt.Println("-------Style1Demo--------")
	timeMachine := TimeMachine{
		FileSystem{},
		BackupDrive{},
	}

	if err := Copy(&timeMachine, 3); err != io.EOF {
		fmt.Println(err)
	}
	fmt.Println("---------------------------------")
}
