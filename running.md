# Running Tests

```bash
go test ./iteration -v 
```

# Running Benchmarks w/ memory
```bash
go test -bench=. -benchmem  ./iteration
```

# Getting Go Test Coverage
```bash
go test -cover ./iteration
```