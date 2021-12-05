
```shell
dlv debug main.go
```

Objective: Print variable values somewhere on the program!

## Commands that look like they can be useful
break main.main
continue

rebuild: starts the program all over

## Going through the program
- next (n): Seems to step through the current function without getting into other function calls
- step (s): goes through the program falling through every single subfunction 
- stepout (so): steps out of the current function

## Printing things
- args: Displays current function arguments!
- set: Sets a variable in your current function (ex: a=3). This is the coolest thing ever!
- print: prints a variable (ex: print myvariable). This is really cool too!

How do I again print where I was in the program?