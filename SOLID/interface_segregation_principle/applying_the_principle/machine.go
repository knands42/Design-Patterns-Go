package applying_the_principle

type Document struct{}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Machine interface {
	Printer
	Scanner
}

type MultiFunctionPrinter struct{}

func (m *MultiFunctionPrinter) Print(d Document) {
	// ...
}

func (m *MultiFunctionPrinter) Scan(d Document) {
	// ...
}

type OldFashionedPrinter struct{}

func (o *OldFashionedPrinter) Print(d Document) {
	// ...
}

type Photocopier struct{}

func (p *Photocopier) Print(d Document) {
	// ...
}

func (p *Photocopier) Scan(d Document) {
	// ...
}
