# knox-go

Connects to knox, finds a solution per agent and sends it.
Future work: Let agents send files around to each other, to maximize
total secrecy value of the solution.

# Running

First build, using

```bash
go build
```
followed by:

```bash
./knox
```

# Tests

(limited) Tests can be run by

```bash
go test
```

# Approach
Each agent logs in to a certain channel and retrieves its name, then it waits for the others to connect.
After this they retrieve their respective list of files (Documents). These get stored in the agent's DocumentList.
The DocumentList is then asked to provide an optimal set of documents, given the agent's bandwidth.
The agent then sends each of the documents in the solution and finishes by sending "done" to the 'agent.WhistleBlower'.

## Solution algorithm outline
1. Create all possible combinations of length 0..len(documentList) which have a cost < agent.BandWidth.
2. Pass each combination to a maximizer function, which will try to add documents not already contained in
   the combination as long as solution.Cost + document.Cost < agent.Bandwidth. This is a solution.
3. Select the solution with the highest Value.

## Goroutines
A total of 34 goroutines are used; 1 main, 1 for each agent, 10 solution generators per agent, 1 solution combinator per agent.

### Agent
Each agent is implemented as a goroutine, with the main process waiting for each of them to finish.

### Solution algorithm
Because creating solutions which are based on first creating all possible combinations of length 1 can be
done concurrent to creating solutions which are based on first creating all possible combinations of length 2,
or length 3, etc... a goroutine is started for each n in 0..len(DocumentList).

Found combinations (with a cost < bandwidth)
are passed to a maximizer function, which extends the solution when possible. After this the solution is send
over a buffered channel (chan Solution).

A final goroutine is ranging over the solutions in this channel, comparing them to the current best solution. If it is
better, it is saved as the new best solution.

When all combinator goroutines are done (indicated using a sync.WaitGroup), the solution channel is closed, causing the
collector goroutine to stop ranging over this channel, and pushing its best solution over its output channel.
