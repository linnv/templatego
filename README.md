### when reload or stop app by `kill -number` the number mustn't be 9
e.g:
```
kill -3 $pid
```
linux send terminated by default, that is `kill -15 `

signal table for reference:

```
hangup 1                                                                                 │
interrupt 2                                                                              │
quit 3                                                                                   │
!!!!!killed 9 !!!don't use it
terminated 15                                                                            │
stopped (signal) 19
```
