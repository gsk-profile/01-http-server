DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users') THEN
        DELETE FROM users WHERE email IN ('demo@example.com', 'admin@example.com');
    END IF;
END $$;