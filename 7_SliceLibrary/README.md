### Slice Library

Create a function to shuffle an []int.
Create a function to "weave" two []int slices into one another. If one slice is longer than the other, append the remaining values to the end of the combined slice. `([ a, b, c] + [t, u, v, w, x, y, z] = [a, t, b, u, c, v, w, x, y, z])`

example input:

```
slice0 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
slice2 := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
shuffle(slice0)
weave(slice1, slice2)
```

example output:

```
[2 7 5 6 3 4 1 8 10 9] // shuffled
[1 11 2 12 3 13 4 14 5 15 6 16 7 17 8 18 9 19 10 20 21 22 23 24] // weaved
```
