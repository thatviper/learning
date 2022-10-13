package erratum


func Use(opener ResourceOpener, input string) (err error) {
	var res Resource
	flag := false
	for {
		res, err = opener()
		switch err.(type) {
		case TransientError:
		default:
			if err != nil {
				return err
			} else {
                defer res.Close()
				flag = true
			}
		}
		if flag {
			break
		}
	}
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				frobErr := r.(FrobError)
				res.Defrob(frobErr.defrobTag)
				err = frobErr
			default:
				err = r.(error)
			}
		}
	}()
	res.Frob(input)
	return nil
}