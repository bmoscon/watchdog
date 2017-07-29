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
	"net/http"
)

func main() {
	url := "http://127.0.0.1:8888/heartbeat?id=demo/demo"
	http.Get(url)
}
