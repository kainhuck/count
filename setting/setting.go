package setting

type Setting struct {
	IgnoreHide      bool
	SpecifiedSuffix []string
	IgnoreSpaceLine bool
	ShowInfo        bool
	ButSuffix       []string
}

func NewDefault() *Setting {
	return New(true, nil, false, false, nil)
}

func New(ignoreHide bool, specifiedSuffix []string, ignoreSpaceLine bool, showInfo bool, butSuffix []string) *Setting {
	return &Setting{
		IgnoreHide:      ignoreHide,
		SpecifiedSuffix: specifiedSuffix,
		IgnoreSpaceLine: ignoreSpaceLine,
		ShowInfo:        showInfo,
		ButSuffix: butSuffix,
	}
}

func NewDefaultBySuffix(suffix []string) *Setting {
	return New(true, suffix, false, false, nil)
}

func (s *Setting) AddSuffix(suffix string) {
	if s.SpecifiedSuffix == nil {
		s.SpecifiedSuffix = make([]string, 0, 1)
	}
	s.SpecifiedSuffix = append(s.SpecifiedSuffix)
}

func (s *Setting) AddButSuffix(suffix string) {
	if s.ButSuffix == nil {
		s.ButSuffix = make([]string, 0, 1)
	}
	s.ButSuffix = append(s.ButSuffix)
}
