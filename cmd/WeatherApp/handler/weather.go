package weatherHandler

import (
	"CliTaskManager/cmd/WeatherApp/Model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetCelciusFromLocation(Location string) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=4cd2714be825aa4660532f7bfe540107", Location)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	weather := &Model.WeatherBody{}
	err := json.Unmarshal(body, &weather)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s sicagligi: %f derece\n", weather.Name, weather.Main.Temp)
}
