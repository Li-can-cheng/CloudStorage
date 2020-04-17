#  Golang Cloud Storage
- Configure MySQL  at config/db.go

  ```bash
  MySQLSource = "user:password@tcp(127.0.0.1:3306)/<database>?charset=utf8"
  ```

- Configure  Ceph Object Gateway at  config/ceph.go

	```bash
	CephAccessKey = <AccessKey>
	CephSecretKey = <SecretKey>
	CephGWEndpoint = "http://127.0.0.1:<PORT>"
	```
	
- Run server
	```bash
	go run main.go
	```