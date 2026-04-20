export type SendEmailDTO = {
    to:string;
    subject:string;
    templateId:string;
    data:Record<string, any>;
}