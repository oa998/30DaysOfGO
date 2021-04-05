### Contact Tracing

Given a list of individuals that were in contact with one another, write a function to find the minimum distance between two given people.

example input:

```
pairs = [][]string{
    []string{"Todd", "Kim"},
    []string{"Todd", "Thomas"},
    []string{"Kim", "Mike"},
    []string{"Susan", "Kim"},
    []string{"Phillip", "Kim"},
    []string{"Mason", "Todd"},
    []string{"Thomas", "Trevor"},
    []string{"Rick", "Thomas"},
}
start, end := "Kim", "Trevor"
```

example output:

```
Distance from Kim to Trevor is 3
```
