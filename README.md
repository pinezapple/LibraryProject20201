# Website Quản lý thư viện trường
## Thông tin đề tài
**Học phần** : IT4556Q-Các phương pháp phát triển phần mềm nhanh

**Mã lớp chuyên ngành** : VUWIT16A

**Thành viên** :
- Nguyễn Thanh Tùng - 20176121
- Phạm Thanh Hằng - 20176085

**Ngôn ngữ và framework sử dụng** :
- Golang, Python, proto3, bash (backend).
- Bootstrap, Javascript (frontend).

**Các opensource softwares sử dụng**:
- [etcd](https://etcd.io/)
- [gRPC](https://grpc.io/)
- [cassandra](https://cassandra.apache.org/)

**Cách cài đặt các dependencies và deploy hệ thống**:
***Cài đặt etcd***:
- Tải các bản release của etcd [tại đây](https://github.com/etcd-io/etcd/releases).
- Chạy etcd:
```
cd /your/etcd/file
./bin/etcd
```
***Cài đặt cassandra***:
- Cassandra được cài đặt như [link](https://cassandra.apache.org/doc/latest/getting_started/installing.html).
- Cassandra sẽ start bằng lệnh:
```
systemctl start cassandra.service
```
***Cài đặt gRPC cho Golang***:
- Để đặt gRPC thực hiện như [link](https://grpc.io/docs/languages/go/quickstart/)

***Deploy hệ thống***:
- Clone code từ github:
```
git clone https://github.com/pinezapple/LibraryProject20201.git
```
- Push config của các service lên etcd:
```
cd /your/cloned/files
cd etcdConfigs && python3 pushConfigs.py
```
- Chạy các service docmanager:
```
cd /your/cloned/files
./docmanager/docmanger -shard_id=`service shard id`
```
- Chạy service portal:
```
cd /your/cloned/files
./portal/portal -shard_number=`your total shards number`
```

**Các tài liệu khác** :
- [Phân tích yêu cầu phần mềm](https://drive.google.com/open?id=1XAD5SHuuKTj9p12JePVix_pfcdow-a6jhYzckpSOWxE)
- [Sprint planning](https://docs.google.com/spreadsheets/d/1rQJ7jLo0bQ3YLJ5W82lS9-sIo_PXcWvu6X4y02v3euc/edit?usp=sharing)
- [QnA File](https://drive.google.com/file/d/1mIAgsAoSo6Jj55SHId-utsZTrNTX_O9B/view?usp=sharing)
- [Tài liệu phát triển phần mềm](https://github.com/pinezapple/LibraryProject20201/tree/master/Documents)
