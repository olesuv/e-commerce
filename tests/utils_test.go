package tests

import (
	"testing"

	"server.go/constants"
	"server.go/graph/model"
	errors "server.go/utils/errors"
	typesConverter "server.go/utils/types_converters"
)

func TestOrderErrors_CheckCreateOrderInput(t *testing.T) {

	orderErrors := errors.OrderErrors{}

	// Test case 1: Empty title
	input1 := model.CreateOrderInput{
		Title:       nil,
		Description: strPtr("Sample description"),
		Price:       floatPtr(10.5),
	}
	err1 := orderErrors.CheckCreateOrderInput(input1)
	if err1 == nil || err1.Error() != "title is required" {
		t.Errorf("Expected 'title is required' error, got: %v", err1)
	}

	// Test case 2: Title length less than 3 characters
	input2 := model.CreateOrderInput{
		Title:       strPtr("ab"),
		Description: strPtr("Sample description"),
		Price:       floatPtr(10.5),
	}
	err2 := orderErrors.CheckCreateOrderInput(input2)
	if err2 == nil || err2.Error() != "minimum length of title is 3 characters" {
		t.Errorf("Expected 'minimum length of title is 3 characters' error, got: %v", err2)
	}

	// Test case 3: Title length greater than 100 characters
	input3 := model.CreateOrderInput{
		Title:       strPtr("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ut urna auctor, lacinia nunc sed, tincidunt nisl. Nulla facilisi. Sed nec nunc auctor, ultrices nisl id, lacinia nunc."),
		Description: strPtr("Sample description"),
		Price:       floatPtr(10.5),
	}
	err3 := orderErrors.CheckCreateOrderInput(input3)
	if err3 == nil || err3.Error() != "maximum length of title is 100 characters" {
		t.Errorf("Expected 'maximum length of title is 100 characters' error, got: %v", err3)
	}

	// Test case 4: Empty description
	input4 := model.CreateOrderInput{
		Title:       strPtr("Sample title"),
		Description: nil,
		Price:       floatPtr(10.5),
	}
	err4 := orderErrors.CheckCreateOrderInput(input4)
	if err4 == nil || err4.Error() != "description is required" {
		t.Errorf("Expected 'description is required' error, got: %v", err4)
	}

	// Test case 5: Description length less than 3 characters
	input5 := model.CreateOrderInput{
		Title:       strPtr("Sample title"),
		Description: strPtr("ab"),
		Price:       floatPtr(10.5),
	}
	err5 := orderErrors.CheckCreateOrderInput(input5)
	if err5 == nil || err5.Error() != "minimum length of descriprion is 3 characters" {
		t.Errorf("Expected 'minimum length of descriprion is 3 characters' error, got: %v", err5)
	}

	// Test case 6: Description length greater than 500 characters
	input6 := model.CreateOrderInput{
		Title: strPtr("Sample title"),
		Description: strPtr(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
                             Fusce ultricies leo at lectus consequat, sit amet condimentum quam aliquam. 
                             Integer sed ullamcorper justo. Proin efficitur vestibulum mi, vel vulputate turpis. 
                             Sed ullamcorper urna id tincidunt auctor. Nullam eget commodo mi. 
                             Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. 
                             Suspendisse potenti. Nullam vitae sem et sem sodales ullamcorper. 
                             Vivamus gravida euismod mi a aliquet. Maecenas sed arcu eget risus tincidunt vulputate. 
                             Sed id dui non turpis pretium tempus. Sed accumsan nec turpis non cursus. 
                             Vivamus tincidunt dictum enim, sit amet sollicitudin purus tincidunt at. 
                             Sed eget lobortis velit. Etiam venenatis, libero non sollicitudin tristique, 
                             sapien erat lobortis elit, vel malesuada purus sapien ac felis. Curabitur nec dui auctor, 
                             pulvinar turpis nec, tempor justo.`),
		Price: floatPtr(10.5),
	}
	err6 := orderErrors.CheckCreateOrderInput(input6)
	if err6 == nil || err6.Error() != "maximum length of title is 500 characters" {
		t.Errorf("Expected 'maximum length of title is 500 characters' error, got: %v", err6)
	}

	// Test case 7: Empty price
	input7 := model.CreateOrderInput{
		Title:       strPtr("Sample title"),
		Description: strPtr("Sample description"),
		Price:       nil,
	}
	err7 := orderErrors.CheckCreateOrderInput(input7)
	if err7 == nil || err7.Error() != "price is required" {
		t.Errorf("Expected 'price is required' error, got: %v", err7)
	}

	// Test case 8: Price less than 0
	input8 := model.CreateOrderInput{
		Title:       strPtr("Sample title"),
		Description: strPtr("Sample description"),
		Price:       floatPtr(-10.5),
	}
	err8 := orderErrors.CheckCreateOrderInput(input8)
	if err8 == nil || err8.Error() != "price must be greater than 0" {
		t.Errorf("Expected 'price must be greater than 0' error, got: %v", err8)
	}

	// Test case 9: Price greater than 1000000
	input9 := model.CreateOrderInput{
		Title:       strPtr("Sample title"),
		Description: strPtr("Sample description"),
		Price:       floatPtr(1000001),
	}
	err9 := orderErrors.CheckCreateOrderInput(input9)
	if err9 == nil || err9.Error() != "this platform is for 'normal' people, not for millionaires" {
		t.Errorf("Expected 'this platform is for 'normal' people, not for millionaires' error, got: %v", err9)
	}
}

func strPtr(s string) *string {
	return &s
}

func floatPtr(f float64) *float64 {
	return &f
}

func TestOrderTypesConverter_ConvertCategoryTypes(t *testing.T) {
	converter := typesConverter.OrderTypesConverter{}

	// Test case 1: Empty category
	input1 := model.CreateOrderInput{}
	expected1 := constants.OrderCategories{constants.Other}
	result1 := converter.ConvertCategoryTypes(input1)
	if !compareOrderCategories(result1, expected1) {
		t.Errorf("Unexpected result. Expected: %v, got: %v", expected1, result1)
	}

	// Test case 2: Single category
	electronicsCategory := model.CategoryElectronics // Get the constant value
	input2 := model.CreateOrderInput{
		Category: []*model.Category{&electronicsCategory}, // Use pointer to Category
	}
	expected2 := constants.OrderCategories{constants.Electronics}
	result2 := converter.ConvertCategoryTypes(input2)
	if !compareOrderCategories(result2, expected2) {
		t.Errorf("Unexpected result. Expected: %v, got: %v", expected2, result2)
	}

	// Test case 3: Multiple categories
	fashionCategory := model.CategoryFashion // Get the constant value
	homeCategory := model.CategoryHome       // Get the constant value
	input3 := model.CreateOrderInput{
		Category: []*model.Category{&fashionCategory, &homeCategory}, // Use pointers to Categories
	}
	expected3 := constants.OrderCategories{constants.Fashion, constants.Home}
	result3 := converter.ConvertCategoryTypes(input3)
	if !compareOrderCategories(result3, expected3) {
		t.Errorf("Unexpected result. Expected: %v, got: %v", expected3, result3)
	}

}

func compareOrderCategories(a, b constants.OrderCategories) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
