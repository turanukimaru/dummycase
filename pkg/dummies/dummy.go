package dummies

import (
	"fmt"
	"github.com/turanukimaru/gormstart/pkg/dummydb"
)

type Dummy struct {
}

func (d *Dummy) Hello() (err error) {
	fmt.Println("Hello Dummies !")
	err = dummydb.DbAccess()
	return
}
