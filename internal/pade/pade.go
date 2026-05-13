package pade

import (
	"errors"
	"math"
	"math/big"
	"sync"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
	mat "github.com/ojsung/basic_stats_calculator/pkg/matrix"
)

// log₂(4/0.1²) = log₂(400). Each [n/n] order contributes this many bits of accuracy
// when |u| ≤ 0.1 after argument reduction (lo=0.9, hi=1.1 thresholds).
var log2_400 = math.Log2(400)

func orderForPrec(prec uint) int {
	return int(math.Ceil(float64(prec) / log2_400))
}

func solveCoefficients(m, n int) (padeCoefficients, error) {
	taylorCoeff := func(k int) *big.Float {
		if k <= 0 {
			return bu.PrecFloat().SetInt64(0)
		}
		num := int64(1)
		if k%2 == 0 {
			num = -1
		}
		return bu.PrecFloat().Quo(
			bu.PrecFloat().SetInt64(num),
			bu.PrecFloat().SetInt64(int64(k)),
		)
	}

	// Build n×n matrix C: C[i][j] = c_{m+i-j}, i,j = 0..n-1
	cRows := make([][]*big.Float, n)
	for i := range n {
		cRows[i] = make([]*big.Float, n)
		for j := range n {
			cRows[i][j] = taylorCoeff(m + i - j)
		}
	}
	cMat, err := mat.NewBigMatrix[*big.Float](cRows)
	if err != nil {
		return padeCoefficients{}, err
	}

	// Build n×1 RHS: d[i] = -c_{m+i+1}
	dRows := make([][]*big.Float, n)
	for i := range n {
		dRows[i] = []*big.Float{bu.PrecFloat().Neg(taylorCoeff(m + i + 1))}
	}
	dMat, err := mat.NewBigMatrix[*big.Float](dRows)
	if err != nil {
		return padeCoefficients{}, err
	}

	inv, singular, err := cMat.Inverse()
	if err != nil {
		return padeCoefficients{}, err
	}
	if singular {
		return padeCoefficients{}, errors.New("pade: coefficient matrix is singular")
	}

	qMat, err := inv.Mul(dMat)
	if err != nil {
		return padeCoefficients{}, err
	}

	// q[0]=1 (normalization), q[1..n] from solver
	qRows := qMat.GetRows()
	q := make([]*big.Float, n+1)
	q[0] = bu.PrecFloat().SetInt64(1)
	for i := range n {
		q[i+1] = qRows[i][0]
	}

	// p[k] = c_k + sum_{j=1}^{min(k,n)} c_{k-j}*q[j], k=0..m
	p := make([]*big.Float, m+1)
	p[0] = bu.PrecFloat().SetInt64(0)
	for k := 1; k <= m; k++ {
		pk := taylorCoeff(k)
		for j := 1; j <= min(k, n); j++ {
			term := bu.PrecFloat().Mul(taylorCoeff(k-j), q[j])
			pk = bu.PrecFloat().Add(pk, term)
		}
		p[k] = pk
	}

	return padeCoefficients{p: p, q: q}, nil
}

func evaluate(coeffs padeCoefficients, u *big.Float) *big.Float {
	m := len(coeffs.p) - 1
	n := len(coeffs.q) - 1

	p := bu.PrecFloat().Copy(coeffs.p[m])
	for k := m - 1; k >= 0; k-- {
		p = bu.PrecFloat().Add(bu.PrecFloat().Mul(p, u), coeffs.p[k])
	}

	q := bu.PrecFloat().Copy(coeffs.q[n])
	for k := n - 1; k >= 0; k-- {
		q = bu.PrecFloat().Add(bu.PrecFloat().Mul(q, u), coeffs.q[k])
	}

	return bu.PrecFloat().Quo(p, q)
}

var cache = &coefficientCache{data: make(map[cacheKey]padeCoefficients)}

var (
	ln2Mu     sync.Mutex
	ln2Cached *big.Float
)

func getLn2(prec uint) (*big.Float, error) {
	ln2Mu.Lock()
	defer ln2Mu.Unlock()
	if ln2Cached != nil && ln2Cached.Prec() >= prec {
		return bu.PrecFloat(prec).Set(ln2Cached), nil
	}
	v, err := reduceLn(bu.PrecFloat(prec).SetInt64(2))
	if err != nil {
		return nil, err
	}
	ln2Cached = v
	return bu.PrecFloat(prec).Set(v), nil
}

func reduceLn(z *big.Float) (*big.Float, error) {
	prec := z.Prec()
	pf := func() *big.Float { return bu.PrecFloat(prec) }

	lo := pf().SetFloat64(0.9)
	hi := pf().SetFloat64(1.1)
	z = pf().Set(z)
	scale := 0
	for z.Cmp(lo) < 0 {
		z = pf().Sqrt(z)
		scale++
	}
	for z.Cmp(hi) >= 0 {
		z = pf().Sqrt(z)
		scale++
	}

	n := orderForPrec(prec)
	coeffs, ok := cache.get(n, n)
	if !ok {
		var err error
		coeffs, err = solveCoefficients(n, n)
		if err != nil {
			return nil, err
		}
		cache.set(n, n, coeffs)
	}

	u := pf().Sub(z, pf().SetInt64(1))
	result := evaluate(coeffs, u)
	multiplier := pf().SetMantExp(pf().SetInt64(1), scale)
	return pf().Mul(result, multiplier), nil
}

// MantExp pre-decomposition bounds the sqrt count in reduceLn to at most ~4
// iterations for any input magnitude.
func ApproximateLn(x *big.Float) (*big.Float, error) {
	if x.Sign() <= 0 {
		return nil, errors.New("pade: argument must be positive")
	}
	prec := x.Prec()
	mant := bu.PrecFloat(prec)
	exp := x.MantExp(mant)

	lnMant, err := reduceLn(mant)
	if err != nil {
		return nil, err
	}
	if exp == 0 {
		return lnMant, nil
	}

	ln2, err := getLn2(prec)
	if err != nil {
		return nil, err
	}
	pf := func() *big.Float { return bu.PrecFloat(prec) }
	return pf().Add(lnMant, pf().Mul(pf().SetInt64(int64(exp)), ln2)), nil
}
