package settings

type Settings struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
	DB_DIALECT string
	SSL string
}

var globalSetting *Settings

func GetGlobalSettings() *Settings {
	if globalSetting != nil {
		return globalSetting
	}
	//
	return globalSetting
}