package ent

type Resultstring struct {
	Value string
	Err   error
}

func ReturnResultstring(v string) Resultstring {
	return Resultstring{
		Value: v,
	}
}

func (m Resultstring) Bind(bl func(string) Resultstring) Resultstring {
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
