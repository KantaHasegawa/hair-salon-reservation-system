create table reservations(
  id varchar(255) not null PRIMARY KEY,
  customer_id varchar(255) not null,
  beautician_id varchar(255) not null,
  menu_id varchar(255) not null,
  start_time datetime not null,
  end_time datetime not null,
  price int not null,
  FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
  FOREIGN KEY (beautician_id) REFERENCES beauticians(id) ON DELETE CASCADE,
  FOREIGN KEY (menu_id) REFERENCES menus(id) ON DELETE CASCADE
);
