## Best example of concurrency

Making a cup off tea, putting the kettle on and then, while I was waiting for it to boil, get the milk out of the fridge, get the tea out of the cupboard, find your favorite mug, put the teabag into the cup and then, when the kettle is done boiling, put the water into the cup


Normally in Go when we call a function 'doSomething()' we wait for it to return (even if it has nothing to return, we want it to finish running)/
- We say this operation is BLOCKING - forces us to wait before finishing
- An operation that does not block in Go will run in a separate process called a goroutine.

A good example of this is like:
Reading down the page of Go code from top to bottom, going 'inside' each function when it gets called to read what it does. When a separate process starts, it's like another reader begins reading inside the function, leaving the original reader to carry on going down the page

## Wrapping Up

- 'goroutines', the basic unit of concurrency in  Go, which let us manage more than one website check request

- 'anonymous functions', we used to start each of the concurrent processes that check websites

- channels, to help organize and control the communication between the different processes, allowing us to avoid RACE CONDITION bug

- the race detector which helped us debug problems with concurrent code
```bash
go test -race ./10-concurrency
```