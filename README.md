This was an attempt to use pygo to get a Python Module version of the jujusvg binary. That failed miserably so instead it's just the example go file from jujusvg with an installable package.

```
make deps
go get -v github.com/marcoceppi/python-jujusvg
$GOPATH/bin/python-jujusvg <bundle-file>
```
