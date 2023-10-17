CREATE TABLE tasks (
    id INT NOT NULL,
    user_id INT NOT NULL,
    title VARCHAR(255),
    description TEXT,
    status VARCHAR(255),
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    KEY `taks_FK` (`user_id`),
    CONSTRAINT `taks_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
)ENGINE = InnoDB;