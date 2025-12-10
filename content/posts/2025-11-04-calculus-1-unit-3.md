---
layout:     post
title:      "Derivatives"
subtitle:   "(the math of \"how fast is it changing right now?\")"
description: "The math of \"how fast is it changing right now?\" All the derivative rules, what they mean, and when to use them."
date:    2025-11-04T12:00:00-04:00
author:     "O.J. Wilcox"
image: "/img/posts/bg-calculus-1-unit-3.png"
tags: [ "Math", "Calculus 1" ]
URL: "/2025/11/04/calculus-1-unit-3/"
categories: [ Calculus ]
draft: false
---

## Derivatives (the math of "how fast is it changing right now?")
_(Calculus 1 - Unit 3)_

### The big idea

A derivative tells you how fast something is changing at a specific instant.

That's it. All the formulas, all the rules, all the notation: just tools for answering "how fast is this changing right now?"

You already think this way. When you check if your phone battery will last until you get home, you're estimating a rate of change. When you notice traffic is "getting worse," you're sensing a derivative. Calculus just gives you the machinery to calculate it precisely.

### What a derivative actually is

Remember limits? Here's where they pay off.

The derivative of f(x) at a point is the slope of the tangent line at that point. And we find that slope by taking the limit of secant line slopes as the two points get infinitely close together.

The formal definition looks like this:

f'(x) = lim(h→0) [f(x+h) - f(x)] / h

In english: take the slope between x and a nearby point x+h, then see what happens as h shrinks to zero. That limiting slope is your derivative.

You could calculate every derivative this way. You'd also want to throw your calculator out a window. That's why we have rules.

### The basic rules (your new best friends)

These rules let you skip the limit definition and just write down derivatives directly.

**Power rule:** The derivative of xⁿ is n·xⁿ⁻¹. Bring the exponent down, subtract one from it. The derivative of x³ is 3x². The derivative of x is 1. The derivative of a constant is 0 (flat line, no change).

**Constant multiple:** The derivative of 5f(x) is 5f'(x). Constants just come along for the ride.

**Sum/difference:** The derivative of f(x) + g(x) is f'(x) + g'(x). Differentiate term by term.

With just these rules, you can handle any polynomial. And that covers a surprising amount of calculus homework.

### Product and quotient rules (when functions get tangled)

What if two functions are multiplied together? You can't just differentiate each piece separately. There's interaction.

**Product rule:** The derivative of f(x)·g(x) is f'(x)·g(x) + f(x)·g'(x).

In words: derivative of the first times the second, plus the first times the derivative of the second. Some people remember "first d-second plus second d-first." Whatever works.

**Quotient rule:** The derivative of f(x)/g(x) is [f'(x)·g(x) - f(x)·g'(x)] / [g(x)]².

In words: low d-high minus high d-low, all over low squared. It's uglier than the product rule. That's just how division is.

### The chain rule (the real workhorse)

Most interesting functions are compositions. Something inside something else. Like sin(x²) or e^(3x) or √(1 + x²).

The chain rule says: differentiate the outside, leave the inside alone, then multiply by the derivative of the inside.

For sin(x²): the outside is sin(stuff), its derivative is cos(stuff). The inside is x², its derivative is 2x. So the answer is cos(x²)·2x.

This rule is everywhere. If you're struggling with derivatives, there's a good chance it's a chain rule issue. When in doubt, ask yourself "is there a function inside another function?" If yes, chain rule.

### Special functions you'll see constantly

**Trig functions:**
- Derivative of sin(x) is cos(x)
- Derivative of cos(x) is -sin(x)
- Derivative of tan(x) is sec²(x)

The others follow patterns. Notice sin and cos just cycle into each other (with a sign flip for cosine).

**Exponentials:**
- Derivative of eˣ is eˣ. Yes, really. It's its own derivative. This is why e is special.
- Derivative of aˣ (any base) is aˣ·ln(a). The natural log shows up because e is the "natural" base.

**Logarithms:**
- Derivative of ln(x) is 1/x. Clean and simple.
- Derivative of log_a(x) is 1/(x·ln(a)). The natural log sneaks in again.

### Higher-order derivatives

The derivative of a derivative is the second derivative. Written f''(x) or d²y/dx².

If the first derivative is velocity (rate of change of position), the second derivative is acceleration (rate of change of velocity). You can keep going: third derivative, fourth derivative, as many as you want.

Second derivatives tell you about curvature. Is the function bending up or down? Is the rate of change itself increasing or decreasing?

### Differentiability: when derivatives don't exist

A function has a derivative wherever its graph is "smooth" in a specific way. No sharp corners, no vertical tangent lines, no jumps or holes.

The classic example: f(x) = |x| at x = 0. The graph makes a sharp V. There's no single tangent line that makes sense there. The slopes from the left and right don't agree. So the derivative doesn't exist at that point.

If a function isn't continuous at a point, it definitely isn't differentiable there. But the reverse isn't true: continuous doesn't guarantee differentiable. That V shape in |x| is continuous but not differentiable.

### Tangent line equations

Once you have a derivative, you can write the equation of the tangent line at any point.

You need: the point (a, f(a)) and the slope f'(a). Then it's just point-slope form:

y - f(a) = f'(a)(x - a)

This tangent line is the best linear approximation of the function near that point. Which leads to...

### Linear approximation (derivatives as prediction tools)

Near any point, a smooth function looks almost like its tangent line. So if you know f(a) and f'(a), you can estimate f(x) for x close to a:

f(x) ≈ f(a) + f'(a)(x - a)

This is shockingly useful. Need √4.1 without a calculator? You know √4 = 2. The derivative of √x is 1/(2√x), which at x=4 is 1/4. So √4.1 ≈ 2 + (1/4)(0.1) = 2.025. The actual value is about 2.0248. Not bad for mental math.

The "differential" notation (dy = f'(x)dx) is just another way to write this approximation. The change in y is approximately the derivative times the change in x.

### Implicit differentiation (when y won't cooperate)

Sometimes you can't solve for y explicitly. Like x² + y² = 25 (a circle). You can't write y as a single function of x without splitting into top and bottom halves.

Implicit differentiation handles this. Differentiate both sides with respect to x, treating y as a function of x. Every time you differentiate a y term, tack on a dy/dx (chain rule). Then solve for dy/dx.

For the circle: 2x + 2y(dy/dx) = 0, so dy/dx = -x/y.

This gives you slopes on curves you couldn't otherwise handle.

### Logarithmic differentiation (a clever trick)

Some functions are ugly to differentiate directly. Like x^x or complicated products with variable exponents.

The trick: take the natural log of both sides, use log properties to simplify, differentiate implicitly, then solve for dy/dx.

Why does this work? Because ln turns products into sums and exponents into coefficients. It untangles things.

### Related rates (the word problems)

Here's where derivatives meet the real world.

You have two quantities that are both changing over time, and they're related by some equation. You know how fast one is changing. You want to find how fast the other is changing.

Classic example: A ladder is sliding down a wall. The bottom is moving away from the wall at 2 feet per second. How fast is the top sliding down when the bottom is 6 feet from the wall?

The setup: x² + y² = L² (Pythagorean theorem, L is ladder length). Differentiate both sides with respect to time t. You get 2x(dx/dt) + 2y(dy/dt) = 0. Plug in what you know, solve for what you don't.

The key insight: both x and y are functions of time, even if nobody said so explicitly. Everything that's changing needs a chain rule with respect to t.

These problems are really just implicit differentiation plus paying attention to what's actually varying.

### Newton's method (finding roots with derivatives)

Want to find where f(x) = 0 but can't solve it algebraically? Newton's method gives you a way to get closer and closer.

Start with a guess x₀. Draw the tangent line there. See where it crosses the x-axis. That's your next guess x₁. Repeat.

The formula: x_{n+1} = x_n - f(x_n)/f'(x_n)

Each iteration usually gets you much closer to the root. It can fail (if you start badly or hit a horizontal tangent), but when it works, it works fast.

This is how calculators actually find square roots and other values. Not magic. Just derivatives and iteration.

---

### For the engineers

#### SRE examples

**Derivatives in monitoring:** Your monitoring system's "rate" metrics are derivatives. Requests per second is the derivative of total requests. Error rate is the derivative of error count. When you alert on "rate of change," you're alerting on a derivative.

**Second derivatives for anomaly detection:** If error rate is increasing, that's concerning. If the rate of increase is itself increasing (positive second derivative), that's "things are getting worse faster." Second derivatives catch acceleration toward failure.

**Chain rule in complex systems:** Latency depends on queue length, which depends on arrival rate, which depends on traffic. To understand how traffic changes affect latency, you're implicitly using chain rule thinking: multiply the rates of change through the chain.

**Linear approximation for capacity planning:** If your system handles 1000 RPS at 50ms latency and the derivative of latency with respect to load is 0.02 ms/RPS at that point, you can estimate that 1100 RPS will give you roughly 50 + 0.02(100) = 52ms. Good enough for planning, fast enough to be useful.

**Related rates in practice:** Your database size is growing. Storage cost depends on size. How fast is cost growing? That's a related rates problem. If size grows at 10GB/month and cost is $0.10/GB, cost grows at $1/month. Simple example, but the pattern scales to complex dependencies.

**Newton's method for optimization:** Finding the load level where response time equals your SLO, or the cache size where hit rate plateaus: these are root-finding problems. Newton's method (or variations of it) is running under the hood of many optimization tools.

**Implicit relationships:** Sometimes you have an equation relating metrics (like Little's Law: L = λW, where L is items in system, λ is arrival rate, W is wait time) and you need to know how one changes when another does. Implicit differentiation handles this without needing to solve for each variable separately.
