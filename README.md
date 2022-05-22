# A project for test toolkit

If you have any questions or requirements, please let me know.
[tao.xu2008@outlook.com -- Tao.Xu](tao.xu2008@outlook.com)

## Install

```shell
go get -u gitlab.xxx.com/test_toolkit
```

## Usage

### 1. Build local develop env

```shell
# 1.1 Get source code from git
git clone git@gitlab.xxx.com:test_toolkit.git
cd test_toolkit

# 1.2 set env GOPATH
vi /etc/profile
export GOROOT=/usr/local/go  # Set the default goroot for go install path
export GOPATH=$HOME/workspace/go   # set default gopath for go src/pkgs path
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
source /etc/profile

# 1.3 Download third-part libs
go mod download
go mod tidy

# 1.4 build a binary for use(Not required)
go build

# 1.5 run with source code
go run main.go -h
./test_toolkit -h
test_toolkit.exe -h
```

### 2. Basic usage

```shell
$ ./test_toolkit.exe -h

```

### 3. Module Details

#### 3.1 tools

```shell
# tools
```

#### 3.1 test

```shell
# test cases
```
