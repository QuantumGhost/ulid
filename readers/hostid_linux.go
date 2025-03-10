// +build linux

package readers

import "io/ioutil"

func readPlatformMachineID() (string, error) {
	b, err := ioutil.ReadFile("/sys/class/dmi/id/product_uuid")
	return string(b), err
}
