// Package challenge10 contains the solution for Challenge 10.
package challenge10

import (
	"fmt"
	"math"
	// Add any necessary imports here
)

// Shape interface defines methods that all shapes must implement
type Shape interface {
	Area() float64
	Perimeter() float64
	fmt.Stringer // Includes String() string method
}

// Rectangle represents a four-sided shape with perpendicular sides
type Rectangle struct {
	Width  float64
	Height float64
}

// NewRectangle creates a new Rectangle with validation
func NewRectangle(width, height float64) (*Rectangle, error) {
	// TODO: Implement validation and construction
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("width and height must be positive")
	}

	var rec Rectangle
	rec.Height = height
	rec.Width = width
	return &rec, nil
}

// Area calculates the area of the rectangle
func (r *Rectangle) Area() float64 {
	// TODO: Implement area calculation

	return r.Height*r.Width
}

// Perimeter calculates the perimeter of the rectangle
func (r *Rectangle) Perimeter() float64 {
	// TODO: Implement perimeter calculation

	return 2*(r.Height+r.Width)
}

// String returns a string representation of the rectangle
func (r *Rectangle) String() string {
	// TODO: Implement string representation

	return fmt.Sprintf(
    "Rectangle{Width: %.2f, Height: %.2f}",
    r.Width,
    r.Height,
)
}

// Circle represents a perfectly round shape
type Circle struct {
	Radius float64
}

// NewCircle creates a new Circle with validation
func NewCircle(radius float64) (*Circle, error) {
	// TODO: Implement validation and construction
	if radius <= 0{
		return nil,fmt.Errorf("the radius is less than 0")
	}
	var c Circle
	c.Radius = radius
	return &c, nil
}

// Area calculates the area of the circle
func (c *Circle) Area() float64 {
	// TODO: Implement area calculation
	  return math.Pow(c.Radius, 2) * math.Pi
}

// Perimeter calculates the circumference of the circle
func (c *Circle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return 2*math.Pi*c.Radius
}

// String returns a string representation of the circle
func (c *Circle) String() string {
	// TODO: Implement string representation
	return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

// Triangle represents a three-sided polygon
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func NewTriangle(a, b, c float64) (*Triangle, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return nil, fmt.Errorf("triangle sides must be positive")
	}

	// Triangle inequality
	if a+b <= c || a+c <= b || b+c <= a {
		return nil, fmt.Errorf("invalid triangle sides")
	}

	return &Triangle{
		SideA: a,
		SideB: b,
		SideC: c,
	}, nil
}

// Area calculates the area of the triangle using Heron's formula
func (t *Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2

	return math.Sqrt(
		s * (s - t.SideA) *
			(s - t.SideB) *
			(s - t.SideC),
	)
}

// Perimeter calculates the perimeter of the triangle
func (t *Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

// String returns a string representation of the triangle
func (t *Triangle) String() string {
	return fmt.Sprintf(
    "Triangle{sides: %.2f, %.2f, %.2f}",
    t.SideA,
    t.SideB,
    t.SideC,
)
}
// ShapeCalculator provides utility functions for shapes
type ShapeCalculator struct{

}

// NewShapeCalculator creates a new ShapeCalculator
func NewShapeCalculator() *ShapeCalculator {
	// TODO: Implement constructor
	return  &ShapeCalculator{}
}

// PrintProperties prints the properties of a shape
func (sc *ShapeCalculator) PrintProperties(s Shape) {
	// TODO: Implement printing shape properties

	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
	fmt.Println(s.String())

}

// TotalArea calculates the sum of areas of all shapes
func (sc *ShapeCalculator) TotalArea(shapes []Shape) float64 {
	// TODO: Implement total area calculation
	var cpt float64

	for _,s := range shapes{
		cpt+=s.Area()
	}
	return cpt
}

// LargestShape finds the shape with the largest area
func (sc *ShapeCalculator) LargestShape(shapes []Shape) Shape {
	// TODO: Implement finding largest shape
	m := shapes[0]
	for i:=1;i<len(shapes);i++{
		if m.Area()<shapes[i].Area(){
			m=shapes[i]
		}
	}
	return m
}

// SortByArea sorts shapes by area in ascending or descending order
func (sc *ShapeCalculator) SortByArea(shapes []Shape, ascending bool) []Shape {
	// TODO: Implement sorting shapes by area
	
	if ascending == true{
		for i := 0; i < len(shapes)-1; i++ {
		// Track if any swaps happen in this pass
		swapped := false
		
		for j := 0; j < len(shapes)-i-1; j++ {
			// Change > to < for descending order
			if shapes[j].Area() > shapes[j+1].Area() {
				// Swap elements using Go's tuple assignment
				shapes[j], shapes[j+1] = shapes[j+1], shapes[j]
				swapped = true
			}
		}
		
		// If no elements were swapped, the array is already sorted
		if !swapped {
			break
		}
	}
	
	}else{
		for i := 0; i < len(shapes)-1; i++ {
		// Track if any swaps happen in this pass
		swapped := false
		
		for j := 0; j < len(shapes)-i-1; j++ {
			// Change > to < for descending order
			if shapes[j].Area() < shapes[j+1].Area() {
				// Swap elements using Go's tuple assignment
				shapes[j], shapes[j+1] = shapes[j+1], shapes[j]
				swapped = true
			}
		}
		
		// If no elements were swapped, the array is already sorted
		if !swapped {
			break
		}
	}
	}

	return shapes
} 