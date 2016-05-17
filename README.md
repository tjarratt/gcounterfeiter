Gomega Matchers for Counterfeiter
=================================
This package provides some helpful [Gomega](https://github.com/onsi/gomega) matchers that help you write effective assertions against [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) fakes.

HaveReceived()
--------------
Verifies that a function was invoked on a counterfeiter fake.

e.g.:

```go
myFake := new(FakeSomething)
myFake.Something("arg1", 0)
Expect(myFake).To(HaveReceived("Something").With(Equal("arg1")).ANdWith(BeEquivalentTo(0)))
```

This actually works with any object that implements an "invocation recording" interface.

```go
type Recorder interface{
  Invocations() map[string][][]interface{}
}
```

Requested Features
------------------
* I should be able to specify multiple arguments at once
  - e.g.: `Expect(myFake).To(HaveReceived("Something").With(Equal("my-arg", Equal(0)))
* I should be able to specify a number of times a function was invoked
  - e.g.: `Expect(myFake).To(HaveReceived("Something").Times(1))
* I should not need to specify a gomega matcher if I just want to check equality
  - e.g.: these should be equivalent:
  - `Expect(myFake).To(HaveReceived("Something").With(Equal("arg1"), Equal(0)))`
  - `Expect(myFake).To(HaveReceived("Something").With("arg1", 0))`
