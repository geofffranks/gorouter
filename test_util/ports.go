package test_util

import (
	. "github.com/onsi/ginkgo/v2"

	"sync"
)

var (
	lastPortUsed int
	portLock     sync.Mutex
	once         sync.Once
)

func NextAvailPort() uint16 {
	portLock.Lock()
	defer portLock.Unlock()

	if lastPortUsed == 0 {
		once.Do(func() {
			const portRangeStart = 25000
			lastPortUsed = portRangeStart + GinkgoParallelNode()
		})
	}

	suiteCfg, _ := GinkgoConfiguration()
	lastPortUsed += suiteCfg.ParallelTotal
	return uint16(lastPortUsed)
}
