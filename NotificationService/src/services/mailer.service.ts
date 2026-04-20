import { nodemailerConfig } from "../config";
import transporter from "../config/mailer.config";
import { InternalServerError } from "../utils/errors/app.error";

export async function MailerService(to:string, subject:string, body:string) {
    try {
        await transporter.sendMail({
            from:nodemailerConfig.SMTP_USER,
            to,
            subject,
            html:body
        })
        
    } catch (error) {
        throw new InternalServerError(`${(error as Error).message}`);
    }
}