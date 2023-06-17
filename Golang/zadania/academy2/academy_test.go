package academy_test

import (
	"testing"

	academy "github.com/grupawp/akademia-programowania/Golang/zadania/academy2"
	"github.com/grupawp/akademia-programowania/Golang/zadania/academy2/mocks"
	"github.com/pkg/errors"
)

func TestGradeYear(t *testing.T) {
	t.Run("Repository list error", func(t *testing.T) {
		testYear := uint8(1)
		mockRepository := mocks.NewRepository(t)

		listError := errors.New("unexpected error")
		mockRepository.On("List", testYear).Return(nil, listError)

		err := academy.GradeYear(mockRepository, testYear)

		if err == nil {
			t.Errorf("Expected error: %v but got no error", listError)
		}
	})
	t.Run("Grade students error", func(t *testing.T) {
		testYear := uint8(1)

		mockRepository := mocks.NewRepository(t)
		namesList := []string{"test"}

		mockRepository.On("List", testYear).Return(namesList, nil)
		getStudentErr := errors.New("unexpected error")
		mockRepository.On("Get", "test").Return(nil, getStudentErr)

		err := academy.GradeYear(mockRepository, testYear)
		if err == nil {
			t.Errorf("Expected error: %v but got not error", getStudentErr)
		}

	})
	t.Run("Grade students student not found error", func(t *testing.T) {
		testYear := uint8(1)

		mockRepository := mocks.NewRepository(t)
		mockStudent := mocks.NewStudent(t)

		namesList := []string{"test"}
		mockRepository.On("List", testYear).Return(namesList, nil)
		studentNotFoundError := academy.ErrStudentNotFound
		mockRepository.On("Get", "test").Return(mockStudent, studentNotFoundError)

		err := academy.GradeYear(mockRepository, testYear)
		if err != nil {
			t.Errorf("Expected no error but received: %v", err)
		}
	})
	t.Run("Grade students no errors", func(t *testing.T) {
		testYear := uint8(1)

		mockRepository := mocks.NewRepository(t)
		mockStudent := mocks.NewStudent(t)

		namesList := []string{"test"}
		mockRepository.On("List", testYear).Return(namesList, nil)
		mockRepository.On("Get", "test").Return(mockStudent, nil)
		mockRepository.On("Save", "test", uint8(2)).Return(nil)
		mockStudent.On("FinalGrade").Return(3)
		mockStudent.On("Name").Return("test")
		mockStudent.On("Year").Return(uint8(1))

		err := academy.GradeYear(mockRepository, testYear)
		if err != nil {
			t.Errorf("Expected no error but received: %v", err)
		}
	})
}

func TestGradeStudent(t *testing.T) {
	t.Run("Repository get error", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		mockRepository.On("Get", "test").Return(nil, errors.New("unexpected error"))

		err := academy.GradeStudent(mockRepository, "test")

		if err == nil {
			t.Errorf("Expected error: %v but got no error", err)
		}

	})
	t.Run("Repository get StudentNotFound error", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		studentNotFoundError := academy.ErrStudentNotFound
		mockRepository.On("Get", "test").Return(nil, studentNotFoundError)

		err := academy.GradeStudent(mockRepository, "test")

		if err != nil {
			t.Errorf("Expected no error, but got error: %v", err)
		}

	})

	gradesOutOfBoundsTests := []struct {
		name  string
		grade int
	}{
		{"Student grade lower than 1", 0},
		{"Student grade above 5", 6},
	}
	for _, tData := range gradesOutOfBoundsTests {

		t.Run(tData.name, func(t *testing.T) {
			mockRepository := mocks.NewRepository(t)
			mockStudent := mocks.NewStudent(t)

			mockRepository.On("Get", "test").Return(mockStudent, nil)
			mockStudent.On("FinalGrade").Return(tData.grade)

			err := academy.GradeStudent(mockRepository, "test")

			if err == nil {
				t.Errorf("Expected: %v, but got no error.", academy.ErrInvalidGrade)
			}
		})
	}

	t.Run("Student grade 1 Repository save error", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		mockStudent := mocks.NewStudent(t)

		mockStudent.On("Name").Return("test")
		mockStudent.On("Year").Return(uint8(3))
		mockStudent.On("FinalGrade").Return(1)

		mockRepository.On("Get", "test").Return(mockStudent, nil)
		saveError := errors.New("unexpected error")
		mockRepository.On("Save", mockStudent.Name(), mockStudent.Year()).Return(saveError)

		err := academy.GradeStudent(mockRepository, "test")

		if err == nil {
			t.Errorf("Expected: %v, but got no error", saveError)
		}
	})

	t.Run("Student year 3 Repository graduate error", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		mockStudent := mocks.NewStudent(t)

		mockStudent.On("FinalGrade").Return(3)
		mockStudent.On("Year").Return(uint8(3))
		mockStudent.On("Name").Return("test")

		mockRepository.On("Get", "test").Return(mockStudent, nil)
		graduateError := errors.New("unexpected error")
		mockRepository.On("Graduate", mockStudent.Name()).Return(graduateError)

		err := academy.GradeStudent(mockRepository, "test")

		if err == nil {
			t.Errorf("Expected to receive %v, but got no error.", graduateError)
		}
	})

	t.Run("Default case Repository save error", func(t *testing.T) {
		studentYear := uint8(2)
		mockRepository := mocks.NewRepository(t)
		mockStudent := mocks.NewStudent(t)

		mockStudent.On("FinalGrade").Return(3)
		mockStudent.On("Name").Return("test")
		mockStudent.On("Year").Return(studentYear)

		mockRepository.On("Get", "test").Return(mockStudent, nil)
		saveError := errors.New("unexpected error")
		mockRepository.On("Save", mockStudent.Name(), uint8(mockStudent.Year()+1)).Return(saveError)

		err := academy.GradeStudent(mockRepository, "test")

		if err == nil {
			t.Errorf("Expected to receive %v, but got no error.", saveError)
		}

	})

	t.Run("No errors", func(t *testing.T) {
		studentYear := uint8(2)
		mockRepository := mocks.NewRepository(t)
		mockStudent := mocks.NewStudent(t)

		mockStudent.On("FinalGrade").Return(3)
		mockStudent.On("Name").Return("test")
		mockStudent.On("Year").Return(studentYear)

		mockRepository.On("Get", "test").Return(mockStudent, nil)
		mockRepository.On("Save", mockStudent.Name(), uint8(mockStudent.Year()+1)).Return(nil)

		err := academy.GradeStudent(mockRepository, "test")

		if err != nil {
			t.Errorf("Expected to receive no errors, but received: %v", err)
		}
	})
}
