import amqp from "amqplib"

let rabbitmqConnection: amqp.ChannelModel | null = null
let rabbitmqChannel: amqp.Channel | null = null

export const connectRabbitMQ = async () => {
    let retryCounter = 5
    let isConnected = false

    while (retryCounter) {
        try {
            const url = `${process.env?.RABBITMQ_USERNAME}:${process.env?.RABBITMQ_PASSWORD}@${process.env?.AMQP_CONNECT_URL}`
            const connection = await amqp.connect(`amqp://${url || ""}`)

            rabbitmqConnection = connection
            rabbitmqChannel = await connection.createChannel()
            isConnected = true

            console.log("[RabbitMQ] RabbitMQ is connected!")
            break
        } catch (error) {
            console.log("[RabbitMQ] Connection has been failed with RabbitMQ :: Attempt: ", retryCounter)
            isConnected = false
            retryCounter--

            await new Promise((res) => setTimeout(res, 5000)) // 5 sec wait
        }
    }

    return { rabbitmqConnection, isConnected }
}

export function getChannel(): amqp.Channel | null {
    return rabbitmqChannel
}