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

type bpf_debugConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpf_debugSecEvent struct {
	Meta     bpf_debugSecEventMetaT
	Filename [256]uint8
	Buf      [2048]uint8
	Type     uint8
	_        [1]byte
	Conn     bpf_debugConnectionInfoT
	_        [2]byte
}

type bpf_debugSecEventMetaT struct {
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

type bpf_debugSockArgsT struct{ Addr uint64 }

// loadBpf_debug returns the embedded CollectionSpec for bpf_debug.
func loadBpf_debug() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_debugBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_debug: %w", err)
	}

	return spec, err
}

// loadBpf_debugObjects loads bpf_debug and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_debugObjects
//	*bpf_debugPrograms
//	*bpf_debugMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_debugObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_debug()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_debugSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugSpecs struct {
	bpf_debugProgramSpecs
	bpf_debugMapSpecs
}

// bpf_debugSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugProgramSpecs struct {
	KprobeDoTaskDead        *ebpf.ProgramSpec `ebpf:"kprobe_do_task_dead"`
	KprobeSysExecve         *ebpf.ProgramSpec `ebpf:"kprobe_sys_execve"`
	KprobeSysExecveat       *ebpf.ProgramSpec `ebpf:"kprobe_sys_execveat"`
	KprobeSysRename         *ebpf.ProgramSpec `ebpf:"kprobe_sys_rename"`
	KprobeSysRenameat       *ebpf.ProgramSpec `ebpf:"kprobe_sys_renameat"`
	KprobeSysUnlink         *ebpf.ProgramSpec `ebpf:"kprobe_sys_unlink"`
	KprobeSysUnlinkat       *ebpf.ProgramSpec `ebpf:"kprobe_sys_unlinkat"`
	KprobeTcpConnect        *ebpf.ProgramSpec `ebpf:"kprobe_tcp_connect"`
	KprobeTcpRcvEstablished *ebpf.ProgramSpec `ebpf:"kprobe_tcp_rcv_established"`
	KprobeWakeUpNewTask     *ebpf.ProgramSpec `ebpf:"kprobe_wake_up_new_task"`
	KretprobeSockAlloc      *ebpf.ProgramSpec `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4     *ebpf.ProgramSpec `ebpf:"kretprobe_sys_accept4"`
	KretprobeSysConnect     *ebpf.ProgramSpec `ebpf:"kretprobe_sys_connect"`
	SocketHttpFilter        *ebpf.ProgramSpec `ebpf:"socket__http_filter"`
	SyscallEnterExecve      *ebpf.ProgramSpec `ebpf:"syscall_enter_execve"`
	SyscallEnterExecveat    *ebpf.ProgramSpec `ebpf:"syscall_enter_execveat"`
}

// bpf_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugMapSpecs struct {
	ActiveAcceptArgs    *ebpf.MapSpec `ebpf:"active_accept_args"`
	ActiveConnectArgs   *ebpf.MapSpec `ebpf:"active_connect_args"`
	ActivePids          *ebpf.MapSpec `ebpf:"active_pids"`
	Events              *ebpf.MapSpec `ebpf:"events"`
	FilteredConnections *ebpf.MapSpec `ebpf:"filtered_connections"`
}

// bpf_debugObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugObjects struct {
	bpf_debugPrograms
	bpf_debugMaps
}

func (o *bpf_debugObjects) Close() error {
	return _Bpf_debugClose(
		&o.bpf_debugPrograms,
		&o.bpf_debugMaps,
	)
}

// bpf_debugMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugMaps struct {
	ActiveAcceptArgs    *ebpf.Map `ebpf:"active_accept_args"`
	ActiveConnectArgs   *ebpf.Map `ebpf:"active_connect_args"`
	ActivePids          *ebpf.Map `ebpf:"active_pids"`
	Events              *ebpf.Map `ebpf:"events"`
	FilteredConnections *ebpf.Map `ebpf:"filtered_connections"`
}

func (m *bpf_debugMaps) Close() error {
	return _Bpf_debugClose(
		m.ActiveAcceptArgs,
		m.ActiveConnectArgs,
		m.ActivePids,
		m.Events,
		m.FilteredConnections,
	)
}

// bpf_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugPrograms struct {
	KprobeDoTaskDead        *ebpf.Program `ebpf:"kprobe_do_task_dead"`
	KprobeSysExecve         *ebpf.Program `ebpf:"kprobe_sys_execve"`
	KprobeSysExecveat       *ebpf.Program `ebpf:"kprobe_sys_execveat"`
	KprobeSysRename         *ebpf.Program `ebpf:"kprobe_sys_rename"`
	KprobeSysRenameat       *ebpf.Program `ebpf:"kprobe_sys_renameat"`
	KprobeSysUnlink         *ebpf.Program `ebpf:"kprobe_sys_unlink"`
	KprobeSysUnlinkat       *ebpf.Program `ebpf:"kprobe_sys_unlinkat"`
	KprobeTcpConnect        *ebpf.Program `ebpf:"kprobe_tcp_connect"`
	KprobeTcpRcvEstablished *ebpf.Program `ebpf:"kprobe_tcp_rcv_established"`
	KprobeWakeUpNewTask     *ebpf.Program `ebpf:"kprobe_wake_up_new_task"`
	KretprobeSockAlloc      *ebpf.Program `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4     *ebpf.Program `ebpf:"kretprobe_sys_accept4"`
	KretprobeSysConnect     *ebpf.Program `ebpf:"kretprobe_sys_connect"`
	SocketHttpFilter        *ebpf.Program `ebpf:"socket__http_filter"`
	SyscallEnterExecve      *ebpf.Program `ebpf:"syscall_enter_execve"`
	SyscallEnterExecveat    *ebpf.Program `ebpf:"syscall_enter_execveat"`
}

func (p *bpf_debugPrograms) Close() error {
	return _Bpf_debugClose(
		p.KprobeDoTaskDead,
		p.KprobeSysExecve,
		p.KprobeSysExecveat,
		p.KprobeSysRename,
		p.KprobeSysRenameat,
		p.KprobeSysUnlink,
		p.KprobeSysUnlinkat,
		p.KprobeTcpConnect,
		p.KprobeTcpRcvEstablished,
		p.KprobeWakeUpNewTask,
		p.KretprobeSockAlloc,
		p.KretprobeSysAccept4,
		p.KretprobeSysConnect,
		p.SocketHttpFilter,
		p.SyscallEnterExecve,
		p.SyscallEnterExecveat,
	)
}

func _Bpf_debugClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_debug_bpfel_arm64.o
var _Bpf_debugBytes []byte