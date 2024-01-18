package article

type Articles []*Article

func (articles *Articles) Validate() error {
	for _, article := range *articles {
		if errValidate := article.Validate(); errValidate != nil {
			return errValidate
		}
	}

	return nil
}
