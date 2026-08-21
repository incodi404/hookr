import amqp from "amqplib"

let rabbitmqConnection: amqp.Connection | null = null
let rabbitmqChannel: amqp.Channel | null = null

export const connectRabbitMQ = () => {
    let retryCounter = 5

    while (retryCounter) {
        try {
            const url = `${process.env?.RABBITMQ_USERNAME}:${process.env?.RABBITMQ_PASSSWORD}@${process.env?.AMQP_CONNECT_URL}`
        } catch (error) {

        }
    }
}