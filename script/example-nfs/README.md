自动创建nfs目录，原始shell

```shell
#!/bin/bash
set -e
echo "当前分享的目录:"
NFS_IP="10.29.0.149"
NFS_DIR="/nfs_ssd"
showmount -e $NFS_IP

dir_name=$1
if [ "$dir_name" == "" ];then
read -p "请输入一个新的共享目录名称, eg: registry_10.29.100.1 :" dir_name
fi
if [ -d $NFS_DIR/$dir_name ] ;then
  echo "dir have been used,try again "
  exit 0
fi


mkdir $NFS_DIR/$dir_name
echo "$NFS_DIR/$dir_name  10.29.0.0/8(rw,sync,no_root_squash)" >>/etc/exports
systemctl restart nfs
showmount -e $NFS_IP
echo "look like successed, use \" showmount -e $NFS_IP \" at clinet to test"
```