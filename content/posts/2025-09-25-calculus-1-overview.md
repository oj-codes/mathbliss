---
layout:     post
title:      "Calculus 1 Overview"
subtitle:   ""
description: "The two big questions calculus answers, and why you already think this way (your phone definitely does)."
date:    2025-09-25T12:00:00-04:00
author:     "O.J. Wilcox"
image: "/img/posts/bg-calculus-1-overview.png"
tags: [ "Math", "Calculus 1" ]
URL: "/2025/09/25/calculus-1-overview/"
categories: [ Calculus ]
draft: false
---

## Calculus 1 Overview
_(Calculus 1 - Introduction)_

### What is calculus, really?

Calculus gets a reputation for being scary. It's the class people warn you about. "Oh, you're taking calc? Good luck." Thanks, super helpful.

Here's the thing though: calculus is just the toolkit for two questions.

1. How fast is something changing right now?
2. How much has piled up over time?

That's it. Seriously. The first question gives you **derivatives**. The second gives you **integrals**. Everything else is just technique for answering those two questions in different situations.

If you've ever checked how fast your phone battery is dying, you've thought about derivatives. If you've ever added up your total screen time for the week and felt judged by your own phone, you've thought about integrals. You already have the instincts. Calculus just gives you the precise tools.

---

### Unit 1: A quick pre-calculus refresh

Before we get to the good stuff, you need to be comfortable with functions and graphs. This unit is the warm-up.

**Lines and steepness:** A line has a slope. That's just "how much does the output change when the input changes a little?" Steep line means big changes. Flat line means not much is happening.

Your monthly rent is a flat line. (Unless your landlord is a monster.) Your heating bill in winter? That's a steep upward line as the temperature drops. You get it.

**Functions:** A function is a rule. Put in a number, get out a number. The graph shows you the shape of that relationship.

Think hours studied vs. test score. More hours generally means a higher score, but it's probably not a straight line. It might flatten out once you've studied "enough." We've all hit that point where more studying just means more staring at the same page.

**Combining functions:** Big things are made of smaller pieces. Understand the pieces, understand the whole.

Your commute time equals walking to the station plus waiting plus train ride plus walking to work. If one piece gets slower, the whole thing gets slower. This is not a profound insight, but calculus formalizes it so you can actually calculate stuff.

---

### Unit 2: Limits & continuity

A limit is "what value are we heading toward" as we get closer and closer to some point.

Why does this matter? Because it's how we move from "average over time" to "right now."

Your average speed on a road trip is easy. Miles divided by hours, done. But your speed at this exact instant? That requires limits. We shrink the time window smaller and smaller until we're basically asking about a single moment. It's a bit mind-bending the first time, but it clicks.

**Continuity** means no surprises. Small input changes produce small output changes. The graph is smooth.

Room temperature is continuous. It doesn't jump from 70° to 40° in a second. (If it does, something is very wrong with your house.) A light switch is not continuous. It's off, then it's on. No in-between. That discontinuity is actually useful information.

---

### Unit 3: Derivatives

A derivative answers "how fast is this changing right now?"

Not the average over the last hour. Not the trend over the last week. Right now, this instant.

**Why it matters:** Knowing something is increasing is useful. Knowing it's increasing *and accelerating* is way more useful. Derivatives let you see not just where things are, but where they're headed. It's the difference between "we're going up" and "we're going up and it's getting faster oh no."

**Everyday example:** Bread dough rising. It rises slowly at first, then faster, then slows down as it peaks. The "how fast is it rising right now" at any moment is the derivative. Bakers know this intuitively. Calculus just lets you put numbers on it.

**Chained effects:** Sometimes A depends on B, and B depends on C. Nudge C, and it ripples through to A.

Your energy depends on your sleep. Your sleep depends on how much coffee you had. Calculus can formalize "if I drink one more cup at 4pm, how will I feel tomorrow?" (The answer is "bad." You don't need calculus for that one. But you could use it.)

**Quick estimates:** Near any point, most curves look almost like a straight line. That means you can make fast, good-enough predictions without heavy math.

"If I leave 10 minutes later, I'll probably hit 5 more minutes of traffic." That's a linear approximation. You do this constantly. Calculus just makes it official.

---

### Unit 4: What derivatives tell you

Once you can measure "rate of change," you can find peaks, valleys, and trends.

**Highs and lows:** Where does the graph peak? Where does it bottom out? At those points, the derivative is zero. Things have momentarily stopped changing before heading the other direction.

Your alertness during the day has peaks and valleys. Somewhere mid-morning, you're sharp. After lunch, you're a zombie. Those turning points are where the "rate of change" crosses zero. The derivative literally tells you when you hit the wall.

**Speeding up or slowing down:** The derivative tells you if something is increasing or decreasing. The derivative *of the derivative* tells you if that change is accelerating or decelerating.

Your savings account: Is it growing? Good. Is the growth itself speeding up because of compound interest? Even better. That's the second derivative at work. (This is also why compound interest is either your best friend or your worst enemy depending on which side of it you're on.)

**Long-term behavior:** What happens as things get very big or very small? Does the function level off? Blow up? Knowing the end behavior helps you set expectations.

Learning a new skill follows this pattern. You improve quickly at first, then progress slows as you approach your ceiling. That ceiling is an asymptote. Accepting this saves you from the frustration of "why am I not improving anymore?" You are. It's just slower now. That's how ceilings work.

---

### Unit 5: Integration

Integration answers "how much has piled up over time?"

It's the flip side of derivatives. If derivatives break things into tiny instantaneous moments, integrals add all those moments back together.

**Area under the curve:** Imagine chopping time into tiny slices and adding up "value times time" for each slice. That's integration.

Water usage during a shower. Your flow rate varies: full blast, then lower while you shampoo, then full blast again to rinse. The total water used is the area under that flow-rate curve. Your water bill doesn't care about the shape. It just cares about the total.

**Definite integrals:** "From 2pm to 3pm, how much accumulated?" This is how you get totals over specific windows.

Calories burned during a workout. Your burn rate changes minute to minute, but the integral gives you the total. Your fitness tracker does this math constantly. Now you know what it's doing.

**The fundamental connection:** Derivatives and integrals undo each other. If you know the rate of change at every moment, you can recover the total. If you know the total as a function of time, you can recover the rate.

Two sides of the same coin. This is the big beautiful idea that ties calculus together, and it's genuinely elegant once it clicks.

---

### For the engineers
#### SRE examples

**Derivatives in monitoring:** Watching error *rate* is useful. Watching how fast that rate is *changing* catches spikes earlier. That's the derivative. Your alerts should care about this.

**Integrals in impact assessment:** Total downtime minutes is the integral of your "up/down" signal over time. A short severe outage might have the same area (same total impact) as a long mild degradation. Integrals let you compare apples to apples.

**Limits in alerting:** A 5-minute average can mask a spike happening right now. Limits (and the derivatives built on them) let you focus on the instant, not just the average. This is why "rate of change" alerts exist.

**Chained effects in capacity planning:** Latency depends on queue length. Queue length depends on arrival rate. The chain rule helps you answer "if traffic increases 10%, how much will latency increase?" Impress your coworkers. Or at least confuse them.

**Asymptotes in optimization:** There's often a best-case latency you can't beat without a redesign. Knowing that floor saves you from chasing impossible improvements. Sometimes the answer is "this is as good as it gets without a rewrite."


---

### TL;DR

**Derivative** = how fast it's changing right now. Useful for spotting trends, spikes, and turning points.

**Integral** = how much has piled up over time. Useful for totals, costs, and cumulative impact.

Calculus helps you see both the moment and the big picture. That's the whole game.
