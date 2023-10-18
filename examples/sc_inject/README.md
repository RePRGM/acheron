# Fork Update
__Extremely__ minimal update. Copying and pasting stageless shellcode (especially formatted differently than the typical C/C++ or C# byte array styles) is an absolute nightmare. So, I made use of Golang's ability to embed files.

## How To
1. Edit main.go and add your raw shellcode file to the `//go:embed` line. Ideally, shellcode file is in the same directory. Otherwise, it can be a relative path to this (inject) directory. It **cannot** be an absolute path.
2. Copy and paste one of the two below commands to compile the payload. 

# sc_inject

Extremely simple shellcode injector PoC, that injects calc shellcode using syscalls for `NtAllocateVirtualMemory`+`NtWriteVirtualMemory`+`NtCreateThreadEx`.

Using build tags, you can compile both the direct and indirect syscall versions of the injector, if you want to run them against defensive tools to check out the detection and compare the IOCs of each technique.

```bash
# indirect syscall version (default)
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o sc_inject_indirect.exe

# direct syscall version
GOOS=windows GOARCH=amd64 go build -tags='direct' -ldflags "-s -w" -o sc_inject_direct.exe
```
