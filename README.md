## PostgresSQL User Management Service Serializer/Migration tool:
Shape large amounts of data into psql inserts. Adjust source & build as needed.

> Example of migration output:
```
INSERT INTO user_member (id, name, email, user_role, phone, image_url, user_secret) VALUES
('0', 'Sofia Lowe', 'corkery.corrine@schneider.com', 'Merchant', '(562) 246-7846', 'https://i.postimg.cc/7hGxR500/28.png', 'a89vsad0ansaodjsa90vas0dmfnfslalkasmdv');
INSERT INTO user_member (id, name, email, user_role, phone, image_url, user_secret) VALUES
('1', 'Erik Lubowitz', 'susan52@yahoo.com', 'Doorman', '(573) 248-0225', 'https://i.postimg.cc/qBSB3GHL/30.png', 'a89vsad0ansaodjsa90vas0dmfnfslalkasmdv');
INSERT INTO user_member (id, name, email, user_role, phone, image_url, user_secret) VALUES
('2', 'Willow Barton', 'darrel.boyer@gmail.com', 'Counsellor', '(205) 107-7163', 'https://i.postimg.cc/vDwdS3Qw/41.png', 'a89vsad0ansaodjsa90vas0dmfnfslalkasmdv');
INSERT INTO user_member (id, name, email, user_role, phone, image_url, user_secret) VALUES
('3', 'Ezequiel Altenwerth', 'bethany78@bradtke.biz', 'Medical Assistant', '(465) 538-6730', 'https://i.postimg.cc/4NcfbhJT/32.png', 'a89vsad0ansaodjsa90vas0dmfnfslalkasmdv');
INSERT INTO user_member (id, name, email, user_role, phone, image_url, user_secret) VALUES
('4', 'Mallie Ruecker', 'kwilderman@gmail.com', 'Catering staff', '(741) 010-7113', 'https://i.postimg.cc/9XRRK0WL/40.png', 'a89vsad0ansaodjsa90vas0dmfnfslalkasmdv');
INSERT INTO user_member (id, name, email, user_role, phone, image_url, user_secret) VALUES
('5', 'Courtney Watsica', 'bsanford@gmail.com', 'Receptionist', '(070) 647-8375', 'https://i.postimg.cc/907hpvHB/39.png', 'a89vsad0ansaodjsa90vas0dmfnfslalkasmdv');

#etc . . .
```
