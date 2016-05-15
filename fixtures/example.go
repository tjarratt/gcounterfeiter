package fixtures

type Example interface {
	Something()
	TakesAParameter(string)
	TakesThreeParameters(string, string, string)
}
