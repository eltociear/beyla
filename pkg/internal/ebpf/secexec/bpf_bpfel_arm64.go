// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64
// +build arm64

package secexec

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfSecEvent struct {
	Meta struct {
		Op       uint8
		_        [3]byte
		Pid      uint32
		Tid      uint32
		Ppid     uint32
		Uid      uint32
		Auid     uint32
		NsPid    uint32
		NsPpid   uint32
		PidNsId  uint32
		_        [4]byte
		TimeNs   uint64
		CapEff   uint64
		CapInh   uint64
		CapPerm  uint64
		CgrpId   uint32
		NetNs    uint32
		CgrpName [128]uint8
		Comm     [16]uint8
	}
	Buf [2048]uint8
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	KprobeDoTaskDead     *ebpf.ProgramSpec `ebpf:"kprobe_do_task_dead"`
	SyscallEnterExecve   *ebpf.ProgramSpec `ebpf:"syscall_enter_execve"`
	SyscallEnterExecveat *ebpf.ProgramSpec `ebpf:"syscall_enter_execveat"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	Events *ebpf.MapSpec `ebpf:"events"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	Events *ebpf.Map `ebpf:"events"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.Events,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	KprobeDoTaskDead     *ebpf.Program `ebpf:"kprobe_do_task_dead"`
	SyscallEnterExecve   *ebpf.Program `ebpf:"syscall_enter_execve"`
	SyscallEnterExecveat *ebpf.Program `ebpf:"syscall_enter_execveat"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.KprobeDoTaskDead,
		p.SyscallEnterExecve,
		p.SyscallEnterExecveat,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel_arm64.o
var _BpfBytes []byte
