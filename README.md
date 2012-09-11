# treap

Another treap in Go. Somewhat optimized, but still sort of generic.

## Performance

I made a set of benchmarks comparing the Treap type with the native map:

    BenchmarkTreap10	 5000000	       440 ns/op
    BenchmarkTreap100	 5000000	       635 ns/op
    BenchmarkTreap1000	 2000000	       865 ns/op
    BenchmarkTreap10000	 1000000	      1128 ns/op
    BenchmarkTreap100000	 1000000	      1450 ns/op
    BenchmarkTreap1000000	 1000000	      1821 ns/op
    BenchmarkMap10	 5000000	       352 ns/op
    BenchmarkMap100	 5000000	       467 ns/op
    BenchmarkMap1000	 5000000	       454 ns/op
    BenchmarkMap10000	 5000000	       471 ns/op
    BenchmarkMap100000	 5000000	       772 ns/op
    BenchmarkMap1000000	 5000000	       765 ns/op

Not too shabby.

Only drawback is that the only key types allowed in Treap is `[]byte`. But I usually have no trouble converting keys to `[]byte` for these cases.

## Documentation

http://go.pkgdoc.org/github.com/zond/treap