# Planet Scale Paste Bin and Link Sharing (as a Service)

## Members
Jachob Dolak  : Infrastructure Engineering   
Will Erdman 	: Core Development  
Bobby Rizzo   :  Integration Development  

## Proposal
We will create a URL shortener and paste bin web service. This will allow users to send API requests to our server to create short and shareable URLS to either a website or text document. 
  

This program will be written in Go and make use of Go’s concurrency constructs. It will provide a RESTful API, be hosted on AWS, and source with build instructions will be available on github.  

In light of the project’s theme and an attempt to make this more interesting, we will also try to make this service extremely scalable, and distributed. This could possibly be multi-cloud, containerized, using kubernetes, with in-memory databases, with horizontal scaling, and edge computing. (and more buzzwords to come)

## Building

This project and is written in go and while a binary can be compiled locally, due to this project's reliance on Redis and to be consistant with deployment, it is recommended to run build it as a container.

First, clone the repo and `cd` into it.
```sh
git clone https://github.com/jdolak/Planet-Scale-Link-Shortener.git
cd Planet-Scale-Link-Shortener
```

Setup your dev environment. This downloads the go modules onto into a `./libs` directory so that when successively build containers we do not need to download new modules every time.
```sh
make init
```

Build and run the container.
```
make
```

