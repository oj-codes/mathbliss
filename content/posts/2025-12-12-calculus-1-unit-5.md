---
layout:     post
title:      "Integrals"
subtitle:   "(undoing derivatives, finding totals, and the grand finale)"
description: "The grand finale: how integrals undo derivatives, find totals, and tie the whole course together."
date:    2025-12-12T07:00:00-04:00
author:     "O.J. Wilcox"
image: "/img/posts/bg-calculus-1-unit-5.png"
tags: [ "Math", "Calculus 1" ]
URL: "/2025/12/12/calculus-1-unit-5/"
categories: [ Calculus ]
draft: false
---

## Integrals (undoing derivatives, finding totals, and the grand finale)
_(Calculus 1 - Unit 5)_

### The big idea

If derivatives answer "how fast is it changing?", integrals answer "how much did it accumulate?"

You've been doing this your whole life. Checking your total screen time. Watching your bank balance grow (or shrink). Counting miles on a road trip. These are all accumulation problems.

Integrals formalize the process. And here's the twist: they turn out to be the reverse of derivatives. That's the punchline of the entire course.

### Area: where it all starts

The integral started as a geometry problem. How do you find the area under a curve?

For rectangles, easy: base times height. For triangles, half that. But what about the area under y = x² from 0 to 1? There's no formula from geometry class.

The calculus trick: chop it into thin rectangles. Add up their areas. Make the rectangles thinner and thinner. Take the limit as they become infinitely thin.

That limiting sum is the integral.

### Sigma notation (a quick detour)

Before we get to integrals, you need to be comfortable with Σ (sigma). It just means "add these up."

Σᵢ₌₁ⁿ i means 1 + 2 + 3 + ... + n. The i is a counter. It starts at 1, ends at n, and you sum every value.

It's fancy notation for addition. You'll see it in the definition of integrals, and it's useful for writing Riemann sums compactly.

### Riemann sums: the rectangles method

Here's how you actually approximate area under a curve.

Pick an interval [a, b]. Chop it into n pieces. In each piece, pick a point and use the function value there as your rectangle height. Multiply by the width. Add them all up.

Left Riemann sum: use the left endpoint of each piece.
Right Riemann sum: use the right endpoint.
Midpoint sum: use the middle.

As n gets larger (more, thinner rectangles), all of these approach the same value. That value is the definite integral.

You probably won't compute many Riemann sums by hand. But understanding that integrals ARE these sums (in the limit) makes everything else make sense.

### The definite integral

The definite integral of f(x) from a to b is written:

∫ₐᵇ f(x) dx

It represents the exact area under the curve (with some caveats about negative values, which we'll get to).

The ∫ is a stretched-out S, for "sum." The dx reminds you that you're summing up infinitely thin slices of width dx. The a and b are your boundaries.

This isn't just notation. It's describing a process: slice, multiply, sum, take the limit.

### When the function goes negative

If f(x) dips below the x-axis, the integral counts that area as negative.

So ∫ₐᵇ f(x) dx gives you the "signed area." Area above the axis minus area below.

If you want total area regardless of sign, you need to integrate |f(x)|, or split the integral at the zeros and handle each piece appropriately.

This isn't a bug. Signed area is often what you want. If velocity goes negative (you're going backwards), the integral of velocity should subtract that distance from your total displacement.

### Properties of definite integrals

Integrals behave nicely.

**Additivity:** ∫ₐᵇ + ∫ᵦᶜ = ∫ₐᶜ. You can break an integral into pieces.

**Reversing limits:** ∫ₐᵇ = -∫ᵦₐ. Flip the bounds, flip the sign.

**Constants pull out:** ∫ₐᵇ c·f(x) dx = c · ∫ₐᵇ f(x) dx.

**Sum/difference:** ∫ₐᵇ [f(x) ± g(x)] dx = ∫ₐᵇ f(x) dx ± ∫ₐᵇ g(x) dx.

These let you break hard integrals into easier pieces.

### Antiderivatives: running derivatives backwards

An antiderivative of f(x) is any function F(x) whose derivative is f(x). You're undoing differentiation.

The antiderivative of 2x is x² (because the derivative of x² is 2x). But also x² + 5. Or x² - 17. Any constant works because constants vanish when you differentiate.

So we write the general antiderivative as F(x) + C, where C is an arbitrary constant. The "+C" matters. Don't forget it.

### Indefinite integrals

The indefinite integral is just notation for "find the antiderivative":

∫ f(x) dx = F(x) + C

No bounds. You're finding a family of functions, not a number.

**Basic rules (reverse the derivative rules):**

∫ xⁿ dx = xⁿ⁺¹/(n+1) + C (as long as n ≠ -1)

∫ 1/x dx = ln|x| + C

∫ eˣ dx = eˣ + C

∫ sin(x) dx = -cos(x) + C

∫ cos(x) dx = sin(x) + C

∫ sec²(x) dx = tan(x) + C

Notice these are just derivative rules read backwards.

### The Fundamental Theorem of Calculus

This is the big one. The theorem that ties the whole course together.

**Part 1:** If you define F(x) = ∫ₐˣ f(t) dt (the integral from a to a variable endpoint), then F'(x) = f(x). The derivative of an integral (with a variable upper limit) gives you back the original function.

**Part 2:** If F is any antiderivative of f, then ∫ₐᵇ f(x) dx = F(b) - F(a).

Part 2 is the computational workhorse. To evaluate a definite integral, find any antiderivative and plug in the bounds. Subtract. Done.

No Riemann sums. No limits of infinite series. Just find F, compute F(b) - F(a), and you have your answer.

This is why antiderivatives matter. This is why we spent Unit 3 learning derivatives. The Fundamental Theorem connects the two big ideas of calculus: instantaneous rate of change and total accumulation.

### u-substitution: the chain rule in reverse

When the function is a composition, you need u-substitution. It's the chain rule, backwards.

If you see something like ∫ 2x·eˣ² dx, notice that 2x is the derivative of x². Let u = x². Then du = 2x dx. The integral becomes ∫ eᵘ du = eᵘ + C = eˣ² + C.

The pattern: find an "inside function" whose derivative is also present (maybe off by a constant). Substitute. Integrate. Substitute back.

This works for powers, trig functions, exponentials. Anywhere you'd use chain rule for derivatives, you'll use u-sub for integrals.

For definite integrals, you can either convert back to x and use the original bounds, or convert the bounds to u-values and never switch back. Both work.

### Area between curves

What if you want the area between two curves, not just under one?

If f(x) ≥ g(x) on [a, b], the area between them is:

∫ₐᵇ [f(x) - g(x)] dx

Top minus bottom. That's it.

If the curves cross (intertwine), you need to find where they intersect, split into pieces where one is consistently on top, and add up the areas.

Draw a picture. Seriously. It helps you see which function is on top in each region.

### Average value of a function

The average value of f(x) on [a, b] is:

(1/(b-a)) · ∫ₐᵇ f(x) dx

In words: the integral divided by the length of the interval. Total accumulation divided by the amount of space. Just like averaging grades: sum divided by count.

### Mean Value Theorem for Integrals

If f is continuous on [a, b], there exists some c in [a, b] where:

f(c) = (1/(b-a)) · ∫ₐᵇ f(x) dx

Translation: somewhere on the interval, the function actually equals its average value.

This is the integral version of the MVT you saw in Unit 4. The average is achieved somewhere.

### Solving differential equations (a preview)

If you're given y' = f(x) and an initial condition y(a) = b, you can solve for y by integrating.

y = ∫ f(x) dx gives you y = F(x) + C. Plug in the initial condition to find C.

This is a baby differential equation. More complex ones are a whole course. But the idea starts here: derivatives describe rates, integrals recover the original quantities.

### Approximating integrals numerically

Sometimes you can't find an antiderivative. The function is too ugly, or you only have data points, not a formula.

Approximation methods:

**Left/Right/Midpoint Riemann sums:** What we started with. Choose your points, add up rectangles.

**Trapezoid rule:** Instead of rectangles, use trapezoids. Generally more accurate for the same number of pieces.

**Simpson's rule:** Use parabolas instead of flat tops or straight lines. Even more accurate.

Computers use these (and fancier versions) all the time. Not every integral has a nice closed form. But every integral can be approximated.

### Using tables (a practical note)

There are tables of integrals. Big books of them. Before computers, this was how people found antiderivatives of complicated functions.

You can still use them. Recognize the form, look it up, apply it. It's not cheating. It's engineering.

---

### For the engineers

#### SRE examples

**Integrals in monitoring are everywhere:** Total requests over an hour is the integral of requests per second. Total errors is the integral of error rate. Any "counter" metric is secretly an integral of a rate.

**Area under the curve for SLOs:** SLO burn rate graphs show error rate over time. The area under that curve is your total error budget consumed. When the area exceeds your budget, you've violated the SLO.

**Signed area and net change:** If your queue length rate goes positive (filling) then negative (draining), the integral gives net change. You might process 1000 items total but end up only 50 ahead of where you started.

**Average value for reporting:** "Average CPU utilization over the past week" is exactly the average value formula. Integrate the utilization, divide by the time span.

**Riemann sums are how your monitoring works:** Prometheus, Datadog, whatever you use. It samples at intervals and sums rectangles. When you see "rate over 5 minutes," that's a Riemann approximation. Smaller intervals = more accurate = more data = more cost. Tradeoffs everywhere.

**Numerical integration for forecasting:** If you're projecting storage growth and all you have is daily snapshots, you're doing numerical integration. Trapezoid rule is often good enough.

**Area between curves for comparison:** How much better was version B than version A? Plot both latency curves over time. The area between them is your total improvement (or regression).

**The Fundamental Theorem for capacity:** If you know your rate of resource consumption (derivative), you can integrate to find total consumption over any period. If you know total consumption at two points, you can work out the average rate between them. The FTC is the bridge.

**u-substitution in log analysis:** If your metric depends on a transformation of another variable (latency as a function of log(requests)), substitution untangles it. The same algebraic pattern applies.