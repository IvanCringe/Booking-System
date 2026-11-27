package service

import (
	"testing"
)

func TestIsValidPhone(t *testing.T) {
	tests := []struct {
		name  string
		phone string
		want  bool
	}{
		{"с плюсом", "+79991234567", true},
		{"с восьмеркой", "89991234567", true},
		{"пустая строка", "", false},
		{"слишком короткий", "8999123456", false},
		{"слишком длинный", "899912345678", false},
		{"буквы внутри", "8999abc4567", false},
		{"плюс вместо цифры", "8+999123456", false},
		{"минус вместо цифры", "8-999123456", false},
		{"семерка без плюса", "79991234567", false},
		{"плюс и восьмерка", "+89991234567", false},
		{"пробелы", "8 999 123 45 67", false},
		{"скобки и дефисы", "8(999)123-45-67", false},
		{"только префикс", "+7", false},
		{"только восьмерка", "8", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isValidPhone(tc.phone)
			if got != tc.want {
				t.Errorf("phone %q: got %v, want %v", tc.phone, got, tc.want)
			}
		})
	}
}
