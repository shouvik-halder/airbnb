import { CreationAttributes, Model, ModelStatic, WhereOptions } from "sequelize";

export abstract class BaseRepository<T extends Model> {
    protected model: ModelStatic<T>;

    constructor(model: ModelStatic<T>) {
        this.model = model;
    }

    async findById(id: number): Promise<T | null> {
        return this.model.findByPk(id);
    }
    
    async findAll(): Promise<T[]> {
        return this.model.findAll();
    }
    
    async create(data: CreationAttributes<T>): Promise<T> {
        return this.model.create(data);
    }

    async update(id: number, data: Partial<CreationAttributes<T>>): Promise<T | null> {
        const instance = await this.model.findByPk(id);
        if(!instance)
            return null;
        await instance.update(data);
        return instance;
    }

    async softDelete(whereOptions: WhereOptions<T>): Promise<boolean> {
        const instance = await this.model.destroy({
            where: {
                ...whereOptions
            }
        });
        return instance>0?true:false;
    }


}

