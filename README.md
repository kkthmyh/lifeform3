# lifeform NFT mint说明
## 1、配置相关信息
修改根目录下的config.yaml文件，按照注释修改
## 2、添加私钥和地址
可以使用批量生成钱包的软件批量生成地址和私钥如

https://cointool.app/createWallet/eth

使用上述工具对每个新钱包进行批量转账,实际测试每次mint消耗0.36u左右，为保证mint成功建议每个钱包转0.5u+

将地址导出粘贴到./data/address.csv中

将私钥导出粘贴到./data/privateKey.csv中

注意:每行填一个地址或一个私钥，私钥不需要加0x


## 3、启动main.go
查看控制台的mint结果

