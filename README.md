# Planet Scale Paste Bin and Link Sharing Network (as a Service)

## Members
Jachob Dolak  : Infrastructure Engineering   
Will Erdman 	: Core Development  
Bobby Rizzo   :  Integration Development  

## Proposal
We will create a URL shortener and paste bin web service. This will allow users to send API requests to our server to create short and shareable URLS to either a website or text document. 
  

This program will be written in Go and make use of Go’s concurrency constructs. It will provide a RESTful API, be hosted on AWS, and source with build instructions will be available on github.  

In light of the project’s theme and an attempt to make this more interesting, we will also try to make this service extremely scalable, and distributed. This could possibly be multi-cloud, containerized, using kubernetes, with in-memory databases, with horizontal scaling, and edge computing. (and more buzzwords to come)
