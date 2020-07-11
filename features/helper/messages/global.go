package messages

// NotFoundDriver : pesan untuk driver
func NotFoundDriver() string {
	return "Driver not found"
}

// NotFoundPlatform : pesan untuk platforms
func NotFoundPlatform() string {
	return "Platform not found"
}

// NotDetected : pesan untuk sesuatu yang tidak terdeteksi karena kosong
func NotDetected() string {
	return "Not detected"
}

// NotValid : pesan untuk sesuatu yang tidak valid
func NotValid(text string) string {
	return "Not a valid " + text
}
