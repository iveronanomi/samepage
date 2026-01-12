package domain

import (
	"bytes"
	"fmt"
	"sync"
)

// Vector is a transient aggregate state.
// It has no identity and no fixed meaning.
type Vector struct {
	_vec
	rw sync.RWMutex
}

type _vec [_axis]*float64

// ZeroVector represents emptiness (śūnyatā).
func ZeroVector() Vector {
	return Vector{}
}

// Set assigns a value to an axis.
// Zero values are still stored intentionally.
func (v *Vector) Set(a int, value float64) {
	if a < 0 || a > int(_axis) {
		return
	}
	val := value
	v._vec[a] = &val
}

// Get returns value for an axis (0 if absent).
func (v *Vector) Get(a int) *float64 {
	v.rw.RLock()
	defer v.rw.RUnlock()

	return v._vec[a]
}

// Add applies delta to the _vec in-place.
func (v *Vector) Add(delta *Vector) {
	for a := 0; a < int(_axis); a++ {
		if delta._vec[a] == nil {
			continue
		}
		if v._vec[a] == nil {
			val := *delta._vec[a]
			v._vec[a] = &val
		} else {
			*v._vec[a] += *delta._vec[a]
		}
	}
}

// Scaled returns a scaled copy (decay/moderation).
func (v *Vector) Scaled(f float64) *Vector {
	out := ZeroVector()
	v.Each(func(i int, val *float64) {
		if val == nil {
			return
		}
		value := *val * f
		out._vec[i] = &value
	})
	return &out
}

func (v *Vector) Each(f func(i int, val *float64)) {
	v.each(func(i int) {
		f(i, v._vec[i])
	})
}

func (v *Vector) each(f func(i int)) {
	v.rw.Lock()
	defer v.rw.Unlock()

	for i := 0; i < int(_axis); i++ {
		f(i)
	}
}

// Clone returns a deep copy of the _vec.
func (v *Vector) Clone() *Vector {
	var out = Ptr(ZeroVector())
	v.each(func(i int) {
		val := v._vec[i]
		if val == nil {
			return
		}
		c := *val
		out._vec[i] = &c
	})
	return out
}

// Sub returns v - other.
func (v *Vector) Sub(other *Vector) *Vector {
	var out = v.Clone()
	out.each(func(i int) {
		val := other._vec[i]
		if val == nil {
			return
		}
		if out._vec[i] == nil {
			c := -(*val)
			out._vec[i] = &c
			return
		}
		*out._vec[i] -= *val
	})
	return out
}

// Dot computes the dot product between two vectors.
// This measures directional alignment only.
func (v *Vector) Dot(vec *Vector) float64 {
	var out float64
	v.each(func(i int) {
		if v._vec[i] == nil || vec._vec[i] == nil {
			return
		}
		out += (*v._vec[i]) * (*vec._vec[i])
	})
	return out
}

func (v *Vector) IsZero() bool {
	for _, val := range v._vec {
		if val != nil && *val != 0 {
			return false
		}
	}
	return true
}

// Magnitude returns L1 magnitude of non-nil axes.
// Used to detect whether a change is meaningful.
func (v *Vector) Magnitude() float64 {
	var out float64
	v.each(func(i int) {
		if v._vec[i] == nil {
			return
		}
		if *v._vec[i] < 0 {
			out -= *v._vec[i]
			return
		}
		out += *v._vec[i]
	})
	return out
}

// Clamp limits all _vec components into [min, max].
// This enforces a bounded state space.
func (v *Vector) Clamp(i int, min, max float64) {
	if v._vec[i] == nil {
		return
	}
	if *v._vec[i] < min {
		*v._vec[i] = min
	}
	if *v._vec[i] > max {
		*v._vec[i] = max
	}
}

const (
	DefaultNormalizedVectorMin = -1.0
	DefaultNormalizedVectorMax = 1.0
)

// Normalize clamps all _vec components into [ DefaultNormalizedVectorMin, DefaultNormalizedVectorMax ].
func (v *Vector) Normalize() {
	v.each(func(i int) {
		v.Clamp(i, DefaultNormalizedVectorMin, DefaultNormalizedVectorMax)
	})
}

func (v *Vector) Debug() string {
	var b bytes.Buffer
	b.WriteString("vector:\n")
	v.each(func(i int) {
		if val := v._vec[i]; val != nil {
			b.WriteString(fmt.Sprintf("  %-14s  %+.2f\n", Axis(i).String(), *val))
		}
	})
	return b.String()
}
