package util

var confMap map[string]string

func init() {
	confMap = readProperties("./conf")

}

func GetConf(item string) string {
	return confMap[item]
}
