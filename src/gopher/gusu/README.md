# cmd

```
# カバレッジのhtml生成
$ go tool cover -html=profile

# テストの実施
$ go test  -v -coverprofile=profile 
=== RUN   TestGusu
=== RUN   TestGusu/+odd
=== RUN   TestGusu/+even
=== RUN   TestGusu/-odd
=== RUN   TestGusu/-even
--- PASS: TestGusu (0.00s)
    --- PASS: TestGusu/+odd (0.00s)
    --- PASS: TestGusu/+even (0.00s)
    --- PASS: TestGusu/-odd (0.00s)
    --- PASS: TestGusu/-even (0.00s)
PASS
coverage: 100.0% of statements
ok      test/gusu      0.292s
```