package main

import (
	"awesomeProject/common"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	//config, err := rest.InClusterConfig()
	//if err != nil {
	//	panic(err.Error())
	//}
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err.Error())
	//}
	//// 创建 ConfigMap 对象，并设置 MySQL 连接参数
	//data := make(map[string]string)
	//data["hostname"] = "mysql-service"
	//data["port"] = "3306"
	//data["username"] = "root"
	//data["password"] = "root"
	//data["dbname"] = "exam"
	//configMap := &corev1.ConfigMap{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name:      "mysql-parameters",
	//		Namespace: "default",
	//	},
	//	Data: data,
	//}
	//// 创建或更新 ConfigMap
	//_, err = clientset.CoreV1().ConfigMaps("default").Update(context.Background(), configMap, metav1.UpdateOptions{})
	//if err != nil {
	//	_, err = clientset.CoreV1().ConfigMaps("default").Create(context.Background(), configMap, metav1.CreateOptions{})
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//}
	//fmt.Println("MySQL 参数存储成功！")
	//初始化连接数据库配置
	InitConfig()
	//连接数据库
	db := common.Initdb()
	defer db.Close() //程序退出关闭数据库
	r := gin.Default()
	//告诉gin框架模板文件去哪里找
	//r.Static("/xxx", "./statics")
	//r.Static("/xx", "./templates")
	r.LoadHTMLGlob("templates/**/*")
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
