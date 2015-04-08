<!-- TOC depth:6 withLinks:1 updateOnSave:1 -->
- [Goals](#goals)
- [Core Principles](#core-principles)
	- [Everything is a system](#everything-is-a-system)
		- [Filters](#filters)
			- [Parsing Filters](#parsing-filters)
	- [Stack Based Precedence](#stack-based-precedence)
		- [Examples](#examples)
		- [Reasoning](#reasoning)
	- [Further ideas](#further-ideas)
		- [Natural Language Processing](#natural-language-processing)
		- [UI](#ui)
		- [Architecture](#architecture)
		- [Attributes of Instruments](#attributes-of-instruments)
			- [Non Power Levels](#non-power-levels)
	- [Modifying Existing Levels](#modifying-existing-levels)
	- [Cues](#cues)
- [FAQ](#faq)
<!-- /TOC -->

# Goals

* Simplicity
    * UI should be intuitive. If you know lighting, you should be able to walk
      up and start creating.
    * The data structures should be easy to conceptualize and as normalized as
      possible. What I mean by this is if a cue is really just a combination of
      a state + some timing + the number it should be stored and represented
      that way. The same would apply for a submaster/group, which is also
      just a state + a label. Why on an ETC (any model) are these represented
      so differently?
* Does what we need
* Extensible (for anyone)

# Core Principles
Lingo:

* **system**: A mapping of lights to levels. The most simple example would be
  a single dimmer at a level, akin to a "dimmer check". In terms of an ETC Eos
  Palettes, Presets, Cues, Channels, and Groups would all have a component of
  them that was a **system**, even if they stored other information as well.

## Everything is a system
Types of systems:

* **dimmer**: A single dimmer at a level.
* **filter**: Uses your patch to dynamically retrieve lights matching certain
  attributes. For example, the filter `US Front W @ 25%` will search the patch
  for all dimmers matching `{type: 'front', position_vertical: 'up', color: 'warm'}`
  and set them to 35%.
* **group**: A custom named set of subsystems
* **cue**: same as a group, but with extra timing information and ordered with
  other cues.

If you want something more concrete, this is a Go implementation:

```go
type Dimmer int
type Level float // or maybe uint8?

type System interface {
    DimmersAt() map[Dimmer]Level
}
```

### Filters
One of the harder systems to figure will be the filters. They represent the
primary way you will build your looks. Right now I am looking at it as if
each dimmer has certain attributes, and if we can just write down each
of the attributes, then we can query based on them.

This works great for a few things, like whether the light is a Front or Back
light. I think the best label for that would be **use**.

Then there is the color
of the lights. I don't mean the actual gel number, but just how you would talk
about the light, regarding it's color, like *warm*, *cool*, *neutral*, etc.
This **color** attribute seems to make sense, and I don't see much of a problem
storing this. However, we also need to think about how this will be displayed.
For example, you would might/probably wanna see "USL Front W", where "warm"
is abbreviated.

Another big one is some type of focus position. Like "USL" or "US". As long
as you don't have more than five zones, we can parse these names. It is really
two attributes here, the US-DS position and the SL-SR position. Is there a name
for those two axis? Width and hieght? Horizontal and vertical? I think the best
thing to do is to store these separately, sort of like storing the x and y
coordinate. So they would be called something like **position_vertical** and
**position_horizontal**. However, this assumes some very normalized positions.
What if you have a special on the "US Couch"? Well I would argue that we need
a **position** keyword as well, that would hold "couch". I am not sure if 
**position** is the right word for any of these, maybe **dest_position** or
**focus_position** is more appropriate. 

This all makes sense in an ideal world, but shows are not going to be like this
all the time. We need room for flexibility. So obviously these keywords shouldn't
be fixed. The user needs some way of saying "Hey I actually have my own keywords".
What would be great is if all this logic could be edited *by* the user. It might
end up being too complicated for that, but that would be ideal. Then they could
decide how to interpret positions and all that. 


#### Parsing Filters
We also need to figure out how we interpret use input that has a mixing of
attributes, that are all unlabeled. For example, if the user types in
"USL Front warm @ full", how do we know what attributes come out of this query?
How do we know "warm" is a `color` and "front" is the `use`? Well we have a
couple of options (that I see):

1. Since we know the patch already, we can go through inputted word and see
   if it exists in any of the patched attributes. So when tell it "Front",
   it will look through all the values of every attribute for all lights and
   see that "Front" appears only as a `use`, so will categorize it like that.
   Problems would come up if there were the same values present for multiple
   attributes. Also, could be a bit slow. It's nice that it allows flexible
   ordering, however. And the speed could be remedied by pre-computing
   that mapping of attributes values to attribute keys.
2. We could allow the user to specify a consistent ordering that they use.
   For example, this could be `position use color`. So then if they wrote
   "USL front warm", it would know which input words match with which attributes.
   However, a problem with this design is that it is stricter, so is worse
   for the user to use. Also, you are going to have optional keywords, so
   how does it know which ones you input?
3. Something for advanced use natural language machine learning.


## Stack Based Precedence
To combine systems you order them in a stack and merge them all together,
where the top of the stack takes precedence. Later systems that
are added will have priority.

### Examples
For example, let's say your stack looked like this:

```
* Dimmer 10 @ 25%
* Dimmer 10 @ 50%
* Dimmer 3 @ 50%
```

Here the final state would be dimmer 10 @ 25% and dimmer 3 @ 50%.

In a more realistic example:

```
* Front USC W @ 30%
* Front W @ 78%
* Cue 1
```

Cue 1 is made from:

```
* Front W @ 50%
* Front C @ 25%
* Top @ 30%
```

So all together the final state will be equivalent to top @ 30%, front C @ 25%,
all front warm besides the USC @ 78%, and the USC front warm @ 30%.

### Reasoning
1. It is easy for the operator to answer *why* levels are at where they are.
2. The stack based approach is unambiguous and uniform, eliminating the need
   for capturing or different ordering approaches.

## Further ideas
### Natural Language Processing
[Annyang](https://www.talater.com/annyang/) provides a way to take human typed
text + speech input and get actions out of it. So you can train it with things
like "When I say 'bring up dimmer 10' I want you to do `{system: dimmmer, number: 10, level: 100}`".
This means we can just add on language processing on top of the standard way of bringing up
dimmers

* "Bring dimmer 10 up to full"
* "Record dimmer 10 as USL front warm"
* "Bring the front cools to full"
* "What are the USL front warms at?"


### UI
I know it will have to contain the current state, as a list of systems that are
up. There should also be a way to drill down into a system, maybe by clicking on it.
Then you can see what it is made up of.


Also there are a  multiple ways of viewing many of the systems, the human
readable notation ("US Front") and the actual representation `{position: "up", use: "front"}`.
I think the human readable should be displayed by default and can be parsed into
the actual representation, but the more precise notation should be easily accessible
if you hover over or click or something.

If you hover over any percentage, or click on it, you should be able to
modify it with the scroll wheel or keyboard.

### Architecture
No lighting systems that I know of use HTTP and web technologies to create their
interfaces. All that I know of render the UI locally on whatever machine
they are running on.

However, I think using HTTP does present some advantages:

1. Access on any device/cross platform
2. CSS is arguably most used layout system ever created. This means tons of
   libraries to build from, as well as plenty of examples.
3. The same goes for JS.
4. Server software can potentially be run on very cheap hardware, just serving
   up HTTP server (with language of choice).
5. Networking is incredibly advanced + cheap for HTTP, based on its reliance
   on TCP/IP. This means we get wireless connections for free.

The largest disadvantage is a possible duplication of logic, because of the
seperation of client and server. For example, we need to duplicate all state
handling on both ends, so that the user can see what is going on and edit it
and the server can render the resaulting levels to DMX.

Another disadvantage is unreliable performance. I think most of the variability
here is from network level latency, not the unpredictability of the application.
Well, we have control over the performance of the client and server code, but
less control over the network, which is worrying. I have not done any testing
yet of performance characteristics under different network conditions, but I
assume tha tusing a dedicated network for lighting is the simplest way to
alleviate fears of other devices creating latency.

For a look at some of the current options, check out the [./architecture.md](./architecture.md)
file.

### Attributes of Instruments
Major complications arise when considering non light emitting outputs, for example
moving lights, color scrollers, mirrors, and gobo spinnners.

While these would all function fine under the above proposed system the workflow
could be a bit ugly.

For example, let's say you have a bunch of backlight with color scrollers on
them.

Let's say your patch looks like this:

```
{
    1: {use: back, pos: left, attribute: power}
    2: {use: back, pos: center, attribute: power}
    3: {use: back, pos: right, attribute: power}
    101: {use: back, pos: left, attribute: color}
    102: {use: back, pos: center, attribute: color}
    102: {use: back, pos: right, attribute: color}
}
```

Then one way to bring up the lights would be to address them seperately

```
* Back Color @ 25%
* Back @ 50%
```

This will translate to two filter systems. We can assume that when you don't
specify an `atribute` you mean `power`.

While this works, there are other ways of arranging this which might be which
might be preferable. Another option is to some sort of **combination filter**
which can select multiple attributes from the same group and give them levels,
like this:

```
* Back @ 50%, Color @ 50%
```

A sample implementation could look something like this:

```go
import (
    "github.com/imdario/mergo"
)

type Tag string
type Value string
type Query map[Tag]Value

type FilterSystem struct {
    Query Query
    Level Level
}

func (fs *FilterSystem) DimmersAt() {
    // assuming this is written, to get the dimmers for that filter
    // returning them at that level
}

type AttributeSystem struct {
    SubQuery Query
    AttributeLevels map[Value]Level // like {"power": .10, "color": .5}
}

// This will iterate through the different attributes and make queries for each,
// and then merge them all together
func (as *AttributeSystem) DimmersAt() {

    dm = make(DimmerMap)
    for v, l := range as.AttributeLevels {
        SubQuery["attribute"] = v
        Mergo.Merge(%dm, FilterSystem{SubQuery, l})
    }
    return dm
}
```

In our case, we would create ours like:

```go
AttributeSystem{
    Query{"use": "back"},
    map[Value]Level{"power": .25, "color": .5}
}
```


A simpler (and probably better) method would just be to create a group from
those two levels, of the power and the color, and list that group in the state.
Then in the UI, instead of just displaying like "Group 1", just look through
the group to find commonalities like that to display that text. This would have
the advantage of not requiring another data structure and pushing more of the
complexity for rendering down to the client. It also turns this problem into
just representation instead of storage, which increases data normalization.

#### Non Power Levels
Another question is how to properly implement things like color and mirror
control. I think the UI to the user is pretty clear in some cases.


```
* Back Color @ Green # alias for 34%
* US Mirror Control @ (13%, 20%)
* DS Mirror Control @ Center Stage # this would be an alias for a coordinate
* Side Strip Color @ (10%, 30%, 50%) # these would be RGB levels, hovering over should allow color wheel selection
```

The first example requires the system to know that for that particular color
scroller, "green" is 34%. So this in turn requires mapping certain lights
to what physical instrument they are. For example, we know, in this example,
that dimmers 101-103 are color scrollers with a green color at level 34%.

We could just have a map of dimmer numbers to their profile, but that seems a bit
messy. For example, it isn't really the *dimmer number* that is that instrument
but the label. Even if the the dimmer number of that scroller changes, it will
still be called the Back Color. So let's make up a new interface, 

Profile/Instrument Type
OK lets do that:

**TODO**: Figure out exactly how this interface will work, how mapping works
          between this instrument patch and the other patch

## Modifying Existing Levels
Often times we don't care about the absolute level of a light, we just want it
to be a bit more than it already is. One way to do this is just to capture the
light higher than it is currently. However, I think it would be more powerful
to be able to create relative levels.

For examples if we have

    Warm Indoor 

And we want to make it a bit cooler, it would be nice if we could do

    Front Cools +@ 10%
    Warm Indoor 
    
Which will increase whatever level the is previously. I think this is preferable
for a few reasons. First it is more obvious what our intention was. Coming back
to look at this cue later, it is obvious what we wanted to do. If we had set
it to an absolute level instead, we would have to do some more digging, before
we knew what we were trying to do. Also, if we have and modify `Warm Indoor`
later, it will still be a bit cooler in this scene.

If we include `+@` it would also make sense to have `-@`, `*@`, and `/@`. 

So we have to make sure of a few things. First that these only work with the
numeric type of levels. It doesn't make sense (so shouldn't be allowed) to
+= a color...

**TODO**: Figure out if we really need this. And if so, when it is useful.

## Cues
**TODO**: Try to understand best way of cueing, and different timings needed.
          Look at existing systems out there.


# FAQ
Q: *But the Ios board already has this type of abstraction with Presets and Palletts.*

A: While they do allow this same type of abstraction, it is much harder to setup and 
   display. The fundamental difference is that they have these abstractions coinciding
   with channel information. In this system, there are no channels. Also the difference
   of the stack based system means that it is clearer what the current state
   is made up of.

Q: *Touchscreens suck.*

A: Yep I agree. Manual sliders are nicer for manual control and channel setting.
   It is a question of trade offs. The benefit of this system is not that it is
   digital, but that by starting from scratch we can implement different data
   schemes. I would love to add support for hardware, or even dedicated hardware,
   but that is a secondary priority right now.

Q: *This will be less reliable, because it is not on a dedicated machine.*

A: So this is a complicated one to answer, because it's reliability is mainly
   based on what the codebase looks like and how much use it has gotten. While
   it will likely be less popular than other systems, there is nothing fundamental
   stopping it from being just as reliable. I see this as running on a stripped
   down linux box that only runs this app. The reliability of the system would
   be comparable to the reliable of any linux based server.

Q: *You can't please everyone, this doesn't address the concerns of x.*

A: Fundamentally I am writing this to scratch *my* itch not yours. I am totally
   open to other perspectives and wants, as long as they don't fundamentally
   alter the premise that allows me to get what I want out of it and don't
   add unneeded complexity.

