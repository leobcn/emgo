package nbl

import (
	"sync/atomic"
	"sync/fence"
)

// Int64 implements int64 value for one writer and multiple readers. It is
// designed for architectures that does not support 64-bit atomic operations.
type Int64 struct {
	val [2]int64
	aba uintptr
}

func (i *Int64) ABA() uintptr {
	return atomic.LoadUintptr(&i.aba)
}

func (i *Int64) TryLoad(aba uintptr) int64 {
	fence.RD_SMP() // load(i.val[aba&1]) depends on load(i.aba).
	return i.val[aba&1]
}

func (i *Int64) CheckABA(aba uintptr) (uintptr, bool) {
	// load(i.aba) must be observed after load(i.val[aba&1]) and any other load
	// between TryLoad and CheckABA.
	fence.R_SMP() 
	aba1 := atomic.LoadUintptr(&i.aba)
	return aba1, aba1 == aba
}

// Load uses ABA, TryLoad and CheckABA to load and return value of i.
//
// BUG: On 32bit CPUs Load does not guarantee that it returns valid value. The
// probability of failure depends on the frequency of updates:
// 1 kHz: aba wraps onece per 1193 houres,
// 1 MHz: aba wraps once per 72 minutes.
// Load fails if aba wraps beetwen ABA and CheckABA or between subsequent
// CheckABA and after that is readed second time with the same value.
func (i *Int64) Load() int64 {
	aba := i.ABA()
	for {
		v := i.TryLoad(aba)
		var ok bool
		if aba, ok = i.CheckABA(aba); ok {
			return v
		}
	}
}

// WriterLoad is more efficient than Load but there should be guarantee that
// only one writer can write to i at the same time.
func (i *Int64) WriterLoad() int64 {
	return i.val[i.aba&1]
}

// WriterStore stores v into i.
// Only ONE writer can call WriterStore at the same time.
func (i *Int64) WriterStore(v int64) {
	aba := i.aba
	aba++
	i.val[aba&1] = v
	fence.W_SMP() // store(i.val[aba&1]) must be observed before store(i.aba).
	atomic.StoreUintptr(&i.aba, aba)
}

// WriterAdd performs i += v and returns new value of i.
// Only ONE writer can call WriterAdd at the same time.
func (i *Int64) WriterAdd(v int64) int64 {
	v += i.WriterLoad()
	i.WriterStore(v)
	return v
}
