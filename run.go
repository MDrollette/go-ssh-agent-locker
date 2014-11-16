package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>
void Run() {
    @autoreleasepool {
        [[NSRunLoop currentRunLoop] runMode:NSDefaultRunLoopMode
            beforeDate:[NSDate dateWithTimeIntervalSinceNow:0.1]];
    }
}
*/
import "C"

import (
	"runtime"
)

var (
	stop = make(chan bool)
)

func init() {
	runtime.LockOSThread()
}

func Run() {
	for {
		select {
		case <-stop:
			return
		default:
			C.Run()
		}
	}
}

func Stop() {
	stop <- true
}
