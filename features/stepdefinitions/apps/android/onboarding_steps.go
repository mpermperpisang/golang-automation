package androidsteps

import (
	android "github.com/golang-automation/features/helper/apps/android/action"
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
)

/*clickBtnMulai function to click Mulai button*/
func clickBtnMulai() error {
	android.ClickByXPath(androidpages.BtnMulai)

	return nil
}
