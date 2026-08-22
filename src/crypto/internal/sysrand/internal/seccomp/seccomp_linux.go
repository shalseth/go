// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seccomp

/*
#include <sys/prctl.h>
#include <sys/syscall.h>
#include <errno.h>
#include <stddef.h>
#include <unistd.h>
#include <stdint.h>

// A few definitions copied from linux/filter.h and linux/seccomp.h,
// which might not be available on all systems.

struct sock_filter {
    uint16_t code;
    uint8_t jt;
    uint8_t jf;
    uint32_t k;
};

struct sock_fprog {
    unsigned short len;
    struct sock_filter *filter;
};

#define BPF_LD	0x00
#define BPF_W	0x00
#define BPF_ABS	0x20
#define BPF_JMP	0x05
#define BPF_JEQ	0x10
#define BPF_K	0x00
#define BPF_RET	0x06

#define BPF_STMT(code, k) { (unsigned short)(code), 0, 0, k }
#define BPF_JUMP(code, k, jt, jf) { (unsigned short)(code), jt, jf, k }

struct seccomp_data {
	int nr;
	uint32_t arch;
	uint64_t instruction_pointer;
	uint64_t args[6];
};

#define SECCOMP_RET_ERRNO 0x00050000U
#define SECCOMP_RET_ALLOW 0x7fff0000U
#define SECCOMP_SET_MODE_FILTER 1

int disable_getrandom() {
    if (prctl(PR_SET_NO_NEW_PRIVS, 1, 0, 0, 0)) {
        return 1;
    }
    struct sock_filter filter[] = {
        BPF_STMT(BPF_LD | BPF_W | BPF_ABS, (offsetof(struct seccomp_data, nr))),
        BPF_JUMP(BPF_JMP | BPF_JEQ | BPF_K, SYS_getrandom, 0, 1),
        BPF_STMT(BPF_RET | BPF_K, SECCOMP_RET_ERRNO | ENOSYS),
        BPF_STMT(BPF_RET | BPF_K, SECCOMP_RET_ALLOW),
    };
    struct sock_fprog prog = {
        .len = sizeof(filter) / sizeof((filter)[0]),
        .filter = filter,
    };
    if (syscall(SYS_seccomp, SECCOMP_SET_MODE_FILTER, 0, &prog)) {
        return -errno;
    }
    return 0;
}
*/
import "C"

import (
	"errors"
	"fmt"
	"syscall"
)

// ErrUnsupported reports that this kernel cannot install a seccomp filter.
// Linux only offers filters on architectures that select
// HAVE_ARCH_SECCOMP_FILTER, which sparc does not.
var ErrUnsupported = errors.New("seccomp filters are not supported on this system")

// DisableGetrandom makes future calls to getrandom(2) fail with ENOSYS. It
// applies only to the current thread and to any programs executed from it.
// Callers should use [runtime.LockOSThread] in a dedicated goroutine.
func DisableGetrandom() error {
	switch rc := C.disable_getrandom(); {
	case rc == 0:
		return nil
	case rc < 0:
		errno := syscall.Errno(-rc)
		if errno == syscall.ENOSYS || errno == syscall.EINVAL {
			return fmt.Errorf("%w: seccomp: %v", ErrUnsupported, errno)
		}
		return fmt.Errorf("failed to disable getrandom: seccomp: %v", errno)
	default:
		return fmt.Errorf("failed to disable getrandom: %v", rc)
	}
}
