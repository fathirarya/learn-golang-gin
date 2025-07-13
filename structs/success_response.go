package structs

type SuccessReaponse struct {
	Succes  bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
