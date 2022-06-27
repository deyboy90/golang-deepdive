package style2

import (
	"fmt"
	"io"
	"math/rand"
)

type Data struct {
	Line string
}

// Puller declares behavior for pulling data
type Puller interface {
	Pull(d *Data) error
}

// Storer declares behavior for storing data
type Storer interface {
	Store(d *Data) error
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

type BackupSystem interface {
	Puller
	Storer
}

// Wraps FileSystem and BackupDrive type into a single system
type TimeMachine struct {
	Puller
	Storer
}

func pull(p Puller, data []Data) (int, error) {
	// Range over the slice of data and share each element with the FileSystem's Pull method.
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Copy knows how to pull and store data from the System.
// The BackupSystem interface is composed of both Puller and Storer interfaces
// which require defining pull and store behaviors respectively
func Copy(bs BackupSystem, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(bs, data)
		if i > 0 {
			if _, err := store(bs, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func Execute() {
	// Initiatlizing a Time Machine struct which requires passing implementators
	// Puller and Storer interface which in this case are FileSystem and BackupDrive
	//
	// timeMachine := TimeMachine{
	// 	Puller: <struct which satisfies Puller interface>,
	// 	Storer: <struct which satisfied Storer interface>,
	// }

	timeMachine := TimeMachine{
		Puller: &FileSystem{},
		Storer: &BackupDrive{},
	}

	if err := Copy(timeMachine, 3); err != io.EOF {
		fmt.Println(err)
	}
}
