package interfaces

type Visitor interface {
	VisitForSquare(Shape)
	VisitForCircle(Shape)
	VisitForTriangle(Shape)
}