import express from "express";
import {serverConfig} from "./config";
import v1Router from "./routers/v1/index.router";
import v2Router from "./routers/v2/index.router";
import { genericErrorMiddleware as genericErrorHandler } from "./middlewares/errors/error.middleware";
import logger from "./config/logger.config";
import { attachCorrelationId } from "./middlewares/correlationid.middleware";
import { setupEmailWorker } from "./workers/mailer.worker";

const app = express();

app.use(express.json());

app.use(attachCorrelationId);

app.use('/api/v1', v1Router);
app.use('/api/v2', v2Router);

app.use(genericErrorHandler);

app.listen(serverConfig.PORT, ()=>{
    logger.info(`Server is running on port ${serverConfig.PORT}`);

    setupEmailWorker();
});