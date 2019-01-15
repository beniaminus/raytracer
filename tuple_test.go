package main

import (
	"fmt"
	"math"
	"raytracer/tuple"

	"github.com/DATA-DOG/godog"
)

var a,
	b,
	p,
	q,
	v,
	w,
	zero,
	norm tuple.Tuple

func aTuple(arg1, arg2, arg3, arg4 float64) error {
	a = tuple.New(arg1, arg2, arg3, arg4)
	return nil
}

func bTuple(arg1, arg2, arg3, arg4 float64) error {
	b = tuple.New(arg1, arg2, arg3, arg4)
	return nil
}
func ax(expected float64) error {
	if a.X != expected {
		return fmt.Errorf(
			"expected %f but found %f",
			expected,
			a.X,
		)
	}

	return nil
}

func ay(expected float64) error {
	if a.Y != expected {
		return fmt.Errorf(
			"expected %f but found %f",
			expected,
			a.Y,
		)
	}

	return nil
}

func az(expected float64) error {
	if a.Z != expected {
		return fmt.Errorf(
			"expected %f but found %f",
			expected,
			a.Z,
		)
	}

	return nil
}

func aw(expected float64) error {
	if a.W != expected {
		return fmt.Errorf(
			"expected %f but found %f",
			expected,
			a.W,
		)
	}

	return nil
}

func aIsAPoint() error {
	if !tuple.IsPoint(a) {
		return fmt.Errorf("expected a to be a point, but it is not")
	}

	return nil
}

func aIsNotAVector() error {
	if tuple.IsVector(a) {
		return fmt.Errorf("expected a to not be a vector, but it is")
	}

	return nil
}

func aIsNotAPoint() error {
	if tuple.IsPoint(a) {
		return fmt.Errorf("expected a to not be a point, but it is")
	}

	return nil
}

func aIsAVector() error {
	if !tuple.IsVector(a) {
		return fmt.Errorf("expected a to be a vector, but it is not")
	}

	return nil
}

func pPoint(arg1, arg2, arg3 float64) error {
	p = tuple.Point(arg1, arg2, arg3)
	return nil
}

func pIsATuple(arg1, arg2, arg3, arg4 float64) error {
	var expected = tuple.New(arg1, arg2, arg3, arg4)
	return isEqual(p, expected)
}

func vVector(arg1, arg2, arg3 float64) error {
	v = tuple.Vector(arg1, arg2, arg3)
	return nil
}

func vIsATuple(arg1, arg2, arg3, arg4 float64) error {
	var expected = tuple.New(arg1, arg2, arg3, arg4)
	return isEqual(v, expected)
}

func sumAPlusBEqualsTuple(arg1, arg2, arg3, arg4 float64) error {
	var result tuple.Tuple
	var expected tuple.Tuple

	result = tuple.Plus(a, b)

	expected = tuple.New(arg1, arg2, arg3, arg4)

	return isEqual(result, expected)
}

func qPoint(arg1, arg2, arg3 float64) error {
	q = tuple.Point(arg1, arg2, arg3)

	return nil
}

func isEqual(result, expected tuple.Tuple) error {
	if !tuple.IsEqual(result, expected) {
		return fmt.Errorf(
			"expected tuple(%f, %f, %f, %f) but found tuple(%f, %f, %f, %f)",
			expected.X, expected.Y, expected.Z, expected.W,
			result.X, result.Y, result.Z, result.W,
		)
	}

	return nil
}

func minusPAndQEqualsVector(arg1, arg2, arg3 float64) error {
	var result,
		expected tuple.Tuple

	result = tuple.Minus(p, q)

	expected = tuple.Vector(arg1, arg2, arg3)

	return isEqual(result, expected)
}

func minusPandVEqualsPoint(arg1, arg2, arg3 float64) error {
	var result,
		expected tuple.Tuple

	result = tuple.Minus(p, v)

	expected = tuple.Point(arg1, arg2, arg3)

	return isEqual(result, expected)
}

func wVector(arg1, arg2, arg3 float64) error {
	w = tuple.Vector(arg1, arg2, arg3)
	return nil
}

func minusVandWEqualsVector(arg1, arg2, arg3 float64) error {
	var result,
		expected tuple.Tuple

	result = tuple.Minus(v, w)

	expected = tuple.Vector(arg1, arg2, arg3)

	return isEqual(result, expected)
}

func zeroVector(arg1, arg2, arg3 float64) error {
	zero = tuple.Vector(arg1, arg2, arg3)

	return nil
}

func minusZeroAndVEqualsVector(arg1, arg2, arg3 float64) error {
	var result,
		expected tuple.Tuple

	result = tuple.Minus(zero, v)

	expected = tuple.Vector(arg1, arg2, arg3)

	return isEqual(result, expected)
}

func negateATuple(arg1, arg2, arg3, arg4 float64) error {
	var result,
		expected tuple.Tuple

	result = tuple.Negate(a)

	expected = tuple.New(arg1, arg2, arg3, arg4)

	return isEqual(result, expected)
}

func mulitplyAByAScalarEqualsTuple(scalar, arg1, arg2, arg3, arg4 float64) error {
	var result,
		expected tuple.Tuple

	result = tuple.Scale(a, scalar)

	expected = tuple.New(arg1, arg2, arg3, arg4)

	return isEqual(result, expected)
}

func divideByScalarEqualsTuple(scalar, arg1, arg2, arg3, arg4 float64) error {
	var result,
		expected tuple.Tuple

	result = tuple.Divide(a, scalar)

	expected = tuple.New(arg1, arg2, arg3, arg4)

	return isEqual(result, expected)
}

func magnitudeVEquals(expected float64) error {
	var result float64

	result = tuple.Magnitude(v)

	if result != expected {
		return fmt.Errorf(
			"expected %f but found %f",
			expected,
			result,
		)
	}

	return nil
}

func magnitudeVEqualsSqrt(expected float64) error {
	var result float64

	result = tuple.Magnitude(v)

	if result != math.Sqrt(expected) {
		return fmt.Errorf(
			"expected %f but found %f",
			math.Sqrt(expected),
			result,
		)
	}

	return nil
}

func normNormalizeV() error {
	norm = tuple.Normalize(v)
	return nil
}

func normalizeVEqualsVector(arg1, arg2, arg3 float64) error {
	var expected tuple.Tuple

	normNormalizeV()

	expected = tuple.Vector(arg1, arg2, arg3)

	return isEqual(norm, expected)
}

func magnitudeNormEquals(expected float64) error {
	var result float64

	result = tuple.Magnitude(norm)

	if result != expected {
		return fmt.Errorf(
			"expected %f but found %f",
			expected,
			result,
		)
	}

	return nil
}

func dotVAndWEquals(expected float64) error {
	var result float64

	result = tuple.Dot(v, w)

	if result != expected {
		return fmt.Errorf(
			"expected %f but found %f",
			expected,
			result,
		)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a ← tuple\((\d+\.\d+), (\-\d+\.\d+), (\d+\.\d+), (\d+\.\d+)\)$`, aTuple)
	s.Step(`^a\.x = (\d+\.\d+)$`, ax)
	s.Step(`^a\.y = (\-\d+\.\d+)$`, ay)
	s.Step(`^a\.z = (\d+\.\d+)$`, az)
	s.Step(`^a\.w = (\d+\.\d+)$`, aw)
	s.Step(`^a is a point$`, aIsAPoint)
	s.Step(`^a is not a vector$`, aIsNotAVector)
	s.Step(`^a is not a point$`, aIsNotAPoint)
	s.Step(`^a is a vector$`, aIsAVector)
	s.Step(`^p ← point\((\d+\.\d+), (\-\d+\.\d+), (\d+\.\d+)\)$`, pPoint)
	s.Step(`^p = tuple\((\d+\.\d+), (\-\d+\.\d+), (\d+\.\d+), (\d+\.\d+)\)$`, pIsATuple)
	s.Step(`^v ← vector\((\-?\d+\.\d+), (\-?\d+\.\d+), (\-?\d+\.\d+)\)$`, vVector)
	s.Step(`^v = tuple\((\d+\.\d+), (\-\d+\.\d+), (\d+\.\d+), (\d+\.\d+)\)$`, vIsATuple)
	s.Step(`^a ← tuple\((\d+), (\-\d+), (\d+), (\d+)\)$`, aTuple)
	s.Step(`^b ← tuple\((\-\d+), (\d+), (\d+), (\d+)\)$`, bTuple)
	s.Step(`^a \+ b = tuple\((\d+), (\d+), (\d+), (\d+)\)$`, sumAPlusBEqualsTuple)
	s.Step(`^p ← point\((\d+), (\d+), (\d+)\)$`, pPoint)
	s.Step(`^q ← point\((\d+), (\d+), (\d+)\)$`, qPoint)
	s.Step(`^p - q = vector\((\-\d+), (\-\d+), (\-\d+)\)$`, minusPAndQEqualsVector)
	s.Step(`^v ← vector\((\d+), (\d+), (\d+)\)$`, vVector)
	s.Step(`^p - v = point\((\-\d+), (\-\d+), (\-\d+)\)$`, minusPandVEqualsPoint)
	s.Step(`^w ← vector\((\d+), (\d+), (\d+)\)$`, wVector)
	s.Step(`^v - w = vector\((\-\d+), (\-\d+), (\-\d+)\)$`, minusVandWEqualsVector)
	s.Step(`^zero ← vector\((\d+), (\d+), (\d+)\)$`, zeroVector)
	s.Step(`^v ← vector\((\d+), (\-\d+), (\d+)\)$`, vVector)
	s.Step(`^zero - v = vector\((\-\d+), (\d+), (\-\d+)\)$`, minusZeroAndVEqualsVector)
	s.Step(`^a ← tuple\((\d+), (\-\d+), (\d+), (\-\d+)\)$`, aTuple)
	s.Step(`^-a = tuple\((\-\d+), (\d+), (\-\d+), (\d+)\)$`, negateATuple)
	s.Step(`^a \* (\d+\.\d+) = tuple\((\d+\.\d+), (\-\d+\.\d+), (\d+\.\d+), (\-\d+\.\d+)\)$`, mulitplyAByAScalarEqualsTuple)
	s.Step(`^a \/ (\d+) = tuple\((\d+\.\d+), (\-\d+\.\d+), (\d+\.\d+), (\-\d+\.\d+)\)$`, divideByScalarEqualsTuple)
	s.Step(`^magnitude\(v\) = (\d+)$`, magnitudeVEquals)
	s.Step(`^magnitude\(v\) = √(\d+)$`, magnitudeVEqualsSqrt)
	s.Step(`^normalize\(v\) = vector\((\-?\d+\.\d+), (\-?\d+\.\d+), (\-?\d+\.\d+)\)$`, normalizeVEqualsVector)
	s.Step(`^normalize\(v\) = approximately vector\((\-?\d+\.\d+), (\-?\d+\.\d+), (\-?\d+\.\d+)\)$`, normalizeVEqualsVector)
	s.Step(`^norm ← normalize\(v\)$`, normNormalizeV)
	s.Step(`^magnitude\(norm\) = (\d+)$`, magnitudeNormEquals)
	s.Step(`^v ← vector\((\-?\d+\.\d+), (\-?\d+\.\d+), (\-?\d+\.\d+)\)$`, vVector)
	s.Step(`^w ← vector\((\-?\d+\.\d+), (\-?\d+\.\d+), (\-?\d+\.\d+)\)$`, wVector)
	s.Step(`^dot\(v, w\) = (\d+)$`, dotVAndWEquals)
}
