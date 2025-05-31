# shuffle

Fast pseudo-random shuffling using a Linear Congruential Generator for high speed in-place shuffling. It's twice as fast as Go's native shuffling, and easier to use.

```go
slice := []anything{1, 2, 3, 4, 5, 6}
shuffle.Shuffle(shuffle.Seed(), slice)
```
