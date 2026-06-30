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

func TestBuildCourseWelcome(t *testing.T) {
	t.Parallel()

	tests := []struct {
		courseName string
		want       string
	}{
		{courseName: "Go backend", want: "Курс: Go backend"},
		{courseName: "Основы Go", want: "Курс: Основы Go"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.courseName, func(t *testing.T) {
			t.Parallel()

			got := BuildCourseWelcome(tt.courseName)
			if got != tt.want {
				t.Fatalf("BuildCourseWelcome(%q) = %q, want %q", tt.courseName, got, tt.want)
			}
		})
	}
}

func TestBuildLessonTitle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		lessonName string
		want       string
	}{
		{lessonName: "терминал, Git и первый Go-проект", want: "Урок 1: терминал, Git и первый Go-проект"},
		{lessonName: "package, import и fmt", want: "Урок 1: package, import и fmt"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.lessonName, func(t *testing.T) {
			t.Parallel()

			got := BuildLessonTitle(tt.lessonName)
			if got != tt.want {
				t.Fatalf("BuildLessonTitle(%q) = %q, want %q", tt.lessonName, got, tt.want)
			}
		})
	}
}

func TestBuildRepositoryPath(t *testing.T) {
	t.Parallel()

	tests := []struct {
		owner string
		repo  string
		want  string
	}{
		{owner: "rinat-course", repo: "homework1", want: "github.com/rinat-course/homework1"},
		{owner: "student", repo: "go-homework", want: "github.com/student/go-homework"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.owner+"/"+tt.repo, func(t *testing.T) {
			t.Parallel()

			got := BuildRepositoryPath(tt.owner, tt.repo)
			if got != tt.want {
				t.Fatalf("BuildRepositoryPath(%q, %q) = %q, want %q", tt.owner, tt.repo, got, tt.want)
			}
		})
	}
}

func TestBuildRunCommand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		packagePath string
		want        string
	}{
		{packagePath: "./cmd/demo", want: "go run ./cmd/demo"},
		{packagePath: "./cmd/app", want: "go run ./cmd/app"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.packagePath, func(t *testing.T) {
			t.Parallel()

			got := BuildRunCommand(tt.packagePath)
			if got != tt.want {
				t.Fatalf("BuildRunCommand(%q) = %q, want %q", tt.packagePath, got, tt.want)
			}
		})
	}
}
