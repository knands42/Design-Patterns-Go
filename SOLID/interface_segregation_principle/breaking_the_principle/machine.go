package breaking_the_principle

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m *MultiFunctionPrinter) Print(d Document) {
	// ...
}
func (m *MultiFunctionPrinter) Fax(d Document) {
	// ...
}
func (m *MultiFunctionPrinter) Scan(d Document) {
	// ...
}

type OldFashionedPrinter struct{}

func (o *OldFashionedPrinter) Print(d Document) {
	// ...
}
func (o *OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

func (o *OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}
