import { Worker } from "bullmq";
import { SendEmailDTO } from "../dtos/notification.dto";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisClient } from "../config/redis.config";
import logger from "../config/logger.config";
import { SEND_EMAIL_PAYLOAD } from "../producers/mailer.producer";
import { BadRequestError } from "../utils/errors/app.error";
import { renderMailerTemplates } from "../templates/template.handlers";
import { MailerService } from "../services/mailer.service";


export const setupEmailWorker = () =>{
    
    const emailProcessor = new Worker<SendEmailDTO>(
        MAILER_QUEUE,
        async (job) =>{
            if(job.name !== SEND_EMAIL_PAYLOAD){
                logger.error(`Unknown job type ${job.name}`);
                throw new BadRequestError(`Unknown job type ${job.name}`);
            }

            const payload = job.data;
            logger.info(`Processing email template`);

            const emailContent = await renderMailerTemplates(payload.templateId, payload.data);

            await MailerService(payload.to, payload.subject, emailContent);
    
        },
        {
            connection: getRedisClient()
        }
    );
    
    emailProcessor.on("completed", (job)=>{
        logger.info(`Email job completed for ${job.data.to}`);
    });
    
    emailProcessor.on("failed", (job, err)=>{
        logger.error(`Email job failed for ${job?.data.to} with error ${err.message}`);
    })
}