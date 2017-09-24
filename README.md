# go-ackermann

This an implementation of the most difficult program to compute. The Ackermann function.

I've added a lookup table to speed up processing a bit, but I haven't managed to compute `ackermann(4, 2)`.

I am using a Intel Core i7 Skylake with 32GB of RAM to compute this, and got pretty fast results until Go's stack size limit.

After increasing the stack size limit I caught the following weird panic that might seem something related to system or kernel protection, 
but got no logged kernel messages.


```
2017/09/24 12:43:38 Calculated `ackermann(4, 1) = 65533` in 34.248921ms
fatal error: unexpected signal during runtime execution
[signal SIGSEGV: segmentation violation code=0x1 addr=0xc3b8280b80 pc=0x452521]

runtime stack:
runtime.throw(0x4cbd46, 0x2a)
	/home/herbert/.go/src/runtime/panic.go:605 +0x95
runtime.sigpanic()
	/home/herbert/.go/src/runtime/signal_unix.go:351 +0x2b8
runtime.memmove(0xc3b8280b68, 0xc63e160368, 0x7ffffc98)
	/home/herbert/.go/src/runtime/memmove_amd64.s:419 +0x501
runtime.copystack(0xc420000180, 0x100000000, 0x7f8532acb601)
	/home/herbert/.go/src/runtime/stack.go:876 +0x141
runtime.newstack(0x0)
	/home/herbert/.go/src/runtime/stack.go:1059 +0x330
runtime.morestack()
	/home/herbert/.go/src/runtime/asm_amd64.s:415 +0x86

goroutine 1 [copystack]:
runtime.memhash128(0xc63e1603f8, 0x54871e72, 0x10)
	/home/herbert/.go/src/runtime/alg.go:62 +0x57 fp=0xc63e160370 sp=0xc63e160368 pc=0x401477
runtime.mapaccess2(0x4accc0, 0xc42006a150, 0xc63e1603f8, 0xc63e160458, 0x54871e72)
	/home/herbert/.go/src/runtime/hashmap.go:416 +0x73 fp=0xc63e1603b8 sp=0xc63e160370 pc=0x407ce3
```
