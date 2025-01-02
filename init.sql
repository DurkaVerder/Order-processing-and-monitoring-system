

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,                          -- Уникальный идентификатор заказа
    customer_name VARCHAR(255) NOT NULL,            -- Имя клиента
    customer_email VARCHAR(255) NOT NULL,           -- Email клиента
    description TEXT,                               -- Описание заказа
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата и время создания заказа
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата и время последнего обновления
    amount NUMERIC(10, 2) NOT NULL                  -- Сумма заказа
);

CREATE TABLE status_orders (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,                          -- Ссылка на id из таблицы orders
    status VARCHAR(255) NOT NULL                    -- Статус заказа
);

CREATE TABLE reports (
    id SERIAL PRIMARY KEY,                          -- Уникальный идентификатор отчета
    status VARCHAR(255) NOT NULL,                   -- Статус отчета
    date_time TIMESTAMP NOT NULL                    -- Дата и время отчета
);
