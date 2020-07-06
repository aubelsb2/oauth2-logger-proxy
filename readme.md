# oauth2 proxy logger 

This is a set of simple proxy loggers for https://github.com/go-oauth2/oauth2 as it can be hard to get a good idea what's
happening. Usage is rather simple, it wraps the component you actually use and logs the input and output. It's a bit 
verbose as it's meant to be used for debugging purposes.

Manager:

        	srv := server.NewDefaultServer(ProxyManager{manager})


ClientStore: 

        	manager.MapClientStorage(ProxyClientStore{dynamo.NewClientStore(c)})


Happy to accept PRs.