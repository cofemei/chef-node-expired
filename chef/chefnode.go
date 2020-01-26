package main

// ExpiredNode is struct
type ExpiredNode struct {
	NodeName string
	IsExpired bool
	ExpiredSec int
	OhaiTime int64
	OS string
	URL string
	Version string
	ipv4 string
}

// ExpiredNodeList for sort
type ExpiredNodeList []ExpiredNode

// for sort
func (a ExpiredNodeList) Len() int           { return len(a) }
func (a ExpiredNodeList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ExpiredNodeList) Less(i, j int) bool { return a[i].OhaiTime < a[j].OhaiTime }

