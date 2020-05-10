# simple-paxos-go
A simple implementation of the paxos consensus algorithm in Golang
This README assumes you have a pre-existing understanding of the paxos algorithmm
this project implements a simple version of Paxos

To run this project simply type
```
$ go run paxos.go
```
while logated in the base directory with axis to the "paxos.go" file

## UNDERSTANDING LOGS
The 3 type of roles that nodes can assume declare themselves as they log updates.

**Proposer**
A proposer log might look like
```
PROPOSER 0: sending on prep_ch: {id:1 val:-1}
```
A proposer declares it's id and then declares the channel it's either sending
or receiving on, finally it will print the value being sent or received
generally the format as seen above
```
PROPOSER <proposer id>: <sending or receiving> on <channel>: {id:<id of message> val:<value of message, -1 if no value>}
```

##Acceptor
There are acceptors
