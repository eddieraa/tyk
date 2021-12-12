//go:build linux
// +build linux

package checkup

import "syscall"

func fileDescriptors() {
	rlimit := &syscall.Rlimit{}

	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, rlimit)
	if err == nil && rlimit.Cur < minFileDescriptors {
		log.Warningf("File descriptor limit %d is too low for production use. A minimum of %d is recommended.\n"+
			"\tThis could have a significant negative impact on performance.\n"+
			"\tPlease refer to the following link for further guidance:\n"+
			"\t\thttps://tyk.io/docs/deploy-tyk-premise-production/#file-handles--file-descriptors",
			rlimit.Cur, minFileDescriptors)
	}
}
