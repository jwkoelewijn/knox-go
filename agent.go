package main

import (
	"fmt"
	"log"
	"net"
)

type Agent struct {
	Name          string
	Bandwidth     int
	Documents     *DocumentList
	Connection    net.Conn
	Whistleblower string
	parser        *Parser
	Solution      Solution
}

func NewAgent() *Agent {
	docs := NewDocumentList()
	agent := &Agent{Documents: docs}
	return agent
}

func (a *Agent) Connect() {
	a.Log("Connecting")
	conn, err := net.Dial("tcp", "knox.nedap.healthcare:80")
	if err != nil {
		panic(err)
	}
	a.Connection = conn
	a.parser = NewParser(a, conn)
	fmt.Fprintf(conn, "jwtje\n")
}

func (a *Agent) LogIn() {
	a.Name = a.parser.ParseAgentName()
	a.Log("Logged in...")
}

func (a *Agent) WaitForOthers() {
	a.parser.WaitForStart()
}

func (a *Agent) Disconnect() {
	a.Log("Disconnecting")
	a.Connection.Close()
}

func (a *Agent) Log(line string) {
	log.Printf("%s: %s", a.Name, line)
}

func (a *Agent) writeLine(line string) {
	a.Log(line)
	fmt.Fprintf(a.Connection, "%s\n", line)
}

func (a *Agent) LoadList() {
	a.writeLine("/list")
	a.parser.ParseList()
}

func (a *Agent) String() string {
	return fmt.Sprintf("%s, BW %dKB (->%s)\n%s\nSolution (%d/%d) [%d]:\n%s\n", a.Name, a.Bandwidth, a.Whistleblower, a.documentString(), a.Solution.Cost, a.Bandwidth, a.Solution.Value(), a.Solution.Documents.String())
}

func (a *Agent) Message(rec, msg string) {
	a.writeLine(fmt.Sprintf("/msg %s %s", rec, msg))
	for a.parser.Scan() {
		a.parser.Text()
	}
}

func (a *Agent) documentString() string {
	return a.Documents.String()
}

func (a *Agent) FindOthers() {
	a.writeLine("/look")
	a.Whistleblower = a.parser.ParseName()
}

func (a *Agent) AddDocument(doc Document) {
	a.Documents = a.Documents.Add(doc)
	log.Printf("added doc, now %+v", *a.Documents)
}

func (a *Agent) Length() int {
	return a.Documents.Length()
}

func (a *Agent) FindSolution() {
	a.Solution = a.Documents.FindBestSolution(a.Bandwidth)
}

func (a *Agent) sendDocument(doc Document) {
	a.writeLine(fmt.Sprintf("/send %s %s", a.Whistleblower, doc.Name))
	if a.parser.Scan() {
		a.parser.Text()
	}
}

func (a *Agent) SendSolution() {
	for _, doc := range a.Solution.Documents {
		a.sendDocument(doc)
	}
}
