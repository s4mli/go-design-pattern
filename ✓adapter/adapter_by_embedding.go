package adapter

type Banner struct {
	str string
}

func (b *Banner) getString() string {
	return "* " + b.str + " *"
}

type Decorator interface {
	Decorate() string
}

type EmbeddedDecorateBanner struct {
	*Banner // magic :) golang doesn't have inheritance
}

func (edb *EmbeddedDecorateBanner) Decorate() string {
	return edb.getString()
}

func NewEmbeddedDecorateBanner(str string) *EmbeddedDecorateBanner {
	// golang doesn't have inheritance
	return &EmbeddedDecorateBanner{&Banner{str}}
}
