Go library to control a light.
Works with the [WebLight](https://github.com/sowbug/weblight).

# Usage

Run `light.Update()` in its own goroutine, perhaps with a short delay (sub-second).

```go
	light.Add(light.Task{
		Color: &light.Red,
	})

	color, err := light.Update()
	if err != nil {
		log.Fatal(err) // or ignore, whatever
	}
	log.Printf("set color: %+v", color)
```

To add colors that take priority expire after a time:

```go
	light.Add(light.Task{
		Priority: 100,     // higher takes precedence
		Color: &light.Red,
		Duration: time.Second,
	})
```

To cancel tasks:

```go
	ref := light.Add(light.Task{
		Priority: 10,
		Color: &light.Color{128, 0, 255},
	})

	// later
	ref.Cancel()
```
