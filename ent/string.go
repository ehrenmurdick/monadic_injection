package ent

type MaybeString struct {
	Value string
	Err   error
}

func (in MaybeString) Within(bl func(string) error) MaybeString {
	if in.Err != nil {
		return in
	} else {
		return MaybeString{
			Err:   bl(in.Value),
			Value: in.Value,
		}
	}
}

func (in MaybeString) Handle(handler func(error) error) MaybeString {
	if in.Err != nil {
		handler(in.Err)
	}
	return in
}
