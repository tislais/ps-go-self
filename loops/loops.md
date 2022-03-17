## For
`for ` without an expression creates an infinite loop

example:
```
for {
    <code>
}
```

## For Range
`for range ` iterates over a list (one entry per iteration of the loop)

```
for i := 0; i < 10; i++ {
    <code>
}
```
## Break 
A single `break` will break out of the current `for` loop. If nested, then you break into the parent `for` loop. In order to break out more than one level up, we use break points. They can be called `banana` if you want. If you define a `banana` tho, you gotta use it, or the app won't run.

```
banana:
    for <expression...> {
        <code>
        for <expression...> {
            <code>
            for <expression...> {
                <code>
                break banana
            }
        }
    }
```

## Continue
The idea is that whenever Go encounters one in a loop, it stops what it's doing and jumps straight to the top for a new evaluation or potentially more iterations through the loop. Post statements are executed as part of the continue.

## Recap

Infinite loops

```
for {
    <code>
}
```

While loops (boolean expression)
```
for pre; expr; post {
    <code>
}
```

Range loops
```
for i := range <list> {
    <code>
}
```