CREATE DATABASE house_manager;



CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(225) UNIQUE NOT NULL,
  email VARCHAR(225) UNIQUE NOT NULL,
  password VARCHAR(225) NOT NULL
);

CREATE TABLE user_session (
  session_id INT AUTO_INCREMENT PRIMARY KEY,  
  user_id INT REFERENCES user(id) ON DELETE CASCADE,                      
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  
  expires_at DATETIME                       
);
