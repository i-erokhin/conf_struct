package conf_struct

func ErrsToString (errs []error) (errsStr string) {
	n := ""
	for _, e := range errs {
		errsStr += n + e.Error()
		n = "\n"
	}
	return
}
