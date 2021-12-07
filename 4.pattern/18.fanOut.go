package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	type result struct {
		id  int
		op  string
		err error
	}

	// insertUser simulates a database operation.
	insertUser := func(id int) result {
		r := result{
			id:  id,
			op:  fmt.Sprintf("INSERT users VALUE (%id)", id),
		}

		if rand.Intn(10) == 0 {
			r.err = fmt.Errorf("unable to insert %d into user table", id)
		}

		return r
	}

	// insertTrans simulates a database operation.
	insertTrans := func(id int) result {
		r := result{
			id: id,
			op: fmt.Sprintf("INSERT trans VALUE (%d)", id),
		}

		if rand.Intn(10) == 0 {
			r.err = fmt.Errorf("unable to insert %d into user table", id)
		}

		return r
	}

	const routines = 10
	const inserts = routines * 2

	ch := make(chan result, inserts)

	waitInserts := inserts

	for i := 0; i < routines; i++ {
		go func(id int) {
			ch <- insertUser(id)
			// We don't need to wait to start the second insert thanks to the buffered channel.
			// The first send will happen immediately.
			ch <- insertTrans(id)
		}(i)
	}

	for waitInserts > 0 {
		r := <-ch

		log.Printf("N: %d ID: %d OP: %s ERR: %v", waitInserts, r.id, r.op, r.err)
		waitInserts--
	}

	log.Println("Inserts Complete")
}
