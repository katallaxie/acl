package runtime

// PatternMatcher ...
type PatternMatcher interface {
	Match(path string) bool
}

type patternMatcher struct {
	fields []string
	params []string
}

var _ PatternMatcher = (*patternMatcher)(nil)

// NewPatternMatcher ...
func NewPatternMatcher(path string) (*patternMatcher, error) {
	m := new(patternMatcher)

	err := m.parse(path)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Match ...
func (p *patternMatcher) Match(path string) bool {
	return false
}

// GetFields ...
func (p *patternMatcher) GetFields() []string {
	return p.fields
}

// GetParams ...
func (p *patternMatcher) GetParams() []string {
	return p.params
}

func (p *patternMatcher) parse(path string) error {
	fields := make([]string, 0)
	params := make([]string, 0)

	p.fields = fields
	p.params = params

	return nil
}
