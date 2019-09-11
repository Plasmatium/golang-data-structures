package rbTree

import "testing"

func TestIsRoot(t *testing.T) {
	root := NewRoot(100)
	_ = root.Insert(15)

	if root.isRoot() != true {
		t.Error("failed: root.isRoot() is false")
	}

	if root.Left.Data != 15 {
		t.Errorf("failed: root.Left.Data should be 15, now is %v", root.Left.Data)
	}

	if root.Left.isRoot() == true {
		t.Error("failed: root.Left.isRoot() is true")
	}

	if root.Right.isRoot() == true {
		t.Error("failed: root.Right.isRoot() is true")
	}
}

func TestDL(t *testing.T) {
	r := NewRoot(100)
	t.Log(r.isRoot() || r.Parent.Parent.Parent.isRoot())
}