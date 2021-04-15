package utils

import (
	"encoding/gob"
	"os"
)

type Gob struct {
	File *os.File
}

func NewGob() *Gob {
	return &Gob{}
}

func (this *Gob) Create(name string) *Gob {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	this.File = file
	return this
}

func (this *Gob) Encode(e interface{}) {
	enc := gob.NewEncoder(this.File)
	err := enc.Encode(e)
	if err != nil {
		panic(err)
	}
}
