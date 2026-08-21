import { Pool } from "pg";

export const pgPool = new Pool({
    host: process.env?.PGSQL_HOST || "",
    port: 5432,
    user: process.env?.PGSQL_USERNAME || "",
    password: process.env?.PGSQL_PASSWORD || "",
    max: 5,
    idleTimeoutMillis: 30000,
    connectionTimeoutMillis: 2000,
});
