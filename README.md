# Rate-limiter-example

Rate-limiter-example is a simple web app to demonstrate the rate limiter functionality

Project is modularised to have separate packages for api and rate-limiter

* Whole project uses **dependency injection patterns** to make it easily testable and extendable. any implementation that binds the dependency interface can be injected
* For Simplicity, we have taken **second** as the smallest unit of time. As extension we can make it milli second or microsecond, which is a small refactoring
* This app use **Factory patterns** to create the dependencies.
* The project has lot of comments explaining the structure, functions and decisions. It also has lot of unimplemented functions :p

### Rate-limiter
* rate-limiter provides a contract that different implementation strategies can implement
* Currently I have only implemented the **sliding window strategy**. But we add different strategies behind the factory. strategies like fixed window, leaky bucket etc can be used provided it implements the RateLimiter interface
* Storage is abstracted into separate interface where we can easily switch the **different data structures**  


#### setup

* please follow https://golang.org/doc/install for instruction to setup go

* Once you have setup. clone the repository under cd $HOME/go/src/github.com/manoharkshetty/ directory.

* repo url: https://github.com/manoharkshetty/rate-limiter-example


#### To Run the project:

```
cd $HOME/go/src/github.com/manoharkshetty/rate-limiter-example
go build 
./rate-limiter-example
```

or 

Just click that cute green button if you have an IDE (goland) :)



