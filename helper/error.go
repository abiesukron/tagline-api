package helper

func PanicErrorHandle(err error) {
	if err != nil {
		panic(err)
	}
}
