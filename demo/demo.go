/* Watchdog Demo
 *
 * Copyright (c) 2017 Bryant Moscon
 *
 * Please see the LICENSE file for the terms and conditions
 * associated with this software.
 *
 */

package main

import (
	"fmt"
	"os"

	"github.com/bmoscon/watchdog/heartbeat"
)

func main() {
	path, err := os.Executable()

	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		heartbeat.Heartbeat("127.0.0.1:8888", path)
	}
}
