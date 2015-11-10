package monads

type ResultItem struct {
	Value Item
	Err   error
}

func (m ResultItem) Bind(bl func(Item) ResultItem) ResultItem {
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
