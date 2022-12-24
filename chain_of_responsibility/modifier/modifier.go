package modifier

type Modifier interface {
	Add(m Modifier)
	Handle()
}
