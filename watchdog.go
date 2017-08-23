/* Watchdog
 *
 * Copyright (c) 2017 Bryant Moscon
 *
 * Please see the LICENSE file for the terms and conditions
 * associated with this software.
 *
 */

package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var services map[string]time.Time

func handler(w http.ResponseWriter, r *http.Request) {
	service := r.URL.Query().Get("id")
	log.Printf("Got data from service %s", service)
	if service != "" {
		services[service] = time.Now()
	} else {
		http.Error(w, "invalid data", http.StatusBadRequest)
	}
}

func restart(name string) {
	cmd := exec.Command(name)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	err := cmd.Start()

	if err != nil {
		log.Printf("Error starting service %s: %v", name, err)
	}
}

func watcher(ticker *time.Ticker) {
	for {
		<-ticker.C
		for name, timestamp := range services {
			if time.Now().Sub(timestamp).Seconds() > 10 {
				delete(services, name)
				log.Printf("Service %s died - restarting", name)
				go restart(name)
			}
		}
	}
}

func main() {
	services = make(map[string]time.Time)

	f, err := os.OpenFile("watchdog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Cannot create logfile: ", err)
	}

	defer f.Close()
	log.SetOutput(f)

	ticker := time.NewTicker(5 * time.Second)
	go watcher(ticker)

	http.HandleFunc("/heartbeat", handler)
	http.ListenAndServe(":8888", nil)
}
