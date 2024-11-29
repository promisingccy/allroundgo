package main

import (
	"net"
	"sync"
)

func main() {

}

// 若以下代码在多个Go程中调用，就会导致 unsafeService 映射产生竞争。
// 对映射的并发读写是不安全的：
var unsafeService map[string]net.Addr

func unsafeRegisterService(name string, addr net.Addr) {
	unsafeService[name] = addr
}

func unsafeLookupService(name string) net.Addr {
	return unsafeService[name]
}

// ----------------- 要保证此代码的安全，需通过互斥锁来保护对它的访问：
var (
	safeService map[string]net.Addr
	serviceMu   sync.Mutex
)

func RegisterService(name string, addr net.Addr) {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	safeService[name] = addr
}

func LookupService(name string) net.Addr {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	return safeService[name]
}
