package model

type ModelUtils interface {
	ToDbModel()
}

func ToDbModel(userModel UserModel) string {
	return ""
}