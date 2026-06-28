package homework

import "testing"

func TestBuildGreeting(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{name: "Мария", want: "Привет, Мария! Добро пожаловать в Go."},
		{name: "Игорь", want: "Привет, Игорь! Добро пожаловать в Go."},
		{name: "Ринат", want: "Привет, Ринат! Добро пожаловать в Go."},
		{name: "Анна-Мария", want: "Привет, Анна-Мария! Добро пожаловать в Go."},
		{name: "Go Team", want: "Привет, Go Team! Добро пожаловать в Go."},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := BuildGreeting(tt.name)
			if got != tt.want {
				t.Fatalf("BuildGreeting(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestIsAdult(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		age  int
		want bool
	}{
		{name: "negative age", age: -1, want: false},
		{name: "zero age", age: 0, want: false},
		{name: "child", age: 7, want: false},
		{name: "teenager", age: 16, want: false},
		{name: "less than 18", age: 17, want: false},
		{name: "exactly 18", age: 18, want: true},
		{name: "greater than 18", age: 26, want: true},
		{name: "older adult", age: 80, want: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := IsAdult(tt.age)
			if got != tt.want {
				t.Fatalf("IsAdult(%d) = %v, want %v", tt.age, got, tt.want)
			}
		})
	}
}

func TestCalculateTotal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		price int
		count int
		want  int
	}{
		{name: "one item", price: 100, count: 1, want: 100},
		{name: "three items", price: 100, count: 3, want: 300},
		{name: "two items", price: 250, count: 2, want: 500},
		{name: "zero price", price: 0, count: 5, want: 0},
		{name: "zero count", price: 500, count: 0, want: 0},
		{name: "large total", price: 999, count: 10, want: 9990},
		{name: "expensive item", price: 12500, count: 4, want: 50000},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := CalculateTotal(tt.price, tt.count)
			if got != tt.want {
				t.Fatalf("CalculateTotal(%d, %d) = %d, want %d", tt.price, tt.count, got, tt.want)
			}
		})
	}
}

func TestFormatCourseProgress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		completed int
		total     int
		want      string
	}{
		{name: "not started", completed: 0, total: 12, want: "Пройдено 0 из 12 уроков"},
		{name: "first lesson", completed: 1, total: 12, want: "Пройдено 1 из 12 уроков"},
		{name: "middle", completed: 3, total: 12, want: "Пройдено 3 из 12 уроков"},
		{name: "half", completed: 6, total: 12, want: "Пройдено 6 из 12 уроков"},
		{name: "completed", completed: 12, total: 12, want: "Пройдено 12 из 12 уроков"},
		{name: "single lesson", completed: 1, total: 1, want: "Пройдено 1 из 1 уроков"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := FormatCourseProgress(tt.completed, tt.total)
			if got != tt.want {
				t.Fatalf(
					"FormatCourseProgress(%d, %d) = %q, want %q",
					tt.completed,
					tt.total,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCanStartBackendBlock(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		goCoreCompleted bool
		homeworkDone    bool
		want            bool
	}{
		{name: "everything done", goCoreCompleted: true, homeworkDone: true, want: true},
		{name: "homework missing", goCoreCompleted: true, homeworkDone: false, want: false},
		{name: "core missing", goCoreCompleted: false, homeworkDone: true, want: false},
		{name: "nothing done", goCoreCompleted: false, homeworkDone: false, want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := CanStartBackendBlock(tt.goCoreCompleted, tt.homeworkDone)
			if got != tt.want {
				t.Fatalf(
					"CanStartBackendBlock(%v, %v) = %v, want %v",
					tt.goCoreCompleted,
					tt.homeworkDone,
					got,
					tt.want,
				)
			}
		})
	}
}
