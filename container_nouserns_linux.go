// +build !go1.4

package libcontainer

import (
	"fmt"
	"syscall"
)

// addUidGidMappings; not available before go 1.4
func (c *linuxContainer) addUidGidMappings(sys *syscall.SysProcAttr) error {
	return fmt.Errorf("User namespace is not supported in golang < 1.4")
}
