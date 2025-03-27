package creation

const TableName string = "sneakers"
const CreateSneakersTableSQL string = `
CREATE TABLE IF NOT EXISTS sneakers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title CHAR(40) UNIQUE NOT NULL,
    description CHAR(100) NOT NULL,
    price DECIMAL NOT NULL, 
    created_at TIMESTAMP DEFAULT now()
);
`
