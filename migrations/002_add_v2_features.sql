-- Migration for V2 features
-- Add new tables for membership, driver, booking type, and driver incentive

-- Create membership table
CREATE TABLE IF NOT EXISTS memberships (
    membership_id SERIAL PRIMARY KEY,
    membership_name VARCHAR(50) NOT NULL,
    discount_percentage DECIMAL(5,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create driver table
CREATE TABLE IF NOT EXISTS drivers (
    driver_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    nik VARCHAR(20) UNIQUE NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    daily_cost DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create booking type table
CREATE TABLE IF NOT EXISTS booking_types (
    booking_type_id SERIAL PRIMARY KEY,
    booking_type VARCHAR(50) NOT NULL,
    description VARCHAR(200),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create driver incentive table
CREATE TABLE IF NOT EXISTS driver_incentives (
    incentive_id SERIAL PRIMARY KEY,
    rental_id INTEGER NOT NULL,
    incentive DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (rental_id) REFERENCES rentals(rental_id) ON DELETE CASCADE
);

-- Alter customers table to add membership_id
ALTER TABLE customers ADD COLUMN IF NOT EXISTS membership_id INTEGER;
ALTER TABLE customers ADD CONSTRAINT fk_customer_membership 
    FOREIGN KEY (membership_id) REFERENCES memberships(membership_id) ON DELETE SET NULL;

-- Alter rentals table to add new columns
ALTER TABLE rentals ADD COLUMN IF NOT EXISTS discount DECIMAL(10,2) DEFAULT 0;
ALTER TABLE rentals ADD COLUMN IF NOT EXISTS booking_type_id INTEGER NOT NULL DEFAULT 1;
ALTER TABLE rentals ADD COLUMN IF NOT EXISTS driver_id INTEGER;
ALTER TABLE rentals ADD COLUMN IF NOT EXISTS total_driver_cost DECIMAL(10,2) DEFAULT 0;

-- Add foreign key constraints to rentals table
ALTER TABLE rentals ADD CONSTRAINT fk_rental_booking_type 
    FOREIGN KEY (booking_type_id) REFERENCES booking_types(booking_type_id);
ALTER TABLE rentals ADD CONSTRAINT fk_rental_driver 
    FOREIGN KEY (driver_id) REFERENCES drivers(driver_id) ON DELETE SET NULL;

-- Insert default data
INSERT INTO memberships (membership_name, discount_percentage) VALUES
('Bronze', 4.00),
('Silver', 7.00),
('Gold', 15.00);

INSERT INTO booking_types (booking_type, description) VALUES
('Car Only', 'Rent Car only'),
('Car & Driver', 'Rent Car and a Driver');