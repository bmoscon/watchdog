## Watchdog

A very simple example of a watchdog microservice in Go. 


To build:
```
make
```

This will produce `watchdog` and `example` binaries. Run `watchdog` in the backgroudnd, then `example`.

```
$ ./watchdog&
$ ./example
```

`example` will register itself with `watchdog` and then exit. `watchdog` will restart `example` (which will simply exit) endlessly.

