CREATE TABLE playlist (
                                 number serial NOT NULL ,
                                 name varchar(70) NOT NULL UNIQUE ,
                                 duration int NOT NULL
);