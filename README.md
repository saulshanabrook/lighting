[![Build Status](https://travis-ci.org/saulshanabrook/lighting.svg?branch=master)](https://travis-ci.org/saulshanabrook/lighting)
# Lighting
*Sorry I couldn't think of a better name*.

**More productive and modern stage lighting control system.**

## User Stories

### Talk into phone.
* "Bring dimmer 10 up to full"
* "Record dimmer 10 as USL front warm"
* "Bring the fron cools to full"
* "What are the USL fron warms at?"

### Cue/Scene Design
We should be talking in terms of groups, not channels or submasters. I don't care what channel are up at what percent, I just want to know how my scene is lit. That means instead of the traditional layout of channel number and levels, instead we need something like this:

```
USL Front Warm: 5%
Front Warm: 10%
Cue 10: 10%
```

The system should only maintain/show/record the most abstract information to represent any certain state. So instead of remembering dimmer numbers or channel numbers for a cue, it should instead use these type of groups.

## Tools
* Use [Go](https://golang.org/).
  *  It is fast.
  *  It is fun.
  *  I want to learn it.
  *  Our project looks like it is right in it's [sweet spot](http://www.reddit.com/r/golang/comments/29ua55/rust_or_go/ciojqcg)
    * fast
    * concurrent
    * not super low level, but not a CMS either
*  Use [Gobot](http://gobot.io/).
  * Guys working on it are awesome.
  * Give back to OS community.
  * Tie our project into larger picture.
* Use [Annyang](https://www.talater.com/annyang/) to recognize speech input.
* Have server running Go code. That has a USB to DMX out. Also has an HTTP API server that talks to web apps/iPhone apps (on local network).
* Maybe use [Gopherjs](http://www.gopherjs.org/) for frontend as well. Would remove a lot of potential code duplication
  * need something like meteorjs for go, so that state can be updated between backend and frontend

## Implementation
This project will actually requie many smaller subprojects to make it up.

* Go Server
    * go <-> dmx
        * provides a way to specify a USB device and then set dimmer levels and send
        * should also provide a way to create a mocked out dmx, for testing
    * lighting
        * this is the core logic
        * holds state of current board
        * understands cues, groups, patch
        * will save state to machine, but this is second priority
    * API
        * uses lighting library to create a HTTP api to serve up interface
        * just basically wraps it
* Client (either web or native mobile)
    * shows current stack
    * shows patch
    * shows cues
    * provides text based interface
    * provides audio dictation interface

