package baz

import "example-module/bar"

func OtherExport() {
	bar.SomeExport()
}
