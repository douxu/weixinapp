package structs

//ItemStruct : type of item struct
type ItemStruct struct {
	ItemTitle string
	ItemType  int
}

//FormStruct : type of form struct
type FormStruct struct {
	FormTitle string
	Items     []ItemStruct
}
