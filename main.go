package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"umarutv/config"
	"umarutv/route"
)

var App *gin.Engine
func main() {
	App = gin.New()

	route.RegisterApiRoute(&App.RouterGroup)

	err := App.Run(config.App.ServerAddr)

	if err != nil {
		log.Fatalln(err)
	}


	//origin, _ := os.OpenFile("/home/ricardo/Desktop/9753371416055175852.png", os.O_RDWR, 0666)
	//target, _ := os.Open("/home/ricardo/Desktop/filename0.ts")
	//merged, _ := os.OpenFile("/home/ricardo/Desktop/merged.png", os.O_CREATE | os.O_WRONLY, 0666)
	//defer func() {
	//	origin.Close()
	//	target.Close()
	//	merged.Close()
	//}()
	//
	//originData, _ := ioutil.ReadAll(origin)
	//targetData, _ := ioutil.ReadAll(target)
	//// 写入文件头
	//merged.Write(originData[:33])
	//// Write tEXT chunk size
	//binary.Write(merged, binary.BigEndian, uint32(len(targetData)))
	//// Write chunk name
	//merged.Write([]byte("tEXT"))
	//// Write data
	//merged.Write(targetData)
	//// Write CRC code
	//binary.Write(merged,binary.BigEndian,crc32.ChecksumIEEE(targetData))
	//// Write other data
	//merged.Write(originData[33:])

	// 跳过标识头
	//_, _ = origin.Seek(8, 0)
	//
	//blockLength := make([]byte, 4)
	//_, err := origin.Read(blockLength)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//// 跳过 IHDR block
	////_, _ = origin.Seek(int64(binary.BigEndian.Uint32(blockLength)+8), 1)
	//block := make([]byte, binary.BigEndian.Uint32(blockLength) + 4)
	//
	//_, _ = origin.Read(block)
	//
	//ieee := crc32.ChecksumIEEE(block)
	//
	//err = binary.Write(origin, binary.BigEndian, ieee)
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//_, _ = origin.Write()



	//for err == nil {
	//	l := binary.BigEndian.Uint32(blockLength)
	//	// block 的实际长度为：标识头（4 bytes） + body + 校验码（4 bytes）
	//	block := make([]byte, l + 8)
	//	_, err = origin.Read(block)
	//
	//	if err != nil {
	//		break
	//	}
	//
	//	blockLength = make([]byte, 4)
	//	_, err = origin.Read(blockLength)
	//}

}
