// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package file_write

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type writeEventData struct {
	EventContext struct {
		Ts        uint64
		StartTime uint64
		Pid       uint32
		Tgid      uint32
		Ppid      uint32
		Glpid     uint32
		Uid       uint32
		Gid       uint32
		Comm      [16]uint8
		Cwd       [32]uint8
		CgroupId  uint64
		NodeInfo  struct {
			Sysname    [65]uint8
			Nodename   [65]uint8
			Release    [65]uint8
			Version    [65]uint8
			Machine    [65]uint8
			Domainname [65]uint8
		}
		MountInfo struct {
			MountId      int32
			MountNsId    uint32
			MountDevname [256]uint8
		}
	}
	_     [2]byte
	Id    int32
	Fd    int32
	Count uint64
	Ret   int64
}

// loadWrite returns the embedded CollectionSpec for write.
func loadWrite() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_WriteBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load write: %w", err)
	}

	return spec, err
}

// loadWriteObjects loads write and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*writeObjects
//	*writePrograms
//	*writeMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadWriteObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadWrite()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// writeSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type writeSpecs struct {
	writeProgramSpecs
	writeMapSpecs
}

// writeSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type writeProgramSpecs struct {
	KprobeWriteEntry   *ebpf.ProgramSpec `ebpf:"kprobe_write_entry"`
	KretprobeWriteExit *ebpf.ProgramSpec `ebpf:"kretprobe_write_exit"`
}

// writeMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type writeMapSpecs struct {
	Event *ebpf.MapSpec `ebpf:"event"`
}

// writeObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadWriteObjects or ebpf.CollectionSpec.LoadAndAssign.
type writeObjects struct {
	writePrograms
	writeMaps
}

func (o *writeObjects) Close() error {
	return _WriteClose(
		&o.writePrograms,
		&o.writeMaps,
	)
}

// writeMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadWriteObjects or ebpf.CollectionSpec.LoadAndAssign.
type writeMaps struct {
	Event *ebpf.Map `ebpf:"event"`
}

func (m *writeMaps) Close() error {
	return _WriteClose(
		m.Event,
	)
}

// writePrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadWriteObjects or ebpf.CollectionSpec.LoadAndAssign.
type writePrograms struct {
	KprobeWriteEntry   *ebpf.Program `ebpf:"kprobe_write_entry"`
	KretprobeWriteExit *ebpf.Program `ebpf:"kretprobe_write_exit"`
}

func (p *writePrograms) Close() error {
	return _WriteClose(
		p.KprobeWriteEntry,
		p.KretprobeWriteExit,
	)
}

func _WriteClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed write_bpfel_x86.o
var _WriteBytes []byte
