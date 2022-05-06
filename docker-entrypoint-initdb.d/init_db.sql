CREATE TABLE request (
    request_id varchar(255),
    request_time varchar(255),
    type varchar(255),
    body text
);

CREATE TABLE response (
    request_id varchar(255),
    response_time varchar(255),
    type varchar(255),
    status varchar(255),
    body text,
    error_message text
);