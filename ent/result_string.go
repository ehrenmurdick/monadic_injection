package ent

//go:generate ../monads/result result_string.go string
// GENERATED
type Resultstring struct {
	Value string
	Err   error
}

func ReturnResultstring(v string) Resultstring {
	return Resultstring{
		Value: v,
	}
}

func (m Resultstring) Bind(bl func(string) string) Resultstring {
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

func (m Resultstring) Handle(bl func(error) error) Resultstring {
	if m.Err != nil {
		return Resultstring{
			Err: bl(m.Err),
		}
	} else {
		return m
	}
}
