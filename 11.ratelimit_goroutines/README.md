# Run with docker to see cpu and memory usage

```
cd 7.limiter
docker run --rm -it --memory 128m $(docker build -q .)
```

This was supposed to be used for rate-limiter to limit number of parallel goroutines.
Kept it here for future retry
