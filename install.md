# Golang Install Notes

Follow guide here: https://www.gopherguides.com/courses/preparing-your-environment-for-go-development/

Preferred method is to download binary from https://golang.org/dl/ and install from package.

Extract archive into `/usr/local`:

```
$ tar -C /usr/local -xzf go1.13.1.linux-amd64.tar.gz
```

Create `GOPATH` in home directory:

```
$ mkdir -p $HOME/go/{bin,pkg,src}
```

Add this line to $HOME/.profile:

```
export PATH=$PATH:/usr/local/go/bin
```

Add this line to .bashrc file:
```
export GOPATH="$HOME/go"
```

##### Restart computer for changes to take place.

## Test Installation

Continue steps here https://golang.org/doc/install?download=go1.13.1.linux-amd64.tar.gz under `Test your Installation`

Basic run steps are:

1.  Make directory in `$HOME/go/src/`

  ```
  $ mkdir -p $HOME/go/src/$ProjectName
  ```

2.  Add code.

3.  Build the binary with go tool.

  ```
  $ cd $HOME/go/src/$ProjectName
  $ go build
  ```

4.  Run binary:

  ```
  $ ./$BinaryName
  ```
