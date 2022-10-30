create table if not exists customers (
customer_id int not null,
customer_name char(50) not null,
primary key(customer_id)
);

create table if not exists shipments (
shipment_id int not null,
name char(50) not null,
primary key(shipment_id)
);

create table if not exists orders (
order_id int not null,
customer_id int references customers(customer_id),
product_id int references shipments(shipment_id),
order_date date not null,
total float8 not null,

primary key(order_id)
);

--alter table orders add constraint fk_customers foreign key(customer_id)
--	references customers(customer_id);
		
--alter table orders constraint fk_shipments foreign key(product_id)
--	references shipments(shipment_id);