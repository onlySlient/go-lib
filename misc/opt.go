package misc

func GetOptName(opts ...string) string {
	for _, o := range opts {
		if o != "" {
			return o
		}
	}

	return ""
}
