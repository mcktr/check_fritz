package perfdata

import "fmt"

// PerformanceData is the data structure for check performance data
type PerformanceData struct {
	Label    string
	Value    float64
	UOM      string
	Warning  string
	Critical string
	Minimum  string
	Maximum  string
}

// CreatePerformanceData creates and returns a new PerformanceData object
func CreatePerformanceData(Label string, Value float64, UOM string) *PerformanceData {
	var pd PerformanceData

	pd.Label = Label
	pd.Value = Value
	pd.UOM = UOM

	return &pd
}

// SetWarning sets the warning value of a PerformanceData object
func (pd *PerformanceData) SetWarning(Warning float64) {
	pd.Warning = fmt.Sprintf("%f", Warning)

}

// SetCritical sets the critical value of a PerformanceData object
func (pd *PerformanceData) SetCritical(Critical float64) {
	pd.Critical = fmt.Sprintf("%f", Critical)

}

// SetMinimum sets the minimum value of a PerformanceData object
func (pd *PerformanceData) SetMinimum(Minimum float64) {
	pd.Minimum = fmt.Sprintf("%f", Minimum)
}

// SetMaximum sets the maximum value of a PerformanceData object
func (pd *PerformanceData) SetMaximum(Maximum float64) {
	pd.Maximum = fmt.Sprintf("%f", Maximum)
}

// GetPerformanceDataAsString returns a PerformanceData object as formatted string
func (pd *PerformanceData) GetPerformanceDataAsString() string {
	return fmt.Sprintf("| '%s'=%f%s;%s;%s;%s;%s", pd.Label, pd.Value, pd.UOM, pd.Warning, pd.Critical, pd.Minimum, pd.Maximum)
}
