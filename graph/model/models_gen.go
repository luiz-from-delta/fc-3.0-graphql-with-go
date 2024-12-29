// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Category struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	Description      *string           `json:"description,omitempty"`
	CourseCategories []*CourseCategory `json:"courseCategories"`
}

type Course struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	Description      *string           `json:"description,omitempty"`
	CourseCategories []*CourseCategory `json:"courseCategories"`
}

type CourseCategory struct {
	CourseID   string    `json:"courseId"`
	Course     *Course   `json:"course,omitempty"`
	CategoryID string    `json:"categoryId"`
	Category   *Category `json:"category,omitempty"`
}

type CreateCategoryInput struct {
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Courses     []string `json:"courses,omitempty"`
}

type CreateCourseInput struct {
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Categories  []string `json:"categories,omitempty"`
}

type CreateStudentInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateSubscriptionInput struct {
	StudentID string             `json:"studentId"`
	CourseID  string             `json:"courseId"`
	Status    SubscriptionStatus `json:"status"`
}

type Mutation struct {
}

type Query struct {
}

type Student struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Subscription struct {
}

type UpdateCategoryInput struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Courses     []string `json:"courses,omitempty"`
}

type UpdateCourseInput struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Categories  []string `json:"categories,omitempty"`
}

type UpdateStudentInput struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type UpdateSubscriptionInput struct {
	Status SubscriptionStatus `json:"status"`
}

type SubscriptionStatus string

const (
	SubscriptionStatusActive   SubscriptionStatus = "ACTIVE"
	SubscriptionStatusInactive SubscriptionStatus = "INACTIVE"
)

var AllSubscriptionStatus = []SubscriptionStatus{
	SubscriptionStatusActive,
	SubscriptionStatusInactive,
}

func (e SubscriptionStatus) IsValid() bool {
	switch e {
	case SubscriptionStatusActive, SubscriptionStatusInactive:
		return true
	}
	return false
}

func (e SubscriptionStatus) String() string {
	return string(e)
}

func (e *SubscriptionStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SubscriptionStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SubscriptionStatus", str)
	}
	return nil
}

func (e SubscriptionStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
