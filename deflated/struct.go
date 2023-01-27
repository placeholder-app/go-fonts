package deflated

type DeflatedFont struct {
	// name of the font
	Name string
	// weight for this variant of the font
	Weight string
	// source of the deflated font
	Source string
	// deflated font bytes
	Bytes []byte
}
