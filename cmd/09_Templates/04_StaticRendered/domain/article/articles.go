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

func (articles *Articles) SaveToPath(folder string) error {
	for _, article := range *articles {
		if _, errSave := article.SaveToPath(folder); errSave != nil {
			return errSave
		}
	}

	return nil
}
