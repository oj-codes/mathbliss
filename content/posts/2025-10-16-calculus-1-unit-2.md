---
layout:     post
title:      "Limits & continuity"
subtitle:   "(getting 'instant' without zooming forever)"
description: "Your speedometer already knows calculus. Here's how limits actually work."
date:    2025-10-17T12:00:00-04:00
author:     "O.J. Wilcox"
image: "/img/posts/bg-calculus-1-unit-2.png"
tags: [ "Math", "Calculus 1" ]
URL: "/2025/10/17/calculus-1-unit-2/"
categories: [ Calculus ]
draft: false
---

## Limits & continuity (getting "instant" without zooming forever)
_(Calculus 1 - Unit 2)_

### The big idea

A limit is just "what value are we heading toward" as we get closer and closer to some point.

We're not asking "what's the value at this exact spot." We're asking "what does everything around it suggest it should be?"

### Instant vs. average: why limits matter

You already understand this intuitively.

Your car's speedometer shows how fast you're going right now. But if I told you "I drove 120 miles in 2 hours," that's your average: 60 mph. You weren't going exactly 60 the whole time. You sped up, slowed down, maybe stopped for gas.

Calculus is the math that lets us move from the average to the instant. And limits are how we do it. We ask "what happens to the average speed as we shrink the time window smaller and smaller and smaller, until we're basically asking about this exact moment?"

That's the heart of it. The rest is just technique.

### From secant to tangent: the visual version

Here's the picture that makes it click.

Draw a curve. Pick two points on it. Connect them with a straight line. That's a secant line, and its slope is your average rate of change between those points. Rise over run. Nothing fancy.

Now slide one of those points closer to the other. The secant line pivots. Keep sliding. The closer the points get, the more that secant line starts hugging the curve at a single spot.

When the points are "infinitely close" (this is where limits come in), you're left with a tangent line. A line that just kisses the curve at exactly one point. The slope of that tangent line is your instantaneous rate of change. The speed right now, not the average over some interval.

This is literally what your speedometer is calculating. Your car's computer is finding the slope of the tangent line to your position curve.

### Three ways to find a limit

So how do you actually figure out what value a limit approaches? You've got options.

**Look at the graph.** Trace the curve from the left and from the right. Where do they meet? That's your limit. This is the "eyeball it" method. Quick and useful for building intuition, but not always precise.

**Build a table.** Plug in values that get closer and closer to the point you care about. If f(x) keeps settling toward some number as x approaches 2, that number is your limit. Tedious but convincing.

**Do the algebra.** This is where you factor, simplify, rationalize, or use other tricks to eliminate whatever's making the function undefined at that point. Often the cleanest method when it works.

Most of calculus class is learning the algebra tricks. But understanding what a limit actually means? That's the graph and table intuition.

### One-sided limits: direction matters

Sometimes you need to ask "what value are we heading toward from the left?" and "from the right?" separately.

Think about the exact moment you flip a light switch. Right at the flip, is the light on or off? It's a jump. There's no smooth "in between." From the left (before the flip), the light is off. From the right (after the flip), it's on.

If the left and right limits don't match, the overall limit doesn't exist at that point. Not a failure. Just information.

You'll see this written as lim(x→2⁻) for "approaching 2 from the left" and lim(x→2⁺) for "approaching from the right." The little minus and plus signs are doing a lot of work there.

### Limit properties: breaking big problems into small ones

Once you know a few basic limits, you can combine them. Limits play nice with arithmetic.

The limit of a sum is the sum of the limits. Same for differences, products, and quotients (as long as you're not dividing by zero). You can pull constants out front. You can take limits of powers and roots.

This isn't deep philosophy. It's just saying "if I know where f(x) and g(x) are each heading, I know where f(x) + g(x) is heading." Common sense, formalized.

These properties let you break complicated limits into simple pieces, solve each piece, then put the answer back together.

### The squeeze theorem: when you're trapped, you're found

Sometimes you can't find a limit directly, but you can trap the function between two other functions that are easier to deal with.

If g(x) ≤ f(x) ≤ h(x), and both g(x) and h(x) approach the same limit L, then f(x) has no choice. It's sandwiched. It also approaches L.

Picture walking down a hallway that's narrowing. The walls on either side are converging to a single point. You're between them. Where are you going to end up?

This shows up more than you'd expect. Especially with trig functions that oscillate wildly but can be bounded.

### Continuity: no surprises

A function is "continuous" at a point if three things are true: the function is defined there, the limit exists there, and they're equal. In plain english: the value is what you'd expect from the surrounding values. No teleportation.

A function is "continuous everywhere" if this is true at every point. The graph is smooth with no jumps, holes, or vertical asymptotes.

Why care? Because continuous functions behave predictably. You can trust that "almost there" means "almost that value." If it's not continuous, there's a surprise hiding somewhere.

### Which functions are continuous?

Good news: most of the functions you'll work with are continuous wherever they're defined.

Polynomials are continuous everywhere. Always. No exceptions.

Rational functions (fractions with polynomials) are continuous everywhere except where the denominator is zero. Common sense.

Trig functions, exponentials, logarithms: all continuous on their domains. Sine doesn't suddenly jump. e^x doesn't have holes.

When you combine continuous functions with addition, subtraction, multiplication, division (no zero denominators), or composition, you get continuous functions.

The intuition: "well-behaved" functions stay well-behaved.

### Intermediate value theorem: no teleporting allowed

Here's a powerful consequence of continuity.

If you have a continuous function, and it equals 3 at some point and 7 at another point, then somewhere in between it must equal 4, 5, 6, and every other value from 3 to 7.

It's like driving from sea level to 5,000 feet elevation. If your path is continuous (no helicopters, no tunnels through the mountain), you hit every elevation in between.

This sounds obvious, but it's actually useful. It guarantees solutions exist. If f(a) is negative and f(b) is positive, and f is continuous, there must be some point between a and b where f equals zero. A root exists. You might not know exactly where, but you know it's there.

### When limits misbehave

Sometimes there's no clean answer. The function might blow up toward infinity, oscillate wildly without settling down, or approach different values depending on which direction you come from.

When that happens, we say the limit "doesn't exist." That's not failure. It's information. It tells you something interesting is happening at that point.

### The formal definition (for the curious)

There's a rigorous definition using Greek letters. Epsilon-delta. It says: for any tiny tolerance ε you give me around the limit value, I can find a tiny neighborhood δ around the input point such that all inputs within δ produce outputs within ε.

It's precise. It's how mathematicians make the handwavy "closer and closer" actually mean something. It's also where most people's eyes glaze over.

For this course, the intuition is what matters. You're asking "what value does this approach?" not "can I prove it with epsilon-delta." If you go on to real analysis, you'll live in that world. For now, know it exists and why it exists: to make limits rigorous instead of relying on "obviously it goes to 5."

---

### For the engineers

#### SRE examples
**Instant vs. average in monitoring:** A 5-minute average error rate might look fine while you're actually getting hammered this second. Alerts on rate-of-change (the derivative, which builds on limits) catch spikes faster than alerts on averages.

**Secant to tangent in practice:** When you calculate "errors per second over the last minute," that's a secant slope. When your monitoring system shows the instantaneous rate, it's approximating a tangent slope. The smaller your time window, the closer you get to the true instantaneous rate.

**Continuity in metrics:** SLO burn rate should move smoothly as load changes. If you see sudden jumps, that's often a broken counter, a deploy, or a data gap. Not real user impact. Discontinuities are worth investigating.

**Intermediate Value Theorem for capacity:** If your system handles 1,000 RPS fine and falls over at 5,000 RPS, there's some threshold in between where things start degrading. Continuous behavior means you can't jump from "fine" to "dead" without passing through "struggling."

**Squeeze theorem intuition:** If you know your p50 latency is between 10ms and 50ms, and both bounds are converging to 30ms as load stabilizes, your p50 is heading to 30ms. Bounding gives you confidence even when exact measurement is noisy.

**Limits that misbehave:** During a canary rollout, if the metric near deploy time is chaotic and won't settle, don't try to fit a clean model yet. Stabilize first, then analyze.

</details>
