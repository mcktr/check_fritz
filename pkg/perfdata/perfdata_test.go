package perfdata

import "testing"

func TestCreatePerformanceData(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.43, "")
	want := "| 'testing'=123.430000;;;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "s")
	want = "| 'testing'=123.430000s;;;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "us")
	want = "| 'testing'=123.430000us;;;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "ms")
	want = "| 'testing'=123.430000ms;;;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}

	pd = CreatePerformanceData("testing", 99.43, "%")
	want = "| 'testing'=99.430000%;;;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "B")
	want = "| 'testing'=123.430000B;;;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "KB")
	want = "| 'testing'=123.430000KB;;;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "MB")
	want = "| 'testing'=123.430000MB;;;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "TB")
	want = "| 'testing'=123.430000TB;;;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "c")
	want = "| 'testing'=123.430000c;;;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}
}

func TestSetWarning(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.43, "")

	pd.SetWarning(62)

	want := "| 'testing'=123.430000;62.000000;;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}
}

func TestSetCritical(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.43, "")

	pd.SetCritical(62)

	want := "| 'testing'=123.430000;;62.000000;;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}
}

func TestSetMinimum(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.43, "")

	pd.SetMinimum(62)

	want := "| 'testing'=123.430000;;;62.000000;"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}
}

func TestSetMaximum(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.43, "")

	pd.SetMaximum(62)

	want := "| 'testing'=123.430000;;;;62.000000"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}
}

func TestGetPerformanceDataAsString(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.43, "s")

	pd.SetWarning(48)
	pd.SetCritical(55)
	pd.SetMinimum(13)
	pd.SetMaximum(875)

	want := "| 'testing'=123.430000s;48.000000;55.000000;13.000000;875.000000"

	if pd.GetPerformanceDataAsString() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.GetPerformanceDataAsString(), want)
	}
}
