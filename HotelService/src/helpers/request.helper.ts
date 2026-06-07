import { AsyncLocalStorage } from "async_hooks";

type AsyncLocalStorageType = {
    correlationId: string;
}

export const asyncLocalStorage = new AsyncLocalStorage<AsyncLocalStorageType>(); // Created an instance of AsyncLocalStorage

export const runWithCorrelationId = <T>(correlationId: string | undefined, callback: () => T) => {
    return asyncLocalStorage.run({
        correlationId: correlationId || 'unknown-error-while-creating-correlation-id',
    }, callback);
}

export const getCorrelationId = () => {
    const asyncStore = asyncLocalStorage.getStore();
    return asyncStore?.correlationId || 'unknown-error-while-creating-correlation-id'; // Default value if not found 
}
