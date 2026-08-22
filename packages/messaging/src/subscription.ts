// These are routing keys.
export const SUBSCRIPTIONS = {
    "email-delivery": [
        "error.*",
        "auth.*"
    ],
    "websocket-notification": [
        "delivery.error.*",
        "delivery.*",
        "error.*"
    ],
    "hook-delivery": [
        "delivery.*"
    ],
    "delay-retry-queue": [
        "delivery.error.*"
    ]
} as const