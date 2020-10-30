package setting

type Setting struct {
	IgnoreHide      bool
	SpecifiedSuffix []string
	IgnoreSpaceLine bool
	ShowInfo        bool
}

func NewDefault() *Setting {
	return New(true, nil, false, false)
}

func New(ignoreHide bool, specifiedSuffix []string, ignoreSpaceLine bool, showInfo bool) *Setting {
	return &Setting{
		IgnoreHide:      ignoreHide,
		SpecifiedSuffix: specifiedSuffix,
		IgnoreSpaceLine: ignoreSpaceLine,
		ShowInfo:        showInfo,
	}
}

func NewDefaultBySuffix(suffix []string) *Setting {
	return New(true, suffix, false, false)
}

func (s *Setting) AddSuffix(suffix string) {
	if s.SpecifiedSuffix == nil {
		s.SpecifiedSuffix = make([]string, 0, 1)
	}
	s.SpecifiedSuffix = append(s.SpecifiedSuffix)
}
