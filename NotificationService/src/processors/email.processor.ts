import { Job, Worker } from "bullmq";
import { NotificationDto } from "../dtos/notification.dto";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producer/email.producer";
import { renderMailTemplate } from "../templates/templates.handler";
import { sendEmail } from "../services/mailer.service";
import logger from "../config/logger.config";

export const setUpMailerWorker = () =>{
    const emailProcessor = new Worker<NotificationDto>(
        MAILER_QUEUE,async (job:Job) =>{
            if(job.name !== MAILER_PAYLOAD){
                throw new Error("Invalid job name")
            }

            // call the service layer from here
            const payload = job.data;
            console.log(`processing email for: ${JSON.stringify(payload)}`);
            const emailContent = await renderMailTemplate(payload.templateId,payload.params);
            await sendEmail(payload.to,payload.subject,emailContent);
            logger.info(`Email sent to ${payload.to} with subject "${payload.subject}"`)

        },// process function
        {
            connection: getRedisConnObject()
        }
    )

    emailProcessor.on("failed",()=>{
        console.log("Email processor failed");
    })

    emailProcessor.on("completed",()=>{
        console.log("Email processing completed successfully")
    })
}

