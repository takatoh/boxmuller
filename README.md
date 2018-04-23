# boxmuller

Implementation of [Box-Muller transform](https://en.wikipedia.org/wiki/Box%E2%80%93Muller_transform) in Golang.

## Install
``` go get github.com/takatoh/boxmuller```

## Usage
```go

bm := boxmuller.New(mu, sigma)
z1, z2 := bm.Rand()

```

## API
### New(mu, sigma)
Create a BoxMuller struct.

### Rand()
Returns 2 values.

## License
MIT
