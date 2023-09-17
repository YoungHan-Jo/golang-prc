package main

type Report interface {
	Report() string
}

type MarketingReport struct {
}

func (m *MarketingReport) Report() string {

	return "marketing report"
}

func SendReport(r Report) {
	if _, ok := r.(*MarketingReport); ok {
		panic("Can't send MarketingReport")
	}
}

func main() {
	var report = &MarketingReport{}
	SendReport(report)
}
