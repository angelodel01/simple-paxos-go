package main

import(
  "fmt"
  "time"
  "sync"
  "os"
)

type Message struct{
  id int
  val int //usually will be the id
}
const num_cycles = 1
const majority = 2
const num_acceptors= 3
const num_proposers = 2
var wg sync.WaitGroup
var phase = 1


func main(){
  prep_ch := make(chan Message, num_acceptors)
  prom_ch := make(chan [2]Message, num_acceptors)
  acc_ch := make(chan Message, num_acceptors)
  acc_r_ch := make(chan Message, num_acceptors)
  acceptor_id := []int{0, 1, 2}
  proposer_id := []int{0, 1}
  counter := 0
  old_acc := Message{id : -1, val : -1}
  for i := 0; i < num_cycles; i++{
    for o := 0; o < num_proposers; o++{
      wg.Add(1)
      counter += 1
      go proposer_phs1(counter, proposer_id[o], prep_ch, prom_ch, acc_r_ch, acc_ch)
    }

    for j := 0; j < num_acceptors; j++{
      wg.Add(1)
      go acceptor_phs1(acceptor_id[j], &old_acc, prep_ch, prom_ch)
    }



    for o := 0; o < num_proposers; o++{
      wg.Add(1)
      go proposer_phs2(counter, proposer_id[o], prep_ch, prom_ch, acc_r_ch, acc_ch)
    }
    for j := 0; j < num_acceptors; j++{
      wg.Add(1)
      go acceptor_phs2(acceptor_id[j], &old_acc, acc_r_ch, acc_ch)
    }

    // proposer_id += 1
    learner(acc_ch)
  }
  wg.Wait()
}


func proposer_phs1(counter int, prop_id int, prep_ch chan Message, prom_ch chan [2]Message,
              acc_r_ch chan Message, acc_ch chan Message){
    defer wg.Done()
    prepare := Message{id: counter, val: -1}
    fmt.Printf("PROPOSER %d sending : %+v\n\n", prop_id, prepare)
    for i := 0; i < num_acceptors; i++{
      prep_ch <- prepare
    }
}


func proposer_phs2(counter int, prop_id int, prep_ch chan Message, prom_ch chan [2]Message,
              acc_r_ch chan Message, acc_ch chan Message){
  defer wg.Done()
  time.Sleep(time.Second)
  sent_flag := true
  init_promise := <-prom_ch
  curr_promise := init_promise
  count := 0
  fmt.Printf("PROPOSER %d receiving off prom_ch: %+v\n\n", prop_id, init_promise)
  max_mess := init_promise[1]
  i := 0
  for (i < num_acceptors) && sent_flag{
    fmt.Printf("PROPOSER %d comparing %+v to %+v\n\n", prop_id, init_promise[0], curr_promise[0])
    if (init_promise[0].id == curr_promise[0].id){
      count++
      if (curr_promise[1].id > max_mess.id){
        max_mess = curr_promise[1]
      }
      if (count >= majority){
        fmt.Printf("PROPOSER %d majority received off prom_ch\n\n", prop_id)
        if (max_mess.val > -1){
          fmt.Printf("in proposer %d promise majority on value %d\n\n", prop_id)
          for k := 0; k < num_acceptors; k++{
            acc_r_ch <- Message{id: curr_promise[0].id, val: max_mess.val}
          }
          sent_flag = false
        }else{
          fmt.Printf("PROPOSER %d promise majority on no value so choosing value : %d\n\n", prop_id, prop_id)
          for k := 0; k < num_acceptors; k++{
            acc_r_ch <- Message{id: curr_promise[0].id, val: prop_id}
          }
          sent_flag = false
        }
      }
      curr_promise = <-prom_ch
      i++
    }
  }
}



func acceptor_phs1(acc_id int, old_acc *Message, prep_ch chan Message, prom_ch chan [2]Message){
  defer wg.Done()
  for i := 0; i < num_proposers; i++{
    var prepare = <-prep_ch
    fmt.Printf("ACCEPTOR %d: received off prep_ch : %+v\n\n", acc_id, prepare)
    if (prepare.id >= (*old_acc).id){
      fmt.Printf("ACCEPTOR %d: sending on prom_ch: %+v\n\n", acc_id, [2]Message{prepare, *old_acc})
      prom_ch <- [2]Message{prepare, *old_acc}
    }else{
      fmt.Printf("ACCEPTOR %d: ignoring %+v as a proposal\n\n", acc_id, prepare)
    }
  }
}

func acceptor_phs2(acc_id int, old_acc *Message, acc_r_ch chan Message, acc_ch chan Message){
  defer wg.Done()
  var acc_request = <-acc_r_ch
  fmt.Printf("ACCEPTOR %d: receiving %+v as a accept-request\n\n", acc_id, acc_request)
  if (acc_request.id >= (*old_acc).id){
    fmt.Printf("ACCEPTOR %d: accepting %+v as a accept-request\n\n", acc_id, acc_request)
    acc_ch <- acc_request
    *old_acc = acc_request
  }else{
    fmt.Printf("ACCEPTOR %d: ignoring %+v as a accept-request\n\n", acc_id, acc_request)
  }
}


func learner(acc_ch chan Message){
  defer wg.Done()
  first_accept := <-acc_ch
  accept := first_accept
  accept_count := 0
  for i := 0; i < num_acceptors; i++{
    if (first_accept.id == accept.id){
      fmt.Printf("LEARNER receiving %+v\n\n", accept)
      accept_count++
      if accept_count >= majority{
        fmt.Printf("WE HAVE REACHED CONCENSUS ON VALUE : %d\n", first_accept.id)
        os.Exit(1)
      }
      accept = <-acc_ch
    }
  }
}
