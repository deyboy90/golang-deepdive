package style3

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

// Copy fn explictly asks for a puller and a storer interface vs asking for PullerStorer
// this makes the api more precise and focused on what exactly it needs to perform it's task
func Copy(p Puller, s Storer, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(p, data)
		if i > 0 {
			if _, err := store(s, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func Style3Demo() {
	fmt.Println("-------Style3Demo--------")
	/**
	Since the copy function explicitly asks  or a puller and storer interface so here
	we construct them independently and pass it in.

	We are also getting rid of PullerStorer interface along with the idea of TimeMachine
	struct as that's not needed anymore.
	 **/

	fs := &FileSystem{}
	bd := &BackupDrive{}

	if err := Copy(fs, bd, 3); err != io.EOF {
		fmt.Println(err)
	}
	fmt.Println("---------------------------------")
}
