package main

import (
	"cloudstorage-server/store/ceph"
	"fmt"
	"os"

	"gopkg.in/amz.v1/s3"
)

func main() {
	bucket := ceph.GetCephBucket("userfile")
	fmt.Println("Connect to bucket userfile")
	err := bucket.PutBucket(s3.PublicRead) // 参数权限

	if err != nil {
		fmt.Println(err.Error())
	}
	d, err := bucket.Get("e03dd1f57358bbca20794f627d76756c27ca680e")
	if err != nil {
		fmt.Printf("get object err: %v\n", err)
		return
	}
	tmpFile, _ := os.Create("/tmp/test_file")
	tmpFile.Write(d)
	return

	// // 创建一个新的bucket
	// err := bucket.PutBucket(s3.PublicRead)
	// fmt.Printf("create bucket err: %v\n", err)

	// //查询这个bucket下面指定条件的object keys
	// res, _ := bucket.List("", "", "", 100)
	// fmt.Printf("object keys: %+v\n", res)

	// // 新上传一个对象
	// //err = bucket.Put("/testupload/a.txt", []byte("just for test"), "octet-stream", s3.PublicRead)
	// err = bucket.Put("/a.txt", []byte("just for tests"), "octet-steam", s3.PublicRead)
	// fmt.Printf("upload err: %+v\n", err)

	// // 查询这个bucket下面指定条件的object keys
	// res, err = bucket.List("", "", "", 100)
	// fmt.Printf("object keys: %+v\n", res)
}
