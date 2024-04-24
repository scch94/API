package structs

type Petition struct {
	mobile        string
	message       string
	useOriginName string
}

// constructor
func NewPetition(mobile string, message string, useoriginame string) *Petition {
	return &Petition{
		mobile,
		message,
		useoriginame,
	}
}

// getters y setters
func (p Petition) SetMobile(mobile string) {
	p.mobile = mobile
}

func (p Petition) Mobile() string {
	return p.mobile
}

func (p Petition) SetMessage(message string) {
	p.message = message
}

func (p Petition) Message() string {
	return p.message
}
func (p Petition) SetUseOriginName(useOriginName string) {
	p.useOriginName = useOriginName
}
func (p Petition) UseOriginName() string {
	return p.useOriginName
}
func (p Petition) ToString() string {
	var data string = "this are the values of the petition \n"
	data += "mobile: " + p.mobile + ",\n"
	data += "message: " + p.message + ",\n"
	data += "useOriginName: " + p.useOriginName
	return data
}
