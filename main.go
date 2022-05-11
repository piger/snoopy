package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Cocoa
#import <Foundation/Foundation.h>
#import <Cocoa/Cocoa.h>

void copyToClipboard(const char *text) {
	NSString *s = [NSString stringWithUTF8String:text];
	NSData *inputData = [s dataUsingEncoding:NSUTF8StringEncoding];

	NSPasteboard* pasteboard = [NSPasteboard generalPasteboard];
	[pasteboard clearContents];
	[pasteboard setData: [@"" dataUsingEncoding:NSUTF8StringEncoding] forType: @"org.nspasteboard.ConcealedType"];
	[pasteboard setData: inputData forType: @"NSStringPboardType"];
}
*/
import "C"

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func run() error {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	payload := strings.TrimRight(string(data), "\n")

	C.copyToClipboard(C.CString(payload))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}
