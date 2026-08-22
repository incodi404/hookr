import { getChannel } from "./config/index.config";
import amqp from "amqplib"

export async function consumer<T>(
    exchange: string,
    queueName: string,
    patterns: [string],
    callback: (message: T, channel: amqp.Channel) => Promise<void>
): Promise<void> {
    const channel = getChannel()
    if (!channel) {
        throw new Error("[RabbitMQ] Channel not found in RabbitMQ. Maybe the connection is not established yet.")
    }

    try {
        await channel.assertExchange(exchange, "topic", { durable: false }) // create exchange
        const queue = await channel.assertQueue(queueName, { exclusive: true }) // create queue
        patterns.map(async pattern => {
            await channel.bindQueue(queue.queue, exchange, pattern)
        }) // bind queue

        // consume
        channel.consume(queue.queue, async function (msg) {
            if (msg?.content) {
                const stringifyMessage = msg?.content?.toString()
                await callback(JSON.parse(stringifyMessage), channel)
            }
        }, { noAck: false })

    } catch (error: unknown) {
        if (error instanceof Error) {
            throw new Error(
                `[RabbitMQ] Failed to consume message: ${error.message}`
            )
        }

        throw new Error("[RabbitMQ] Failed to consume message.")
    }
}