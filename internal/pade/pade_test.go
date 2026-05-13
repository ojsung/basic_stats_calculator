package pade

import (
	"math/big"
	"testing"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
)

func Test_solveCoefficients(t *testing.T) {
	t.Run("[2/2] matches known analytical solution", func(t *testing.T) {
		coeffs, err := solveCoefficients(2, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		cases := []struct {
			name     string
			got      *big.Float
			expected string
		}{
			{"p[0]", coeffs.p[0], "0.0000000000000000"},
			{"p[1]", coeffs.p[1], "1.0000000000000000"},
			{"p[2]", coeffs.p[2], "0.5000000000000000"},
			{"q[0]", coeffs.q[0], "1.0000000000000000"},
			{"q[1]", coeffs.q[1], "1.0000000000000000"},
			{"q[2]", coeffs.q[2], "0.1666666666666667"},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				got := tc.got.Text('f', 16)
				if got != tc.expected {
					t.Errorf("got %v, want %v", got, tc.expected)
				}
			})
		}
	})
}

func Test_coefficientCache(t *testing.T) {
	t.Run("returns false on cache miss", func(t *testing.T) {
		c := &coefficientCache{data: make(map[cacheKey]padeCoefficients)}
		_, ok := c.get(3, 3)
		if ok {
			t.Fatal("expected cache miss, got hit")
		}
	})
	t.Run("returns stored coefficients after set", func(t *testing.T) {
		c := &coefficientCache{data: make(map[cacheKey]padeCoefficients)}
		stored := padeCoefficients{
			p: []*big.Float{big.NewFloat(0), big.NewFloat(1)},
			q: []*big.Float{big.NewFloat(1), big.NewFloat(2)},
		}
		c.set(2, 2, stored)
		got, ok := c.get(2, 2)
		if !ok {
			t.Fatal("expected cache hit, got miss")
		}
		if len(got.p) != 2 || len(got.q) != 2 {
			t.Fatalf("unexpected lengths: p=%d q=%d", len(got.p), len(got.q))
		}
	})
	t.Run("different keys are independent", func(t *testing.T) {
		c := &coefficientCache{data: make(map[cacheKey]padeCoefficients)}
		c.set(2, 2, padeCoefficients{p: []*big.Float{big.NewFloat(1)}, q: []*big.Float{big.NewFloat(1)}})
		_, ok := c.get(3, 3)
		if ok {
			t.Fatal("key (3,3) should be absent after setting (2,2)")
		}
	})
}

func Test_evaluate(t *testing.T) {
	t.Run("evaluates P(u)/Q(u) via Horner at u=0.5 with [2/2] coefficients", func(t *testing.T) {
		// [2/2] Padé for ln(1+u): P(u)=u+u²/2, Q(u)=1+u+u²/6
		// P(0.5)/Q(0.5) = (5/8) / (37/24) = 15/37 = 0.405405405405...
		q2 := new(big.Float).SetPrec(256).Quo(big.NewFloat(1), big.NewFloat(6))
		coeffs := padeCoefficients{
			p: []*big.Float{
				big.NewFloat(0),   // p[0]
				big.NewFloat(1),   // p[1]
				big.NewFloat(0.5), // p[2]
			},
			q: []*big.Float{
				big.NewFloat(1), // q[0]
				big.NewFloat(1), // q[1]
				q2,              // q[2] = 1/6
			},
		}
		u := big.NewFloat(0.5)
		got := evaluate(coeffs, u).Text('f', 16)
		want := "0.4054054054054054"
		if got != want {
			t.Errorf("evaluate() = %v, want %v", got, want)
		}
	})
}

func Test_orderForPrec(t *testing.T) {
	cases := []struct {
		name string
		prec uint
		want int
	}{
		{"64-bit", 64, 8},
		{"128-bit", 128, 15},
		{"256-bit", 256, 30},
		{"9-bit", 9, 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := orderForPrec(tc.prec); got != tc.want {
				t.Errorf("orderForPrec(%d) = %d, want %d", tc.prec, got, tc.want)
			}
		})
	}
}

func Test_reduceLn(t *testing.T) {
	t.Run("ln(0.5) low-edge reduction", func(t *testing.T) {
		got, err := reduceLn(bu.StrToFloat("0.5"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Text('f', 16) != "-0.6931471805599453" {
			t.Errorf("got %v, want -0.6931471805599453", got.Text('f', 16))
		}
	})
	t.Run("ln(1.5) hi-edge reduction", func(t *testing.T) {
		got, err := reduceLn(bu.StrToFloat("1.5"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Text('f', 16) != "0.4054651081081644" {
			t.Errorf("got %v, want 0.4054651081081644", got.Text('f', 16))
		}
	})
	t.Run("ln(2.0) hi-edge reduction", func(t *testing.T) {
		got, err := reduceLn(bu.StrToFloat("2.0"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Text('f', 16) != "0.6931471805599453" {
			t.Errorf("got %v, want 0.6931471805599453", got.Text('f', 16))
		}
	})
	t.Run("ln(0.9) at lo boundary", func(t *testing.T) {
		got, err := reduceLn(bu.StrToFloat("0.9"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Text('f', 16) != "-0.1053605156578263" {
			t.Errorf("got %v, want -0.1053605156578263", got.Text('f', 16))
		}
	})
	t.Run("ln(2.0) 256-bit precision to 22 decimal places", func(t *testing.T) {
		// Old [12/12] at u=1.0 gives ~15 digits. Dynamic [30/30] at |u|≤0.1 gives ~77.
		z := new(big.Float).SetPrec(256).SetFloat64(2.0)
		got, err := reduceLn(z)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "0.6931471805599453094172"
		if got.Text('f', 22) != want {
			t.Errorf("got %v, want %v", got.Text('f', 22), want)
		}
	})
}

func Test_ApproximateLn(t *testing.T) {
	tests := []struct {
		name     string
		x        string
		expected string
		wantErr  bool
	}{
		{"x=0 returns error", "0", "", true},
		{"x=0.1", "0.1", "-2.3025850929940457", false},
		{"x=0.05", "0.05", "-2.9957322735539910", false},
		{"x=1.9", "1.9", "0.6418538861723948", false},
		{"x=2.0", "2.0", "0.6931471805599453", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ApproximateLn(bu.StrToFloat(tt.x))
			if (err != nil) != tt.wantErr {
				t.Fatalf("ApproximateLn(%v) error = %v, wantErr %v", tt.x, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			got := result.Text('f', 16)
			if got != tt.expected {
				t.Errorf("ApproximateLn(%v) = %v, want %v", tt.x, got, tt.expected)
			}
		})
	}
}

func Test_ApproximateLn_MantExp(t *testing.T) {
	prec := uint(256)
	ln2Str := "0.69314718055994530941723212145817656807550013436025"

	t.Run("x=2^(-101) extreme small", func(t *testing.T) {
		x := new(big.Float).SetPrec(prec).SetMantExp(
			new(big.Float).SetPrec(prec).SetFloat64(0.5), -100,
		)
		ln2Ref, _, _ := new(big.Float).SetPrec(prec).Parse(ln2Str, 10)
		want := new(big.Float).SetPrec(prec).Mul(
			new(big.Float).SetPrec(prec).SetInt64(-101), ln2Ref,
		)
		result, err := ApproximateLn(x)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Text('f', 16) != want.Text('f', 16) {
			t.Errorf("got %v, want %v", result.Text('f', 16), want.Text('f', 16))
		}
	})

	t.Run("x=2^101 extreme large", func(t *testing.T) {
		x := new(big.Float).SetPrec(prec).SetMantExp(
			new(big.Float).SetPrec(prec).SetFloat64(0.5), 102,
		)
		ln2Ref, _, _ := new(big.Float).SetPrec(prec).Parse(ln2Str, 10)
		want := new(big.Float).SetPrec(prec).Mul(
			new(big.Float).SetPrec(prec).SetInt64(101), ln2Ref,
		)
		result, err := ApproximateLn(x)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Text('f', 16) != want.Text('f', 16) {
			t.Errorf("got %v, want %v", result.Text('f', 16), want.Text('f', 16))
		}
	})
}

func Test_getLn2(t *testing.T) {
	t.Run("returns ln(2) to 16 decimal places", func(t *testing.T) {
		v, err := getLn2(256)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got := v.Text('f', 16); got != "0.6931471805599453" {
			t.Errorf("getLn2(256) = %v, want 0.6931471805599453", got)
		}
	})
	t.Run("cache hit returns consistent result", func(t *testing.T) {
		v1, _ := getLn2(256)
		v2, _ := getLn2(256)
		if v1.Text('f', 16) != v2.Text('f', 16) {
			t.Errorf("getLn2 not idempotent: %v vs %v", v1.Text('f', 16), v2.Text('f', 16))
		}
	})
	t.Run("higher precision after lower precision cache still correct", func(t *testing.T) {
		getLn2(64)
		v, err := getLn2(256)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got := v.Text('f', 16); got != "0.6931471805599453" {
			t.Errorf("getLn2(256) after 64-bit prime = %v, want 0.6931471805599453", got)
		}
	})
}
