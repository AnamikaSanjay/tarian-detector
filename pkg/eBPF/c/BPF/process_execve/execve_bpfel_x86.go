// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package process_execve

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type execveEventData struct {
	Id           int32
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
		NodeInfo  struct {
			Sysname    [65]uint8
			Nodename   [65]uint8
			Release    [65]uint8
			Version    [65]uint8
			Machine    [65]uint8
			Domainname [65]uint8
		}
	}
	_              [2]byte
	Ret            int32
	BinaryFilepath [4096]uint8
	UserComm       [256][4096]uint8
	EnvVars        [256][4096]uint8
}

// loadExecve returns the embedded CollectionSpec for execve.
func loadExecve() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ExecveBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load execve: %w", err)
	}

	return spec, err
}

// loadExecveObjects loads execve and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*execveObjects
//	*execvePrograms
//	*execveMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadExecveObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadExecve()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// execveSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execveSpecs struct {
	execveProgramSpecs
	execveMapSpecs
}

// execveSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execveProgramSpecs struct {
	KprobeExecveEntry   *ebpf.ProgramSpec `ebpf:"kprobe_execve_entry"`
	KretprobeExecveExit *ebpf.ProgramSpec `ebpf:"kretprobe_execve_exit"`
}

// execveMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execveMapSpecs struct {
	Event *ebpf.MapSpec `ebpf:"event"`
}

// execveObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadExecveObjects or ebpf.CollectionSpec.LoadAndAssign.
type execveObjects struct {
	execvePrograms
	execveMaps
}

func (o *execveObjects) Close() error {
	return _ExecveClose(
		&o.execvePrograms,
		&o.execveMaps,
	)
}

// execveMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadExecveObjects or ebpf.CollectionSpec.LoadAndAssign.
type execveMaps struct {
	Event *ebpf.Map `ebpf:"event"`
}

func (m *execveMaps) Close() error {
	return _ExecveClose(
		m.Event,
	)
}

// execvePrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadExecveObjects or ebpf.CollectionSpec.LoadAndAssign.
type execvePrograms struct {
	KprobeExecveEntry   *ebpf.Program `ebpf:"kprobe_execve_entry"`
	KretprobeExecveExit *ebpf.Program `ebpf:"kretprobe_execve_exit"`
}

func (p *execvePrograms) Close() error {
	return _ExecveClose(
		p.KprobeExecveEntry,
		p.KretprobeExecveExit,
	)
}

func _ExecveClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed execve_bpfel_x86.o
var _ExecveBytes []byte
