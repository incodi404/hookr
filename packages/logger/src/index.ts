import pino from "pino";

export const logger = (dest: string) => pino({
    level: "info"
}, pino.destination(dest))