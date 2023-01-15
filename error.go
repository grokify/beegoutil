package beegoutil

const pgLastInsertIDError string = "no LastInsertId available"

// ProcErrPg removes the "no LastInsertId available" error. See more
// here: https://github.com/beego/beego/issues/3070 .
func ProcErrPg(err error) error {
	if err.Error() == pgLastInsertIDError {
		return nil
	}
	return err
}
