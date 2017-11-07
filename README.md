# Directory to container packer

Simple utility to pack and unpack contents as docker images. 

### Get cpacker

```
go get  github.com/sajayantony/cpacker
```

### To pack a directory as a container

```
cpacker pack <SRC_DIR> cpackertest
```

### To unpack the contents to a directory

```
cpacker unpack cpackertest ./contents
```
