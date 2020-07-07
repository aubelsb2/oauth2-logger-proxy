# oauth2 proxy logger 

This is a set of simple proxy loggers for https://github.com/go-oauth2/oauth2 as it can be hard to get a good idea what's
happening. Usage is rather simple, it wraps the component you actually use and logs the input and output. It's a bit 
verbose as it's meant to be used for debugging purposes.

Manager:

        	srv := server.NewDefaultServer(oauth2_logger_proxy.ProxyManager{manager})


ClientStore: 

        	manager.MapClientStorage(oauth2_logger_proxy.ProxyClientStore{dynamo.NewClientStore(c)})


Random injector functions:


            srv.SetClientInfoHandler(oauth2_logger_proxy.ProxyClientFormHandler(ClientFormHandler))


Happy to accept PRs.