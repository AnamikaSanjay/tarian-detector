// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package file_read

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type readEventData struct {
	Id    int32
	E_ctx struct {
		Ts        uint64
		StartTime uint64
		Pid       uint32
		Tgid      uint32
		Ppid      uint32
		Glpid     uint32
		Uid       uint32
		Gid       uint32
		Nodename  [65]uint8
		Comm      [16]uint8
		Cwd       [32]uint8
	}
	_     [3]byte
	Fd    uint32
	_     [4]byte
	Count uint64
	Ret   uint64
}

// loadRead returns the embedded CollectionSpec for read.
func loadRead() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ReadBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load read: %w", err)
	}

	return spec, err
}

// loadReadObjects loads read and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*readObjects
//	*readPrograms
//	*readMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadReadObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadRead()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// readSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type readSpecs struct {
	readProgramSpecs
	readMapSpecs
}

// readSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type readProgramSpecs struct {
	KprobeReadEntry   *ebpf.ProgramSpec `ebpf:"kprobe_read_entry"`
	KretprobeReadExit *ebpf.ProgramSpec `ebpf:"kretprobe_read_exit"`
}

// readMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type readMapSpecs struct {
	Event *ebpf.MapSpec `ebpf:"event"`
}

// readObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadReadObjects or ebpf.CollectionSpec.LoadAndAssign.
type readObjects struct {
	readPrograms
	readMaps
}

func (o *readObjects) Close() error {
	return _ReadClose(
		&o.readPrograms,
		&o.readMaps,
	)
}

// readMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadReadObjects or ebpf.CollectionSpec.LoadAndAssign.
type readMaps struct {
	Event *ebpf.Map `ebpf:"event"`
}

func (m *readMaps) Close() error {
	return _ReadClose(
		m.Event,
	)
}

// readPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadReadObjects or ebpf.CollectionSpec.LoadAndAssign.
type readPrograms struct {
	KprobeReadEntry   *ebpf.Program `ebpf:"kprobe_read_entry"`
	KretprobeReadExit *ebpf.Program `ebpf:"kretprobe_read_exit"`
}

func (p *readPrograms) Close() error {
	return _ReadClose(
		p.KprobeReadEntry,
		p.KretprobeReadExit,
	)
}

func _ReadClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed read_bpfel_x86.o
var _ReadBytes []byte
