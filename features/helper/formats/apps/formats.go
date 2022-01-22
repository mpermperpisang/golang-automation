package formats

func WebPath(path string) string {
	return "%s/test/" + path + "/web/%s/"
}

func AppsPath(path, platform string) string {
	return "%s/test/" + path + "/apps/" + platform + "/"
}
