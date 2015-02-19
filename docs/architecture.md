# Architecture

Below are the options for different technologies at all layers of the stack.
Before deciding, we need to narrow down our speed requirement, cause this
will help us eliminate some.

I heard from my professor that human's feel like ~200ms is "instant",
so anything slower than this is too slow. I want to aim for < 50ms for our test,
which gives the backend 150ms to do other things when it get's more complicated.

- [Backend Language](#)
	- [Go](#)
	- [Python](#)
- [Backend Web Framework (if using Python)](#)
	- [Falcon](#)
	- [Flask](#)
	- [Django](#)
- [Communication Protocal](#)
	- [HTTP](#)
	- [Websockets](#)
- [Frontend System](#)
	- [Web](#)
	- [Mobile](#)
- [Frontend Language (if creating app for web)](#)
	- [Javascript](#)
	- [GopherJS (if using Go for backend)](#)
- [Frontend Web Framework (if using JS for frontend)](#)
	- [AngularJS (1.x)](#)
	- [AngularJS (2.x)](#)
	- [polymer](#)
	- [React (or other flux) + friends](#)

## Backend Language
### Go
*Pros*
* Fast
* Type safe
* simple language design
* exciting new tooling
* built in concurrency support
* easy to deploy (one static binary)

*Cons*
* Greater learning curve
* sometimes forced design decisions might lead to more verbose code
* less library support (serial)

### Python
*Pros*
* Very familiar with language
* Great library support
* PyPy is rather fast
* Fast to prototype ideas
* mature web frameworks

*Cons*
* Could be too slow response times
* not type safe -> requires more testing

## Backend Web Framework (if using Python)
### [Falcon](http://falconframework.org/)
*Pros*
* seems to be the fastest
* actively developed

*Cons*
* not as much plugin support

### [Flask](http://flask.pocoo.org/)
*Pros*
* relatively fast
* lots of tutorial/support

*Cons*
* not as fast as it could be

### [Django](https://www.djangoproject.com/)
*Pros*
* well designed layout of code
* lots of library support for creating API
* lots of testing support

*Cons*
* slow
* huge

### [uWSGI](http://uwsgi-docs.readthedocs.org/en/latest/WebSockets.html)
*Pros*
* websockets support
* bare bones
* so fast?

*cons*
* doesnt give you much for free

## Communication Protocal
[HTTP vs Websockts on stack overflow](http://stackoverflow.com/questions/14703627/websockets-protocol-vs-http)
### HTTP

*Pros*
* Supported by all web frameworks
* Easy to test
* simple model to understand

*Cons*
* potentially more latency

### Websockets

*Pros*
* lower latency

*Cons*
* not suported by all web frameworks
* harder to unit test


## Frontend System
### Web

*Pros*
* compatible on all devices (phones, computers)
* lots of open source tooling
* can do integration testing

*Cons*
* harder to make nice looking UI
* potentially slower

### Mobile

*Pros*
* better UI
* ??

*Cons*
* would have to develop for multiple devices
* harder to deploy



## Frontend Language (if creating app for web)
### Javascript

*Pros*
* faster
* less abstraction -> less leaky abstraction
* easy support for many libraries

*Cons*
* duplicate backend code
* not fun language

### [GopherJS](http://www.gopherjs.org/) (if using Go for backend)

*Pros*
* reduce duplication between client and server
* cleaner code

*Cons*
* potentially slower
* can't use JS client frameworks

## Frontend Web Framework (if using JS for frontend)
### [AngularJS](https://angularjs.org/) (1.x)

*Pros*
* lots of examples, libraries
* battle tested
* fast enough
* strong design model

*Cons*
* hard to know best way to do things/where to put code
* not hot/new (2.0 coming out soon)

### [AngularJS](https://angularjs.io/) (2.x)

*Pros*
* new
* good ideas behind

*Cons*
* beta

### [polymer](https://www.polymer-project.org/)

*Pros*
* built on future web standards
* different conceptual model (web components) that could potentially improve abstractions
* easy to create stand alone elements

*Cons*
* not in production much
* not as many tutorial on how to use it
* would run into more rough patches

### [React](http://facebook.github.io/react/) ([or other flux](https://github.com/voronianski/flux-comparison)) + [friends](https://tuxedojs.org/)

*Pros*
* New
* Facebook uses it (?)
* people are tired of angular and moving to this
* potentially faster rendering because of virtual DOm

*Cons*
* steeper learning curve (i dont know it)
* not as many examples
* too many libraries to choose from to combine it with
