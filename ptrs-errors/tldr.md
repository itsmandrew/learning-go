
## Pointers
- Go copies values when pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the
the thing you want to change

- The fact that Go takes a copy of values if usefful a lot of the time but sometimes you won't want your system to make a copy of something. Examples include referencing very large data structures, db connection pools, etc


## nil
- Pointers can be nil

- When a function returns a ptr to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help here

- Useful for when you want to describe a value that could be missing


## Errors
- Errors are the way to signiffy failure when calling function/method

- By listening to ours tests we conclude tthat checking for a string in an error would result in a flaky test.