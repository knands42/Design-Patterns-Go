package breaking_the_principle

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {

	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func (r *Rectangle) Area() int {
	return r.width * r.height
}

func UseIt(sized Sized) {
	w := sized.GetWidth()
	sized.SetHeight(10)

	expectedArea := 10 * w
	actualArea := sized.GetWidth() * sized.GetHeight()
	if actualArea != expectedArea {
		panic("Actual area is not equal to expected area")
	}
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := &Square{}
	sq.width = size
	sq.height = size
	return sq
}

func (s *Square) SetWidth(width int) {

	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}
