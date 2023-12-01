//go:build !linux
// +build !linux

package rakeshjack

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
