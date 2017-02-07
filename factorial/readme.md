# factorial

```bash
$ go test -bench .
BenchmarkLoop-8                 30000000                44.1 ns/op
BenchmarkRecursion-8             5000000               278 ns/op
BenchmarkTailRecursion-8         5000000               253 ns/op
PASS
ok      github.com/phedoreanu/go-playground/factorial   4.628s

```

```bash
$ gocyclo -top 3 -avg .
2 factorial Recursion factorial.go:17:1
2 factorial Loop factorial.go:9:1
2 factorial TailRecursion factorial.go:24:1
Average: 1.29
```

```bash
$ mccabe-cyclomatic -p github.com/phedoreanu/go-playground/factorial
4
```