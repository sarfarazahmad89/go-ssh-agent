## Go SSH Agent as executables ##
SSH agent, derived from golang's rich SSH libraries, turned into standalone executables for Windows and Linux. 

You might prefer this over Microsoft's port of OpenSSH as that one lacks a few things like, [_timed expiration  when loading new credentials_ ](https://github.com/openssh/libopenssh/blob/master/ssh/PROTOCOL.agent#L101)


### BUILD
##### Linux or Windows
````
$ git clone https://github.com/sarfarazahmad89/go-ssh-agent
$ cd go-ssh-agent
$ go build .
````
#####

##### Cross build (Linux -> Windows) 
````
$ git clone https://github.com/sarfarazahmad89/go-ssh-agent
$ cd go-ssh-agent
$ GOOS="windows" GOARCH="amd64" go build .
````
#####
