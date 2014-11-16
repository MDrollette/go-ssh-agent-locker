package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Security
#import <Foundation/Foundation.h>
extern OSStatus go_keychain_locked(SecKeychainEvent keychainEvent, SecKeychainCallbackInfo *info, void *context);
*/
import "C"

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"unsafe"
)

//export go_keychain_locked
func go_keychain_locked(keychainEvent C.SecKeychainEvent, info *C.SecKeychainCallbackInfo, context unsafe.Pointer) C.OSStatus {
	bufferSize := C.UInt32(4096)
	cstring := make([]C.char, bufferSize)
	keychainPath := &cstring[0]
	C.SecKeychainGetPath(info.keychain, &bufferSize, keychainPath)

	log.Println("Keychain locked:", C.GoString(keychainPath))

	log.Println("Stopping ssh-agent...")
	err := stopAgent()
	if nil != err {
		log.Println("Error stopping ssh-agent", err)
		return C.OSStatus(1)
	}

	return C.OSStatus(0)
}

func stopAgent() error {
	cmd := exec.Command("/bin/launchctl", "stop", "org.openbsd.ssh-agent")
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

func main() {
	C.SecKeychainAddCallback((*[0]byte)(C.go_keychain_locked), C.kSecLockEventMask, nil)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c
		log.Printf("Got signal %v. Shutting down...\n", sig)
		Stop()
	}()

	Run()
}
