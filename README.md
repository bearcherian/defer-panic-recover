# Defer-Panic-Recover

This is a small program that shows how defer, panic, and recover work in Go.

## what it does

The program prompts for some input, any input, and then repeats it back to the screen.

If the user enters `q`, the program exits. 

If the user enters `!!!`, the program panics.

This panic goes to the calling `start()` function, where we have deferred a call to `recoverMe()`, which checks `recover()`, and, when it identifies a panicked state, calls `start()` again. 

## running the code

```sh
$ go build
$ ./defer-panic-recover
```
