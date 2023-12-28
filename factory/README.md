# Factory Design Pattern
The Factory design pattern in Go provides a way to create objects without specifying their exact type within the creating method. It falls under the category of creational design patterns and is useful when you want to delegate the responsibility of object creation to a separate method or factory.

In Go, you can implement the Factory design pattern using interfaces and concrete types. Here's an example to illustrate this concept:

```go
package main

import (
	"fmt"
)

// Shape interface defines the methods that all shapes should implement.
type Shape interface {
	Draw() string
}

// Concrete types implementing the Shape interface.
type Circle struct{}

func (c Circle) Draw() string {
	return "Circle drawn"
}

type Square struct{}

func (s Square) Draw() string {
	return "Square drawn"
}

// ShapeFactory creates different shapes based on the input.
func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{}
	case "square":
		return Square{}
	default:
		return nil // Or return an error for an unsupported shape
	}
}

func main() {
	circle := ShapeFactory("circle")
	square := ShapeFactory("square")

	fmt.Println(circle.Draw()) // Output: Circle drawn
	fmt.Println(square.Draw()) // Output: Square drawn
}
```

Here, `Shape` is the interface that defines the behavior of shapes, and `Circle` and `Square` are concrete types implementing the `Shape` interface. The `ShapeFactory` function takes in a parameter (like a string) to decide which type of shape to create and returns an object implementing the `Shape` interface.

This way, the client code (`main()` function in this case) doesn't need to know the exact type of the object being created; it relies on the factory method to provide the appropriate object based on the input.

You can extend this pattern by adding more shapes or modifying the factory method to accommodate additional types without changing the client code, making it more flexible and maintainable.