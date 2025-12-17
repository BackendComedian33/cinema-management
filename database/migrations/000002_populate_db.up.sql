
-- USERS (ADMIN / STAFF)
INSERT INTO users (password, name, email)
VALUES
  ('$2a$12$hu3bCIozthUvPFFglZ0u7eEWc7lyclAnoFy6Em6EfE6ji8sc8jVK2', 'Admin One', 'admin1@cinema.com'),
  ('$2a$12$hu3bCIozthUvPFFglZ0u7eEWc7lyclAnoFy6Em6EfE6ji8sc8jVK2', 'Admin Two', 'admin2@cinema.com');

-- CLIENT USERS (CUSTOMERS)
INSERT INTO client_users (password, name, email)
VALUES
  ('$2a$12$hu3bCIozthUvPFFglZ0u7eEWc7lyclAnoFy6Em6EfE6ji8sc8jVK2', 'John Doe', 'john.doe@email.com'),
  ('$2a$12$hu3bCIozthUvPFFglZ0u7eEWc7lyclAnoFy6Em6EfE6ji8sc8jVK2', 'Jane Smith', 'jane.smith@email.com'),
  ('$2a$12$hu3bCIozthUvPFFglZ0u7eEWc7lyclAnoFy6Em6EfE6ji8sc8jVK2', 'Alice Brown', 'alice.brown@email.com');

-- CINEMAS
INSERT INTO cinemas (name, address)
VALUES
  ('Downtown Cinema', '123 Main Street'),
  ('City Mall Cinema', '456 Mall Avenue');

-- STUDIOS
-- Cinema 1 has 2 studios, Cinema 2 has 1 studio
INSERT INTO studios (cinema_id, name, total_rows, total_columns)
VALUES
  (1, 'Studio 1', 5, 5),
  (1, 'Studio 2', 6, 6),
  (2, 'Studio A', 4, 5);

-- STUDIO SEATS
-- Studio 1 (5x5)
INSERT INTO studio_seat (studio_id, seat_row, seat_column, status)
SELECT
  1,
  r,
  c,
  'AVAILABLE'
FROM generate_series(1,5) r,
     generate_series(1,5) c;

-- Studio 2 (6x6)
INSERT INTO studio_seat (studio_id, seat_row, seat_column, status)
SELECT
  2,
  r,
  c,
  'AVAILABLE'
FROM generate_series(1,6) r,
     generate_series(1,6) c;

-- Studio A (4x5)
INSERT INTO studio_seat (studio_id, seat_row, seat_column, status)
SELECT
  3,
  r,
  c,
  'AVAILABLE'
FROM generate_series(1,4) r,
     generate_series(1,5) c;

-- MOVIES
INSERT INTO movies (title, duration_minutes, rating)
VALUES
  ('The Great Adventure', 120, 'PG-13'),
  ('Romance in Paris', 95, 'PG'),
  ('Haunted Nights', 110, 'R'),
  ('Animated Dreams', 88, 'G');

