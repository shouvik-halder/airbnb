import {v4 as uuidV4} from 'uuid';
export function GenerateIdempotencyKey():string {
    return uuidV4();
}