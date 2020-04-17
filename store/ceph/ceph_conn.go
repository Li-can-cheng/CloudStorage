package ceph

import (
	cfg "cloudstorage-server/config"
	"fmt"

	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)

var cephConn *s3.S3

// // GetCephUrl : 获取ceph Url连接
func GetCephUrl(bucket string, filename string) string {
	cephUrl := fmt.Sprintf("%s/%s/%s", cfg.CephGWEndpoint, bucket, filename)
	return cephUrl
}

// //Download CephData : 取回ceph数据
// func DownloadCephData(bucketString string, filename string, w http.ResponseWriter) {
// 	bucket := GetCephBucket(bucketString)
// 	_ = bucket.PutBucket(s3.PublicRead)
// 	d, err := bucket.Get(filename)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/octect-stream")
// 	// attachment表示文件将会提示下载到本地，而不是直接在浏览器中打开
// 	w.Header().Set("content-disposition", "attachment; filename=\""+filename+"\"")
// 	w.Write([]byte(d))
// }

// GetCephConnection : 获取S3 ceph连接
func GetCephConnection() *s3.S3 {
	if cephConn != nil {
		return cephConn
	}
	// 1. 初始化ceph的一些信息

	auth := aws.Auth{
		AccessKey: cfg.CephAccessKey,
		SecretKey: cfg.CephSecretKey,
	}

	curRegion := aws.Region{
		Name:                 "default",
		EC2Endpoint:          cfg.CephGWEndpoint,
		S3Endpoint:           cfg.CephGWEndpoint,
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
		Sign:                 aws.SignV2,
	}

	// 2. 创建S3类型的连接
	return s3.New(auth, curRegion)
}

// GetCephBucket : 获取指定的bucket对象
func GetCephBucket(bucket string) *s3.Bucket {
	conn := GetCephConnection()
	return conn.Bucket(bucket)
}

// PutObject : 上传文件到ceph集群
func PutObject(bucket string, path string, data []byte) error {
	return GetCephBucket(bucket).Put(path, data, "octet-stream", s3.PublicRead)
}
