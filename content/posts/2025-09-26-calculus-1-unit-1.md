---
layout:     post
title:      "Calculus 1 - Unit 1"
subtitle:   "A quick pre-calculus refresh"
description: "Lines, functions, graphs, and all the other stuff you need before calculus makes sense. No, you can't skip this."
date:    2025-09-26T12:00:00-04:00
author:     "O.J. Wilcox"
image: "/img/posts/bg-calculus-1-unit-1.png"
tags: [ "Math", "Calculus 1" ]
URL: "/2025/09/26/calculus-1-unit-1-a-quick-pre-calculus-refresh/"
categories: [ Calculus ]
draft: false
---

## Calculus 1: Unit 1
_(Calculus 1 - Precalculus refresh)_

### Why are we doing this?

Precalculus is the toolbox you'll reach for constantly in Calculus I. Lines, functions, graphs, composition, transformations. All of it shows up later.

If you're rusty on this stuff, now's the time to fix that. If you're solid, skim through and make sure there are no weird gaps. Either way, this won't take long.

---

### 1) Lines (and distance)

A line is the simplest possible trend. The **slope** tells you how fast y changes when x changes a little.

Steep line? Big changes. Flat line? Nothing much happening. Vertical line? That's not a function (you'd divide by zero, and math gets angry).

**Distance** is just "how far apart are two points." On a number line, count the steps. On a grid, think of a little right triangle between the points and use the Pythagorean theorem. You've done this before.

Road grade signs like "6% grade" are slope. That's literally what they're telling you: for every 100 feet forward, you go up (or down) 6 feet. Truckers care about this a lot. So do cyclists. So does anyone who's ever underestimated a hill.

---

### 2) Circles

A circle is "all points the same distance from a center." That's it. The center is where it lives, the radius is how big it is.

Think of a compass drawing a circle around a pushpin. The pushpin is the center. The distance to the pencil tip is the radius.

You'll see equations that describe circles. They'll mention a center point and a radius. Recognize the pattern and you can read off both instantly.

---

### 3) Slope between two points (average change)

Pick any two points on a graph. The slope between them is the **average rate of change** over that span.

Town population goes from 3,192 to 5,362 in 10 years. That's about 217 people per year on average. The town didn't add exactly 217 people every single year, but that's the average trend.

This is useful for a first cut. It can hide spikes and dips, but it tells you the general direction. "Are we going up or down, and roughly how fast?"

**Types to recognize:** positive slope (rising), negative slope (falling), zero slope (flat), undefined slope (vertical, which again, not a function).

---

### 4) Equations of lines

Two popular ways to write a line:

**Slope-intercept:** y = (slope) · x + (where it crosses the y-axis)

**Point-slope:** Start from a known point, move by the slope.

Both describe the same line. Use whichever one matches the information you have.

Phone plans work this way. Base fee plus per-GB cost. That's a line. Cloud computing costs work this way too. Base cost plus per-instance pricing. Once you know a point and a slope, you can write the whole equation and predict costs at any scale.

---

### 5) What is a function?

A **function** is a rule that gives exactly one output for each input.

Put 16 into "square root," you get 4. Always 4. Not sometimes 4, sometimes -4. (Technically -4 squared is also 16, but when we say "the square root function," we mean the positive one. Math has conventions. Roll with it.)

The key test: one input, one output. If the same input could give you multiple outputs, it's not a function.

Why does this matter? Because functions are predictable. You can build on them, combine them, analyze them. Unpredictable things are harder to work with.

---

### 6) Evaluate functions (use the rule)

Function notation looks scarier than it is. **f(x)** just means "apply the rule called f to the input x."

If f(x) = x² + 3, then f(5) = 25 + 3 = 28. You're just plugging in values.

You'll see functions named all kinds of things. A(L, W) for area from length and width. D(t) for distance after time t. E(t) for error count at minute t. The letter doesn't matter. The rule matters.

---

### 7) The difference quotient

Fancy name, simple job: "How much did the output change when I nudged the input by a little bit?"

It's the same idea as slope, just written with a tiny step size. Change in output divided by change in input.

Oven temp goes from 325° to 350° in 10 minutes. That's +2.5° per minute on average. That's a difference quotient.

On tests, you'll simplify expressions involving this. Conceptually, just remember: it's measuring sensitivity. "If I change the input a little, how much does the output change?"

---

### 8) Functions from tables and graphs

A function doesn't have to be a formula. It can be a table of values or a graph.

A spreadsheet of "Year → Revenue" is a function. A Grafana dashboard showing "time → CPU usage" is a function. You don't need an equation to have a function.

**Two skills to practice:**

1. Build a small table of values and sketch the graph.
2. Use the **vertical line test**: if a vertical line hits the curve twice at the same x-value, it's not a function. (Same input, two outputs. Fails the rule.)

---

### 9) Reading graphs (carefully)

Graphs tell two kinds of stories:

**Specific points:** "At 9 AM, the temperature was about 68°F."

**Shape and trends:** Rising, falling, leveling off, sudden drops, peaks.

Don't just read the number. Read the shape. A smooth rise and a sudden cliff tell very different stories, even if they end at the same place.

Think of a diver off a board. Height starts high, drops, hits water, maybe comes back up a bit. The shape of that graph tells you about acceleration, impact, everything. The individual numbers are less interesting than the overall pattern.

---

### 10) Piecewise functions

"If x is in this range, use rule A. Otherwise, use rule B."

Overtime pay works this way. $24/hr up to 40 hours, then $36/hr for anything beyond. Same input (hours worked), different rules depending on the range.

Tax brackets work this way too. You're not taxed at one flat rate. Different portions of your income hit different rates. Piecewise.

**Graphing tip:** Each piece lives on its own restricted domain. Use open or closed dots at the boundaries so the graph still passes the vertical line test. This is a common place to mess up on tests. Don't be that person.

---

### 11) Composition (functions feeding into functions)

Do g first, then feed its result into f. Written as **f(g(x))**, pronounced "f of g of x."

Time gives you the radius of a ripple spreading on water. Radius gives you the area of that ripple. Compose them and you go straight from time to area.

Why care? Because real systems are layered. The output of one thing becomes the input of the next. Composition is how you trace cause and effect through a chain.

Load affects queue length. Queue length affects latency. Compose them and you get "how does load affect latency?" in one step. That's useful.

---

### 12) Shifting and stretching graphs

Small edits to a formula cause predictable moves in the graph:

- Add a constant outside the function → shift up or down
- Add a constant inside (to x) → shift left or right (and yes, the direction is counterintuitive: +h inside shifts *left*)
- Multiply the output → stretch or compress vertically
- Negative sign → flip over the x-axis

Once you know the base shape of a function, you can sketch variations quickly without plotting a bunch of points.

Alarm clock example: if your "alertness over time" graph shifts 2 hours to the right, you woke up later. Same shape, different position.

---

### 13) Absolute value functions

|x| means "how far x is from zero." Always non-negative.

The basic graph is a V shape. That's it. V for value. (Okay, that's not why it's V-shaped, but it's a decent mnemonic.)

"5 miles east" and "5 miles west" are both distance 5 from home. Absolute value doesn't care about direction. Just magnitude.

**Twist you'll see:** y = |f(x)| takes any parts of the original graph that dip below the x-axis and flips them above it. Because absolute value can't be negative, so everything gets mirrored up.

---

### Quick checklist (test brain mode)

- **Lines:** recognize slope types, write a line fast given a point and slope
- **Distance:** steps apart on a number line, right triangle on a grid
- **Functions:** one input → one output, vertical line test
- **Evaluate:** plug in carefully, don't rush
- **Difference quotient:** average change over a small step
- **Tables & graphs:** make a table, sketch a graph, read points AND shape
- **Piecewise:** pick the right rule for the input's range, watch the boundary dots
- **Composition:** g first, then f, outputs become inputs
- **Shifts/stretches:** up/down/left/right, stretch/flip, sketch from the base shape
- **Absolute value:** V-shape, negative parts flip above the axis

---

### For the engineers

#### SRE examples

**Error slope guardrail:** "If error trend > X per minute for Y minutes → auto-rollback." That's just slope with a threshold.

**Piecewise autoscaling:** "If CPU < 60% → do nothing. If ≥ 60% for 5 min → add 1 pod. If ≥ 85% → add 2." Different rules for different ranges. Piecewise policy.

**Composition for sensitivity analysis:** Load → queue length → latency. Compose them to explain "how a small load change propagates to latency." Great for capacity planning conversations.

**Graph shape for incident review:** Don't just read the numbers. Read the shape. Smooth rise vs. sudden cliff tells you different things about root cause. "When did it actually get bad? When did we actually recover?" The shape answers that.

**Absolute value for anomaly detection:** |error_delta| treats up-spikes and down-spikes as "magnitude of change." Useful for anomaly scores when you care about deviation in either direction.
