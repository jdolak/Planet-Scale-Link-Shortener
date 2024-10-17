# Planet Scale Paste Bin and Link Sharing (as a Service)

## Members
Jachob Dolak  : Infrastructure & Core Engineering   
Will Erdman 	: Core Development  
Bobby Rizzo   :  Integration Development 

## Goal

Planet Scale Paste Bin and Link Sharing (as a Service) is a URL shortener and paste bin web service.  
It provides both a Web GUI and RESTful API for creating links.

Features include the ability to create custom link expirations and gather analytics like view count and creation IP.  

Our goal is to make sharing data as easy as possible. A simple, no fuss, way to allow others to download your text or find your website.

## Using Go

We takes advantage of the unique Concurrency features of Go to implement our link expiration. 
Utilizing a Goroutine, links are periodically scanned in the DB. This is concurrent to the normal function of our web server. We also capitalize on Go's features for handling data across networks. We are especially fond of marshalling and unmarshalling json into predefined structs.

## Building

### Prerequisites

Docker is required to run this project. Instructions for how to download it can be found [here](https://docs.docker.com/get-started/get-docker/).  

Installing Go is recommend for development, but not required for deployment. This is because the container contains everything required to build the software. If you would like to install Go, it can be found [here](https://go.dev/doc/install).

### Instructions

This project and is written in go and while a binary can be compiled locally, to be consistant with deployment, it is recommended to run build it as a container.

First, clone the repo and `cd` into it.
```sh
git clone https://github.com/jdolak/Planet-Scale-Link-Shortener.git
cd Planet-Scale-Link-Shortener
```

Setup your dev environment. This downloads the go modules onto into a `./libs` directory so that successive containers builds do not need to download new modules every time. 
This only needs to be run once.
```sh
make init
```

Build and run the container. Run this every time a change is made.
```sh
make
```

## Proposal
We will create a URL shortener and paste bin web service. This will allow users to send API requests to our server to create short and shareable URLS to either a website or text document. 
  

This program will be written in Go and make use of Go’s concurrency constructs. It will provide a RESTful API, be hosted on AWS, and source with build instructions will be available on github.  

In light of the project’s theme and an attempt to make this more interesting, we will also try to make this service extremely scalable, and distributed. This could possibly be multi-cloud, containerized, using kubernetes, with in-memory databases, with horizontal scaling, and edge computing. (and more buzzwords to come)

