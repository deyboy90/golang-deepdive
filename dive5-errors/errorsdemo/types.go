package errorsdemo

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type BulkWriteError struct {
	lineNum int
}

func (bw *BulkWriteError) Error() string {
	return "failed to perform bulk write, item: " + strconv.Itoa(bw.lineNum) + " failed"
}

func bulkWrite(entries []string) error {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	switch r.Intn(4) {
	case 2:
		return &BulkWriteError{lineNum: r.Intn(10)}
	default:
		return nil
	}
}

func CustomErrorsDemo() {
	if err := bulkWrite([]string{"meh"}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}
