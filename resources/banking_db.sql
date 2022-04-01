CREATE TABLE customers (
                            customer_ID INT PRIMARY KEY NOT NULL,
                            name char(100) NOT NULL,
                            date_of_birth date NOT NULL,
                            city char(100) NOT NULL,
                            zipcode char(10) NOT NULL,
                            status int NOT NULL DEFAULT '1'
);

INSERT INTO customers VALUES
                            (2000,'Steve','1978-12-15','Delhi','110075',1),
                            (2001,'Arian','1988-05-21','Newburgh, NY','12550',1),
                            (2002,'Hadley','1988-04-30','Englewood, NJ','07631',1),
                            (2003,'Ben','1988-01-04','Manchester, NH','03102',0),
                            (2004,'Nina','1988-05-14','Clarkston, MI','48348',1),
                            (2005,'Osman','1988-11-08','Hyattsville, MD','20782',0);



CREATE TABLE accounts (
                            account_id int primary key NOT NULL,
                            customer_id int NOT NULL,
                            opening_date date NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            account_type varchar(10) NOT NULL,
                            amount decimal(10,2) NOT NULL,
                            status int NOT NULL DEFAULT '1'
                    );

select * from accounts;

INSERT INTO accounts VALUES
                           (95470,2000,'2020-08-22 10:20:06', 'saving', 6823.23, 1),
                           (95471,2002,'2020-08-09 10:27:22', 'checking', 3342.96, 1),
                           (95472,2001,'2020-08-09 10:35:22', 'saving', 7000, 1),
                           (95473,2001,'2020-08-09 10:38:22', 'saving', 5861.86, 1);


CREATE TABLE transactions
(
    transaction_id   int primary key NOT NULL,
    account_id       int             NOT NULL,
    amount           decimal(10, 2)  NOT NULL,
    transaction_type varchar(10)     NOT NULL,
    transaction_date date            NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
                         username varchar(20) NOT NULL,
                         password varchar(20) NOT NULL,
                         role varchar(20) NOT NULL,
                         customer_id int DEFAULT NULL,
                         created_on date NOT NULL DEFAULT CURRENT_TIMESTAMP
);
select * from users;
INSERT INTO users VALUES
                        ('admin','abc123','admin', NULL, '2020-08-09 10:27:22'),
                        ('2001','abc123','user', 2001, '2020-08-09 10:27:22'),
                        ('2000','abc123','user', 2000, '2020-08-09 10:27:22');

select * from refresh_token_store;
CREATE TABLE refresh_token_store (
                                       refresh_token varchar(300) primary key NOT NULL,
                                       created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);