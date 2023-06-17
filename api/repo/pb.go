package repo

import (
	"api/core"
	"bytes"
	"encoding/gob"
	"github.com/cockroachdb/pebble"
	"io"
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

func (p *PbTest) Get(k string) (error, *core.Data) {
	val, closer, err := p.db.Get([]byte(k))
	defer func(closer io.Closer) {
		_ = closer.Close()
	}(closer)

	if err != nil {
		return err, nil
	}

	var d core.Data
	dec := gob.NewDecoder(bytes.NewBuffer(val))
	err = dec.Decode(&d)
	if err != nil {
		return err, nil
	}

	return nil, &d
}
