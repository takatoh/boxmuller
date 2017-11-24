package boxmuller

import (
	"math"
	"math/rand"
	"time"
)

type BoxMuller struct {
	r     *rand.Rand
	mu    float64
	sigma float64
}

func NewBoxMuller(mu, sigma float64) *BoxMuller {
	p := new(BoxMuller)
	p.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	p.mu = mu
	p.sigma = sigma
	return p
}

func (bm *BoxMuller) Rand() (float64, float64) {
	x := bm.r.Float64()
	y := bm.r.Float64()	
	z1 := math.Sqrt(-2.0 * math.Log(x)) * math.Cos(2.0 * math.Pi * y)
	z2 := math.Sqrt(-2.0 * math.Log(x)) * math.Sin(2.0 * math.Pi * y)
	return bm.sigma * z1 + bm.mu, bm.sigma * z2 + bm.mu
}
