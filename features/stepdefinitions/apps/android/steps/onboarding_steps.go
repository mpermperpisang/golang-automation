package step

import (
	android "github.com/golang-automation/features/helper/apps/android/action"
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
)

/*ClickBtnMulai : click Mulai button*/
func ClickBtnMulai() error {
	android.ClickByXPath(androidpages.BtnMulai)

	return nil
}
