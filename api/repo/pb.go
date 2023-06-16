package repo

import (
	"api/core"
	"bytes"
	"encoding/gob"
	"github.com/cockroachdb/pebble"
)

type PbTest struct {
	db *pebble.DB
}

func NewPbTest(pb *pebble.DB) *PbTest {
	return &PbTest{db: pb}
}

func (p *PbTest) Create(d *core.Data) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(d)
	if err != nil {
		return err
	}
	return p.db.Set([]byte(d.C1), buf.Bytes(), pebble.Sync)
}
