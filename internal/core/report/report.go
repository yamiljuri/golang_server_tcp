package report

type Report interface {
	Save(data interface{}) error
}

type report struct {
}

func NewReport() Report {
	return &report{}
}

func (r *report) Save(data interface{}) error {

}
