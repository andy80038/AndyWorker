package persist

import (
	"testing"
)

func TestSave(t *testing.T) {
	// item := model.Profile{
	// 	Age:        34,
	// 	Height:     162,
	// 	Weight:     57,
	// 	Income:     "3000-5000元",
	// 	Gender:     "女",
	// 	Name:       "安靜ㄉ雪",
	// 	Xinzuo:     "牧羊座",
	// 	Occupation: "人事/行政",
	// 	Marriage:   "離異",
	// 	House:      "以購房",
	// }
	// id, err := Save(item)
	// if err != nil {
	// 	panic(err)
	// }
	// //TODO : try to strat up elastic search
	// //here using docker

	// client, err := elastic.NewClient(elastic.SetSniff(false)) //Must turn off sniff in docker
	// if err != nil {
	// 	panic(err)
	// }
	// resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }

	// //t.Logf("%+v", resp.Source)
	// var actual model.Profile
	// err = json.Unmarshal([]byte(*resp.Source), &actual)
	// if err != nil {
	// 	panic(err)
	// }
	// if item != actual {
	// 	panic(err)
	// }

}
