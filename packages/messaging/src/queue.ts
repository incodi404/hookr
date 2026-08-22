export const QUEUES = {
    EMAIL: "email-delivery",
    WS: "websocket-notification",
    DELIVERY: "hook-delivery",
    DLQ: "delay-retry-queue"
} as const