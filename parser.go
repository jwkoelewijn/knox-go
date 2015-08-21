package main

import (
	"bufio"
	"net"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct {
	agent   *Agent
	scanner *bufio.Scanner
}

func NewParser(agent *Agent, conn net.Conn) *Parser {
	reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(reader)
	return &Parser{agent: agent, scanner: scanner}
}

func (p *Parser) log(line string) {
	p.agent.Log(line)
}

func (p *Parser) Scan() bool {
	return p.scanner.Scan()
}

func (p *Parser) Text() string {
	txt := p.scanner.Text()
	p.agent.Log(txt)
	return txt
}

func (p *Parser) ParseAgentName() (agentName string) {
	for agentName == "" && p.Scan() {
		line := p.Text()
		if strings.Index(line, "-->") > 0 {
			agentName = agentParser(line)
		}
	}
	return
}

func agentParser(line string) string {
	if idx := strings.Index(line, "|"); idx > 0 {
		tmpLine := line[idx+2:]
		tmpLine = tmpLine[0:strings.Index(tmpLine, " ")]
		return tmpLine
	} else {
		return ""
	}
}

func (p *Parser) WaitForStart() {
	for p.Scan() {
		line := p.Text()
		if strings.Index(line, "starting...") > 0 {
			p.log("All arrived")
			return
		}
	}
}

func (p *Parser) ParseName() string {
	var line string
	for idx := 0; idx < 6; idx += 1 {
		p.Scan()
		line = p.Text()
	}
	return parseNameFromLine(line)
}

func parseNameFromLine(line string) string {
	pat := `\A.*\| ((\w|-)+)\z`
	re := regexp.MustCompile(pat)
	match := re.FindStringSubmatch(line)
	if match == nil {
		panic("Expected a name in this line")
	}
	return match[1]
}

func (p *Parser) ParseList() {
	p.agent.Bandwidth = p.parseBandWidth()
	// scan once for header line
	p.Scan()
	for p.agent.Length() < 10 && p.Scan() {
		line := p.Text()
		if strings.Index(line, "list") > 0 {
			doc, err := ParseDocument(line)
			if err != nil {
				panic(err)
			}
			p.agent.AddDocument(doc)
		} else {
			p.log("done")
			return
		}
	}
}

func (p *Parser) parseBandWidth() int {
	pattern := `\A.*Bandwidth:\s+(\d+) KB\z`
	re := regexp.MustCompile(pattern)

	for p.Scan() {
		curLine := p.Text()
		if strings.Index(curLine, "Remaining") > 0 {
			match := re.FindStringSubmatch(curLine)
			if match == nil {
				panic("I was expecting to find something")
			}
			bw, err := strconv.Atoi(match[1])
			if err != nil {
				return 0
			} else {
				return bw
			}
		}
	}
	return 0
}
