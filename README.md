# Lighting
*Sorry I couldn't think of a better name*.

## Goals
Talk into phone. "Bring dimmer 10 up to full" -> Dimmer 10 comes up at full.

## Initial Design
* Create a webserver with [Flask](http://flask.pocoo.org/), that provides a form
  to speak into.
* Use [Annyang](https://www.talater.com/annyang/) to recognize speach input.
* Send input back to Flask, figure out what it should do.
* Send DMX output to Arduino.
