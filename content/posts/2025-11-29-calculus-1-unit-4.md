---
layout:     post
title:      "Maximums, minimums, and the shape of things"
subtitle:   "The optimization unit (finally)"
description: "Where derivatives get practical: finding peaks and valleys, understanding curves, and optimizing things that matter."
date:    2025-11-29T12:00:00-04:00
author:     "O.J. Wilcox"
image: "/img/posts/bg-calculus-1-unit-4.png"
tags: [ "Math", "Calculus 1" ]
URL: "/2025/11/29/calculus-1-unit-4/"
categories: [ Calculus ]
draft: false
---

## Maximums, minimums, and the shape of things
_(Calculus 1 - Unit 4)_

### The big idea

Derivatives tell you how fast something is changing. But here's the real power: they also tell you where things peak, bottom out, and change direction.

Finding the maximum profit. The minimum cost. The fastest route. The optimal dose. These aren't abstract math problems. They're the reason calculus exists.

This unit is where derivatives become genuinely useful.

### What even is a maximum or minimum?

A maximum is a peak. A point where the function is higher than everything nearby. A minimum is a valley. Lower than everything nearby.

"Local" max/min means it's the highest/lowest in its neighborhood. Maybe there's something bigger elsewhere, but right here, it's the champ.

"Global" or "absolute" max/min means it's the highest/lowest everywhere. The actual winner, no qualifications.

Your goal is usually to find these points. And derivatives are the tool.

### Critical numbers: where interesting things happen

Here's the key insight: at a local max or min, the derivative is either zero or doesn't exist.

Think about it. At a peak, the function stops going up and starts going down. For an instant, it's flat. The slope is zero. Same at a valley, just flipped.

A critical number is any x value where f'(x) = 0 or f'(x) doesn't exist. These are your candidates for maxes and mins. Not every critical number is a max or min, but every max or min is at a critical number (or at an endpoint).

So step one is always: find the critical numbers.

### Finding the actual max and min

Once you have your critical numbers, you need to figure out which are maxes, which are mins, and which are neither.

**For a closed interval [a, b]:** Evaluate the function at every critical number AND at both endpoints. Compare all the values. The biggest is your absolute max. The smallest is your absolute min. Brute force, but it works.

**First derivative test:** Check what f'(x) does on either side of the critical number. If f' goes from positive to negative (increasing then decreasing), you have a local max. If f' goes from negative to positive, local min. If it doesn't change sign, neither.

**Second derivative test:** Evaluate f''(x) at the critical number. If f'' is negative there, the function is concave down, curving like a frown. That's a local max. If f'' is positive, concave up like a smile. Local min. If f'' is zero, this test is useless and you need the first derivative test instead.

The second derivative test is faster when it works. The first derivative test always works.

### The Extreme Value Theorem (a guarantee)

If your function is continuous on a closed interval [a, b], it's guaranteed to have an absolute max and an absolute min somewhere on that interval.

This sounds obvious, but it matters. It means the search will succeed. You won't be chasing a maximum that doesn't exist. Continuous function, closed interval, extremes exist. Period.

The "closed interval" part is crucial. On an open interval (a, b), the function might approach a value at the endpoints without ever reaching it. No guarantee.

### Rolle's Theorem (the simple case)

If f(a) = f(b) (same height at both endpoints), and f is continuous on [a, b] and differentiable on (a, b), then somewhere between a and b, the derivative equals zero.

Picture it: you start and end at the same height. Unless you stayed flat the entire time, you must have had a peak or valley in between. At that peak or valley, the tangent line is horizontal. Slope zero.

This feels almost too obvious to state, but it's the foundation for something more powerful.

### Mean Value Theorem (the real deal)

This is Rolle's Theorem generalized.

If f is continuous on [a, b] and differentiable on (a, b), then there's at least one point c in between where the instantaneous rate of change equals the average rate of change.

f'(c) = [f(b) - f(a)] / (b - a)

In english: somewhere on the trip, you were going exactly the average speed.

If you drove 120 miles in 2 hours, your average speed was 60 mph. The Mean Value Theorem says at some moment during that drive, your speedometer read exactly 60. Maybe just for an instant, but it happened.

This theorem gets used to prove other things more than you'll use it directly. But the intuition is gold: averages are achieved somewhere.

### Concavity: which way is it curving?

Concave up means the function curves like a bowl. Holds water. Smiley face. The derivative is increasing.

Concave down means it curves like a dome. Spills water. Frowny face. The derivative is decreasing.

The second derivative tells you which: f'' > 0 means concave up. f'' < 0 means concave down.

Why care? Because concavity tells you about acceleration. If your position function is concave up, your velocity is increasing. You're speeding up.

### Inflection points: where the bend changes

An inflection point is where concavity switches. The curve goes from bowl-shaped to dome-shaped, or vice versa.

At an inflection point, f'' = 0 (or doesn't exist). But not every place where f'' = 0 is an inflection point. You need the concavity to actually change on either side.

Inflection points are where the "mood" of the function shifts. Going from accelerating to decelerating. From things getting better faster to things getting better slower.

### Asymptotes: behavior at the extremes

What happens as x gets huge? Or hugely negative? Or approaches some forbidden value?

**Horizontal asymptotes:** If f(x) approaches some constant L as x → ∞ (or x → -∞), the line y = L is a horizontal asymptote. The function levels off. Gets closer and closer to that height but never quite settles.

**Vertical asymptotes:** If f(x) blows up to ±∞ as x approaches some value a, the line x = a is a vertical asymptote. The function goes haywire near that input. Usually because you're dividing by something approaching zero.

**Oblique (slant) asymptotes:** Sometimes as x → ∞, the function doesn't level off to a constant but instead approaches a slanted line. This happens with rational functions where the numerator's degree is exactly one more than the denominator's.

Asymptotes are about end behavior. The big picture of what the function is doing far from the origin.

### Limits at infinity (zooming out)

To find horizontal asymptotes, you take limits as x → ∞ or x → -∞.

For rational functions (polynomial divided by polynomial), there's a quick rule:
- If the top degree is less than the bottom, limit is 0.
- If degrees are equal, limit is the ratio of leading coefficients.
- If top degree is greater, it blows up (no horizontal asymptote).

For other functions, you might need to do more work. Factor out dominant terms. Use L'Hôpital's rule. Get creative.

### L'Hôpital's Rule (unsticking stuck limits)

Sometimes a limit gives you 0/0 or ∞/∞. These are "indeterminate forms." You can't just plug in.

L'Hôpital's Rule says: if you get 0/0 or ∞/∞, take the derivative of the top and the derivative of the bottom separately, then try the limit again.

lim f(x)/g(x) = lim f'(x)/g'(x)

(assuming that second limit exists or is ±∞)

You can apply it repeatedly if you keep getting indeterminate forms. It's weirdly powerful.

**Other indeterminate forms:** 0·∞, ∞ - ∞, 0⁰, 1^∞, ∞⁰. These can all be rewritten to use L'Hôpital's Rule. Usually by converting products to fractions or taking logs of expressions with weird exponents.

### Putting it all together: sketching curves

With all these tools, you can sketch a function's graph without plotting a million points.

**The checklist:**
1. Domain: where is f defined?
2. Intercepts: where does it cross the axes?
3. Symmetry: is it even, odd, or neither?
4. Asymptotes: horizontal, vertical, oblique?
5. First derivative: where is f increasing/decreasing? Critical numbers? Local max/min?
6. Second derivative: where is f concave up/down? Inflection points?
7. Plot key points and connect the dots with the right curvature.

This is detective work. Each piece of information constrains what the graph can look like. By the end, there's usually only one possibility.

### Applied optimization (the payoff)

This is why anyone outside of math class cares about calculus.

You have a quantity you want to maximize or minimize. Profit, area, time, cost, material. You write it as a function. You find the derivative. You set it to zero. You check that your answer is actually a max or min. Done.

The hard part is usually setting up the function, not the calculus. Read the problem carefully. Draw a picture. Identify what you're optimizing and what constraints you have.

Constraints often let you reduce to a single variable. "The perimeter is 100 feet" lets you write one dimension in terms of the other. Then you have a function of one variable, and you know how to handle that.

---

### For the engineers

#### SRE examples

**Finding optimal configurations:** What cache size maximizes hit rate per dollar? What thread pool size minimizes latency? These are optimization problems. Model the metric as a function of the parameter, find where the derivative is zero, verify it's a minimum or maximum.

**Critical points in system behavior:** The load level where your system transitions from "handling it fine" to "struggling" often corresponds to a critical point in some performance metric. Latency might be flat until a certain load, then spike. Finding that critical point helps you set capacity limits.

**Mean Value Theorem for rate monitoring:** If a counter increased by 1000 over 10 seconds, your average rate was 100/sec. MVT says at some instant, the actual rate was exactly 100/sec. This is why average-rate alerts can miss spikes: the average might be fine even if the instantaneous rate was terrible for a moment.

**Concavity in growth curves:** User growth that's concave up is accelerating. Exciting. Concave down means growth is slowing. Still growing, but the rate of growth is decreasing. Inflection point is where it shifts. Investors watch for this.

**Asymptotic behavior in scaling:** Most systems have horizontal asymptotes in their performance curves. Add more servers, latency drops, but it levels off. There's a floor you can't go below due to network round trips, serialization, etc. Understanding that asymptote helps you avoid throwing money at diminishing returns.

**L'Hôpital for weird metric ratios:** Error rate during deploy might be 0/0 (zero errors, zero requests) for a moment. Ratio of two growing metrics might be ∞/∞ as both explode. L'Hôpital-style thinking (look at the rates of change instead of the raw values) gives you meaningful numbers.

**Vertical asymptotes as system limits:** As you approach 100% CPU or memory, response times don't just increase linearly. They often shoot toward infinity. That's a vertical asymptote in the latency-vs-utilization curve. Stay away from it.

**Optimization with constraints:** Minimize cost while meeting SLO. Maximize throughput without exceeding latency budget. These are constrained optimization problems. The constraint defines a relationship between variables, reducing your degrees of freedom. Then you optimize what's left.