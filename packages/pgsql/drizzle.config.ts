import { defineConfig } from "drizzle-kit";
import dotenv from "dotenv";

dotenv.config({
  path: "../../.env",
});

export default defineConfig({
  schema: "./src/schema/**/*.{ts,js}",
  dialect: "postgresql",
  out: "./migrations",
  dbCredentials: {
    host: process.env.PGSQL_HOST || "",
    port: 5432,
    user: process.env.PGSQL_USERNAME || "",
    password: process.env.PGSQL_PASSWORD || "",
    database: process.env.PGSQL_DB || "",
    ssl: false,
  },
});
