package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)
	isValid := form.Valid()

	if !isValid {
		t.Error("t is invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form shows valid when required fields are missing")
	}
	postedData := url.Values{}
	postedData.Add("a", "A")
	postedData.Add("b", "B")
	postedData.Add("c", "C")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Forms shows does not have the requiree fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	field := "a"
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	has := form.Has(field)
	if has {
		t.Errorf("Should have return false as there is no field in request, has: %t", has)
	}

	postedData := url.Values{}
	postedData.Add("a", "aa")
	form = New(postedData)
	has = form.Has(field)
	if !has {
		t.Errorf("Should have not raise an error because field is filled, has: %t", has)
	}

}

func TestForm_MinLength(t *testing.T) {
	field := "a"
	length := 2

	invalidDatas := url.Values{}
	invalidDatas.Add("a", "a")
	form := New(invalidDatas)
	isMinLength := form.MinLength(field, length)

	if form.Valid() {
		t.Errorf("Should have raised an error but did'nt, isMinLength %t", isMinLength)
	}

	validDatas := url.Values{}
	validDatas.Add("a", "aaa")

	form = New(validDatas)

	if !form.Valid() {
		t.Errorf("Should have raised no error but did, isMinLength %t", isMinLength)
	}
}
func TestForm_IsEmail(t *testing.T) {
	field := "email"

	invalidDatas := url.Values{}
	invalidDatas.Add("email", "aa")

	r := httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = invalidDatas
	form := New(r.PostForm)

	isEmail := form.IsEmail(field)
	if isEmail {
		t.Errorf("Show valid when should have raise an error, isEmail: %t", isEmail)
	}

	validDatas := url.Values{}
	validDatas.Add("email", "aa@aa.com")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = validDatas
	form = New(r.PostForm)

	isEmail = form.IsEmail(field)
	if !isEmail {
		t.Errorf("Should have raise no error but got one. isEmail: %t", isEmail)
	}

}
