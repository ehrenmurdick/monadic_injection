package monads

type ResultItemRepo struct {
	Value ItemRepo
	Err   error
}

func ReturnResultItemRepo(v ItemRepo) ResultItemRepo {
	return ResultItemRepo{
		Value: v,
	}
}

func (m ResultItemRepo) Bind(bl func(ItemRepo) ResultItemRepo) ResultItemRepo {
	if m.Err != nil {
		return ResultItemRepo{
			Err: m.Err,
		}
	} else {
		return ResultItemRepo{
			Value: bl(m.Value),
		}
	}
}
