# trimpath reproduction

This repository only exists, to reproduce the issues we face in [editorconfig-checker/editorconfig-checker#397](https://github.com/editorconfig-checker/editorconfig-checker/issues/397).

This example is pared down to the minimal possible example demonstrating the problematic behaviour.

## Static Demonstration

```text
kandre@mainframe(pts/10)[main] ~/git-repos/trimpath-repro % ./demonstrate-issue.sh
first some setup: we create a directory named github.com that even the current user is not allowed to write to
the github.com directory would be automatically created by go-snaps under -trimpath
+ install -d -m 555 github.com
+ set +x


now we prove, that the test work perfectly, if GOFLAGS is empty

+ export GOFLAGS=
+ GOFLAGS=
+ go test -count=1 -v ./...
=== RUN   TestRun
--- PASS: TestRun (0.00s)
PASS
ok      github.com/klaernie/trimpath-repro      0.003s
+ set +x


and now we can prove, that if we add -trimpath to GOFLAGS, the tests start to break

+ export GOFLAGS=-trimpath
+ GOFLAGS=-trimpath
+ go test -count=1 -v ./...
=== RUN   TestRun
    main_test.go:10: mkdir github.com/klaernie: permission denied
--- FAIL: TestRun (0.00s)
FAIL
FAIL    github.com/klaernie/trimpath-repro      0.003s
FAIL
+ set +x
kandre@mainframe(pts/10)[main] ~/git-repos/trimpath-repro %
```
