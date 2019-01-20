/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.7.24-log : Database - bank_big_data
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`bank_big_data` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `bank_big_data`;

/*Table structure for table `s_ecif_ecif_cert_info_2350000` */

DROP TABLE IF EXISTS `s_ecif_ecif_cert_info_2350000`;

CREATE TABLE `s_ecif_ecif_cert_info_2350000` (
  `bankid` varchar(20) DEFAULT NULL COMMENT '法人编号',
  `cod_cust_id` varchar(20) NOT NULL DEFAULT '-' COMMENT '客户号',
  `cert_type` varchar(20) DEFAULT NULL COMMENT '证件类型',
  `cert_num` varchar(90) DEFAULT NULL COMMENT '证件号码',
  `sign_date` varchar(8) DEFAULT NULL COMMENT '签发日期',
  `effe_date` varchar(8) DEFAULT NULL COMMENT '有效日期',
  `is_fist` varchar(20) DEFAULT NULL COMMENT '是否首选证件',
  `effe_type` varchar(20) DEFAULT NULL COMMENT '有效期类型',
  `cert_dec` varchar(512) DEFAULT NULL COMMENT '证件描述',
  `etl_f` varchar(1) DEFAULT NULL COMMENT '(null)',
  `src_sys_cd` varchar(10) DEFAULT NULL COMMENT '源系统代码',
  `crut_no` varchar(10) DEFAULT NULL COMMENT '法人机构号',
  `busi_dt` varchar(8) DEFAULT NULL COMMENT '业务日期',
  `biz_dt` varchar(8) DEFAULT NULL COMMENT '平台日期',
  `etl_load_time` timestamp NULL DEFAULT NULL COMMENT '数据加载时间戳',
  PRIMARY KEY (`cod_cust_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `s_ecif_ecif_cert_info_2350000` */

/*Table structure for table `s_loan_dk_2350000` */

DROP TABLE IF EXISTS `s_loan_dk_2350000`;

CREATE TABLE `s_loan_dk_2350000` (
  `xd_col1` varchar(22) DEFAULT NULL COMMENT '贷款号',
  `xd_col2` varchar(100) DEFAULT NULL COMMENT '客户名称',
  `xd_col3` timestamp NULL DEFAULT NULL COMMENT '会计日期',
  `xd_col4` timestamp NULL DEFAULT NULL COMMENT '贷款日期',
  `xd_col5` timestamp NULL DEFAULT NULL COMMENT '到期日期',
  `xd_col6` decimal(20,2) DEFAULT NULL COMMENT '贷款金额',
  `xd_col7` decimal(20,2) DEFAULT NULL COMMENT '结欠金额',
  `xd_col8` decimal(14,7) DEFAULT NULL COMMENT '合同利率',
  `xd_col9` decimal(14,7) DEFAULT NULL COMMENT '逾期利率',
  `xd_col10` varchar(20) DEFAULT NULL COMMENT '利率类别',
  `xd_col11` varchar(20) DEFAULT NULL COMMENT '复利标志',
  `xd_col12` varchar(20) DEFAULT NULL COMMENT '出帐种类',
  `xd_col13` varchar(50) DEFAULT NULL COMMENT '编号',
  `xd_col14` varchar(50) DEFAULT NULL COMMENT '客户号',
  `xd_col15` varchar(20) DEFAULT NULL COMMENT '证件类型',
  `xd_col16` varchar(40) DEFAULT NULL COMMENT '证件号码',
  `xd_col17` varchar(20) DEFAULT NULL COMMENT '贷款用途种类',
  `xd_col18` varchar(40) DEFAULT NULL COMMENT '贷款用途',
  `xd_col19` varchar(20) DEFAULT NULL COMMENT '贷款类型',
  `xd_col20` varchar(20) DEFAULT NULL COMMENT '贷款种类',
  `xd_col21` varchar(20) DEFAULT NULL COMMENT '四级分类',
  `xd_col22` varchar(20) DEFAULT NULL COMMENT '五级分类',
  `xd_col23` timestamp NULL DEFAULT NULL COMMENT '展期日期',
  `xd_col24` decimal(10,0) DEFAULT NULL COMMENT '现金转帐',
  `xd_col25` decimal(14,7) DEFAULT NULL COMMENT '贴息比例',
  `xd_col26` varchar(3) DEFAULT NULL COMMENT '结息周期',
  `xd_col27` varchar(20) DEFAULT NULL COMMENT '是否自动扣收本息',
  `xd_col28` timestamp NULL DEFAULT NULL COMMENT '申请日期',
  `xd_col29` varchar(50) DEFAULT NULL COMMENT '贷款帐号',
  `xd_col30` varchar(50) DEFAULT NULL COMMENT '存款帐号',
  `xd_col31` varchar(190) DEFAULT NULL COMMENT '合同号',
  `xd_col32` varchar(20) DEFAULT NULL COMMENT '担保方式',
  `xd_col33` varchar(20) DEFAULT NULL COMMENT '担保状态',
  `xd_col34` varchar(20) DEFAULT NULL COMMENT '助学种类',
  `xd_col35` varchar(20) DEFAULT NULL COMMENT '按揭种类',
  `xd_col36` varchar(20) DEFAULT NULL COMMENT '按揭还款方式',
  `xd_col37` decimal(10,0) DEFAULT NULL COMMENT '按揭贷款期限',
  `xd_col38` decimal(10,0) DEFAULT NULL COMMENT '还款间隔',
  `xd_col39` varchar(20) DEFAULT NULL COMMENT '贴息种类',
  `xd_col40` varchar(20) DEFAULT NULL COMMENT '贴息资金来源',
  `xd_col41` varchar(50) DEFAULT NULL COMMENT '自动扣款账号',
  `xd_col42` varchar(50) DEFAULT NULL COMMENT '自动扣款账户户名',
  `xd_col43` varchar(20) DEFAULT NULL COMMENT '是否自动调形态',
  `xd_col44` varchar(20) DEFAULT NULL COMMENT '资金来源',
  `xd_col45` varchar(20) DEFAULT NULL COMMENT '专项资金来源',
  `xd_col46` varchar(3) DEFAULT NULL COMMENT '五级分类主体',
  `xd_col47` varchar(20) DEFAULT NULL COMMENT '五级分类类型',
  `xd_col48` timestamp NULL DEFAULT NULL COMMENT '登记日期',
  `xd_col49` varchar(50) DEFAULT NULL COMMENT '借据号',
  `xd_col50` varchar(40) DEFAULT NULL COMMENT '分区一',
  `xd_col51` varchar(40) DEFAULT NULL COMMENT '分区二',
  `xd_col52` varchar(40) DEFAULT NULL COMMENT '分区三',
  `xd_col53` varchar(20) DEFAULT NULL COMMENT '贷款期限种类',
  `xd_col54` timestamp NULL DEFAULT NULL COMMENT '起息日期',
  `xd_col55` varchar(16) DEFAULT NULL COMMENT '第一责任人',
  `xd_col56` varchar(16) DEFAULT NULL COMMENT '经办人',
  `xd_col57` varchar(20) DEFAULT NULL COMMENT '经营类型',
  `xd_col58` varchar(50) DEFAULT NULL COMMENT '营业执照号',
  `xd_col59` varchar(20) DEFAULT NULL COMMENT '企业特征',
  `xd_col60` varchar(20) DEFAULT NULL COMMENT '职业分类',
  `xd_col61` varchar(20) DEFAULT NULL COMMENT '科目号',
  `xd_col62` varchar(20) DEFAULT NULL COMMENT '贷款投向',
  `xd_col63` varchar(20) DEFAULT NULL COMMENT '客户类型',
  `xd_col64` varchar(20) DEFAULT NULL COMMENT '行业类型',
  `xd_col65` varchar(20) DEFAULT NULL COMMENT '发放类型',
  `xd_col66` varchar(20) DEFAULT NULL COMMENT '贷款方式',
  `xd_col67` varchar(3) DEFAULT NULL COMMENT '信贷产品名称',
  `xd_col68` varchar(20) DEFAULT NULL COMMENT '法人代表',
  `xd_col69` varchar(50) DEFAULT NULL COMMENT '贷款合同号',
  `xd_col70` varchar(50) DEFAULT NULL COMMENT '原贷款号',
  `xd_col71` timestamp NULL DEFAULT NULL COMMENT '原贷款日期',
  `xd_col72` varchar(20) DEFAULT NULL COMMENT '业务种类',
  `xd_col73` varchar(3) DEFAULT NULL COMMENT '新帐',
  `xd_col74` decimal(20,2) DEFAULT NULL COMMENT '应收未收利息',
  `xd_col75` varchar(1) DEFAULT NULL COMMENT '导入',
  `xd_col76` varchar(3) DEFAULT NULL COMMENT '分类1',
  `xd_col77` varchar(3) DEFAULT NULL COMMENT '分类2',
  `xd_col78` varchar(3) DEFAULT NULL COMMENT '分类3',
  `xd_col79` varchar(3) DEFAULT NULL COMMENT '分类4',
  `xd_col80` varchar(3) DEFAULT NULL COMMENT '分类5',
  `xd_col81` varchar(16) DEFAULT NULL COMMENT '客户经理',
  `xd_col82` varchar(16) DEFAULT NULL COMMENT '配合人',
  `xd_col83` varchar(1) DEFAULT NULL COMMENT '文件',
  `xd_col84` varchar(240) DEFAULT NULL COMMENT '备注',
  `xd_col85` varchar(20) DEFAULT NULL COMMENT '机构号',
  `xd_col86` varchar(16) DEFAULT NULL COMMENT '操作员',
  `xd_col87` decimal(10,0) DEFAULT NULL COMMENT 'AUTOID',
  `xd_col88` timestamp NULL DEFAULT NULL COMMENT 'LASTUPDATE',
  `xd_col89` varchar(50) DEFAULT NULL COMMENT '出票人名称',
  `xd_col90` varchar(50) DEFAULT NULL COMMENT '出票人账号',
  `xd_col91` varchar(50) DEFAULT NULL COMMENT '收款人名称',
  `xd_col92` varchar(50) DEFAULT NULL COMMENT '收款人账号',
  `xd_col93` varchar(50) DEFAULT NULL COMMENT '收款人开户行',
  `xd_col94` varchar(50) DEFAULT NULL COMMENT '收款人行号',
  `xd_col95` varchar(50) DEFAULT NULL COMMENT '付款行名称',
  `xd_col96` varchar(50) DEFAULT NULL COMMENT '付款行行号',
  `xd_col97` varchar(50) DEFAULT NULL COMMENT '付款行地址',
  `xd_col98` varchar(50) DEFAULT NULL COMMENT '承兑协议编号',
  `xd_col99` timestamp NULL DEFAULT NULL COMMENT '汇票出票日',
  `xd_col100` timestamp NULL DEFAULT NULL COMMENT '汇票到期日',
  `xd_col101` decimal(20,2) DEFAULT NULL COMMENT '出票金额',
  `xd_col102` varchar(50) DEFAULT NULL COMMENT '查询回复号',
  `xd_col103` timestamp NULL DEFAULT NULL COMMENT '查询回复日',
  `xd_col104` varchar(200) DEFAULT NULL COMMENT '汇票说明',
  `xd_col105` decimal(20,2) DEFAULT NULL COMMENT 'YSLX',
  `xd_col106` decimal(20,2) DEFAULT NULL COMMENT '手续费',
  `xd_col107` decimal(14,7) DEFAULT NULL COMMENT '管理费率',
  `xd_col108` varchar(20) DEFAULT NULL COMMENT '利息算法1',
  `xd_col109` varchar(20) DEFAULT NULL COMMENT '利息算法2',
  `xd_col110` varchar(20) DEFAULT NULL COMMENT '利息算法3',
  `xd_col111` varchar(20) DEFAULT NULL COMMENT '扣划方式',
  `xd_col112` decimal(10,0) DEFAULT NULL COMMENT '结算日',
  `xd_col113` timestamp NULL DEFAULT NULL COMMENT '交易时间',
  `xd_col114` varchar(30) DEFAULT NULL COMMENT '客户编码',
  `xd_col115` varchar(30) DEFAULT NULL COMMENT '卡号',
  `xd_col116` timestamp NULL DEFAULT NULL COMMENT '原到期日',
  `xd_col117` varchar(30) DEFAULT NULL COMMENT '凭证号',
  `xd_col118` decimal(10,0) DEFAULT NULL COMMENT '打印次数',
  `xd_col119` varchar(50) DEFAULT NULL COMMENT '自动扣款银行',
  `xd_col120` varchar(30) DEFAULT NULL COMMENT '账号序号',
  `xd_col121` decimal(10,0) DEFAULT NULL COMMENT '期数',
  `xd_col122` decimal(20,2) DEFAULT NULL COMMENT '首付款',
  `xd_col123` decimal(20,2) DEFAULT NULL COMMENT '预留尾款',
  `xd_col124` decimal(14,7) DEFAULT NULL COMMENT '违约金比例',
  `xd_col125` decimal(20,2) DEFAULT NULL COMMENT '最低违约金',
  `xd_col126` decimal(14,7) DEFAULT NULL COMMENT '贴息利率',
  `xd_col127` decimal(20,2) DEFAULT NULL COMMENT '贴息金额',
  `xd_col128` decimal(14,7) DEFAULT NULL COMMENT '基准利率浮动',
  `xd_col129` decimal(14,7) DEFAULT NULL COMMENT '逾期利率浮动',
  `xd_col130` varchar(20) DEFAULT NULL COMMENT '经销商编号',
  `xd_col131` varchar(40) DEFAULT NULL COMMENT '监管账号',
  `xd_col132` varchar(20) DEFAULT NULL COMMENT '利率执行方式',
  `xd_col133` decimal(20,2) DEFAULT NULL COMMENT '结欠应收未收',
  `xd_col134` decimal(20,2) DEFAULT NULL COMMENT '考察费',
  `xd_col135` varchar(3) DEFAULT NULL COMMENT '验证结果',
  `xd_col136` varchar(20) DEFAULT NULL COMMENT '涉农',
  `xd_col137` varchar(20) DEFAULT NULL COMMENT '涉政',
  `xd_col138` varchar(20) DEFAULT NULL COMMENT '社团',
  `xd_col139` varchar(8) DEFAULT NULL COMMENT '信贷产品',
  `xd_col140` timestamp NULL DEFAULT NULL COMMENT 'FIRSTDATE',
  `xd_col141` varchar(6) DEFAULT NULL COMMENT '成因',
  `xd_col142` varchar(20) DEFAULT NULL COMMENT '支付方式',
  `xd_col143` varchar(20) DEFAULT NULL COMMENT '客户分类',
  `xd_col144` varchar(120) DEFAULT NULL COMMENT '居住地址',
  `xd_col145` varchar(16) DEFAULT NULL COMMENT '复核人',
  `xd_col146` varchar(50) DEFAULT NULL COMMENT 'KEYID',
  `xd_col147` varchar(60) DEFAULT NULL COMMENT '手机号码',
  `xd_col148` varchar(60) DEFAULT NULL COMMENT '家庭电话',
  `xd_col149` varchar(60) DEFAULT NULL COMMENT '办公电话',
  `xd_col150` varchar(60) DEFAULT NULL COMMENT '手机号码2',
  `xd_col151` varchar(3) DEFAULT NULL COMMENT '融资平台',
  `xd_col152` varchar(20) DEFAULT NULL COMMENT '代理机构号',
  `xd_col153` varchar(2) DEFAULT NULL COMMENT '还款计划',
  `xd_col154` varchar(20) DEFAULT NULL COMMENT '项目号',
  `xd_col155` varchar(20) DEFAULT NULL COMMENT '币种',
  `xd_col156` decimal(20,2) DEFAULT NULL COMMENT '减值准备',
  `xd_col157` varchar(30) DEFAULT NULL COMMENT '自动扣款卡号',
  `xd_col158` varchar(20) DEFAULT NULL COMMENT '联社机构号',
  `xd_col159` varchar(20) DEFAULT NULL COMMENT '放款账号科目号',
  `xd_col160` varchar(20) DEFAULT NULL COMMENT '助学贷款类型',
  `xd_col161` varchar(20) DEFAULT NULL COMMENT '允许优惠标志',
  `xd_col162` varchar(20) DEFAULT NULL COMMENT '利率调整方式',
  `xd_col163` varchar(20) DEFAULT NULL COMMENT '利率定价方式',
  `xd_col164` varchar(50) DEFAULT NULL COMMENT '核销帐号',
  `xd_col165` decimal(10,0) DEFAULT NULL COMMENT '减值变化',
  `xd_col166` varchar(8) DEFAULT NULL COMMENT 'HX科目号',
  `xd_col167` varchar(20) NOT NULL DEFAULT '-' COMMENT 'BANKCD',
  `etl_f` varchar(1) DEFAULT NULL COMMENT '(null)',
  `src_sys_cd` varchar(10) DEFAULT NULL COMMENT '源系统代码',
  `crut_no` varchar(10) DEFAULT NULL COMMENT '法人机构号',
  `busi_dt` varchar(8) DEFAULT NULL COMMENT '业务日期',
  `biz_dt` varchar(8) DEFAULT NULL COMMENT '平台日期',
  `etl_load_time` timestamp NULL DEFAULT NULL COMMENT '数据加载时间戳',
  PRIMARY KEY (`xd_col167`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `s_loan_dk_2350000` */

/*Table structure for table `s_ofcr_ch_acct_mast_2350000` */

DROP TABLE IF EXISTS `s_ofcr_ch_acct_mast_2350000`;

CREATE TABLE `s_ofcr_ch_acct_mast_2350000` (
  `cod_cc_brn` varchar(10) DEFAULT NULL COMMENT '分行号',
  `cod_prod` varchar(20) DEFAULT NULL COMMENT '产品号',
  `cod_acct_no` varchar(16) NOT NULL DEFAULT '-' COMMENT '帐号',
  `cod_officr_id` varchar(36) DEFAULT NULL COMMENT '客户经理ID',
  `dat_acct_open` varchar(8) DEFAULT NULL COMMENT '开户日',
  `cod_acct_title` varchar(360) DEFAULT NULL COMMENT '户名',
  `cod_cust` varchar(10) DEFAULT NULL COMMENT '客户代码',
  `nam_cust_shrt` varchar(60) DEFAULT NULL COMMENT '客户简称',
  `nam_nominee` varchar(120) DEFAULT NULL COMMENT '受益人名字',
  `cod_ccy` varchar(20) DEFAULT NULL COMMENT '币种代码',
  `cod_acct_stat` varchar(20) DEFAULT NULL COMMENT '帐户状态',
  `cod_tds` varchar(10) DEFAULT NULL COMMENT '税代码',
  `bal_acct_min_reqd` decimal(20,2) DEFAULT NULL COMMENT '最小余额要求',
  `flg_joint_acct` varchar(20) DEFAULT NULL COMMENT '是否为联名帐户',
  `flg_stp_pmnt` varchar(20) DEFAULT NULL COMMENT '止付标志',
  `flg_si` varchar(20) DEFAULT NULL COMMENT '预约转账标志',
  `flg_swpout` varchar(1) DEFAULT NULL COMMENT '转出指令标志',
  `flg_idd_auth` varchar(1) DEFAULT NULL COMMENT '直接借记授权标志',
  `flg_tds` varchar(1) DEFAULT NULL COMMENT 'TDS标志',
  `flg_vat` varchar(1) DEFAULT NULL COMMENT '增值税标志',
  `amt_exempt_limit` decimal(20,2) DEFAULT NULL COMMENT '免税额',
  `flg_grp_bonus` varchar(1) DEFAULT NULL COMMENT '组红利标志',
  `flg_atm` varchar(10) DEFAULT NULL COMMENT 'ＡＴＭ标志',
  `ctr_int_sc_xfer` varchar(10) DEFAULT NULL COMMENT '利息手续费转账柜台',
  `rat_var_cr_int` decimal(20,0) DEFAULT NULL COMMENT '贷记利率浮动',
  `flg_dr_auth` varchar(20) DEFAULT NULL COMMENT '借记授权标志',
  `flg_psbook` varchar(20) DEFAULT NULL COMMENT '存折标志',
  `amt_od_limit` decimal(20,2) DEFAULT NULL COMMENT '透支限额',
  `bal_book` decimal(20,2) DEFAULT NULL COMMENT '账面余额',
  `amt_hld` decimal(20,2) DEFAULT NULL COMMENT '冻结金额',
  `bal_available` decimal(20,2) DEFAULT NULL COMMENT '可用余额',
  `amt_dr_today` decimal(20,2) DEFAULT NULL COMMENT '今日记入借方金额',
  `amt_cr_today` decimal(20,2) DEFAULT NULL COMMENT '今日记入贷方金额',
  `amt_dr_mtd` decimal(20,2) DEFAULT NULL COMMENT '本月至今借记金额',
  `amt_cr_mtd` decimal(20,2) DEFAULT NULL COMMENT '本月至今贷记金额',
  `amt_dr_ytd` decimal(20,2) DEFAULT NULL COMMENT '本年至今借记金额',
  `amt_cr_ytd` decimal(20,2) DEFAULT NULL COMMENT '本年至今贷记金额',
  `amt_last_dr` decimal(20,2) DEFAULT NULL COMMENT '上次借记金额',
  `dat_last_dr` varchar(8) DEFAULT NULL COMMENT '上次借记日期',
  `amt_last_cr` decimal(20,2) DEFAULT NULL COMMENT '上次贷记金额',
  `dat_last_cr` varchar(8) DEFAULT NULL COMMENT '上次贷记日期',
  `dat_last_stmnt` varchar(8) DEFAULT NULL COMMENT '上次对账单日期',
  `bal_last_stmnt` decimal(20,2) DEFAULT NULL COMMENT '上次余额结单',
  `dat_last_int_cap` varchar(8) DEFAULT NULL COMMENT '上次结息日期',
  `dat_last_sc_cap` varchar(8) DEFAULT NULL COMMENT '上次手续费结息日期',
  `amt_cr_int_accr_last_cap` decimal(20,2) DEFAULT NULL COMMENT '上一结息日至今贷记利息计提',
  `amt_dr_int_accr_last_cap` decimal(20,2) DEFAULT NULL COMMENT '上一结息日至今借记利息计提',
  `amt_adb_product` decimal(20,2) DEFAULT NULL COMMENT '日均余额',
  `dat_last_int_comp` varchar(8) DEFAULT NULL COMMENT '上次计息日',
  `bal_last_int_cap` decimal(20,2) DEFAULT NULL COMMENT '上次结息余额',
  `amt_tot_dr` decimal(20,2) DEFAULT NULL COMMENT '借记额总计',
  `amt_tot_cr` decimal(20,2) DEFAULT NULL COMMENT '贷记额总计',
  `flg_int_waiver` varchar(20) DEFAULT NULL COMMENT '利息豁免标志',
  `flg_sc_waiver` varchar(20) DEFAULT NULL COMMENT '手续费豁免标志',
  `cod_int_sc_acct_no` varchar(16) DEFAULT NULL COMMENT '利息手续费账号',
  `amt_ytd_int_paid` decimal(20,2) DEFAULT NULL COMMENT '本年至今付利息',
  `amt_ytd_int_recd` decimal(20,2) DEFAULT NULL COMMENT '本年至今所收利息',
  `amt_ytd_sc` decimal(20,2) DEFAULT NULL COMMENT '本年至今手续费额',
  `amt_dr_int_accr` decimal(30,10) DEFAULT NULL COMMENT '借记利息计提额',
  `amt_cr_int_accr` decimal(30,10) DEFAULT NULL COMMENT '贷记利息计提额',
  `amt_dr_int_comp` decimal(30,10) DEFAULT NULL COMMENT '借记利息计息额',
  `amt_cr_int_comp` decimal(30,10) DEFAULT NULL COMMENT '贷记利息计息额',
  `amt_last_cr_int_accr` decimal(30,10) DEFAULT NULL COMMENT '上次贷记利息计提额',
  `amt_last_dr_int_accr` decimal(30,10) DEFAULT NULL COMMENT '上次借记利息计提额',
  `ctr_psbk_line_no` decimal(5,0) DEFAULT NULL COMMENT '存折行号计数器',
  `bal_lst_psbk` decimal(20,2) DEFAULT NULL COMMENT '上次存折余额',
  `flg_acct_close` varchar(20) DEFAULT NULL COMMENT '销户标志',
  `amt_sweepin_lien` decimal(20,2) DEFAULT NULL COMMENT '转入留置额',
  `flg_sweepin` varchar(20) DEFAULT NULL COMMENT '转入标志',
  `bal_previous` decimal(20,2) DEFAULT NULL COMMENT '前余额',
  `flg_prev_bal_stat` varchar(1) DEFAULT NULL COMMENT '前余额状态',
  `flg_mnt_status` varchar(20) DEFAULT NULL COMMENT '维护状态标志',
  `amt_cum_od_limit` decimal(20,2) DEFAULT NULL COMMENT '累计透支限额额度',
  `amt_od_limit_prev` decimal(20,2) DEFAULT NULL COMMENT '前透支限额额度',
  `amt_instal` decimal(20,2) DEFAULT NULL COMMENT '分期付款额',
  `dat_maturity` varchar(8) DEFAULT NULL COMMENT '存款到期日',
  `acct_term_days` decimal(10,0) DEFAULT NULL COMMENT '账户存期天数',
  `frq_instal` varchar(1) DEFAULT NULL COMMENT '分期频率',
  `rat_int_rd` decimal(14,7) DEFAULT NULL COMMENT '零存整取利率',
  `amt_rd_advance` decimal(20,2) DEFAULT NULL COMMENT '零存整取垫款额',
  `amt_maturity_value` decimal(20,2) DEFAULT NULL COMMENT '存款到期日金额',
  `rat_penalty` decimal(14,7) DEFAULT NULL COMMENT '罚息利率',
  `dat_last_penalty` varchar(8) DEFAULT NULL COMMENT '上次罚息日',
  `grace_days` decimal(10,0) DEFAULT NULL COMMENT '宽限天数',
  `flg_grace_period` varchar(10) DEFAULT NULL COMMENT '宽限期限类型',
  `frq_cr_int_cap` varchar(1) DEFAULT NULL COMMENT '贷记结息频率',
  `frq_dr_int_cap` varchar(1) DEFAULT NULL COMMENT '借记结息频率',
  `cr_int_cap_day` varchar(8) DEFAULT NULL COMMENT '贷记结息日',
  `dr_int_cap_day` varchar(8) DEFAULT NULL COMMENT '借记结息日',
  `dat_last_cr_int_cap` varchar(8) DEFAULT NULL COMMENT '上一次贷记结息日',
  `dat_last_dr_int_cap` varchar(8) DEFAULT NULL COMMENT '上一次借记结息日',
  `dat_next_cr_int_cap` varchar(8) DEFAULT NULL COMMENT '下一次贷记结息日',
  `dat_next_dr_int_cap` varchar(8) DEFAULT NULL COMMENT '下一次借记结息日',
  `amt_cr_int_accr_adj` decimal(20,2) DEFAULT NULL COMMENT '利息调整贷记计提',
  `amt_dr_int_accr_adj` decimal(20,2) DEFAULT NULL COMMENT '利息调整借记计提',
  `cod_int_rate_tier_plan` varchar(20) DEFAULT NULL COMMENT '利息利率层计划代码',
  `cod_acct_no_rd_drwdwn` varchar(16) DEFAULT NULL COMMENT '零存整取自扣账号',
  `flg_allow_drwdwn` varchar(20) DEFAULT NULL COMMENT '自扣允许标志',
  `cod_sc_pkg` varchar(5) DEFAULT NULL COMMENT '服务费包编号',
  `cod_iban_no` varchar(40) DEFAULT NULL COMMENT '国际银行账号',
  `cod_sc_xfer_acct_no` varchar(16) DEFAULT NULL COMMENT '手续费扣款账户',
  `cod_cr_int_xfer_acct_no` varchar(16) DEFAULT NULL COMMENT '贷记利息入账账户',
  `cod_dr_int_xfer_acct_no` varchar(16) DEFAULT NULL COMMENT '借记利息入账账户',
  `cod_entity_vpd` decimal(10,0) DEFAULT NULL COMMENT '实体代码',
  `amt_extra_payin` decimal(20,0) DEFAULT NULL COMMENT 'Extra Amount paid towards RD account',
  `dat_extra_payin` varchar(8) DEFAULT NULL COMMENT 'Date on which extra amount is paid',
  `amt_cr_int_comp_extrapay` decimal(20,0) DEFAULT NULL COMMENT 'Interest computed in extra amount',
  `amt_cr_int_accr_extrapay` decimal(20,0) DEFAULT NULL COMMENT 'Interest accrued on extra amount',
  `amt_last_cr_int_accr_extrapay` decimal(20,0) DEFAULT NULL COMMENT 'Accrued interest on extra amount as of previous accrual',
  `dat_int_comp_extrapay` varchar(8) DEFAULT NULL COMMENT 'Date on which interest on extra amount was computed',
  `etl_f` varchar(1) DEFAULT NULL COMMENT '(null)',
  `src_sys_cd` varchar(10) DEFAULT NULL COMMENT '源系统代码',
  `crut_no` varchar(10) DEFAULT NULL COMMENT '法人机构号',
  `busi_dt` varchar(8) DEFAULT NULL COMMENT '业务日期',
  `biz_dt` varchar(8) DEFAULT NULL COMMENT '平台日期',
  `etl_load_time` timestamp NULL DEFAULT NULL COMMENT '数据加载时间戳',
  PRIMARY KEY (`cod_acct_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `s_ofcr_ch_acct_mast_2350000` */

/*Table structure for table `table_column_config` */

DROP TABLE IF EXISTS `table_column_config`;

CREATE TABLE `table_column_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '表列自增id',
  `table_name` char(150) DEFAULT NULL COMMENT '表名',
  `column` char(80) DEFAULT NULL COMMENT '列名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=305 DEFAULT CHARSET=utf8;

/*Data for the table `table_column_config` */

insert  into `table_column_config`(`id`,`table_name`,`column`) values (1,'s_loan_dk','xd_col1'),(2,'s_loan_dk','xd_col2'),(3,'s_loan_dk','xd_col3'),(4,'s_loan_dk','xd_col4'),(5,'s_loan_dk','xd_col5'),(6,'s_loan_dk','xd_col6'),(7,'s_loan_dk','xd_col7'),(8,'s_loan_dk','xd_col8'),(9,'s_loan_dk','xd_col9'),(10,'s_loan_dk','xd_col10'),(11,'s_loan_dk','xd_col11'),(12,'s_loan_dk','xd_col12'),(13,'s_loan_dk','xd_col13'),(14,'s_loan_dk','xd_col14'),(15,'s_loan_dk','xd_col15'),(16,'s_loan_dk','xd_col16'),(17,'s_loan_dk','xd_col17'),(18,'s_loan_dk','xd_col18'),(19,'s_loan_dk','xd_col19'),(20,'s_loan_dk','xd_col20'),(21,'s_loan_dk','xd_col21'),(22,'s_loan_dk','xd_col22'),(23,'s_loan_dk','xd_col23'),(24,'s_loan_dk','xd_col24'),(25,'s_loan_dk','xd_col25'),(26,'s_loan_dk','xd_col26'),(27,'s_loan_dk','xd_col27'),(28,'s_loan_dk','xd_col28'),(29,'s_loan_dk','xd_col29'),(30,'s_loan_dk','xd_col30'),(31,'s_loan_dk','xd_col31'),(32,'s_loan_dk','xd_col32'),(33,'s_loan_dk','xd_col33'),(34,'s_loan_dk','xd_col34'),(35,'s_loan_dk','xd_col35'),(36,'s_loan_dk','xd_col36'),(37,'s_loan_dk','xd_col37'),(38,'s_loan_dk','xd_col38'),(39,'s_loan_dk','xd_col39'),(40,'s_loan_dk','xd_col40'),(41,'s_loan_dk','xd_col41'),(42,'s_loan_dk','xd_col42'),(43,'s_loan_dk','xd_col43'),(44,'s_loan_dk','xd_col44'),(45,'s_loan_dk','xd_col45'),(46,'s_loan_dk','xd_col46'),(47,'s_loan_dk','xd_col47'),(48,'s_loan_dk','xd_col48'),(49,'s_loan_dk','xd_col49'),(50,'s_loan_dk','xd_col50'),(51,'s_loan_dk','xd_col51'),(52,'s_loan_dk','xd_col52'),(53,'s_loan_dk','xd_col53'),(54,'s_loan_dk','xd_col54'),(55,'s_loan_dk','xd_col55'),(56,'s_loan_dk','xd_col56'),(57,'s_loan_dk','xd_col57'),(58,'s_loan_dk','xd_col58'),(59,'s_loan_dk','xd_col59'),(60,'s_loan_dk','xd_col60'),(61,'s_loan_dk','xd_col61'),(62,'s_loan_dk','xd_col62'),(63,'s_loan_dk','xd_col63'),(64,'s_loan_dk','xd_col64'),(65,'s_loan_dk','xd_col65'),(66,'s_loan_dk','xd_col66'),(67,'s_loan_dk','xd_col67'),(68,'s_loan_dk','xd_col68'),(69,'s_loan_dk','xd_col69'),(70,'s_loan_dk','xd_col70'),(71,'s_loan_dk','xd_col71'),(72,'s_loan_dk','xd_col72'),(73,'s_loan_dk','xd_col73'),(74,'s_loan_dk','xd_col74'),(75,'s_loan_dk','xd_col75'),(76,'s_loan_dk','xd_col76'),(77,'s_loan_dk','xd_col77'),(78,'s_loan_dk','xd_col78'),(79,'s_loan_dk','xd_col79'),(80,'s_loan_dk','xd_col80'),(81,'s_loan_dk','xd_col81'),(82,'s_loan_dk','xd_col82'),(83,'s_loan_dk','xd_col83'),(84,'s_loan_dk','xd_col84'),(85,'s_loan_dk','xd_col85'),(86,'s_loan_dk','xd_col86'),(87,'s_loan_dk','xd_col87'),(88,'s_loan_dk','xd_col88'),(89,'s_loan_dk','xd_col89'),(90,'s_loan_dk','xd_col90'),(91,'s_loan_dk','xd_col91'),(92,'s_loan_dk','xd_col92'),(93,'s_loan_dk','xd_col93'),(94,'s_loan_dk','xd_col94'),(95,'s_loan_dk','xd_col95'),(96,'s_loan_dk','xd_col96'),(97,'s_loan_dk','xd_col97'),(98,'s_loan_dk','xd_col98'),(99,'s_loan_dk','xd_col99'),(100,'s_loan_dk','xd_col100'),(101,'s_loan_dk','xd_col101'),(102,'s_loan_dk','xd_col102'),(103,'s_loan_dk','xd_col103'),(104,'s_loan_dk','xd_col104'),(105,'s_loan_dk','xd_col105'),(106,'s_loan_dk','xd_col106'),(107,'s_loan_dk','xd_col107'),(108,'s_loan_dk','xd_col108'),(109,'s_loan_dk','xd_col109'),(110,'s_loan_dk','xd_col110'),(111,'s_loan_dk','xd_col111'),(112,'s_loan_dk','xd_col112'),(113,'s_loan_dk','xd_col113'),(114,'s_loan_dk','xd_col114'),(115,'s_loan_dk','xd_col115'),(116,'s_loan_dk','xd_col116'),(117,'s_loan_dk','xd_col117'),(118,'s_loan_dk','xd_col118'),(119,'s_loan_dk','xd_col119'),(120,'s_loan_dk','xd_col120'),(121,'s_loan_dk','xd_col121'),(122,'s_loan_dk','xd_col122'),(123,'s_loan_dk','xd_col123'),(124,'s_loan_dk','xd_col124'),(125,'s_loan_dk','xd_col125'),(126,'s_loan_dk','xd_col126'),(127,'s_loan_dk','xd_col127'),(128,'s_loan_dk','xd_col128'),(129,'s_loan_dk','xd_col129'),(130,'s_loan_dk','xd_col130'),(131,'s_loan_dk','xd_col131'),(132,'s_loan_dk','xd_col132'),(133,'s_loan_dk','xd_col133'),(134,'s_loan_dk','xd_col134'),(135,'s_loan_dk','xd_col135'),(136,'s_loan_dk','xd_col136'),(137,'s_loan_dk','xd_col137'),(138,'s_loan_dk','xd_col138'),(139,'s_loan_dk','xd_col139'),(140,'s_loan_dk','xd_col140'),(141,'s_loan_dk','xd_col141'),(142,'s_loan_dk','xd_col142'),(143,'s_loan_dk','xd_col143'),(144,'s_loan_dk','xd_col144'),(145,'s_loan_dk','xd_col145'),(146,'s_loan_dk','xd_col146'),(147,'s_loan_dk','xd_col147'),(148,'s_loan_dk','xd_col148'),(149,'s_loan_dk','xd_col149'),(150,'s_loan_dk','xd_col150'),(151,'s_loan_dk','xd_col151'),(152,'s_loan_dk','xd_col152'),(153,'s_loan_dk','xd_col153'),(154,'s_loan_dk','xd_col154'),(155,'s_loan_dk','xd_col155'),(156,'s_loan_dk','xd_col156'),(157,'s_loan_dk','xd_col157'),(158,'s_loan_dk','xd_col158'),(159,'s_loan_dk','xd_col159'),(160,'s_loan_dk','xd_col160'),(161,'s_loan_dk','xd_col161'),(162,'s_loan_dk','xd_col162'),(163,'s_loan_dk','xd_col163'),(164,'s_loan_dk','xd_col164'),(165,'s_loan_dk','xd_col165'),(166,'s_loan_dk','xd_col166'),(167,'s_loan_dk','xd_col167'),(168,'s_loan_dk','etl_f'),(169,'s_loan_dk','src_sys_cd'),(170,'s_loan_dk','crut_no'),(171,'s_loan_dk','busi_dt'),(172,'s_loan_dk','biz_dt'),(173,'s_loan_dk','etl_load_time'),(174,'s_ofcr_ch_acct_mast','cod_cc_brn'),(175,'s_ofcr_ch_acct_mast','cod_prod'),(176,'s_ofcr_ch_acct_mast','cod_acct_no'),(177,'s_ofcr_ch_acct_mast','cod_officr_id'),(178,'s_ofcr_ch_acct_mast','dat_acct_open'),(179,'s_ofcr_ch_acct_mast','cod_acct_title'),(180,'s_ofcr_ch_acct_mast','cod_cust'),(181,'s_ofcr_ch_acct_mast','nam_cust_shrt'),(182,'s_ofcr_ch_acct_mast','nam_nominee'),(183,'s_ofcr_ch_acct_mast','cod_ccy'),(184,'s_ofcr_ch_acct_mast','cod_acct_stat'),(185,'s_ofcr_ch_acct_mast','cod_tds'),(186,'s_ofcr_ch_acct_mast','bal_acct_min_reqd'),(187,'s_ofcr_ch_acct_mast','flg_joint_acct'),(188,'s_ofcr_ch_acct_mast','flg_stp_pmnt'),(189,'s_ofcr_ch_acct_mast','flg_si'),(190,'s_ofcr_ch_acct_mast','flg_swpout'),(191,'s_ofcr_ch_acct_mast','flg_idd_auth'),(192,'s_ofcr_ch_acct_mast','flg_tds'),(193,'s_ofcr_ch_acct_mast','flg_vat'),(194,'s_ofcr_ch_acct_mast','amt_exempt_limit'),(195,'s_ofcr_ch_acct_mast','flg_grp_bonus'),(196,'s_ofcr_ch_acct_mast','flg_atm'),(197,'s_ofcr_ch_acct_mast','ctr_int_sc_xfer'),(198,'s_ofcr_ch_acct_mast','rat_var_cr_int'),(199,'s_ofcr_ch_acct_mast','flg_dr_auth'),(200,'s_ofcr_ch_acct_mast','flg_psbook'),(201,'s_ofcr_ch_acct_mast','amt_od_limit'),(202,'s_ofcr_ch_acct_mast','bal_book'),(203,'s_ofcr_ch_acct_mast','amt_hld'),(204,'s_ofcr_ch_acct_mast','bal_available'),(205,'s_ofcr_ch_acct_mast','amt_dr_today'),(206,'s_ofcr_ch_acct_mast','amt_cr_today'),(207,'s_ofcr_ch_acct_mast','amt_dr_mtd'),(208,'s_ofcr_ch_acct_mast','amt_cr_mtd'),(209,'s_ofcr_ch_acct_mast','amt_dr_ytd'),(210,'s_ofcr_ch_acct_mast','amt_cr_ytd'),(211,'s_ofcr_ch_acct_mast','amt_last_dr'),(212,'s_ofcr_ch_acct_mast','dat_last_dr'),(213,'s_ofcr_ch_acct_mast','amt_last_cr'),(214,'s_ofcr_ch_acct_mast','dat_last_cr'),(215,'s_ofcr_ch_acct_mast','dat_last_stmnt'),(216,'s_ofcr_ch_acct_mast','bal_last_stmnt'),(217,'s_ofcr_ch_acct_mast','dat_last_int_cap'),(218,'s_ofcr_ch_acct_mast','dat_last_sc_cap'),(219,'s_ofcr_ch_acct_mast','amt_cr_int_accr_last_cap'),(220,'s_ofcr_ch_acct_mast','amt_dr_int_accr_last_cap'),(221,'s_ofcr_ch_acct_mast','amt_adb_product'),(222,'s_ofcr_ch_acct_mast','dat_last_int_comp'),(223,'s_ofcr_ch_acct_mast','bal_last_int_cap'),(224,'s_ofcr_ch_acct_mast','amt_tot_dr'),(225,'s_ofcr_ch_acct_mast','amt_tot_cr'),(226,'s_ofcr_ch_acct_mast','flg_int_waiver'),(227,'s_ofcr_ch_acct_mast','flg_sc_waiver'),(228,'s_ofcr_ch_acct_mast','cod_int_sc_acct_no'),(229,'s_ofcr_ch_acct_mast','amt_ytd_int_paid'),(230,'s_ofcr_ch_acct_mast','amt_ytd_int_recd'),(231,'s_ofcr_ch_acct_mast','amt_ytd_sc'),(232,'s_ofcr_ch_acct_mast','amt_dr_int_accr'),(233,'s_ofcr_ch_acct_mast','amt_cr_int_accr'),(234,'s_ofcr_ch_acct_mast','amt_dr_int_comp'),(235,'s_ofcr_ch_acct_mast','amt_cr_int_comp'),(236,'s_ofcr_ch_acct_mast','amt_last_cr_int_accr'),(237,'s_ofcr_ch_acct_mast','amt_last_dr_int_accr'),(238,'s_ofcr_ch_acct_mast','ctr_psbk_line_no'),(239,'s_ofcr_ch_acct_mast','bal_lst_psbk'),(240,'s_ofcr_ch_acct_mast','flg_acct_close'),(241,'s_ofcr_ch_acct_mast','amt_sweepin_lien'),(242,'s_ofcr_ch_acct_mast','flg_sweepin'),(243,'s_ofcr_ch_acct_mast','bal_previous'),(244,'s_ofcr_ch_acct_mast','flg_prev_bal_stat'),(245,'s_ofcr_ch_acct_mast','flg_mnt_status'),(246,'s_ofcr_ch_acct_mast','amt_cum_od_limit'),(247,'s_ofcr_ch_acct_mast','amt_od_limit_prev'),(248,'s_ofcr_ch_acct_mast','amt_instal'),(249,'s_ofcr_ch_acct_mast','dat_maturity'),(250,'s_ofcr_ch_acct_mast','acct_term_days'),(251,'s_ofcr_ch_acct_mast','frq_instal'),(252,'s_ofcr_ch_acct_mast','rat_int_rd'),(253,'s_ofcr_ch_acct_mast','amt_rd_advance'),(254,'s_ofcr_ch_acct_mast','amt_maturity_value'),(255,'s_ofcr_ch_acct_mast','rat_penalty'),(256,'s_ofcr_ch_acct_mast','dat_last_penalty'),(257,'s_ofcr_ch_acct_mast','grace_days'),(258,'s_ofcr_ch_acct_mast','flg_grace_period'),(259,'s_ofcr_ch_acct_mast','frq_cr_int_cap'),(260,'s_ofcr_ch_acct_mast','frq_dr_int_cap'),(261,'s_ofcr_ch_acct_mast','cr_int_cap_day'),(262,'s_ofcr_ch_acct_mast','dr_int_cap_day'),(263,'s_ofcr_ch_acct_mast','dat_last_cr_int_cap'),(264,'s_ofcr_ch_acct_mast','dat_last_dr_int_cap'),(265,'s_ofcr_ch_acct_mast','dat_next_cr_int_cap'),(266,'s_ofcr_ch_acct_mast','dat_next_dr_int_cap'),(267,'s_ofcr_ch_acct_mast','amt_cr_int_accr_adj'),(268,'s_ofcr_ch_acct_mast','amt_dr_int_accr_adj'),(269,'s_ofcr_ch_acct_mast','cod_int_rate_tier_plan'),(270,'s_ofcr_ch_acct_mast','cod_acct_no_rd_drwdwn'),(271,'s_ofcr_ch_acct_mast','flg_allow_drwdwn'),(272,'s_ofcr_ch_acct_mast','cod_sc_pkg'),(273,'s_ofcr_ch_acct_mast','cod_iban_no'),(274,'s_ofcr_ch_acct_mast','cod_sc_xfer_acct_no'),(275,'s_ofcr_ch_acct_mast','cod_cr_int_xfer_acct_no'),(276,'s_ofcr_ch_acct_mast','cod_dr_int_xfer_acct_no'),(277,'s_ofcr_ch_acct_mast','cod_entity_vpd'),(278,'s_ofcr_ch_acct_mast','amt_extra_payin'),(279,'s_ofcr_ch_acct_mast','dat_extra_payin'),(280,'s_ofcr_ch_acct_mast','amt_cr_int_comp_extrapay'),(281,'s_ofcr_ch_acct_mast','amt_cr_int_accr_extrapay'),(282,'s_ofcr_ch_acct_mast','amt_last_cr_int_accr_extrapay'),(283,'s_ofcr_ch_acct_mast','dat_int_comp_extrapay'),(284,'s_ofcr_ch_acct_mast','etl_f'),(285,'s_ofcr_ch_acct_mast','src_sys_cd'),(286,'s_ofcr_ch_acct_mast','crut_no'),(287,'s_ofcr_ch_acct_mast','busi_dt'),(288,'s_ofcr_ch_acct_mast','biz_dt'),(289,'s_ofcr_ch_acct_mast','etl_load_time'),(290,'s_ecif_ecif_cert_info','bankid'),(291,'s_ecif_ecif_cert_info','cod_cust_id'),(292,'s_ecif_ecif_cert_info','cert_type'),(293,'s_ecif_ecif_cert_info','cert_num'),(294,'s_ecif_ecif_cert_info','sign_date'),(295,'s_ecif_ecif_cert_info','effe_date'),(296,'s_ecif_ecif_cert_info','is_fist'),(297,'s_ecif_ecif_cert_info','effe_type'),(298,'s_ecif_ecif_cert_info','cert_dec'),(299,'s_ecif_ecif_cert_info','etl_f'),(300,'s_ecif_ecif_cert_info','src_sys_cd'),(301,'s_ecif_ecif_cert_info','crut_no'),(302,'s_ecif_ecif_cert_info','busi_dt'),(303,'s_ecif_ecif_cert_info','biz_dt'),(304,'s_ecif_ecif_cert_info','etl_load_time');

/*Table structure for table `table_config` */

DROP TABLE IF EXISTS `table_config`;

CREATE TABLE `table_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `table_name` char(150) DEFAULT NULL COMMENT '表名',
  `title` char(150) DEFAULT NULL COMMENT '中文名',
  `type` char(5) DEFAULT NULL COMMENT '类型',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=301 DEFAULT CHARSET=utf8;

/*Data for the table `table_config` */

insert  into `table_config`(`id`,`table_name`,`title`,`type`) values (1,'s_atmp_devinfo','自助设备信息表','all'),(2,'s_atmp_rvs','设备交易冲正表','add'),(3,'s_atmp_t_device','自助设备信息表','all'),(4,'s_atmp_t_devicestate','设备状态表','all'),(5,'s_atmp_trnj','交易流水表','add'),(6,'s_cets_auth_reg','授权登记簿表','add'),(7,'s_cets_auth_regdetail','授权明细表','add'),(8,'s_cups_bank_txn_detail','对账明细表','add'),(9,'s_cups_cups_51c','报文51C','add'),(10,'s_cups_cups_acom','交易流水ACOM','add'),(11,'s_cups_cups_icom','交易流水ICOM','add'),(12,'s_cups_cups_mer_incom','中国银联直联收单机构商户费用收支汇总表','add'),(13,'s_cups_cups_tfl','转账交易流水','add'),(14,'s_cups_inst','CUPS机构信息表','all'),(15,'s_cups_mer_yl','银联商户信息表','all'),(16,'s_cups_tbl_area_inf','机构与银联地区对应信息','all'),(17,'s_cups_tbl_trans_log','交易流水表','add'),(18,'s_ebnk_cb_accinf','企业普通账号信息','add'),(19,'s_ebnk_cb_open','企业客户开户信息','add'),(20,'s_ebnk_cb_payagencyflow','企业网银代发指令流水表','add'),(21,'s_ebnk_cb_payment_flow','企业网银跨行快汇流水表','add'),(22,'s_ebnk_cb_tranflow','企业客户交易流水信息(当天查证)','add'),(23,'s_ebnk_cb_wage_flow','企业代发工资差旅流水表','add'),(24,'s_ebnk_pb_accinf','个人普通账号信息','add'),(25,'s_ebnk_pb_accinf_vip','个人VIP账号信息','add'),(26,'s_ebnk_pb_channelinf','个人注册客户渠道信息表','add'),(27,'s_ebnk_pb_channelinf_history','个人客户渠道信息表','add'),(28,'s_ebnk_pb_fee_record','个人支付缴费信息','add'),(29,'s_ebnk_pb_fixed_deposit_tranflow','理财业务流水表','add'),(30,'s_ebnk_pb_log','个人操作日志','add'),(31,'s_ebnk_pb_open','个人开户信息','add'),(32,'s_ebnk_pb_payment_flow','个人网银跨行快汇流水表','add'),(33,'s_ebnk_pb_tranflow','个人客户交易流水信息(当天查证)','add'),(34,'s_ebnk_t_ncctrjn','超级网银','add'),(35,'s_ebnk_t_nnstrjn','网上银行、手机银行','add'),(36,'s_ecif_dw_certtype','证件类型表','all'),(37,'s_ecif_dw_countrycode','国家代码信息','all'),(38,'s_ecif_ecif_black_list','黑名单登记薄','all'),(39,'s_ecif_ecif_cert_info','证件信息表','all'),(40,'s_ecif_ecif_custmast','客户信息表','all'),(41,'s_ecif_ecif_prdt_sign','产品签约表','all'),(42,'s_finc_bd_accsubj','科目表','all'),(43,'s_finc_bd_corp','公司目录','all'),(44,'s_finc_bd_deptdoc','部门档案','all'),(45,'s_finc_dwf_gl_daily','网点科目日余额表','add'),(46,'s_finc_dwf_gl_daily_after','网点科目日余额表结转后','add'),(47,'s_finc_dwf_gl_daily_sum','法人机构科目余额表','add'),(48,'s_finc_dwf_gl_daily_sum_after','法人机构科目余额表结转后','add'),(49,'s_fsms_sale_cus_info','客户信息','all'),(50,'s_fsms_sale_cust_sign_card_info','客户签约卡信息','all'),(51,'s_fsms_sale_cust_vol_delt_acct','客户持仓明细表','all'),(52,'s_fsms_sale_prd_baseinfo','产品基本要素表','all'),(53,'s_fsms_sale_prd_userdeal','客户交易流水表','all'),(54,'s_fsms_sale_sys_organization','机构管理信息','all'),(55,'s_fsms_tis_act_feedeal','账户费用交易表','add'),(56,'s_fsms_tis_prd_attr_ref','产品录入选项属性关联表','all'),(57,'s_fsms_tis_prd_baseinfo','产品基本要素表','all'),(58,'s_fsms_tis_prd_prindeal','产品存续期交易表','all'),(59,'s_fsms_tis_src_secschedule','债券付息还本计划表','all'),(60,'s_fsms_tis_sttc_datadict','参数表','all'),(61,'s_fsms_tis_sys_organization','机构管理信息','all'),(62,'s_fsms_tis_tru_deal','自定义资产交易信息表','all'),(63,'s_hrms_tb_org_job','职务基本信息','all'),(64,'s_hrms_tb_org_orgunit','机构信息','all'),(65,'s_hrms_tb_org_position','岗位信息','all'),(66,'s_hrms_tb_org_unitrelation','组织单元隶属关系','all'),(67,'s_hrms_tb_sta_emp','员工基本信息','all'),(68,'s_hrms_tb_sta_emp_org','员工组织信息','all'),(69,'s_iccs_tbl_acm_accbsc','财务明细基本信息表','all'),(70,'s_iccs_tbl_acm_acchst','财务历史表','add'),(71,'s_iccs_tbl_acm_txndtl','入账交易明细表','add'),(72,'s_iccs_tbl_csm_crdbsc','卡片基本信息表','all'),(73,'s_iccs_tbl_csm_cstbsc','客户基本信息表','all'),(74,'s_iccs_tbl_oam_mattxn','已匹配交易流水表','add'),(75,'s_iccs_tbl_pcm_brhinf','分支机构ID表','all'),(76,'s_icds_icd_dcacctdzxj','IC卡电子现金账户关联表','add'),(77,'s_icds_icd_ecashtrjn','电子现金交易登记簿','all'),(78,'s_icds_icd_hsecashtrjn','电子现金交易历史登记簿','all'),(79,'s_icds_icd_ichsmsg','IC借记卡历史交易登记簿','all'),(80,'s_icds_icd_icmsg','IC借记卡历史交易登记簿','all'),(81,'s_loan_dhjc','贷后检查','all'),(82,'s_loan_dk','贷款','all'),(83,'s_loan_dk_bb','贷款扩展表','all'),(84,'s_loan_dkbzr','贷款保证人','all'),(85,'s_loan_dkdy','贷款抵押','all'),(86,'s_loan_dksqgtjk','贷款申请共同借款人','all'),(87,'s_loan_dksqspx','贷款审批信息','all'),(88,'s_loan_dksqxx','贷款申请信息','all'),(89,'s_loan_dkyq','贷款延期','all'),(90,'s_loan_dkzy','贷款质押','all'),(91,'s_loan_e_dkhx','贷款核销','all'),(92,'s_loan_e_jtxx','集团信息','all'),(93,'s_loan_e_khsx','客户授信','add'),(94,'s_loan_e_wyqk','违约情况','all'),(95,'s_loan_e_xmxx','项目信息','all'),(96,'s_loan_ggyd','公共约定','all'),(97,'s_loan_hk','还款信息','add'),(98,'s_loan_hkjh','还款计划','add'),(99,'s_loan_jgb','机构表','all'),(100,'s_loan_khbb_bill','客户信息报表管理','all'),(101,'s_loan_khbb_data','客户信息报表数据','all'),(102,'s_loan_khbb_name','客户报表','all'),(103,'s_loan_khns','客户年审','add'),(104,'s_loan_khxx','客户信息','add'),(105,'s_loan_khxx_bb','客户信息扩展表','all'),(106,'s_loan_khxxfzqk','客户负债情况','all'),(107,'s_loan_khxxgl','客户关联人信息','add'),(108,'s_loan_khxxsrqk','客户收入情况','add'),(109,'s_loan_khxxzcmx','客户支出明细','add'),(110,'s_loan_khxxzcqk','客户资产情况','add'),(111,'s_loan_kjkmb','会计科目表','all'),(112,'s_loan_rmbdkjzll','人民币贷款基准利率资料','all'),(113,'s_loan_ryb','人员表','all'),(114,'s_loan_twjfl','调五级分类','add'),(115,'s_loan_txff','贴现买入','all'),(116,'s_loan_txmc','贴现卖出','all'),(117,'s_loan_txsh','贴现收回','all'),(118,'s_loan_xcdmx1','新承兑明细1','all'),(119,'s_loan_xcdmx2','新承兑明细2','all'),(120,'s_loan_xcdmx3','新承兑明细3','all'),(121,'s_loan_xcz','乡村组','all'),(122,'s_loan_xczcy','乡村组成员','all'),(123,'s_loan_xtxmx','贴现明细','all'),(124,'s_mbnk_mb_accinf','账户额度控制表','add'),(125,'s_mbnk_mb_cstinf','客户信息表','add'),(126,'s_mbnk_mb_cstinf_his','注销客户信息表','add'),(127,'s_mbnk_mb_tranflow','交易流水表','add'),(128,'s_mstp_devinfo','自助设备信息表','all'),(129,'s_mstp_t_device','设备表','all'),(130,'s_mstp_trnj','交易流水表','add'),(131,'s_ncs2_ncs_bankinfo','支付行信息','all'),(132,'s_ncs2_ncs_bill','银行汇票','add'),(133,'s_ncs2_ncs_bill_pay','银行汇票解付登记簿','add'),(134,'s_ncs2_ncs_billquery','银行汇票查询查复书','all'),(135,'s_ncs2_ncs_freemsg','电子汇兑自由格式查询书','add'),(136,'s_ncs2_ncs_pymtqry_reply','电子汇兑查复书','add'),(137,'s_ncs2_ncs_pymtquery','电子汇兑查询书','add'),(138,'s_ncs2_ncs_recvlist','电子汇兑来帐','all'),(139,'s_ncs2_ncs_sendlist','电子汇兑往帐','all'),(140,'s_ncs2_ncs_trjn','通存通兑','add'),(141,'s_ncs2_ncs_trjn_extdata','通存通兑附加数据表','add'),(142,'s_nibs_agent_drls','实时当日流水','all'),(143,'s_nibs_app_dwcpxy','单位产品协议表','all'),(144,'s_nibs_app_dwxx','单位信息表','all'),(145,'s_nibs_app_jyxx','交易码表','all'),(146,'s_nibs_app_khcpqy','客户签约信息','all'),(147,'s_nibs_app_khxx','客户信息表','all'),(148,'s_nibs_app_qdxx','渠道信息表','all'),(149,'s_nibs_ywty_hnbt_csb','惠农补贴业务参数表','all'),(150,'s_nibs_ywty_hnbt_lsmx','惠农一折通代发业务当前明细信息表','all'),(151,'s_nibs_ywty_hnbt_mx','惠农补贴代发业务当期明细表','all'),(152,'s_nibs_ywty_jmjkk_kyydj','居民健康卡应用登记表','all'),(153,'s_nibs_ywty_pldf_djb','批量代发登记簿','all'),(154,'s_nibs_ywty_pldf_mxdjb','批量代发明细登记簿','add'),(155,'s_nibs_ywty_pldf_qyxx','批量代发签约信息','all'),(156,'s_nibs_ywty_pldk_djb','批量代扣登记簿','all'),(157,'s_nibs_ywty_pldk_mxdjb','批量代扣明细登记簿','add'),(158,'s_nibs_ywty_pldk_qyxx','批量代扣签约信息','all'),(159,'s_nibs_ywty_plkh_djb','批量开户登记簿','all'),(160,'s_nibs_ywty_plkh_mxdjb','批量开户明细登记簿','all'),(161,'s_nibs_ywty_secur_pldj','社保批量登记登记表','all'),(162,'s_nibs_ywty_secur_plmx','社保批量代发代扣明细登记表','all'),(163,'s_nibs_ywty_xgd_plmx','新供电批量代发代扣明细','add'),(164,'s_nibs_ywty_xhx_csb','新核心业务参数表','all'),(165,'s_nibs_ywty_xhx_jyls','新核心交易流水表(非资金操作)','add'),(166,'s_nibs_ywty_xhx_ywls','新核心业务流水表','all'),(167,'s_nibs_ywty_xysgd_pldj','兴义市供电批量登记','all'),(168,'s_nibs_ywty_xysgd_plmx','兴义市批量交易时记录明细数据','all'),(169,'s_nxtp_tbl_fina_tran_log','金融交易流水表','add'),(170,'s_nxtp_tbl_mer_info','普通商户基本信息表','all'),(171,'s_nxtp_tbl_tem_info','终端基本信息表','all'),(172,'s_ofcr_ac_agreement_dtls','贷款合同借据明细表','add'),(173,'s_ofcr_ba_acct_status','账户状态','all'),(174,'s_ofcr_ba_cc_brn_mast','分行主表','all'),(175,'s_ofcr_ba_ccy_code','币种编号','all'),(176,'s_ofcr_ba_ccy_rate','币种汇率','all'),(177,'s_ofcr_ba_dormancy_table','睡眠户登记簿','all'),(178,'s_ofcr_ba_ext_branch_xref','内外部机构映射','all'),(179,'s_ofcr_ba_hold_funds','冻结登记簿','all'),(180,'s_ofcr_ba_int_indx_mast','利率索引维护','all'),(181,'s_ofcr_ba_int_indx_rate','利率索引','all'),(182,'s_ofcr_ba_prod_type_mast','产品类型维护','all'),(183,'s_ofcr_ch_acct_mast','活期账户主表','all'),(184,'s_ofcr_ch_int_rate_tier_plan','利率层级计划表','all'),(185,'s_ofcr_ch_nobook','活期帐户交易明细','add'),(186,'s_ofcr_ch_prod_mast','活期产品主表','all'),(187,'s_ofcr_ch_sweep_out','转出日志表','all'),(188,'s_ofcr_ch_sweep_out_txn_logs','转出交易日志表','all'),(189,'s_ofcr_ci_cust_types','客户类型','all'),(190,'s_ofcr_ci_custmast','客户主表','all'),(191,'s_ofcr_ci_ic_types','证件类型','all'),(192,'s_ofcr_ext_tf_batchagent','代发工资明细','add'),(193,'s_ofcr_ext_tf_batchagent_ddct','代扣明细','all'),(194,'s_ofcr_gl_table','科目定义表','all'),(195,'s_ofcr_glm_acct_bal_his','内部账历史余额记录表','all'),(196,'s_ofcr_glm_acct_mast','内部账登记簿（内部账分户账）','all'),(197,'s_ofcr_glm_com_item','内部科目表','all'),(198,'s_ofcr_glm_nobook','内部账交易流水','add'),(199,'s_ofcr_glm_prod_mast','内部账产品主表','all'),(200,'s_ofcr_gltb_gl_bal','总账科目余额','add'),(201,'s_ofcr_gltm_glmaster','总账科目主表','all'),(202,'s_ofcr_iv_mca_xref','凭证信息','add'),(203,'s_ofcr_ln_acct_attributes','贷款账户属性表','add'),(204,'s_ofcr_ln_acct_balances','贷款账户余额表','all'),(205,'s_ofcr_ln_acct_dtls','贷款账户信息表','all'),(206,'s_ofcr_ln_acct_int_balance_dtls','贷款利息余额表','all'),(207,'s_ofcr_ln_acct_int_balances','贷款利息余额表(视图)','all'),(208,'s_ofcr_ln_acct_payinstrn','贷款还款表','all'),(209,'s_ofcr_ln_acct_pricing_rate_detl','贷款利率定义表','add'),(210,'s_ofcr_ln_acct_rates','贷款账户利率表','add'),(211,'s_ofcr_ln_acct_rates_detl','贷款利率明细表','add'),(212,'s_ofcr_ln_acct_schedule','贷款还款计划信息','add'),(213,'s_ofcr_ln_acct_schedule_detls','贷款还款计划信息','add'),(214,'s_ofcr_ln_arrear_txn_hist','账户还款历史明细表','add'),(215,'s_ofcr_ln_arrears_table','贷款欠款表','add'),(216,'s_ofcr_ln_daily_txn_log_hist','贷款账户交易明细','add'),(217,'s_ofcr_ln_mat_ext_instr','贷款展期信息','all'),(218,'s_ofcr_ln_prod_int_attr','产品与会计科目定义表','all'),(219,'s_ofcr_ln_prod_mast','贷款产品主表','all'),(220,'s_ofcr_ln_tf_rcc_drv_crr_movement','农信版日终90天、正常状态处理驱动表','add'),(221,'s_ofcr_mc_acct_mast','MC_ACCT账户表','all'),(222,'s_ofcr_mc_acct_xref','MCA账户和分户账关联表','all'),(223,'s_ofcr_mc_pkg_mast','产品表','all'),(224,'s_ofcr_rec_txn_log','柜员操作信息','add'),(225,'s_ofcr_sm_user_profile','用户基本信息表','all'),(226,'s_ofcr_td_acct_mast','账户主表','all'),(227,'s_ofcr_td_dep_mast','定期账户主表','all'),(228,'s_ofcr_td_deposit_details','存款明细','add'),(229,'s_ofcr_td_nobook','定期登记薄','add'),(230,'s_ofcr_td_prod_mast','定期产品主表','all'),(231,'s_ofcr_td_prod_rates','产品分档利率','all'),(232,'s_ofcr_tf_acct_int_dtls','账户利息历史记录','add'),(233,'s_ofcr_tf_acct_int_var_log','账户利息浮动历史','add'),(234,'s_ofcr_tf_acct_open_close_log','开销户登记簿','all'),(235,'s_ofcr_tf_acct_stat_nobook','对账主文件表','add'),(236,'s_ofcr_tf_brn_clearing_xref','机构清算关系表','all'),(237,'s_ofcr_tf_cash_deposit_book','保证金登记薄','all'),(238,'s_ofcr_tf_cc_brn_mast','机构定义表','all'),(239,'s_ofcr_tf_channel_constant','交易渠道定义表','all'),(240,'s_ofcr_tf_cif_corp_add_inf','对公客户基本信息（附加）表','all'),(241,'s_ofcr_tf_cif_corp_inf_modify','对公账户资料修改','add'),(242,'s_ofcr_tf_cif_cust_modify','对私客户维护登记簿','add'),(243,'s_ofcr_tf_cif_indv_add_inf','对私客户附加信息表','add'),(244,'s_ofcr_tf_cif_pub_black_list','黑名单','all'),(245,'s_ofcr_tf_cif_rel_corp','对公客户关联表','all'),(246,'s_ofcr_tf_cif_stock_holder_inf','股东信息表','all'),(247,'s_ofcr_tf_cif_vip_officer','VIP客户经理','all'),(248,'s_ofcr_tf_date_sweep_reg','约定转存登记薄','all'),(249,'s_ofcr_tf_force_deduction','司法扣划登记簿','all'),(250,'s_ofcr_tf_hold_amt','账户冻结金额总额','all'),(251,'s_ofcr_tf_indv_depo_agmt_inf','个人协定存款','all'),(252,'s_ofcr_tf_inventory_attribute','凭证代码表','all'),(253,'s_ofcr_tf_inventory_mast_detail','凭证明细表','all'),(254,'s_ofcr_tf_judicial_query','有权机关查询登记薄','all'),(255,'s_ofcr_tf_large_dep_cert_transfer_reg','大额存单转让登记薄','all'),(256,'s_ofcr_tf_legacy_acctno_xref','新旧账号对照表','add'),(257,'s_ofcr_tf_ln_acct_crr','贷款户状态表','all'),(258,'s_ofcr_tf_ln_acct_crr_log','转减值流水登记簿','all'),(259,'s_ofcr_tf_ln_acct_info','贷款账户信息表','all'),(260,'s_ofcr_tf_ln_entrust_bod_payment','委托贷款BOD还款登记簿','all'),(261,'s_ofcr_tf_ln_nobook','贷款明细表','all'),(262,'s_ofcr_tf_ln_rcc_acct_crr_movement','贷款正常、90天互转登记簿','all'),(263,'s_ofcr_tf_maincard_limit_dtls','卡号累计金额登记簿','add'),(264,'s_ofcr_tf_make_card_parm','卡BIN信息表','all'),(265,'s_ofcr_tf_mc_old_new_cust','客户合并登记簿','all'),(266,'s_ofcr_tf_mdm_ac_rel','账号定义表','all'),(267,'s_ofcr_tf_mdm_attr','介质属性表','all'),(268,'s_ofcr_tf_mdm_fcrac_rel','账号关联表','all'),(269,'s_ofcr_tf_mdm_sttl_type','账户关联表','all'),(270,'s_ofcr_tf_medium_change_log','换折登记簿','add'),(271,'s_ofcr_tf_medium_loss_report','挂失登记簿','all'),(272,'s_ofcr_tf_sc_log','手续费明细','add'),(273,'s_ofcr_tf_si_inf','卡内预约转存登记薄','all'),(274,'s_ofcr_tf_sms_sign_mast','客户短信签约主表','add'),(275,'s_ofcr_tf_sms_sign_txnlog','短信签约流水表','add'),(276,'s_ofcr_tf_sms_telno_lst','签约手机号列表','add'),(277,'s_ofcr_tf_stop_inf','账户止付表','all'),(278,'s_ofcr_tf_wrapper_message','交易定义表','all'),(279,'s_ofcr_xf_stcap_gl_txns_mmdd','交易定义表','add'),(280,'s_ofcr_xface_addl_details_txnlog','交易流水表','add'),(281,'s_pmts_beps_crdt_rcvlt','小额贷记来账','add'),(282,'s_pmts_beps_crdt_sndlt','小额贷记往账','add'),(283,'s_pmts_beps_dbit_rcvlt','小额借记申请(来账)','add'),(284,'s_pmts_beps_dbit_recv_replylt','小额借记申请回执(来账)','add'),(285,'s_pmts_beps_dbit_snd_replylt','小额借记申请回执(往账)','add'),(286,'s_pmts_beps_dbit_sndlt','小额借记申请(往账)','add'),(287,'s_pmts_ccms_bankinfo','行名行号信息表','all'),(288,'s_pmts_ccms_pymtqry_reply','查复书','add'),(289,'s_pmts_ccms_pymtquery','查询书','add'),(290,'s_pmts_hvps_recvlist','大额来帐明细表:保存大额来帐交易明细','add'),(291,'s_pmts_hvps_rt_recvlist','大额即时转账业务表','all'),(292,'s_pmts_hvps_sendlist','大额往账','add'),(293,'s_upss_t_id_chk_book','证件核查表','all'),(294,'s_upss_t_ups_book','通存通兑登记簿','add'),(295,'s_upss_t_ups_cbook','通存通兑回执信息','add'),(296,'s_upss_t_ups_qry','查询查复书','add'),(297,'t_dis_com_lar_dep_cust','特殊处理结果表','all'),(298,'t_dis_com_loan_five_cust','特殊处理结果表','all'),(299,'t_dis_com_new_add_dep_cust','特殊处理结果表','all'),(300,'t_dis_com_v_gl_pzls_to_ods','特殊处理结果表','all');

/*Table structure for table `table_task_time` */

DROP TABLE IF EXISTS `table_task_time`;

CREATE TABLE `table_task_time` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '任务执行自增id',
  `date` date DEFAULT NULL COMMENT '日期文件夹',
  `start_time` datetime DEFAULT NULL COMMENT '任务开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '任务截止时间',
  `state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '任务是否结束',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `table_task_time` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
