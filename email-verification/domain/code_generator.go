package domain

type CodeGenerator interface {
	Generate() (string, error)
	Hash(code string) string
}
