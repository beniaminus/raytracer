package tuple

import "math"

// Tuple is a structure to hold the X,Y,Z co-ordinates and a flag indicating if this is a vector or point.
type Tuple struct {
	X float64

	Y float64

	Z float64

	W float64
}

// Color is a structure to hold red green and blue values for a pixel.
type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

const epsilon = 0.00001

func approximatelyEqual(a, b float64) bool {
	if math.Abs(a-b) < epsilon {
		return true
	}
	return false
}

// New returns a Tuple initialised with x, y, z and w values.
func New(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

// Point returns a Tuple initialised with x, y, z and a point flag.
func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

// Vector returns a Tuple initialised with x, y, z and a vector flag.
func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

// IsPoint returns true if the tuple represents a point
func IsPoint(a Tuple) bool {
	return a.W == 1.0
}

// IsVector returns true if the tuple represents a vector
func IsVector(a Tuple) bool {
	return a.W == 0.0
}

// IsEqual compares two tuples and returns true if they are equal
func IsEqual(a, b Tuple) bool {
	return (approximatelyEqual(a.X, b.X) &&
		approximatelyEqual(a.Y, b.Y) &&
		approximatelyEqual(a.Z, b.Z) &&
		approximatelyEqual(a.W, b.W))
}

// Plus adds two tuples and returns the result
func Plus(a, b Tuple) Tuple {
	return New(a.X+b.X, a.Y+b.Y, a.Z+b.Z, a.W+b.W)
}

// Minus subtracts two tuples and returns the result
func Minus(a, b Tuple) Tuple {
	return New(a.X-b.X, a.Y-b.Y, a.Z-b.Z, a.W-b.W)
}

// Scale returns the scaled value of a tuple
func Scale(a Tuple, scalar float64) Tuple {
	return New(scalar*a.X, scalar*a.Y, scalar*a.Z, scalar*a.W)
}

// Negate returns the negated value of a tuple
func Negate(a Tuple) Tuple {
	return Scale(a, -1)
}

// Divide returns the fractionally scaled value of a tuple
func Divide(a Tuple, scalar float64) Tuple {
	return Scale(a, (1.0 / scalar))
}

// Magnitude returns the length (magnitude) of a vector
func Magnitude(a Tuple) float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W)
}

// Normalize returns the vector normalized to unit size
func Normalize(v Tuple) Tuple {
	var magnitude float64
	magnitude = Magnitude(v)
	return New(v.X/magnitude, v.Y/magnitude, v.Z/magnitude, v.W/magnitude)
}

// Dot returns the product of two vectors
func Dot(v, w Tuple) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z + v.W*w.W
}

// Cross returns the result of a cross product between two vectors
func Cross(v, w Tuple) Tuple {
	return Vector(v.Y*w.Z-v.Z*w.Y,
		v.Z*w.X-v.X*w.Z,
		v.X*w.Y-v.Y*w.X)
}
