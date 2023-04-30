CREATE TABLE users(
	userID SERIAL PRIMARY KEY,
	userFullName VARCHAR(256) NOT NULL,
	userEmail VARCHAR(256) NOT NULL
);

CREATE TABLE passwords(
	passwordID SERIAL PRIMARY KEY,
	encryptedData TEXT NOT NULL
);

CREATE TABLE orders(
	orderID SERIAL PRIMARY KEY,
	orderNote VARCHAR(2048)
);

CREATE TABLE product(
	productID SERIAL PRIMARY KEY,
	productPrice NUMERIC NOT NULL,
	productName VARCHAR(256) NOT NULL,
	productDescription VARCHAR(2048) NOT NULL
);

CREATE TABLE roles(
	roleID SERIAL PRIMARY KEY,
	roleName VARCHAR(256) UNIQUE NOT NULL
);

CREATE TABLE events(
	eventID SERIAL PRIMARY KEY,
	eventTimestamp TIMESTAMP NOT NULL,
	eventDescription VARCHAR(2048) NOT NULL
);

CREATE TABLE orderItem(
	orderItemID SERIAL PRIMARY KEY
);

CREATE TABLE userPassword(
	userPasswordID SERIAL PRIMARY KEY,
	passwordID INTEGER REFERENCES passwords (passwordID) NOT NULL,
	userID INTEGER REFERENCES users (userID) NOT NULL
);

CREATE TABLE productItem(
	productItemID SERIAL PRIMARY KEY,
	orderItemID INTEGER REFERENCES orderItem (orderItemID) NOT NULL,
	productID INTEGER REFERENCES product (productID) NOT NULL,
	productItemQuantity INTEGER CHECK (productItemQuantity > 0),
	productItemPrice NUMERIC NOT NULL
);

CREATE TABLE placedOrder(
	placedOrderID SERIAL PRIMARY KEY,
	orderID INTEGER REFERENCES orders (orderID) NOT NULL,
	productItemID INTEGER REFERENCES productItem (productItemID) NOT NULL,
	userID INTEGER REFERENCES users (userID) NOT NULL
);

CREATE TABLE addedProduct(
	addedProductID SERIAL PRIMARY KEY,
	productID INTEGER REFERENCES product (productID) NOT NULL,
	userID INTEGER REFERENCES users (userID) NOT NULL
);

CREATE TABLE userRole(
	userRoleID SERIAL PRIMARY KEY,
	userID INTEGER REFERENCES users (userID) NOT NULL,
	roleID INTEGER REFERENCES roles (roleID) NOT NULL
);

CREATE TABLE orderPlacementEvent(
	orderPlacementEventID SERIAL PRIMARY KEY,
	eventID INTEGER REFERENCES events (eventID) NOT NULL,
	placedOrderID INTEGER REFERENCES placedOrder (placedOrderID) NOT NULL
);

CREATE TABLE productAddEvent(
	productAddEventID SERIAL PRIMARY KEY,
	eventID INTEGER REFERENCES events (eventID) NOT NULL,
	addedProductID INTEGER REFERENCES addedProduct (addedProductID) NOT NULL
);

CREATE TABLE userRegisterEvent(
	userRegisterEventID SERIAL PRIMARY KEY,
	eventID INTEGER REFERENCES events (eventID) NOT NULL,
	userID INTEGER REFERENCES users (userID) NOT NULL
);

CREATE TABLE userLoginEvent(
	userLoginEventID SERIAL PRIMARY KEY,
	eventID INTEGER REFERENCES events (eventID) NOT NULL,
	userID INTEGER REFERENCES users (userID) NOT NULL
);

CREATE TABLE roleAssignmentEvent(
	roleAssignmentEventID SERIAL PRIMARY KEY,
	eventID INTEGER REFERENCES events (eventID) NOT NULL,
	userID INTEGER REFERENCES users (userID) NOT NULL,
	userRoleID INTEGER REFERENCES userRole (userRoleID) NOT NULL
);



/*Inserts*/

/*Users*/
INSERT INTO users (userFullName, userEmail) VALUES
('Dummy User 1', 'dummyuser1@example.com'),
('Dummy User 2', 'dummyuser2@example.com'),
('Dummy User 3', 'dummyuser3@example.com'),
('Dummy User 4', 'dummyuser4@example.com'),
('Dummy User 5', 'dummyuser5@example.com');

/*Orders*/
INSERT INTO orders (orderNote) VALUES
('Dummy Order 1'),
('Dummy Order 2'),
('Dummy Order 3'),
('Dummy Order 4'),
('Dummy Order 5'),
('Dummy Order 6'),
('Dummy Order 7'),
('Dummy Order 8'),
('Dummy Order 9'),
('Dummy Order 10');

/*Products*/
INSERT INTO product (productPrice, productName, productDescription) VALUES 
(10.99, 'Dummy Product 1', 'This is the first dummy product.'),
(25.99, 'Dummy Product 2', 'This is the second dummy product.'),
(8.49, 'Dummy Product 3', 'This is the third dummy product.'),
(15.99, 'Dummy Product 4', 'This is the fourth dummy product.'),
(19.99, 'Dummy Product 5', 'This is the fifth dummy product.'),
(12.49, 'Dummy Product 6', 'This is the sixth dummy product.'),
(11.99, 'Dummy Product 7', 'This is the seventh dummy product.'),
(29.99, 'Dummy Product 8', 'This is the eighth dummy product.'),
(7.99, 'Dummy Product 9', 'This is the ninth dummy product.'),
(22.99, 'Dummy Product 10', 'This is the tenth dummy product.');

/*Roles*/
INSERT INTO roles (roleName) VALUES
('Admin'),
('Customer');

/*Events*/
INSERT INTO events (eventTimestamp, eventDescription) VALUES
('2019-01-01 00:00:00', 'Dummy Event 1'),
('2019-01-02 00:00:00', 'Dummy Event 2'),
('2019-01-03 00:00:00', 'Dummy Event 3'),
('2019-01-04 00:00:00', 'Dummy Event 4'),
('2019-01-05 00:00:00', 'Dummy Event 5'),
('2019-01-06 00:00:00', 'Dummy Event 6'),
('2019-01-07 00:00:00', 'Dummy Event 7'),
('2019-01-08 00:00:00', 'Dummy Event 8'),
('2019-01-09 00:00:00', 'Dummy Event 9'),
('2019-01-10 00:00:00', 'Dummy Event 10');

/*Order Items*/
INSERT INTO orderItem DEFAULT VALUES;
INSERT INTO orderItem DEFAULT VALUES;
INSERT INTO orderItem DEFAULT VALUES;
INSERT INTO orderItem DEFAULT VALUES;
INSERT INTO orderItem DEFAULT VALUES;
INSERT INTO orderItem DEFAULT VALUES;
INSERT INTO orderItem DEFAULT VALUES;

/*Product Items*/
INSERT INTO productItem (orderItemID, productID, productItemQuantity, productItemPrice) VALUES
(1, 1, 1, 10.99),
(2, 2, 1, 25.99),
(3, 3, 1, 8.49),
(4, 4, 1, 15.99),
(5, 5, 1, 19.99),
(6, 6, 1, 12.49),
(7, 7, 1, 11.99),
(1, 8, 1, 29.99),
(2, 9, 1, 7.99),
(3, 10, 1, 22.99);

/*Placed Orders*/
INSERT INTO placedOrder (orderID, productItemID, userID) VALUES
(1, 1, 1),
(2, 2, 2),
(3, 3, 3),
(4, 4, 4),
(5, 5, 5),
(6, 6, 1),
(7, 7, 2),
(8, 8, 3),
(9, 9, 4),
(10, 10, 5);

/*Added Products*/
INSERT INTO addedProduct (productID, userID) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 1),
(7, 2),
(8, 3),
(9, 4),
(10, 5);

/*User Roles*/
INSERT INTO userRole (userID, roleID) VALUES
(1, 1),
(2, 2),
(3, 2),
(4, 2),
(5, 2);

/*Order Placement Events*/
INSERT INTO orderPlacementEvent (eventID, placedOrderID) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 6),
(7, 7),
(8, 8),
(9, 9),
(10, 10);

/*Product Add Events*/
INSERT INTO productAddEvent (eventID, addedProductID) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 6),
(7, 7),
(8, 8),
(9, 9),
(10, 10);

/*User Register Events*/
INSERT INTO userRegisterEvent (eventID, userID) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 1),
(7, 2),
(8, 3),
(9, 4),
(10, 5);

/*User Login Events*/
INSERT INTO userLoginEvent (eventID, userID) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 1),
(7, 2),
(8, 3),
(9, 4),
(10, 5);

/*Role Assignment Events*/
INSERT INTO roleAssignmentEvent (eventID, userID, userRoleID) VALUES
(1, 1, 1),
(2, 2, 2),
(3, 3, 2),
(4, 4, 2),
(5, 5, 2),
(6, 1, 1),
(7, 2, 2),
(8, 3, 2),
(9, 4, 2),
(10, 5, 2);
