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
	// ����һ������������ͨ��
	exitChan := make(chan int)
	// ����������������
	go server("127.0.0.1:7001", exitChan)
	// ͨ������, �ȴ����շ���ֵ
	code := <-exitChan
	// ��ǳ��򷵻�ֵ���˳�
	os.Exit(code)
}

func slice()  {
	// ����Ԫ������Ϊ1000
	const elementCount = 20
	// Ԥ�����㹻���Ԫ����Ƭ
	srcData := make([]int, elementCount)

	// ������Ƭ����
	refData := srcData
	// ����Ƭ��ֵ
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	copyData := make([]int, elementCount)
	copy(copyData, srcData[2:5])
	fmt.Println(refData, srcData)
	fmt.Println(copyData)

	seq := []string{"a", "b", "c", "d", "e"}
	// ָ��ɾ��λ��
	index := 2
	// �鿴ɾ��λ��֮ǰ��Ԫ�غ�֮���Ԫ��
	fmt.Println(seq[:index], seq[index+1:])
	// ��ɾ����ǰ���Ԫ����������
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
	dic := make(map[int]string)
	dic[1] = "a"
	a, ok := dic[1]
	fmt.Println(a, ok)
}

func point()  {
	// ׼��һ���ַ�������
	var house = "Malibu Point 10880, 90265"
	// ���ַ���ȡ��ַ, ptr����Ϊ*string
	ptr := &house
	// ��ӡptr������
	fmt.Printf("ptr type: %T\n", ptr)
	// ��ӡptr��ָ���ַ
	fmt.Printf("address: %p\n", ptr)
	// ��ָ�����ȡֵ����
	value := *ptr
	// ȡֵ�������
	fmt.Printf("value type: %T\n", value)
	// ָ��ȡֵ�����ָ�������ֵ
	fmt.Printf("value: %s\n", value)
}

type Weapon int

func enumMock(){
	const (
		Arrow Weapon = iota    // ��ʼ����ö��ֵ, Ĭ��Ϊ0
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	// �������ö��ֵ
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	// ʹ��ö�����Ͳ�����ֵ
	var weapon Weapon = Blower
	fmt.Println(weapon)
}

func sinPic() {
	// ͼƬ��С
	const size = 300
	// ���ݸ�����С�����Ҷ�ͼ
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// ����ÿ������
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// ���Ϊ��ɫ
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	// ��0�������������x����
	for x := 0; x < size; x++ {
		// ��sin��ֵ�ķ�Χ��0~2Pi֮��
		s := float64(x) * 2 * math.Pi / size
		// sin�ķ���Ϊһ������ء�����ƫ��һ�����ز���ת
		y := size/2 - math.Sin(s)*size/2
		// �ú�ɫ����sin�켣
		pic.SetGray(x, int(y), color.Gray{0})
	}
	// �����ļ�
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	// ʹ��png��ʽ������д���ļ�
	png.Encode(file, pic)
	//��image��Ϣд���ļ���
	// �ر��ļ�
	file.Close()
}