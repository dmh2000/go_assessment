package goassessment

// fix the Sort interface for PersonSlice
func (p PersonSlice) Len() int               { return 0 }
func (p PersonSlice) Less(i int, j int) bool { return true }
func (p PersonSlice) Swap(i int, j int)      {}

func sendNotification(n Notifier) string {
	return ""
}
