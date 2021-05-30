CREATE TABLE `blocks` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `blockidentifier` varchar(150) NOT NULL DEFAULT '' COMMENT '当前区块',
  `parentblock` varchar(100) NOT NULL DEFAULT '' COMMENT '上一个区块',
  `transactiohash` varchar(100) NOT NULL DEFAULT '' COMMENT '转账hash值',
  `mblockheight` varchar(100) NOT NULL DEFAULT '' COMMENT '区块metadata高度',
  `mmemo` varchar(200) NOT NULL DEFAULT '' COMMENT '区块备注',
  `mtimestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '区块metadata时间',
  `blocktimestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '当前区块时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='区块转账记录表';