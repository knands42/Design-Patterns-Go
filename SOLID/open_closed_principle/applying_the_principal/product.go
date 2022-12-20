package applying_the_principal

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	Name  string
	Color Color
	Size  Size
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type Filter struct{}

func (f *Filter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i := range products {
		if spec.IsSatisfied(&products[i]) {
			result = append(result, &products[i])
		}
	}
	return result
}

type ColorSpecification struct {
	Color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.Color == c.Color
}

type SizeSpecification struct {
	Size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.Size == s.Size
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}
