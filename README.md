# simple-paxos-go
A simple implementation of the paxos consensus algorithm in Golang
This README assumes you have a pre-existing understanding of the paxos algorithmm
this project implements a simple version of Paxos

To run this project simply type
```
$ go run paxos.go
```
while logated in the base directory with axis to the "paxos.go" file

## Understanding Logs
The 3 type of roles that nodes can assume declare themselves as they log updates.

**Proposer**
a proposer log might look like
```
PROPOSER 0: sending on prep_ch: {id:1 val:-1}
```
A proposer declares it's id and then declares the channel it's either sending
or receiving on, finally it will print the value being sent or received
generally the format as seen above
```
PROPOSER <proposer id>: <sending or receiving> on <channel>: {id:<id of message> val:<value of message, -1 if no value>}
```

**Acceptor**
an acceptor log might look like
```
ACCEPTOR 0: receiving off acc_r_ch : {id:1 val:0}
```
An acceptor declares it's id and then declares the channel it's either sending
or receiving on, finally it will print the value being sent or received
generally the format as seen above
```
ACCEPTOR <acceptor id>: <sending or receiving> on <channel>: {id:<id of message> val:<value of message, -1 if no value>}
```

**Learner**
a Learner only receives so a learner log might look like
```
LEARNER receiving {id:1 val:0}
```
Learners don't say their Id since their Ids aren't particularly important
they simply log what they receive a learner log generally has the format
seen above
```
LEARNER receiving {id:<id received> val:<value received>}
```


**Concensus**
Once the paxos run has finished you will see a printout like:

```
WE HAVE REACHED CONCENSUS ON VALUE : 0
```
generally with that format:
```
WE HAVE REACHED CONCENSUS ON VALUE : <consensus value>
```

**Miscellaneous other logs**
There are a few other logs where acceptors will let the user know if they are
ignoring certain messages, and proposers will let the user know if they've
reached majority on promises

## Channels

As you can see in the logs I have implemented several different channels
which are outlined with context to the paxos algorithm here.

**prep_ch** = the channel where proposers send their initial prepare request

**prom_ch** = the channel where acceptors send their receival of the initial prepare request

**acc_r_ch** = the channel where proposers can send accept-requests after receiving a majority of promises on ids from acceptors

**acc_ch** = the channel where acceptors broadcast out their decision on a value
