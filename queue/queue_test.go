package queue

import (
	"testing"
)

func TestQueueOperations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		operations     []func(q *Queue[int])
		wantSize       uint32
		wantFront      int
		expectFrontErr bool
		wantBack       int
		expectBackErr  bool
	}{
		{
			name:           "empty queue",
			operations:     []func(q *Queue[int]){},
			wantSize:       0,
			expectFrontErr: true,
			expectBackErr:  true,
		},
		{
			name: "push single element",
			operations: []func(q *Queue[int]){
				func(q *Queue[int]) { q.Push(42) },
			},
			wantSize:       1,
			wantFront:      42,
			expectFrontErr: false,
			wantBack:       42,
			expectBackErr:  false,
		},
		{
			name: "push multiple elements",
			operations: []func(q *Queue[int]){
				func(q *Queue[int]) { q.Push(10) },
				func(q *Queue[int]) { q.Push(20) },
				func(q *Queue[int]) { q.Push(30) },
			},
			wantSize:       3,
			wantFront:      10,
			expectFrontErr: false,
			wantBack:       30,
			expectBackErr:  false,
		},
		{
			name: "push and pop one element",
			operations: []func(q *Queue[int]){
				func(q *Queue[int]) { q.Push(55) },
				func(q *Queue[int]) { q.Pop() },
			},
			wantSize:       0,
			expectFrontErr: true,
			expectBackErr:  true,
		},
		{
			name: "push multiple and pop one",
			operations: []func(q *Queue[int]){
				func(q *Queue[int]) { q.Push(5) },
				func(q *Queue[int]) { q.Push(15) },
				func(q *Queue[int]) { q.Push(25) },
				func(q *Queue[int]) { q.Pop() },
			},
			wantSize:       2,
			wantFront:      15,
			expectFrontErr: false,
			wantBack:       25,
			expectBackErr:  false,
		},
		{
			name: "pop from empty queue",
			operations: []func(q *Queue[int]){
				func(q *Queue[int]) { q.Pop() },
			},
			wantSize:       0,
			expectFrontErr: true,
			expectBackErr:  true,
		},
		{
			name: "mixed push and pop",
			operations: []func(q *Queue[int]){
				func(q *Queue[int]) { q.Push(1) },
				func(q *Queue[int]) { q.Push(2) },
				func(q *Queue[int]) { q.Pop() },
				func(q *Queue[int]) { q.Push(3) },
			},
			wantSize:       2,
			wantFront:      2,
			expectFrontErr: false,
			wantBack:       3,
			expectBackErr:  false,
		},
		{
			name: "push zero",
			operations: []func(q *Queue[int]){
				func(q *Queue[int]) { q.Push(0) },
			},
			wantSize:       1,
			wantFront:      0,
			expectFrontErr: false,
			wantBack:       0,
			expectBackErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New[int]()
			for _, op := range tt.operations {
				op(&q)
			}

			if q.Size() != tt.wantSize {
				t.Errorf("Size() = %d, want %d", q.Size(), tt.wantSize)
			}

			front, err := q.Front()
			if tt.expectFrontErr {
				if err == nil {
					t.Errorf("Front() expected error, got nil")
				}
				if front != 0 {
					t.Errorf("Front() with error returned %d, want zero value 0", front)
				}
			} else {
				if err != nil {
					t.Errorf("Front() unexpected error: %v", err)
				} else if front != tt.wantFront {
					t.Errorf("Front() = %d, want %d", front, tt.wantFront)
				}
			}

			back, err := q.Back()
			if tt.expectBackErr {
				if err == nil {
					t.Errorf("Back() expected error, got nil")
				}
				if back != 0 {
					t.Errorf("Back() with error returned %d, want zero value 0", back)
				}
			} else {
				if err != nil {
					t.Errorf("Back() unexpected error: %v", err)
				} else if back != tt.wantBack {
					t.Errorf("Back() = %d, want %d", back, tt.wantBack)
				}
			}
		})
	}
}
