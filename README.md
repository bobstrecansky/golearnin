# golearnin
Bob, Will, and Ravi work on a go project together.

### Installation:
Install the golang binary:
```
wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
tar xvf go1.11.5.linux-amd64.tar.gz
mv go /usr/local/
```
make sure you have `/usr/local` in your $PATH

create a $GOPATH
`export GOPATH=$HOME/go`

Install dep (dependency management tool for golang):
`curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh`

Clone this repo:
```
mkdir ~/go/src/github.com/bobstrecansky/
cd ~/go/src/github.com/bobstrecansky/ && git clone https://github.com/bobstrecansky/golearnin.git
```
Ensure you have all the correct dependencies in your repo:
`dep ensure -v`

### Build the binary:
`go build cmd/main.go`

### Run the binary:
`./main.go`
