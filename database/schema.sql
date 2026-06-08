CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,

    deleted_at DATETIME NULL,
    
    PRIMARY KEY (id),
    UNIQUE KEY uk_email (email)
);

CREATE TABLE artists (
    id INT NOT NULL AUTO_INCREMENT,
    api_id VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,

    deleted_at DATETIME NULL,
    
    PRIMARY KEY (id),
    UNIQUE KEY uk_artist_name (name)
);

CREATE TABLE product_types (
    id INT NOT NULL AUTO_INCREMENT,
    type_name VARCHAR(100) NOT NULL,

    deleted_at DATETIME NULL,
    
    PRIMARY KEY(id),
    UNIQUE KEY uk_product_types_type (type_name)
);

INSERT INTO product_types (type_name) VALUES 
('cd'),
('vinyl'),
('vhs'),
('tape');

CREATE TABLE products (
    id INT NOT NULL AUTO_INCREMENT,
    api_id VARCHAR(50) NOT NULL,
    type_id INT NOT NULL,
    release_date DATE NULL,
    title VARCHAR(255) NOT NULL,
    artist_id INT NOT NULL,
    price INT NOT NULL,
    cover VARCHAR(255) NULL,

    deleted_at DATETIME NULL,
    
    PRIMARY KEY (id),
    FOREIGN KEY (type_id) REFERENCES product_types(id) ON DELETE RESTRICT,
    FOREIGN KEY (artist_id) REFERENCES artists(id) ON DELETE RESTRICT
);

CREATE TABLE inventory (
    id INT NOT NULL AUTO_INCREMENT,
    product_id INT NOT NULL,
    stock INT NOT NULL,

    deleted_at DATETIME NULL,

    PRIMARY KEY(id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE RESTRICT
);

CREATE TABLE inventory_log_type (
    id INT NOT NULL AUTO_INCREMENT,
    type_name VARCHAR(15) NOT NULL,

    deleted_at DATETIME NULL,

    PRIMARY KEY(id)
);

INSERT INTO inventory_log_type (type_name) VALUES 
('entry'),
('exit');


CREATE TABLE inventory_log (
    id INT NOT NULL AUTO_INCREMENT,
    product_id INT NOT NULL,
    amount INT NOT NULL,
    type_id INT NOT NULL,

    deleted_at DATETIME NULL,

    PRIMARY KEY(id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE RESTRICT,
    FOREIGN KEY (type_id) REFERENCES inventory_log_type(id) ON DELETE RESTRICT
);

CREATE TABLE sales (
    id INT NOT NULL AUTO_INCREMENT,
    date_time DATETIME DEFAULT CURRENT_TIMESTAMP,

    deleted_at DATETIME NULL,
    
    PRIMARY KEY(id)
);

CREATE TABLE sale_item (
    id INT NOT NULL AUTO_INCREMENT,
    sale_id INT NOT NULL,
    product_id INT NOT NULL,
    product_amount INT NOT NULL,

    deleted_at DATETIME NULL,
    
    PRIMARY KEY (id),
    FOREIGN KEY (sale_id) REFERENCES sales(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE RESTRICT
);

-- INDEXES
CREATE INDEX idx_artist_name ON artists(name, deleted_at);
CREATE INDEX idx_product_title ON products(title, deleted_at);
