# go-ackermann

This an implementation of the most difficult program to compute. The Ackermann function.

I've added a lookup table to speed up processing a bit, but I haven't managed to compute `ackermann(4, 2)`.

I am using a Intel Core i7 Skylake with 32GB of RAM to compute this, and got pretty fast results until Go's stack size limit.

After increasing the stack size limit I caught a weird panic that might seem something related to system or kernel protection, 
but got no logged kernel messages.
