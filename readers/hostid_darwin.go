// +build darwin

package readers

import "syscall"

func readPlatformMachineID() (string, error) {
	return syscall.Sysctl("kern.uuid")
}
