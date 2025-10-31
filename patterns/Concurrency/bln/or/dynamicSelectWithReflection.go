package or

import "reflect"

func use2() {
	channel := make(chan string)
	vch := reflect.ValueOf(channel)

	succeed := vch.TrySend(reflect.ValueOf(100))
	println(succeed)

	branches := []reflect.SelectCase{
		{Dir: reflect.SelectDefault},
		{Dir: reflect.SelectRecv, Chan: vch},
	}

	println(reflect.Select(branches))

}
