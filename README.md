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
We should be talking in terms of groups, not channels or submasters. I don't care what channel are up at what percent, I just want to know how my scene is lit. That means instead of the tradional layout of channel number and levels, instead we need something like this:

```
USL Front Warm: 5%
Front Warm: 10%
Cue 10: 10%
```

The system should only maintain/show/record the most abstract information to represent any certain state. So instead of rememebering dimmer numbers or channel numbers for a cue, it should instead use these type of groups.


## Design
* Use [Go](https://golang.org/).
  *  It is fast.
  *  It is fun.
  *  I want to learn it.
*  Use [Gobot](http://gobot.io/).
  * Guys working on it are awesome.
  * Give back to OS community.
  * Tie our project into larger picture.
* Use [Annyang](https://www.talater.com/annyang/) to recognize speach input.
* Have server running Go code. That has a USB to DMX out. Also has an HTTP API server that talks to web apps/iPhone apps (on local network).
