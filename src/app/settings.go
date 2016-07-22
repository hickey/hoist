package app

var log = GetLogger()

var settings map[string]string

func init() {
	settings = make(map[string]string)
	settings["verbose"] = "false"
	settings["debug"] = "false"
	settings["logging.level"] = "WARNING"

}

func GetSetting(key string) string {
	return settings[key]
}

func GetSettingBool(key string) bool {
	value := settings[key]
	if value == "true" {
		return true
	}
	return false
}
