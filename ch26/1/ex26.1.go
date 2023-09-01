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

type ReportSender struct {
}

func (s *ReportSender) SendReport(report Report) {
	fmt.Printf("send report, report : %s \n", report.Report())
}

func main() {

	mr := &MarketingReport{"marketingReport"}
	fr := &FinanceReport{"financeReport"}

	reportSender := &ReportSender{}
	reportSender.SendReport(mr)
	reportSender.SendReport(fr)

}
