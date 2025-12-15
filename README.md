This repo provides a proof-of-concept to demonstrate the locking bug in [graph-gophers/dataloader](https://github.com/graph-gophers/dataloader).

To run the example, pull down the repo and run `make test`!

**Example Output:**

```shell
go test ./... -v -timeout 5s

=== RUN   TestLoaderWithTracing
panic: test timed out after 5s
	running tests:
		TestLoaderWithTracing (5s)

goroutine 38 [running]:
testing.(*M).startAlarm.func1()
	/usr/local/go/src/testing/testing.go:2682 +0x2b0
created by time.goFunc
	/usr/local/go/src/time/sleep.go:215 +0x38

goroutine 1 [chan receive]:
testing.(*T).Run(0x14000100380, {0x104f89888?, 0x14000116b38?}, 0x105009fe0)
	/usr/local/go/src/testing/testing.go:2005 +0x378
testing.runTests.func1(0x14000100380)
	/usr/local/go/src/testing/testing.go:2477 +0x38
testing.tRunner(0x14000100380, 0x14000116c68)
	/usr/local/go/src/testing/testing.go:1934 +0xc8
testing.runTests(0x1400012a018, {0x1050fcd20, 0x1, 0x1}, {0x14000130160?, 0x7?, 0x105106c00?})
	/usr/local/go/src/testing/testing.go:2475 +0x3b8
testing.(*M).Run(0x1400011a140)
	/usr/local/go/src/testing/testing.go:2337 +0x530
main.main()
	_testmain.go:45 +0x80

goroutine 35 [chan receive]:
github.com/graph-gophers/dataloader/v7.(*Loader[...]).Load.func1()
	/Users/aiden/go/pkg/mod/github.com/graph-gophers/dataloader/v7@v7.1.2/dataloader.go:263 +0x15c
github.com/aidenwallis/graph-gophers-dataloader-locking-bug-poc.(*Tracer[...].TraceLoad.func1()
	/Users/aiden/dev/graph-gophers-dataloader-locking-bug-poc/tracer.go:15 +0x24
github.com/graph-gophers/dataloader/v7.(*Loader[...]).Load(0x10500dd60, {0x10500c328, 0x10512d1c0}, {0x104f85516, 0x3})
	/Users/aiden/go/pkg/mod/github.com/graph-gophers/dataloader/v7@v7.1.2/dataloader.go:308 +0x600
github.com/aidenwallis/graph-gophers-dataloader-locking-bug-poc.TestLoaderWithTracing(0x14000100540)
	/Users/aiden/dev/graph-gophers-dataloader-locking-bug-poc/main_test.go:12 +0x4c
testing.tRunner(0x14000100540, 0x105009fe0)
	/usr/local/go/src/testing/testing.go:1934 +0xc8
created by testing.(*T).Run in goroutine 1
	/usr/local/go/src/testing/testing.go:1997 +0x364

goroutine 36 [chan receive]:
github.com/graph-gophers/dataloader/v7.(*batcher[...]).batch(0x10500d280, {0x10500c328, 0x10512d1c0})
	/Users/aiden/go/pkg/mod/github.com/graph-gophers/dataloader/v7@v7.1.2/dataloader.go:475 +0xc4
created by github.com/graph-gophers/dataloader/v7.(*Loader[...]).Load in goroutine 35
	/Users/aiden/go/pkg/mod/github.com/graph-gophers/dataloader/v7@v7.1.2/dataloader.go:289 +0x4e4

goroutine 37 [sync.Mutex.Lock]:
internal/sync.runtime_SemacquireMutex(0x14000100bf0?, 0x0?, 0x10500a558?)
	/usr/local/go/src/runtime/sema.go:95 +0x28
internal/sync.(*Mutex).lockSlow(0x14000168048)
	/usr/local/go/src/internal/sync/mutex.go:149 +0x170
internal/sync.(*Mutex).Lock(...)
	/usr/local/go/src/internal/sync/mutex.go:70
sync.(*Mutex).Lock(...)
	/usr/local/go/src/sync/mutex.go:46
github.com/graph-gophers/dataloader/v7.(*Loader[...]).sleeper(0x10500dd60, 0x1400011e690, 0x1400010c230)
	/Users/aiden/go/pkg/mod/github.com/graph-gophers/dataloader/v7@v7.1.2/dataloader.go:546 +0xc8
created by github.com/graph-gophers/dataloader/v7.(*Loader[...]).Load in goroutine 35
	/Users/aiden/go/pkg/mod/github.com/graph-gophers/dataloader/v7@v7.1.2/dataloader.go:292 +0x594
FAIL	github.com/aidenwallis/graph-gophers-dataloader-locking-bug-poc	5.203s
FAIL
```
