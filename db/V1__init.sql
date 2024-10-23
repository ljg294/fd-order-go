DROP TABLE IF EXISTS sales_order;

#-----------------------------------------------------------------------------------------------------------------------
#-- Table 명 : sales_order (주문)
#-----------------------------------------------------------------------------------------------------------------------
CREATE TABLE sales_order
(
    sales_order_id            BIGINT                   NOT NULL  AUTO_INCREMENT  COMMENT '주문ID',
    user_id                   BIGINT                   NOT NULL                  COMMENT '주문자ID',
    order_datetime            DATETIME         	       NOT NULL                  COMMENT '주문일시',
    delete_yn                 varchar(1)               NOT NULL                  COMMENT '삭제여부',
    created_user_id           BIGINT           	       NOT NULL                  COMMENT '등록자ID',
    created_datetime          DATETIME         	       NOT NULL                  COMMENT '등록일시',
    modified_user_id          BIGINT           	       NOT NULL                  COMMENT '수정자ID',
    modified_datetime         DATETIME         	       NOT NULL                  COMMENT '수정일시',

    CONSTRAINT pk_sales_order PRIMARY KEY (sales_order_id)
) ENGINE = INNODB DEFAULT CHARSET=utf8mb4 COMMENT='주문';
