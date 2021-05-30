CREATE TABLE `details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `Tranidentifier` varchar(200) NOT NULL DEFAULT '' COMMENT '转账hash值',
  `oindex` varchar(20) NOT NULL DEFAULT '' COMMENT '当前这笔转账的hash索引，如果数据index为多个则分出多条数据',
  `otype` varchar(100) NOT NULL DEFAULT '' COMMENT '转账类型',
  `ostatus` varchar(100) NOT NULL DEFAULT '' COMMENT '转账状态',
  `oaccountaddress` varchar(180) NOT NULL DEFAULT '' COMMENT '转账账户地址',
  `oamountvalue` varchar(250) NOT NULL DEFAULT '' COMMENT '转账价值，可能为负数',
  `oamountcurrencysymbol` varchar(50) NOT NULL DEFAULT '' COMMENT '转账代币缩写',
  `oamountcurrencydecimals` varchar(50) NOT NULL DEFAULT '' COMMENT '转账代币位数',
  `mblockheight` varchar(100) NOT NULL DEFAULT '' COMMENT '区块metadata高度',
  `mmemo` varchar(200) NOT NULL DEFAULT '' COMMENT '区块备注',
  `mtimestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '区块metadata时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='区块转账详情表';