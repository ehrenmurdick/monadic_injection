package PACKAGE

type ResultTYPE struct {
	Value TYPE
	Err   error
}

func ReturnResultTYPE(v TYPE) ResultTYPE {
	return ResultTYPE{
		Value: v,
	}
}

func (m ResultTYPE) Bind(bl func(TYPE) ResultTYPE) ResultTYPE {
	if m.Err != nil {
		return ResultTYPE{
			Err: m.Err,
		}
	} else {
		return ResultTYPE{
			Value: bl(m.Value),
		}
	}
}

func (m ResultTYPE) FmapString(bl func(TYPE) Resultstring) Resultstring {
	if m.Err != nil {
		return Resultstring{
			Err: m.Err,
		}
	} else {
		return Resultstring{
			Value: bl(m.Value),
		}
	}
}

func (m ResultTYPE) Handle(bl func(error) error) ResultTYPE {
	if m.Err != nil {
		return ResultTYPE{
			Err: bl(m.Err),
		}
	} else {
		return m
	}
}
