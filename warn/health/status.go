package health

type Status string

const (
	StatusGreen  Status = "green"  // perfect health
	StatusYellow Status = "yellow" // bad but not fatal leath
	StatusRed    Status = "red"    // crashed, inactive
	StatusBlack  Status = "black"  // inactive, not crashed
)
