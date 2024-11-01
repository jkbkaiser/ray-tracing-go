package vec3

import "testing"

func TestVectorAdd(t *testing.T) {
	v1 := Vec3{1, 2, 3}
	v2 := Vec3{0, -2, 100}
	expected := Vec3{1, 0, 103}

	v3 := v1.Add(v2)

	if v3 != expected {
		t.Errorf("Expected %v, but got %v", expected, v3)
	}
}

func TestVectorSubtract(t *testing.T) {
	v1 := Vec3{1, 2, 3}
	v2 := Vec3{1, -1, 4}
	expected := Vec3{0, 3, -1}

	v3 := v1.Subtract(v2)

	if v3 != expected {
		t.Errorf("Expected %v, but got %v", expected, v3)
	}
}

func TestVectorDot(t *testing.T) {
	v1 := Vec3{1, 2, 3}
	v2 := Vec3{-4, 1, 4}
	expected := 10.

	v3 := v1.Dot(v2)

	if v3 != expected {
		t.Errorf("Expected %v, but got %v", expected, v3)
	}
}

func TestVectorCross(t *testing.T) {
	v1 := Vec3{2, 3, 4}
	v2 := Vec3{5, 6, 7}
	expected := Vec3{-3, 6, -3}

	v3 := v1.Cross(v2)

	if v3 != expected {
		t.Errorf("Expected %v, but got %v", expected, v3)
	}
}

func TestVectorNorm(t *testing.T) {
	v1 := Vec3{10, 0, 0}
	expected := Vec3{1, 0, 0}

	v3 := v1.Norm()

	if v3 != expected {
		t.Errorf("Expected %v, but got %v", expected, v3)
	}
}

func TestVectorScale(t *testing.T) {
	v1 := Vec3{1, 2, 3}
	v2 := -2.
	expected := Vec3{-2, -4, -6}

	v3 := v1.Scale(v2)

	if v3 != expected {
		t.Errorf("Expected %v, but got %v", expected, v3)
	}
}
