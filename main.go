package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	agentChannel := make(chan Agent)
	agents := []Agent{}

	wg.Add(3)

	go agent(agentChannel, &wg)
	go agent(agentChannel, &wg)
	go agent(agentChannel, &wg)

	go func() {
		for ag := range agentChannel {
			agents = append(agents, ag)
		}
	}()

	wg.Wait()

	for _, agent := range agents {
		log.Println(agent.String())
		// kan niet, bij de 3e is de eerste al disconnected, duh
		//agent.Message(agent.Whistleblower, "done")
		agent.Disconnect()
	}
}

func agent(channel chan Agent, wg *sync.WaitGroup) {

	agent := NewAgent()
	agent.Connect()
	agent.LogIn()
	agent.WaitForOthers()
	agent.LoadList()
	agent.FindOthers()
	agent.FindSolution()
	agent.SendSolution()
	agent.Message(agent.Whistleblower, "done")
	channel <- *agent
	wg.Done()
}
