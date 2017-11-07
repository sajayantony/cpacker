# Directory to container packer

Simple utility to pack and unpack contents as docker images. 

### Get cpacker

```
go install  github.com/sajayantony/cpacker
```

or clone the repository and `go build` 

### To pack a directory as a container

```
cpacker pack <SRC_DIR> cpackertest
```

### To unpack the contents to a directory

```
