package stack

import (
	"testing"
)

func TestStackOperations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		operations    []func(s *Stack[int])
		wantLen       uint32
		wantTop       int
		expectTopErr  bool
		expectIsEmpty bool
	}{
		{
			name: "push single element",
			operations: []func(s *Stack[int]){
				func(s *Stack[int]) { s.Push(42) },
			},
			wantLen:       1,
			wantTop:       42,
			expectTopErr:  false,
			expectIsEmpty: false,
		},
		{
			name: "push multiple elements",
			operations: []func(s *Stack[int]){
				func(s *Stack[int]) { s.Push(10) },
				func(s *Stack[int]) { s.Push(20) },
				func(s *Stack[int]) { s.Push(30) },
			},
			wantLen:       3,
			wantTop:       30,
			expectTopErr:  false,
			expectIsEmpty: false,
		},
		{
			name: "push and pop one element",
			operations: []func(s *Stack[int]){
				func(s *Stack[int]) { s.Push(55) },
				func(s *Stack[int]) { s.Pop() },
			},
			wantLen:       0,
			expectTopErr:  true,
			expectIsEmpty: true,
		},
		{
			name: "push multiple and pop one",
			operations: []func(s *Stack[int]){
				func(s *Stack[int]) { s.Push(5) },
				func(s *Stack[int]) { s.Push(15) },
				func(s *Stack[int]) { s.Push(25) },
				func(s *Stack[int]) { s.Pop() },
			},
			wantLen:       2,
			wantTop:       15,
			expectTopErr:  false,
			expectIsEmpty: false,
		},
		{
			name: "pop from empty stack",
			operations: []func(s *Stack[int]){
				func(s *Stack[int]) { s.Pop() },
			},
			wantLen:       0,
			expectTopErr:  true,
			expectIsEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New[int]()
			for _, op := range tt.operations {
				op(&s)
			}

			if s.Len() != tt.wantLen {
				t.Errorf("Len() = %d, want %d", s.Len(), tt.wantLen)
			}

			if isStackEmpty := s.Empty(); isStackEmpty != tt.expectIsEmpty {
				t.Errorf("Expected %v, got %v", tt.expectIsEmpty, isStackEmpty)
			}

			top, err := s.Top()
			if tt.expectTopErr {
				if err == nil {
					t.Errorf("Top() expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Top() unexpected error: %v", err)
				} else if top != tt.wantTop {
					t.Errorf("Top() = %d, want %d", top, tt.wantTop)
				}
			}
		})
	}
}
