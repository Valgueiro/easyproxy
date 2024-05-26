# easyproxy

I (@valgueiro) had a really hard time while searching for proxy frameworks to find an easy to setup tool. To be honest, I just needed a tool that would allow me to setup allow/block lists of URLs and accept basic authentication, but I had to study a bunch to make it work as I expected. During my investigations using the proxy that I've created, I could also see that there is a bunch of libraries that says that support proxies, but that always had new issues complaining about that.

With this in mind, and also to improve my golang skills, I decided to create easyproxy, a CLI tool that can help you spin up a simple but effective forward proxy. It can help in analysis to understand if your infrastructure is ready to be behind a proxy,  and can even be used by other frameworks in their integration tests to make sure that their proxy configurations are working as expected. 

## TODO list
* [x] Create a cli
* [ ] Read about the proxy protocol to understand what needs to be 
* [ ] Start accepting a simple HTTP proxy request
* [ ] Enable setting allowed hosts
* [ ] Add the possibility to add basic auth
* [ ] Create a run command that creates a template with the options and spins up an envoy
* [ ] (?) Try to create a test with it to check if it can be used properly and think about this requirement 
