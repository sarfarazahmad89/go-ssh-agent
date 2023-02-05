## Go SSH Agent as executables ##
SSH agent, derived from golang's rich SSH libraries, turned into standalone executables for Windows and Linux. 

You might prefer this over Microsoft's port of OpenSSH as that one lacks a few things like, [_timed expiration  when loading new credentials_ ](https://github.com/openssh/libopenssh/blob/master/ssh/PROTOCOL.agent#L101)


### BUILD
##### Linux or Windows
````
$ go build .
````
#####

##### Cross build (Linux -> Windows) 
````
GOOS="windows" GOARCH="amd64" go build .
````
#####
