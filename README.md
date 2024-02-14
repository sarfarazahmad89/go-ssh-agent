## Go SSH Agent as executables ##
SSH agent, derived from golang's rich SSH libraries, turned into standalone executables for Windows and Linux. 

You might prefer this over Microsoft's port of OpenSSH as that one lacks a few things like, [_timed expiration  when loading new credentials_ ](https://github.com/openssh/libopenssh/blob/master/ssh/PROTOCOL.agent#L101)

### INSTALL
Generally you can just install this with and it should just install to your $GOPATH/bin/,
```bash
go install github.com/sarfarazahmad89/go-ssh-agent
```

### Alternatively, BUILD, like
##### Linux or Windows
````
$ git clone https://github.com/sarfarazahmad89/go-ssh-agent
$ cd go-ssh-agent
$ go build .
$ ./ssh_agent -h
Usage of ./ssh_agent:
  -sshpipe string
        UNIX socket for the OpenSSH agent (default "/home/ahmad/.ssh/ssh-auth-sock")
$ ./ssh_agent
2023/02/12 11:31:21 started ssh agent on `/home/user1/.ssh/ssh-auth-sock`
````
#####

##### Cross build (Linux -> Windows) 
````
$ git clone https://github.com/sarfarazahmad89/go-ssh-agent
$ cd go-ssh-agent
$ GOOS="windows" GOARCH="amd64" go build .
````
#####
