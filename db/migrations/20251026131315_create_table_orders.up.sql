CREATE TABLE orders
(
    id           VARCHAR(255) NOT NULL,
    product_id   VARCHAR(255) NOT NULL,
    total_price  BIGINT       NOT NULL,
    status       VARCHAR(255) NOT NULL,
    created_at   BIGINT       NOT NULL,

    PRIMARY KEY (id)
);
