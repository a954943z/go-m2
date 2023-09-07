package main

import (
	"bytes"
	"encoding/json"
	"log"
	"test1/src/count"
	"test1/src/redistest"
	"time"

	"github.com/gin-gonic/gin"
)

// "fmt"
// "test1/src/d0608"
// "test1/src/work"
// "github.com/rs/xid"

// func main() {
// 	// newId := xid.New()
// 	// _ = newId

// 	// test1 := []int{4, 5, 6}
// 	// for index, v := range test1 {
// 	// 	fmt.Printf("index: %v, v: %v", index, v)
// 	// 	fmt.Println("")
// 	// // }
// 	// aaaa.Forfor()
// 	// l1 := []int{1, 2, 3}
// 	// l2 := []string{"4", "5", "6"}

// 	// _, qq := bbb.Maptest(l1, l2)

// 	// fmt.Printf("www: %v ", qq)
// 	// fmt.Println("")
// 	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 3}
//     home := work.Homework1(list1)
// 	fmt.Println(home)
// 	// aa := d0608.Test1(list1)
// 	// fmt.Print(aa)

// }

// // func maptest(){

// // }
// // ///
// // func aaa() {
// // 	// newId := xid.New()
// // 	// _ = newId

// // 	test1 := make([]int, 0, 3) //預建記憶體空間
// // 	test1 = append(test1, 1)
// // 	test1 = append(test1, 2)
// // 	test1 = append(test1, 3)
// // 	test1 = append(test1, 4)
// // 	for i := 0; i < 4; i++ {

// // 		fmt.Printf("test1: %v", test1[i])
// // 		fmt.Println("")
// // 	}
// // 	fmt.Println("for range")
// // 	for index, v := range test1 {
// // 		fmt.Printf("test1: index:%v, v:%v", index, v)
// // 		fmt.Println("")
// // 	}
// // 	// for index, v := range test1 {
// // 	// 	fmt.Printf("index: %v, v: %v", index, v)
// // 	// 	fmt.Println("")
// // 	// }
// // }

// type function1 struct {
// 	funget string `form:"get"`
// }
// type function2 struct {
// 	funpost string `form:"post"`
// }

// func main() {
// 	route := gin.Default()
// 	route.GET("/get1", startPage)
// 	route.POST("/post1", startPage)
// 	route.Run(":8085")
// }

// func startPage(c *gin.Context) {
// 	var f1 function1
// 	var f2 function2
// 	// If `GET`, only `Form` binding engine (`query`) used.
// 	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
// 	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
// 	if c.ShouldBind(&f1) == nil {
// 		log.Println(f1.funget)
// 		fmt.Printf(f1.funget)
// 	}

// 	if c.ShouldBind(&f2) == nil {
// 		log.Println(f2.funpost)

// 	}

// 	c.String(200, "Success")
// }

type A1 struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

type A2 struct {
	A string `form:"f1"`
	B string `form:"f2"`
	C string `form:"f3"`
}

type A3 struct {
	K1 string `json:"a1"`
	K2 string `json:"a2"`
}
type A4 struct {
	Q1 string `uri:"gg" `
	Q2 string `json:"a1"`
	Q3 string `json:"a2" `
}
type RDBS struct {
	ADDVALUE int `form:ADDVALUE`
}
type Bank struct {
	Amount   int64  `json:"amount"`
	PlayerID string `json:"playerid"`
}
type F_Bank struct {
	Amount   float64 `json:"amount"`
	PlayerID string  `json:"playerid"`
}

func main() {
	//初始化
	count.Init()
	router := gin.Default()
	// router.GET("/test1", function1)
	router.POST("/test3", function3)
	router.POST("/test4/:gg", function4)
	// function5(router)
	// GetRdb := redistest.InitRedisClinet()
	// GetRdb.Example()
	router.GET("/test1", function1)
	router.GET("/test2", function2)
	router.GET("/v1/getredis", GetRedis)
	router.POST("/v1/addredis", PostRedis)
	router.POST("/bank", PostRedisBody)
	router.POST("/F_bank", F_PostRedisBody)
	router.POST("/SetValue", SetRedisValue)
	router.POST("/CountsAPI/:playerID", CountsAPI)
	router.POST("/CountsAPI2/:playerID", CountsAPI)
	router.POST("/CountsAPI3/:playerID", CountsAPI)
	// router.Use(Logger1())
	// router.Use(Middleware1())

	// function6(router)
	// function7(router)
	// r:=gin.New()
	// router.POST("/test8", function8)

	router.Run(":8080")
}

// 8/23作業
func CountsAPI(c *gin.Context) {

	playerID := c.Param("playerID")
	//saveData := count.Save(playerID)
	GetRdb := redistest.InitRedisClinet()
	data := GetRdb.PlayerCounts(playerID)
	println(playerID)
	println(data)

	// 調用 SetTimes 函數，設置計時器，60 表示 60 秒
	timesUp := GetRdb.SetTimes(playerID, 10)
	if timesUp != nil {
		c.JSON(500, gin.H{"error": "無法設置計時器"})
		return
	}

	if data <= 10 {
		c.JSON(200, gin.H{"playerID": playerID, "counts": data})
	} else {
		c.JSON(200, "不可超過十次")
	}

}

// 第一題：用GET，不管內容值，只回傳OK
func function1(c *gin.Context) {
	var person A1
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	c.String(200, "ok")
}

// 第二題：用GET，回傳json a b c 及內容值
func function2(c *gin.Context) {
	var data2 A2
	if c.ShouldBind(&data2) == nil {
		log.Println(data2.A)
		log.Println(data2.B)
		log.Println(data2.C)
	}
	//叫用Redis
	GetRdb := redistest.InitRedisClinet()
	GetRdb.Example()
	c.JSON(200, gin.H{"a": data2.A, "b": data2.B, "c": data2.C})

}

func GetRedis(c *gin.Context) {

	GetRdb := redistest.InitRedisClinet()
	res := GetRdb.Example()
	c.JSON(200, gin.H{"KEY1": res})
}

func PostRedis(c *gin.Context) {
	//建立連線
	GetRdb := redistest.InitRedisClinet()
	//建立可容納解析用的物件（struct）
	var rdbsdata RDBS
	//解析Request
	c.ShouldBind(&rdbsdata)
	//叫用INCRBY的方法並帶入前端的值
	res := GetRdb.INCRBYExample(rdbsdata.ADDVALUE)
	//Response
	c.JSON(200, gin.H{"KEY2": res})
}

func PostRedisBody(c *gin.Context) {
	GetRdb := redistest.InitRedisClinet()
	var rdbsdata Bank
	c.BindJSON(&rdbsdata)
	res := GetRdb.INCRBYbank(int(rdbsdata.Amount), rdbsdata.PlayerID)
	c.JSON(200, gin.H{"Balance": res, "Before_Balance": res - rdbsdata.Amount})

}

//080作業，練習銀行可存小數點Float，用新的ＡＰＩ

func F_PostRedisBody(c *gin.Context) {
	GetRdb := redistest.InitRedisClinet()
	var rdbsdata F_Bank
	c.BindJSON(&rdbsdata)
	res := GetRdb.INCRBYFLOAT(float64(rdbsdata.Amount), rdbsdata.PlayerID)
	c.JSON(200, gin.H{"Blance": res, "BeforeBalance": res - rdbsdata.Amount})
}

// 0808 作業2，練習用ＳＥＴ
func SetRedisValue(c *gin.Context) {
	GetRdb := redistest.InitRedisClinet()
	var rdbsdata F_Bank
	c.BindJSON((&rdbsdata))
	res := GetRdb.SetValue(int64(rdbsdata.Amount), rdbsdata.PlayerID)
	c.JSON(200, gin.H{"Blance": res})
}

// 第三題：用POST，Body帶參數k1 k2，回傳json a1 a2（a1為k1的內容值，a2為k2的內容值）
func function3(c *gin.Context) {
	var data3 A3
	if c.ShouldBindJSON(&data3) == nil {
		log.Println(data3.K1)
		log.Println(data3.K2)
	}
	c.JSON(200, gin.H{
		"K1": data3.K1,
		"K2": data3.K2,
	})
}

// 第四題：用POST，router帶參數q1 q2 q3，Body帶參數k1 k2，回傳json q1 q2 q3 k1 k2
func function4(c *gin.Context) {
	var data4 A4
	if c.ShouldBindUri(&data4) == nil {
		log.Println(data4)
	}

	c.ShouldBindJSON(&data4)
	c.JSON(200, gin.H{
		"QQ1": data4.Q1,
		"QQ2": data4.Q2,
		"QQ3": data4.Q3,
	})
}

// 第五題.group =>  參考官網 ,fun名稱自訂，router名稱和官網一樣，responce回router名稱,前面都加v1，用Get
func function5(router *gin.Engine) {
	// router :=gin.Default()
	v1 := router.Group("v1")
	{
		v1.GET("/login", function5_1)
		v1.GET("/submit", function5_2)
		v1.GET("/read", function5_3)
	}
}
func function5_1(c *gin.Context) {
	LL := c.MustGet("latency").(string)
	c.JSON(200, gin.H{"v1/login": "", "Middleware": LL})
}
func function5_2(c *gin.Context) {
	c.JSON(200, gin.H{"v1/submit": ""})
}
func function5_3(c *gin.Context) {
	c.JSON(200, gin.H{"v1/read": ""})
}

func Logger1() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("yyyy", 5555)
		c.Next()
	}
}

// 第六題.加冒號KEY，回傳為冒號KEY內容
func function6(router *gin.Engine) {
	v2 := router.Group("v2")
	{
		v2.GET(":login", function6_1)
	}

}
func function6_1(c *gin.Context) {
	example := c.MustGet("yyyy").(int)
	c.JSON(200, gin.H{"v2/login": "", "yyyy": example})

}

// 第七題.三層的group=>r1,r2,r3
func function7(router *gin.Engine) {
	rr1 := router.Group("r1")
	{
		rr2 := rr1.Group("/r2")
		{
			rr3 := rr2.Group("/r3")
			{
				rr3.GET("Index", function7_1)
				rr3.POST("Home", function7_2)
				rr3.PATCH("router", function7_3)
			}
		}
	}
}
func function7_1(c *gin.Context) {
	example := c.MustGet("yyyy").(int)
	c.JSON(200, gin.H{"r1/r2/r3/Index": "", "yyyy": example})
}
func function7_2(c *gin.Context) {
	example := c.MustGet("yyyy").(int)
	c.JSON(200, gin.H{"r1/r2/r3/Home": "", "yyyy": example})
}
func function7_3(c *gin.Context) {
	example := c.MustGet("yyyy").(int)
	c.JSON(200, gin.H{"r1/r2/r3/router": "", "yyyy": example})
}

// 第八題. QueryMap 參考並製作出相同練習
func function8(c *gin.Context) {

	//params
	ids := c.QueryMap("ids")

	//body
	names := c.PostFormMap("names")

	c.JSON(200, gin.H{
		"ids1":  ids,
		"name1": names,
	})
}

// 第九集：Middleware(使用時間)
type BodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func Middleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		wb := &BodyWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}
		c.Writer = wb

		c.Next()

		latency := time.Since(t)
		log.Print(latency)

		data := wb.body.String()

		obj := make(map[string]interface{})

		json.Unmarshal([]byte(data), &obj)

		obj["時間"] = latency.String()
		obj["work"] = "NEW_WORK"

		updatedBody, _ := json.Marshal(obj)

		wb.ResponseWriter.WriteString(string(updatedBody))
		wb.body.Reset()

	}
}
