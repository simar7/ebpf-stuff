# ebpf-stuff

### Introduction
This repo hosts a bunch eBPF programs with a runner `charlie` and catcher `snoopy` written in Go.
A few examples include:
- Running a shell inside a container

### How to run?
- Charlie: `make clean all run`
- Snoopy (must be run as root): `sudo -E ./snoopy.bin [options]`

### Sample output
- Charlie:
```
$ make run
Hello, eBPF World from within a container!
2019/10/21 23:39:22 writing file with contents:  A quick brown fox jumps over the lazy little dog.
2019/10/21 23:39:22 file written to:  /charlie-481061718.tmp
2019/10/21 23:39:22 reading file from:  /charlie-481061718.tmp
2019/10/21 23:39:22 read file contents:  A quick brown fox jumps over the lazy little dog.
2019/10/21 23:39:22 spawning a new shell within the container...
2019/10/21 23:39:22 shell successfully spawned.
```

- Snoopy (running in execve mode, scanning for spawn of new shell to a container):
```
$ sudo -E ./snoopy.bin -t -l /bin/sh
TIME(s) PCOMM            PID    PPID   RET ARGS
2.207   sh               16254  16194    0 /bin/sh
```