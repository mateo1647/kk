package options

type SearchOptions struct {
	AllNamespaces bool
	Namespace     string
	Selector      string
	FieldSelector string
}

// NewSearchOptions - genericclioptions wrapper for searchOptions
func NewSearchOptions() *SearchOptions {
	return &SearchOptions{}
}
