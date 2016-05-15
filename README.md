Gomega Matchers for Counterfeiter
=================================
This package provides some helpful [Gomega](https://github.com/onsi/gomega) matchers that help you write effective assertions against [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) fakes.

HaveReceived()
--------------
Verifies that a function was invoked on a counterfeiter fake.

e.g.:

```go
myFake := new(FakeSomething)
myFake.Something("arg1", "arg2")
Expect(myFake).To(HaveReceived("Something").With("arg1", "arg2"))
```

This actually works with any object that implements an "invocation recording" interface.

```go
type Recorder interface{
  Invocations() map[string][][]interface{}
}
```
