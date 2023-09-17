package main

import "fmt"

type Report interface {
	Report() string
}

type FinanceReport struct {
	report string
}

func (fr *FinanceReport) Report() string {
	return fr.report
}

type MarketingReport struct {
	report string
}

func (mr *MarketingReport) Report() string {
	return mr.report
}

type ReportSender interface {
	Send(r Report)
}

type EmailSender struct{}

func (e *EmailSender) Send(r Report) {
	fmt.Printf("sent by Email, report : %s \n", r.Report())
}

type FaxSender struct{}

func (f *FaxSender) Send(r Report) {
	fmt.Printf("sent by Fax, report : %s \n", r.Report())
}

var sender ReportSender

func main() {

	mr := &MarketingReport{"marketingReport"}
	fr := &FinanceReport{"financeReport"}

	// sender = &EmailSender{}
	sender = &FaxSender{}

	sender.Send(mr)
	sender.Send(fr)

}
