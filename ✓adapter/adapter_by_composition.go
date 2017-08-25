package adapter

type CompositionDecorateBanner struct {
	banner *Banner
}

func (cdb *CompositionDecorateBanner) Decorate() string {
	return cdb.banner.getString()
}

func NewCompositionDecorateBanner(str string) *CompositionDecorateBanner {
	return &CompositionDecorateBanner{&Banner{str}}
}
