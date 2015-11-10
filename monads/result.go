package monads

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
