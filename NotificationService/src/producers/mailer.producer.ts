import logger from "../config/logger.config";
import { SendEmailDTO } from "../dtos/notification.dto";
import { mailerQueue } from "../queues/mailer.queue";

export const SEND_EMAIL_PAYLOAD = `mailer-payload`;

export const addEmailToQueue = async (emailPayload:SendEmailDTO) => {
    mailerQueue.add(SEND_EMAIL_PAYLOAD, emailPayload);
    logger.info(`Email added to queue for ${emailPayload.to}`);
}