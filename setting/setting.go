package setting

type Setting struct {
	IgnoreHide      bool
	SpecifiedSuffix []string
	IgnoreSpaceLine bool
	ShowInfo        bool
}

func NewDefault()*Setting{
	return &Setting{
		IgnoreHide:      true,
		SpecifiedSuffix: []string{"go"},
		IgnoreSpaceLine: true,
		ShowInfo: true,
	}
}