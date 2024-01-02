# boxmuller

Implementation of [Box-Muller transform](https://en.wikipedia.org/wiki/Box%E2%80%93Muller_transform) in Golang.

## Install

Add the following line to your `go.mod` file.

```
require github.com/takatoh/boxmuller v1.1.0
```

## Usage

```go
bm := boxmuller.New(mu, sigma)
z1, z2 := bm.Rand()
```

## API

### New(mu, sigma)

Create a BoxMuller struct.

### (bm *BoxMuller) Rand()

Returns 2 values.

## License

MIT License
