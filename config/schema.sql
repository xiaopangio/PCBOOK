create database if not exists pcbook character set utf8;

use pcbook;
create table if not exists role(
    role_id int primary key auto_increment ,
    role_name varchar(10) not null
);
create table if not exists user(
    user_id int primary key auto_increment,
    username varchar(15) not null ,
    hashed_password varchar(100) not null ,
    role_id int,
    key r_id (role_id),
    constraint r_id foreign key (role_id) references role (role_id)
);
create table if not exists cpu(
    cpu_id int primary key auto_increment,
    brand varchar(10) not null ,
    name varchar(10) not null ,
    numbers_cores int,
    number_threads int,
    min_ghz float,
    max_ghz float
);
create table if not exists memory_unit(
    memory_unit_id int primary key auto_increment,
    unit_name varchar(10)
);
create table if not exists memory(
    memory_id int primary key auto_increment,
    value int ,
    unit_id int,
    constraint unit_id foreign key (unit_id) references memory_unit(memory_unit_id)
);
create table if not exists gpu(
    gpu_id int primary key auto_increment,
    brand varchar(10),
    name varchar(10),
    min_ghz float,
    max_ghz float,
    memory_id int,
    key m_id (memory_id),
    constraint m_id foreign key (memory_id) references memory(memory_id)
);
create table if not exists laptap_gpu(
    laptap_id varchar(20) ,
    gpu_id int,
    primary key (laptap_id,gpu_id)
);
create table if not exists storage_driver(
  storage_driver_id int primary key,
  name varchar(10)
);
create table if not exists storage(
    storage_id int primary key,
    driver_id int,
    memory_id int,
    constraint driver foreign key (driver_id) references storage_driver(storage_driver_id),
    constraint memory foreign key (memory_id) references memory (memory_id)
);
create table if not exists laptap_storage(
    laptap_id varchar(20),
    storage_id int,
    primary key (laptap_id,storage_id)
);
create table if not exists screen_resolution(
    screen_resolution_id int primary key,
    width int,
    height int
);
create table if not exists screen_panel(
    panel_id int primary key,
    name varchar(10)
);
create table if not exists screen (
    screen_id int primary key auto_increment,
    size_inch float,
    resolution_id int,
    panel_id int,
    multitouch int,
    constraint resolution foreign key (resolution_id)references screen_resolution(screen_resolution_id),
    constraint panel foreign key (panel_id)references screen_panel(panel_id)
);
create table if not exists keyboard_layout(
    layout_id int primary key ,
    name varchar(10)
);
create table if not exists keyboard(
    keyboard_id int primary key auto_increment,
    layout_id int,
    backlit bool,
    constraint layout foreign key (layout_id)references keyboard_layout(layout_id)
);
create table if not exists weight_unit(
    weight_unit_id int ,
    type varchar(5),
    primary key (weight_unit_id,type)
);
create table if not exists weight(
    weight_id int primary key auto_increment,
    value float,
    weight_unit int,
    constraint weight_unit foreign key (weight_id)references weight_unit(weight_unit_id)
);
create table if not exists laptap(
    laptap_id varchar(20) primary key ,
    brand varchar(10),
    name varchar(10),
    cpu_id int,
    ram_id int,
    screen_id int,
    keyboard_id int,
    weight_id int,
    price_usd float,
    release_year int,
    update_at datetime,
    constraint cpu foreign key (cpu_id)references cpu(cpu_id),
    constraint ram foreign key (ram_id)references memory(memory_id),
    constraint screen foreign key (screen_id)references screen(screen_id),
    constraint keyboard foreign key (keyboard_id)references storage(storage_id),
    constraint weight foreign key (weight_id)references weight(weight_id)
)
