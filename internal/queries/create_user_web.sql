create user 'web'@'localhost';

grant select, insert, update, delete ON snippetbox.* TO 'web'@'localhost';

alter user 'web'@'localhost' identified by 'golang';
