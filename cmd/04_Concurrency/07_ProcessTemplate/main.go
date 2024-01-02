package main

const noWorkers = 10

func main() {
	serv := newService()
	defer serv.cleanUp()

	go dispatchWork(noWorkers, serv)

	receiveWork(serv)
}
