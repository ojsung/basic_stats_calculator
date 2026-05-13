package pade

import (
	"math/big"
	"sync"
)

type cacheKey struct{ m, n int }

type padeCoefficients struct {
	p []*big.Float
	q []*big.Float
}

type coefficientCache struct {
	mu   sync.RWMutex
	data map[cacheKey]padeCoefficients
}

func (c *coefficientCache) get(m, n int) (padeCoefficients, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[cacheKey{m, n}]
	return v, ok
}

func (c *coefficientCache) set(m, n int, coeffs padeCoefficients) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[cacheKey{m, n}] = coeffs
}
