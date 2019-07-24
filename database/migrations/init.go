package migrations

import (

	goodsPB "github.com/gomsa/goods/proto/goods"
	db "github.com/gomsa/goods/providers/database"
)

func init() {
	barcode()
	goods()
	brand()
	category()
	firm()
	unspsc()
	taxcode()
}

// barcode 商品条码数据迁移
func barcode() {
	barcode := &goodsPB.Barcode{}
	if !db.DB.HasTable(&barcode) {
		db.DB.Exec(`
			CREATE TABLE barcodes (
			id int(36) NOT NULL COMMENT '商品条码',
			goods_id varchar(36) DEFAULT NULL COMMENT '商品UUID',
			stock_id varchar(36) DEFAULT NULL COMMENT '库存自动UUID(多条码可以使用一个库存[禁止不同价格使用同一库存UUID])',
			price int(32) DEFAULT NULL COMMENT '价格',
			width int(32) DEFAULT NULL COMMENT '宽(毫米)',
			height int(32) DEFAULT NULL COMMENT '高(毫米)',
			depth int(32) DEFAULT NULL COMMENT '深(毫米)',
			unit varchar(16) DEFAULT NULL COMMENT '单位(件、个、斤...)',
			spec varchar(64) DEFAULT NULL COMMENT '规格',
			grossweight int(64) DEFAULT NULL COMMENT '总重(克)',
			netweight int(64) DEFAULT NULL COMMENT '净重(克)',
			img varchar(128) DEFAULT NULL COMMENT '主图',
			barcode_img varchar(128) DEFAULT NULL COMMENT '条码图',
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// goods 商品数据迁移
func goods() {
	good := &goodsPB.Good{}
	if !db.DB.HasTable(&good) {
		db.DB.Exec(`
			CREATE TABLE goods (
			id varchar(36) NOT NULL COMMENT '商品ID(UUID)',
			name varchar(64) DEFAULT NULL COMMENT '商品名称',
			eng_name varchar(64) DEFAULT NULL COMMENT '英文名称',
			desc varchar(255) DEFAULT NULL COMMENT '商品描述',
			cess varchar(8) DEFAULT NULL COMMENT '税率',
			status  int(11) DEFAULT NULL COMMENT '商品状态(下架-1、上架1、线上2、外卖3)',
			brand_id int(11) DEFAULT NULL COMMENT '品牌ID',
			category_id int(11) DEFAULT NULL COMMENT '商品分类ID',
			department_id int(11) DEFAULT NULL COMMENT '部门分类ID',
			firm_id int(11) DEFAULT NULL COMMENT '所属公司ID',
			unspsc_id int(11) DEFAULT NULL COMMENT '商品及服务编码',
			taxcode_id int(11) DEFAULT NULL COMMENT '税收分类编码',
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// brand 商品商标数据迁移
func brand() {
	brand := &goodsPB.Brand{}
	if !db.DB.HasTable(&brand) {
		db.DB.Exec(`
			CREATE TABLE brands (
			id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '品牌ID',
			name varchar(36) DEFAULT NULL COMMENT '品牌名称',
			logo varchar(64) DEFAULT NULL COMMENT '品牌LOGO',
			desc varchar(64) DEFAULT NULL COMMENT '品牌描述',
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// category 商品分类数据迁移
func category() {
	category := &goodsPB.Category{}
	if !db.DB.HasTable(&category) {
		db.DB.Exec(`
			CREATE TABLE categorys (
			id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
			parent int(11) DEFAULT NULL COMMENT '上级分类ID,
			name varchar(36) DEFAULT NULL COMMENT '分类名称',
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// firm 商品公司数据迁移
func firm() {
	firm := &goodsPB.Firm{}
	if !db.DB.HasTable(&firm) {
		db.DB.Exec(`
			CREATE TABLE firms (
			id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '公司ID',
			name varchar(36) DEFAULT NULL COMMENT '公司名称',
			address varchar(128) DEFAULT NULL COMMENT '公司地址',
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// unspsc 商品及服务编码数据迁移
func unspsc() {
	unspsc := &goodsPB.Unspsc{}
	if !db.DB.HasTable(&unspsc) {
		db.DB.Exec(`
			CREATE TABLE unspscs (
			id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
			parent int(11) DEFAULT NULL COMMENT '上级分类ID,
			name varchar(36) DEFAULT NULL COMMENT '分类名称',
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// taxcode 税收分类编码数据迁移
func taxcode() {
	taxcode := &goodsPB.Taxcode{}
	if !db.DB.HasTable(&taxcode) {
		db.DB.Exec(`
			CREATE TABLE taxcodes (
			id int(32) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
			parent int(32) DEFAULT NULL COMMENT '上级分类ID,
			is_parent int(1) DEFAULT NULL COMMENT '是否是上级',
			parent int(32) DEFAULT NULL COMMENT '上级分类ID,
			cess varchar(8) DEFAULT NULL COMMENT '税率',
			duty_free varchar(64) DEFAULT NULL COMMENT '税收减免情况',
			duty_free_desc varchar(255) DEFAULT NULL COMMENT '税收减免情况详情',
			code int(32) DEFAULT NULL COMMENT '编码',
			name varchar(64) DEFAULT NULL COMMENT '名称',
			instruction varchar(255) DEFAULT NULL COMMENT '说明', 
			keyword varchar(128) DEFAULT NULL COMMENT '关键词', 
			version varchar(32) DEFAULT NULL COMMENT '版本',
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}
