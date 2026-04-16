CREATE DATABASE bookingservicedb IF NOT EXISTS;
CREATE DATABASE hotelservicedb IF NOT EXISTS;

GRANT ALL PRIVILEGES ON bookingservicedb.* TO 'user'@'%';
GRANT ALL PRIVILEGES ON hotelservicedb.* TO 'user'@'%';

FLUSH PRIVILEGES;