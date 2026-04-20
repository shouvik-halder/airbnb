import nodemailer from 'nodemailer';
import { nodemailerConfig } from '.';

const transporter = nodemailer.createTransport({
  service: 'gmail',
  auth: {
    user: nodemailerConfig.SMTP_USER,
    pass: nodemailerConfig.SMTP_PASS,
  },
});

export default transporter;