package main

import (
	//"goAdventure/concurr"
	//"fmt")
	// "goAdventure/error"

	//"goAdventure/intry"
	//"goAdventure/myjsonpack"
	"goAdventure/server"
)

func main() {
	// body := error.Myhttp()
	// fmt.Println(body)
	// error.OsErr()
	// error.AgeCheck()

	/*
		var todo myjsonpack.Myjson
		todo, err := myjsonpack.JsonReqHandler()

		if err != nil {
			return
		}
		fmt.Println(todo.UserID)
		fmt.Println(todo.Title)
		fmt.Println(todo.ID)
		fmt.Println(todo.Completed)

		myjsonpack.My()
	*/

	/*				channels				*/
	/*
		mychan1 := make(chan int)
		mychan2 := make(chan int)
		go concurr.Channel1("naber moruq", mychan1)
		go concurr.Channel1("naber moruqu", mychan2)
		evenorOdd, evenorOdd1 := <-mychan1, <-mychan2
		print(evenorOdd + evenorOdd1)
	*/

	server.Actual()
}
