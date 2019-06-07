package readers

import (
	"crypto/md5"
	cryptoRand "crypto/rand"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"math/rand"
	"os"
	"sync/atomic"
)

var (
	machineId = readMachineID()
	pid       = os.Getpid()
)

var ErrIncorrectBufferSize = errors.New("buffer size is not 10")

func readMachineID() []byte {
	id := make([]byte, 4)
	hid, err := readPlatformMachineID()
	if err != nil || len(hid) == 0 {
		hid, err = os.Hostname()
	}
	if err == nil && len(hid) != 0 {
		hw := md5.New()
		hw.Write([]byte(hid))
		copy(id, hw.Sum(nil))
	} else {
		// Fallback to rand number if machine id can't be gathered
		if _, randErr := cryptoRand.Reader.Read(id); randErr != nil {
			panic(fmt.Errorf("xid: cannot get hostname nor generate a random number: %v; %v", err, randErr))
		}
	}
	return id
}

type xidReader struct {
	inc uint32
}

func NewXidReader() io.Reader {
	return &xidReader{inc: rand.Uint32()}
}

func (r *xidReader) Read(p []byte) (n int, err error) {
	if len(p) != 10 {
		return 0, ErrIncorrectBufferSize
	}
	p[0] = machineId[0]
	p[1] = machineId[1]
	p[2] = machineId[2]
	p[3] = machineId[3]
	p[4] = byte(pid >> 8)
	p[5] = byte(pid)
	i := atomic.AddUint32(&r.inc, 1)
	p[6] = byte(i >> 24)
	p[7] = byte(i >> 16)
	p[8] = byte(i >> 8)
	p[9] = byte(i)
	return len(p), nil
}

type monolithReader struct {
	hi int32
	lo uint64
}

func NewMonolithReader() io.Reader {
	return &monolithReader{hi: rand.Int31(), lo: rand.Uint64()}
}

func (r *monolithReader) Read(p []byte) (n int, err error) {
	if len(p) != 10 {
		return 0, ErrIncorrectBufferSize
	}
	hi := atomic.LoadInt32(&r.hi)
	lo := atomic.AddUint64(&r.lo, 1)
	if lo == 0 {
		atomic.AddInt32(&r.hi, 1)
	}
	p[0] = byte(hi >> 8)
	p[1] = byte(hi)
	for i := 7; i >= 0; i-- {
		p[9-i] = byte(lo >> (8 * uint(i)))
	}

	return len(p), nil
}
