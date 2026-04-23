--drop table public.subscriptions;
CREATE USER effective WITH PASSWORD '~!@effective';
grant all privileges on schema public to effective;
CREATE TABLE IF NOT EXISTS public.subscriptions (
subscription_id BIGSERIAL PRIMARY KEY,
service_name varchar(256) NOT NULL UNIQUE,
price int,
user_id UUID,
start_date TIMESTAMP,
create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
GRANT ALL PRIVILEGES ON public.subscriptions TO effective;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO effective;
insert into public.subscriptions(service_name, price, user_id, 
start_date)
values
('test', 400, '60601fee-2bf1-4721-ae6f-7636e79a0cba', CURRENT_TIMESTAMP)
