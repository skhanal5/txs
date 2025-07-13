CREATE TABLE IF NOT EXISTS currency_codes (
    code VARCHAR(3) NOT NULL UNIQUE PRIMARY KEY,
    description TEXT
)

INSERT INTO currency_codes (code, description) VALUES
  ('USD', 'United States Dollar');

CREATE TABLE IF NOT EXISTS account_types (
    code VARCHAR(12) NOT NULL UNIQUE PRIMARY KEY,
    description TEXT
)

INSERT INTO account_types (code, description) VALUES
  ('saving', 'Saving Account'),
  ('checking', 'Checking Account');

CREATE TABLE IF NOT EXISTS account_statuses (
    code VARCHAR(12) NOT NULL UNIQUE PRIMARY KEY,
    description TEXT
)

INSERT INTO account_statuses (code, description) VALUES
  ('active', 'Account is active and operational'),
  ('frozen', 'Account is temporarily frozen'),
  ('closed', 'Account is closed');

CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) NOT NULL ON DELETE CASCADE,
    balance NUMERIC(12, 2) NOT NULL DEFAULT 0.00,
    currency_code VARCHAR(3) REFERENCES currency_codes(code) NOT NULL,
    status VARCHAR(12) REFERENCES account_statuses(code) NOT NULL DEFAULT 'active',
    type VARCHAR(12) REFERENCES account_types(code) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_updated_at 
    BEFORE UPDATE ON accounts
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();
