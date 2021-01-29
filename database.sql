CREATE DATABASE IF NOT EXISTS WebScraper;

USE WebScraper

CREATE TABLE Advertisement (
    id INT NOT NULL AUTO_INCREMENT ,
    district VARCHAR (255) NOT NULL ,
    category VARCHAR (255) NOT NULL ,
    title VARCHAR (255) ,
    description VARCHAR (255) ,
    price VARCHAR (255) ,
    url VARCHAR (255) ,
    postedOn VARCHAR (255) ,
    forSaleBy VARCHAR (255) ,
    meta TEXT ,
    fullDescription TEXT (255) ,
    CONSTRAINT PRIMARY KEY (id)
)ENGINE=INNODB;

