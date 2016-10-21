// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// time a simple command.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

// TODO: need to handle signals

func main() {
	if len(os.Args) < 2 {
		return
	}
	start := time.Now()
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	d := time.Since(start)
	min := int64(d.Minutes())
	sec := d.Seconds() - float64(60*min)
	fmt.Printf("\nreal\t%dm%.3fs\n", min, sec)
	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			os.Exit(status.ExitStatus())
		}
	}
}
