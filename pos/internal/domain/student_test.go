package domain

import "testing"

func TestNewStudent(t *testing.T) {
	student, err := NewStudent("001", "John")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if student.ID() != "001" {
		t.Errorf("expected ID 001, got %s", student.ID())
	}

	if student.Name() != "John" {
		t.Errorf("expected name John, got %s", student.Name())
	}
}

func TestNewStudent_EmptyID(t *testing.T) {
	_, err := NewStudent("", "John")

	if err == nil {
		t.Fatal("expected error for empty ID")
	}
}

func TestNewStudent_EmptyName(t *testing.T) {
	_, err := NewStudent("001", "")

	if err == nil {
		t.Fatal("expected error for empty name")
	}
}

func TestNewStudent_EmptyIDAndName(t *testing.T) {
	_, err := NewStudent("      ", "     ")

	if err == nil {
		t.Fatal("expected error for empty ID and name")
	}
}

func TestNewStudent_TrimsWhitespace(t *testing.T) {
	student, err := NewStudent("  001  ", "  Ivan  ")

	if err != nil {
		t.Fatal(err)
	}

	if student.ID() != "001" {
		t.Errorf("expected ID 001, got %q", student.ID())
	}

	if student.Name() != "Ivan" {
		t.Errorf("expected name Ivan, got %q", student.Name())
	}
}
