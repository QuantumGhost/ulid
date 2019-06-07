// +build freebsd

package readers

import "syscall"

func readPlatformMachineID() (string, error) {
	return syscall.Sysctl("kern.hostuuid")
}
