### Permissions

Given an integer (uint8 is what I used), determine what constants that integer is bit-wise **true** with using `const` golang variables, `iota`, and conditionals. An exercise in some basic operations.

5 made-up masks, similar to the unix file system permissions:

> Read, Write, Delete, Execute, Sudo

input:

`23 // 0010111 in binary`

output:

```
0010111 me
0000001 Read ✅
---
0010111 me
0000010 Write ✅
---
0010111 me
0000100 Delete ✅
---
0010111 me
0001000 Exectute ❌
---
0010111 me
0010000 Super ✅
---
```
