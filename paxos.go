package main

type Message struct{
  id int
  val int //usually will be the id
}
const num_cycles = 1
const majority = 2


func main(){
  prep_ch := make(chan Message)
  prom_ch := make(chan [2]Message)
  acc_ch := make(chan Message)
  acc_r_ch := make(chan Message)

}

func proposer(prop_id int, prep_ch chan Message, prom_ch chan Message,
              acc_r_ch chan Message, acc_ch chan Message)){
  time := 0
  for (i := 0; i < num_cycles; i++){
    prepare := Message{id = 1, val = -1}
    prep_ch <- prepare
    init_promise := <-prom_ch
    curr_promise := init_promise
    count := 0
    max_mess := init_promise[1]
    while (init_promise[0].id == curr_promise[0].id){
      count++
      if (curr_promise[1].id > max_mess.id){
        max_mess = curr_promise[1]
      }
      if (count >= majority){
        if (max_mess.val > -1){
          acc_r_ch <- Message(id: curr_promise[0].id, val: max_mess.val)
        }else{
          acc_r_ch <- Message(id: curr_promise[0].id, val: prop_id)
        }
      }
      curr_promise = <-prom_ch
    }
  }
}

func learner(acc_ch chan Message){
  first_accept := <-acc_ch
  accept := first_accept
  accept_count := 0
  while
}


func acceptor(prep_ch chan Message, prom_ch chan Message,
              acc_r_ch chan Message, acc_ch chan Message){
  old_acc := Message{id = -1, val = -1}
  go acceptor_phs1(&old_acc, prep_ch, prom_ch)
  go acceptor_phs2(&old_acc, acc_r_ch, acc_ch)
}

func acceptor_phs1(old_acc *Message, prep_ch chan Message, prom_ch chan Message){
  for (i := 0; i < num_cycles; i++){
    var prepare = <-prep_ch
    if (prepare.id >= *old_acc.id){
      prom_ch <- [2]Message{prepare, *old_acc}
    }
  }
}

func acceptor_phs2(old_acc *Message, acc_r_ch chan Message, acc_ch chan Message){
  for (i := 0; i < num_cycles; i++){
    var acc_request = <-acc_r_ch
    if (acc_request.id >= *old_acc.id){
      acc_ch <- acc_request
      *old_acc = acc_request
    }
  }
}
