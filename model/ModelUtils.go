package model

type ModelUtils interface {
	ToDbModel()
}

func (modelUtils ModelUtils) ToDbModel(userModel UserModel) string {
	return ""
}