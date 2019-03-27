package main
import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)
func main() {
	//telnetServer()
	race()
}

func telnetServer()  {
	// 创建一个程序结束码的通道
	exitChan := make(chan int)
	// 将服务器并发运行
	go server("127.0.0.1:7001", exitChan)
	// 通道阻塞, 等待接收返回值
	code := <-exitChan
	// 标记程序返回值并退出
	os.Exit(code)
}

func slice()  {
	// 设置元素数量为1000
	const elementCount = 20
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)

	// 引用切片数据
	refData := srcData
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	copyData := make([]int, elementCount)
	copy(copyData, srcData[2:5])
	fmt.Println(refData, srcData)
	fmt.Println(copyData)

	seq := []string{"a", "b", "c", "d", "e"}
	// 指定删除位置
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
	dic := make(map[int]string)
	dic[1] = "a"
	a, ok := dic[1]
	fmt.Println(a, ok)
}

func point()  {
	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)
}

type Weapon int

func enumMock(){
	const (
		Arrow Weapon = iota    // 开始生成枚举值, 默认为0
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	// 输出所有枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon)
}

func sinPic() {
	// 图片大小
	const size = 300
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	// 从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		// 让sin的值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	// 创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, pic)
	//将image信息写入文件中
	// 关闭文件
	file.Close()
}