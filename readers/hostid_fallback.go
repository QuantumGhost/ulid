// +build !darwin,!linux,!freebsd,!windows

package readers

import "errors"

func readPlatformMachineID() (string, error) {
	return "", errors.New("not implemented")
}
