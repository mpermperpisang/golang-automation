package androidsteps

import (
	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/objectabstractions/apps/androidpages"
)

/*clickBtnMulai function to click Mulai button*/
func clickBtnMulai() error {
	android.ClickByXPath(androidpages.BtnMulai)

	return nil
}
