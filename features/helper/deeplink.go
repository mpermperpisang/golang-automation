package helper

import "github.com/golang-automation/features/helper/apps/action"

// OpenDeeplink : open page by deeplink
func OpenDeeplink(URL string) error {
	action.Device.Navigate(URL)

	return nil
}
