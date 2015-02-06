# Business (logic)

This whole system is about bringing up lights for different reasons. And then
combining those lights in a smart way

## Systems
The types of systems you might want to bring up include:

1. bringing one dimmer up for a dimmer check
	1. `Dimmer 1`
	2. `Dimmer 3-10 @ 0`
	3. `Dimmer 9, 11 @ 50%`
2. Bring a "smart selection" of lights up
	1. `USL fronts`
	2. `USL, USR front warms @ 50%`
	3. `light special @ 20%`
	4. `backlight scrollers @ green`
3. Bring up "scene" of lights, that was previously saved
	1. `blueout`
	2. `Blackout`
	3. `indoor warm`
	4. `indoor cool`
5. Cues
	1. `Cue \#1` [changes from current cue to cue 1 over certain time]
	2. `Cue "indoor scene" (\#1)` [alias for cue 1]
	3. `Cue "blackout before dance"` [alias for 1 - cue of dance]

If we look at these examples, we can make som abstractions. First, let's say that
every system has a level, which is hidden if it is full. This is a helpful
assumption to make, because it reduces differences between systems and makes
sense practically.

You also notice that some systems are actually made up of many subsystems.
This includes both `dimmer 9, 11` and `indoor warm`.

So we have established two things
1. All systems have a level
2. Some systems are built from subsystems 

## Levels
You will also notice that while most of the levels are percentages, one is "green".
One option is to actually store this as a level, and just show the color,
but I think it makes more sense to store the color and just look up what level
it corresponds to each time we want to render it. This is helpful, because if
we change the scroller particulars later, we just have to update the color once
and all instances will change.

So we need a way to map certain colors to certain levels, for a certain "profile".
For example, on one scroller green could be 26% while on another it could be 30%.

So now we need anther system mapping color profiles to percents. And then
one more mapping dimmers to color profiles.


## Combining Systems
if you have two different systems up at once, there needs to be a way to combine
the. For example, if you have "USL fronts" and "dimmer 1 @ 50%", which 
takes precedence?

We have opted to be explicit and use an order to decide. Later systems that
are added will have priority.

So added the fronts then the dimmer will result in a list like this:

	Dimmer 1 @ 50%
	USL fronts

Since the dimmer system is above the front light system it will take priority.


### Modifying Existing Levels
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
+=