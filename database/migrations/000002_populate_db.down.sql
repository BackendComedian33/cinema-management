
-- 1. DEPENDENT TABLES FIRST (FK SAFE ORDER)

-- Tickets depend on orders & seat_showtimes
DELETE FROM tickets;

-- Refunds depend on orders
DELETE FROM refunds;

-- Orders depend on client_users & showtimes
DELETE FROM orders;

-- Seat showtimes depend on showtimes & studio_seat
DELETE FROM seat_showtimes;

-- Showtimes depend on movies & studios
DELETE FROM showtimes;

-- 2. CORE DOMAIN DATA

-- Studio seats depend on studios
DELETE FROM studio_seat;

-- Studios depend on cinemas
DELETE FROM studios;

-- Movies
DELETE FROM movies;

-- Cinemas
DELETE FROM cinemas;

-- 3. USERS

-- Client users (customers)
DELETE FROM client_users;

-- Admin / staff users
DELETE FROM users;

-- Optional: reset sequences (PostgreSQL only)
ALTER SEQUENCE users_id_seq RESTART WITH 1;
ALTER SEQUENCE client_users_id_seq RESTART WITH 1;
ALTER SEQUENCE cinemas_id_seq RESTART WITH 1;
ALTER SEQUENCE studios_id_seq RESTART WITH 1;
ALTER SEQUENCE studio_seat_id_seq RESTART WITH 1;
ALTER SEQUENCE movies_id_seq RESTART WITH 1;
ALTER SEQUENCE showtimes_id_seq RESTART WITH 1;
ALTER SEQUENCE seat_showtimes_id_seq RESTART WITH 1;
ALTER SEQUENCE orders_id_seq RESTART WITH 1;
ALTER SEQUENCE tickets_id_seq RESTART WITH 1;
ALTER SEQUENCE refunds_id_seq RESTART WITH 1;

