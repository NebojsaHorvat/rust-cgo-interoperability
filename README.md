Repo is used as a demo for calling dynamic library written in Rust from go code. 
Rust code also has a callback function which it call in go before finishing it's execution.
Go callback function is wrapped with C function which is acctually sent to Rust function as argument

Steps to recreate:

Create go module if it does not exist:
```go mod init main```

Build Rust lib:
```rustc --crate-type=cdylib lib.rs```

Build go project:
```go build```
