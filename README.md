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
## example output

```sh
$ ./defer-panic-recover 
What is the location?
abcd
OK, 'abcd'
Whatchu want?
1234
OK, '1234'
Enter something fun
XYZ
OK, 'XYZ'
Enter something fun
alpha789 omega 3
OK, 'alpha789 omega 3'
Are you being served?
!!!
2019/06/10 18:28:23 AAGGGHH!!!
2019/06/10 18:28:23 recovering %!s(main.fn=0x1097cf0) (AAGGGHH!!!)
Whatchu want?
123
OK, '123'
YO!
q
2019/06/10 18:28:32 Bye-Bye!
```
