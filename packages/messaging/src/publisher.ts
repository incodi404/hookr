import { getChannel } from "./config/index.config";

export async function publishMessage<T>(exchange: string, routingKey: string, msg: T): Promise<void> {
    const channel = getChannel()
    if (!channel) {
        throw new Error("[RabbitMQ] Channel not found in RabbitMQ. Maybe the connection is not established yet.")
    }

    try {
        // create exchange
        await channel.assertExchange(exchange, "topic", { durable: false })

        let stringifyMessage = ""
        if (typeof msg !== "string") {
            stringifyMessage = JSON.stringify(msg)
        } else {
            stringifyMessage = msg
        }

        // publish
        channel.publish(exchange, routingKey, Buffer.from(stringifyMessage))
    } catch (error: unknown) {
        if (error instanceof Error) {
            throw new Error(
                `[RabbitMQ] Failed to publish message: ${error.message}`
            )
        }

        throw new Error("[RabbitMQ] Failed to publish message.")
    }
}