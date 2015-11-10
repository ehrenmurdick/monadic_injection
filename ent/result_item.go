package ent

func (i ResultItem) GetTitle() Resultstring {
	return i.Fmapstring(Item.GetTitle)
}

//go:generate ../monads/result result_item.go Item
//go:generate ../monads/fmap result_item.go Item string
// GENERATED
type ResultItem struct {
	Value Item
	Err   error
}

func ReturnResultItem(v Item) ResultItem {
	return ResultItem{
		Value: v,
	}
}

func (m ResultItem) Bind(bl func(Item) Item) ResultItem {
	if m.Err != nil {
		return ResultItem{
			Err: m.Err,
		}
	} else {
		return ResultItem{
			Value: bl(m.Value),
		}
	}
}

func (m ResultItem) Handle(bl func(error) error) ResultItem {
	if m.Err != nil {
		return ResultItem{
			Err: bl(m.Err),
		}
	} else {
		return m
	}
}

func (m ResultItem) Fmapstring(bl func(Item) (string, error)) (out Resultstring) {
	if m.Err != nil {
			out.Err = m.Err
	} else {
		out.Value, out.Err = bl(m.Value)
	}
	return
}
