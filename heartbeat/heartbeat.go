/* Watchdog Heartbeat
 *
 * Copyright (c) 2017 Bryant Moscon
 *
 * Please see the LICENSE file for the terms and conditions
 * associated with this software.
 *
 */

package heartbeat

import (
	"net/http"
)

func Heartbeat(addr string, name string) {
	url := "http://" + addr + "/heartbeat?id=" + name
	http.Get(url)
}
