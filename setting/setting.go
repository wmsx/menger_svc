package setting

const XConfURL = "http://sx-xconf-micro:8090"
const DevXConfURL = "http://192.168.0.199:8090"

func SetUp(appName, env string) error {
	if err := setUpDB(appName, env); err != nil {
		return err
	}
	//if err := setUpJwt(appName, env); err  != nil {
	//	return err
	//}
	return nil
}
