package main

import (
	"fmt"
	"github.com/dengpju/higo-utils/utils/fileutil"
)

type Test struct {
	a string `json:"a"`
	B int    `json:"b"`
}

var s = `{"0":[{"AlwaysShow":false,"component":"Layout","frontRouteId":1,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"affix":false,"icon":"fas-home","noClosable":false,"noKeepAlive":"","tabHidden":true,"title":"首页"},"name":"Root","parentId":0,"path":"/","redirect":"/index","sort":0},{"alwaysShow":false,"component":"Layout","frontRouteId":3,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":true,"title":"儿童评价库"},"name":"ChildrenEvaluation","parentId":0,"path":"/childrenEvaluation","redirect":"/childrenEvaluation/index","sort":0},{"alwaysShow":false,"component":"Layout","frontRouteId":6,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":true,"title":"统计分析"},"name":"Statistics","parentId":0,"path":"/statistics","redirect":"/statistics/index","sort":0},{"alwaysShow":false,"component":"Layout","frontRouteId":11,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":true,"title":"成长档案"},"name":"File","parentId":0,"path":"/growthFile","redirect":"/growthFile/platform","sort":0},{"alwaysShow":false,"component":"Layout","frontRouteId":23,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dynamicNewTab":false,"icon":"","noKeepAlive":true,"tabHidden":true,"title":"一日流程"},"name":"OneDayProcess","parentId":0,"path":"/oneDayProcess","redirect":"/oneDayProcess/oneDayEdit","sort":0},{"alwaysShow":false,"component":"Layout","frontRouteId":17,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":true,"title":"管理"},"name":"Management","parentId":0,"path":"/management","redirect":"/management/manageRoles","sort":1}],"1":[{"alwaysShow":false,"component":"@views/index","frontRouteId":2,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":true,"badge":"","dot":false,"dynamicNewTab":true,"icon":"icon-home-3-line","noClosable":true,"noKeepAlive":true,"tabHidden":false,"title":"首页"},"name":"Index","parentId":1,"path":"index","redirect":"","sort":0}],"11":[{"alwaysShow":false,"component":"@/views/growthFile/park","frontRouteId":12,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"园区信息管理"},"name":"Park","parentId":11,"path":"park","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/growthFile/class","frontRouteId":13,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"班级信息管理"},"name":"Class","parentId":11,"path":"class","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/growthFile/list","frontRouteId":14,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"档案列表"},"name":"list","parentId":11,"path":"FileList","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/growthFile/fileEditor","frontRouteId":15,"hidden":true,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"档案编辑"},"name":"FileEditor","parentId":11,"path":"fileEditor","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/growthFile/archivePreview","frontRouteId":16,"hidden":true,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"档案预览"},"name":"ArchivePreview","parentId":11,"path":"archivePreview","redirect":"","sort":0}],"17":[{"alwaysShow":false,"component":"@/views/management/evaluationRuleEdit","frontRouteId":18,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"评价规则编辑"},"name":"EvaluationRuleEdit","parentId":17,"path":"evaluationRuleEdit","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/management/supportStrategy","frontRouteId":19,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"支持策略查看"},"name":"SupportStrategy","parentId":17,"path":"supportStrategy","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/management/manageRoles","frontRouteId":20,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"角色管理"},"name":"ManageRoles","parentId":17,"path":"manageRoles","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/management/manageUsers","frontRouteId":21,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"用户管理"},"name":"ManageUsers","parentId":17,"path":"manageUsers","redirect":"","sort":0}],"23":[{"alwaysShow":false,"component":"@/views/oneDayProcess/oneDayEdit","frontRouteId":24,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dynamicNewTab":false,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"流程编辑"},"name":"OneDayEdit","parentId":23,"path":"oneDayEdit","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/oneDayProcess/oneDayDetail","frontRouteId":25,"hidden":true,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dynamicNewTab":false,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"流程编辑"},"name":"OneDayDetail","parentId":23,"path":"oneDayDetail","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/oneDayProcess/kdiEdit","frontRouteId":26,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dynamicNewTab":false,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"KDI编辑"},"name":"KdiEdit","parentId":23,"path":"kdiEdit","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/oneDayProcess/audit","frontRouteId":27,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dynamicNewTab":false,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"模板审核"},"name":"Audit","parentId":23,"path":"audit","redirect":"","sort":0}],"3":[{"alwaysShow":false,"component":"@/views/childrenEvaluation","frontRouteId":4,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"儿童评价库"},"name":"Evaluation","parentId":3,"path":"index","redirect":"","sort":1},{"alwaysShow":false,"component":"@/views/childrenEvaluation/myEvaluation","frontRouteId":5,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"我的评价库"},"name":"MyEvaluation","parentId":3,"path":"myEvaluation","redirect":"","sort":4}],"6":[{"alwaysShow":false,"component":"@/views/statistics","frontRouteId":7,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"学区统计"},"name":"SchoolDistrict","parentId":6,"path":"index","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/statistics/park","frontRouteId":8,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"园区统计"},"name":"ParkStatistics","parentId":6,"path":"parkStatistics","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/statistics/class","frontRouteId":9,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"班级统计"},"name":"ClassStatistics","parentId":6,"path":"classStatistics","redirect":"","sort":0},{"alwaysShow":false,"component":"@/views/statistics/student","frontRouteId":10,"hidden":false,"interiorSort":"","menuHidden":false,"meta":{"activeMenu":"","affix":false,"badge":"","dot":false,"dynamicNewTab":true,"icon":"","noKeepAlive":true,"tabHidden":false,"title":"学生统计"},"name":"StudentStatistics","parentId":6,"path":"studentStatistics","redirect":"","sort":0}]}`

func main() {
	/**
	fmt.Println(utils.GmtTime("2021-08-25T16:52:21+08:00"))
	//fmt.Println(utils.JsonKeyToCase(s))
	return
	*/
	/**
	//map
	fmt.Println(utils.Dir("./utils").Scan().Suffix("go").Get())
	result := utils.MapOperation(make(map[string]interface{})).Put("key1", "value1").Put("key2", "value2")
	fmt.Println(result)
	result.ForEach(func(key string, value interface{}) {
		fmt.Println(key, value)
	})
	result.Replace("key1", "k")
	fmt.Println(result.Get("key1"))
	result.Remove("k")
	result.Remove("key2")
	fmt.Println(result.Len())
	fmt.Println(result)
	result.Clear()
	result.ForEach(func(key string, value interface{}) {
		fmt.Println(key, value, "for")
	})
	*/
	/**
	a := utils.Array().Push("11").Push("22").Push("33")
	fmt.Println(a)
	fmt.Println(a.Current())
	fmt.Println(a)
	fmt.Println(a.End())
	fmt.Println(a)
	a.Remove("1")
	fmt.Println(a)
	fmt.Println(a.Value())
	s := make([]string, 0)
	s = append(s, "1")
	s = append(s, "2")
	s = append(s, "3")
	fmt.Println(fmt.Sprintf("%s", s))

	*/

	/**
	// Rsa
	rsa := utils.NewRsa().SetBits(1024).Build()
	fmt.Println("当前时间戳", utils.Time())
	rsa.SetExpired(utils.Time() + 6)
	rsa.SetLimen(10)
	fmt.Println(rsa.GetFlag())

	// 公钥加密
	e := utils.PubEncrypt(rsa, []byte("123"))
	fmt.Println("公钥加密===")
	fmt.Println(base64.StdEncoding.EncodeToString(e))
	fmt.Println("私钥解密===")
	// 私钥解密
	s := utils.PriDecrypt(rsa, e).String()
	fmt.Println(s)
	fmt.Println("=====")
	time.Sleep(5 * time.Second)
	if i := utils.SecretContainer.Len(); i > 0 {
		fmt.Println("有", i, "个秘钥对")
	}
	utils.SecretContainer.ForEach(func(key string, value interface{}) {
		fmt.Println(key, value.(*utils.Rsa).GetExpired())
		if utils.Time() >= value.(*utils.Rsa).GetExpired() {
			fmt.Println("秘钥对过期了")
			utils.SecretContainer.Remove(key)
			fmt.Println("删除过期秘钥对")
		}
	})
	utils.SecretExpiredClear()
	if i := utils.SecretContainer.Len(); i <= 0 {
		fmt.Println("有", i, "个秘钥对")
	}

	// 私钥加密
	s1 := utils.PriEncrypt(rsa, []byte("789")).Base64Encode()
	fmt.Println(s1)
	fmt.Println(s1.Base64Decode())
	fmt.Println(s1.Base64Decode().Base64Encode().Base64Decode())
	// 公钥解密
	ss := utils.PubDecrypt(rsa, s1.Base64Decode()).String()
	fmt.Println(ss)
	fmt.Println(utils.SecretContainer.Exist(rsa.GetFlag()))
	if utils.SecretContainer.Exist(rsa.GetFlag()) {
		fmt.Println(utils.SecretContainer.Get(rsa.GetFlag()).(*utils.Rsa).GetFlag())
	}
	*/

	/**
	//Strtotime
	fmt.Println(utils.Time())
	fmt.Println(utils.Date(utils.Time(), "Y/m/d"))
	fmt.Println(utils.Strtotime("2021-03-08 22:19:30"))
	fmt.Println(utils.Strtotime("-2day 1hour +1 minute"))

	month := time.Now().Month()
	year := time.Now().Year()

	switch month {
	case time.April, time.June, time.September, time.November:
		//if day > 30 {
		//	return false
		//}
		fmt.Println(30)
	case time.February:
		// leap year
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			//if day > 29 {
			//	return false
			//}
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	default:
		fmt.Println(31)
	}

	//fmt.Println(time.Second * time.Duration(1 * 60))

	*/

	/**
	//flysnowRegexp := regexp.MustCompile(`(\+|\-|)\s*(\d+)\s*(year|mouth|day|hour|minute|second)`)
	flysnowRegexp := regexp.MustCompile(`(\+|\-|)\s*(\d+)\s*(year|mouth|day|hour|minute|second)`)
	params := flysnowRegexp.FindAllStringSubmatch("-2day +1 hour +1 minute", -1)

	for _, param := range params {
		for _, p := range param {
			fmt.Println(p)
		}
	}

	*/

	// File
	f := &fileutil.File{Name: "fff"}
	_ = f.ForEach(func(line int, b []byte) {

	})
	fmt.Println(f.SplitFunc == nil)
	/**
	d := utils.Dir("./test/yy/dd")
	d.Create()
	return
	fmt.Println(utils.Basename("testweb\\home.php", ".php"))
	fmt.Println(utils.Dirname("\\00\\testweb\\home.php"))
	fmt.Println(utils.Dirname("\\00\\11\\testweb"))
	fmt.Println(utils.Dirslice(".\\00\\11\\testweb\\home.php"))
	fmt.Println(utils.DirBasename(".\\00\\11\\testweb\\home.php"))
	fmt.Println(utils.Mkdir(".\\00\\11\\testweb", 0666))
	//fmt.Println(utils.Remove(".\\00\\11\\11"))
	fmt.Println(utils.Emdir(".\\00\\11"))
	//fmt.Println(utils.Rmdir(".\\00\\11"))
	fmt.Println(utils.Pathstring([]string{".", "data", "user"}))
	utils.NewFile(".\\00\\22").File().WriteString("ggg")
	fmt.Println(utils.DirExist(".\\00\\111"))
	*/

	/**
	// Random
	fmt.Println(utils.Random().Int(1000))
	fmt.Println(utils.Random().Int64(1000))
	fmt.Println(utils.Random().BetweenInt(800, 1000))
	fmt.Println(utils.Random().BetweenInt64(800, 1000))

	*/

	/**
	// Slice
	sl := utils.NewSliceString()
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.Append("2")
	sl.Insert(0, "1")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.Append("3")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) {
		fmt.Println(index, value)
	})
	sl.Insert(1, "11")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) {
		fmt.Println(index, value)
	})
	sl.Insert(3, "114")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) {
		fmt.Println(index, value)
	})
	sl.Remove("114")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) {
		fmt.Println(index, value)
	})
	sl.Delete(1)
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) {
		fmt.Println(index, value)
	})
	sl.Replace("1", "11")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) {
		fmt.Println(index, value)
	})
	sl.Reverse()
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) {
		fmt.Println(index, value)
	})
	fmt.Println(sl.Exist("3"))
	fmt.Println(sl.Value().([]string)[2:])
	sl.Empty()
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) {
		fmt.Println(index, value)
	})

	*/

}
