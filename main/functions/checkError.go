package functions

//Func for checking error
//Created to avoid same code
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}