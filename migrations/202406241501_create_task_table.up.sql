CREATE TABLE IF NOT EXISTS public.tasks (
    task_id SERIAL PRIMARY KEY,
    task_name VARCHAR(255) COLLATE pg_catalog."default",
    task_description TEXT COLLATE pg_catalog."default",
    solution_code TEXT COLLATE pg_catalog."default",
    test_cases JSONB
);
