BEGIN;

-- USERS
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  password VARCHAR NOT NULL,
  name VARCHAR NOT NULL,
  email VARCHAR NOT NULL UNIQUE,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

-- CLIENT USERS
CREATE TABLE client_users (
  id BIGSERIAL PRIMARY KEY,
  password VARCHAR NOT NULL,
  name VARCHAR NOT NULL,
  email VARCHAR NOT NULL UNIQUE,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

-- CINEMAS
CREATE TABLE cinemas (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  address VARCHAR
);

-- STUDIOS
CREATE TABLE studios (
  id SERIAL PRIMARY KEY,
  cinema_id INT NOT NULL,
  name VARCHAR NOT NULL,
  total_rows INT NOT NULL,
  total_columns INT NOT NULL,
  CONSTRAINT fk_studios_cinema
    FOREIGN KEY (cinema_id)
    REFERENCES cinemas(id)
    ON DELETE CASCADE
);

-- STUDIO SEATS
CREATE TABLE studio_seat (
  id BIGSERIAL PRIMARY KEY,
  studio_id INT NOT NULL,
  seat_row INT NOT NULL,
  seat_column INT NOT NULL,
  status VARCHAR NOT NULL, -- AVAILABLE, UNAVAILABLE
  CONSTRAINT fk_studio_seat_studio
    FOREIGN KEY (studio_id)
    REFERENCES studios(id)
    ON DELETE CASCADE,
  CONSTRAINT unique_seat_position
    UNIQUE (studio_id, seat_row, seat_column)
);

-- MOVIES
CREATE TABLE movies (
  id SERIAL PRIMARY KEY,
  title VARCHAR NOT NULL,
  duration_minutes INT NOT NULL,
  rating VARCHAR
);

-- SHOWTIMES
CREATE TABLE showtimes (
  id BIGSERIAL PRIMARY KEY,
  movie_id INT NOT NULL,
  studio_id INT NOT NULL,
  show_date DATE NOT NULL,
  start_time TIME NOT NULL,
  duration_minutes int not null,
  status VARCHAR NOT NULL, -- ACTIVE, CANCELLED
  CONSTRAINT fk_showtimes_movie
    FOREIGN KEY (movie_id)
    REFERENCES movies(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_showtimes_studio
    FOREIGN KEY (studio_id)
    REFERENCES studios(id)
    ON DELETE CASCADE
);

-- SEAT SHOWTIMES
CREATE TABLE seat_showtimes (
  id BIGSERIAL PRIMARY KEY,
  showtime_id BIGINT NOT NULL,
  studio_seat_id BIGINT NOT NULL,
  status VARCHAR NOT NULL, -- AVAILABLE, LOCKED, SOLD
  locked_until TIMESTAMP,
  CONSTRAINT fk_seat_showtimes_showtime
    FOREIGN KEY (showtime_id)
    REFERENCES showtimes(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_seat_showtimes_studio_seat
    FOREIGN KEY (studio_seat_id)
    REFERENCES studio_seat(id)
    ON DELETE CASCADE,
  CONSTRAINT unique_seat_per_showtime
    UNIQUE (showtime_id, studio_seat_id)
);

-- ORDERS
CREATE TABLE orders (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  showtime_id BIGINT NOT NULL,
  order_status VARCHAR NOT NULL, -- PENDING, PAID, CANCELLED, REFUNDED
  total_price DECIMAL(10,2) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
  payment_method VARCHAR,
  payment_status VARCHAR,
  paid_at TIMESTAMP,
  CONSTRAINT fk_orders_user
    FOREIGN KEY (user_id)
    REFERENCES client_users(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_orders_showtime
    FOREIGN KEY (showtime_id)
    REFERENCES showtimes(id)
    ON DELETE CASCADE
);

-- TICKETS
CREATE TABLE tickets (
  id BIGSERIAL PRIMARY KEY,
  order_id BIGINT NOT NULL,
  seat_showtime_id BIGINT NOT NULL,
  ticket_code VARCHAR NOT NULL UNIQUE,
  CONSTRAINT fk_tickets_order
    FOREIGN KEY (order_id)
    REFERENCES orders(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_tickets_seat_showtime
    FOREIGN KEY (seat_showtime_id)
    REFERENCES seat_showtimes(id)
    ON DELETE CASCADE
);

-- REFUNDS
CREATE TABLE refunds (
  id BIGSERIAL PRIMARY KEY,
  order_id BIGINT NOT NULL,
  refund_reason VARCHAR,
  refund_status VARCHAR NOT NULL, -- REQUESTED, COMPLETED
  refunded_at TIMESTAMP,
  CONSTRAINT fk_refunds_order
    FOREIGN KEY (order_id)
    REFERENCES orders(id)
    ON DELETE CASCADE
);

COMMIT;
