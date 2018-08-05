package tbm

type field struct{
	name  string
	value interface{}
	fType int
}

func CreateField(name string,value interface{},fType int) *field {
	return &field{name,value,fType}
}
