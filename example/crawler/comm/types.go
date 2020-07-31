package comm

type Request struct {
	Url        string
	ParserName string
}

type Result struct {
	Requests []Request
	Profile  Profile
}

type Profile struct {
	Id   string
	Data interface{}
}
