create table notifications(
    id bigserial primary key,
    title text not null,
    body text not null,
    userId int
)